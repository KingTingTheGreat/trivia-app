package shared

import (
	"bytes"
	"errors"
	"math/rand"
	"strings"
	"sync"
	"time"
	"trivia-app/dlog"
	"trivia-app/handlers"

	"github.com/gorilla/websocket"
)

// generates a random 64 character token
func generateToken() string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*_+-=<>?")
	var token string
	for i := 0; i < 64; i++ {
		token += string(chars[rand.Intn(len(chars))])
	}
	return token
}

type Player struct {
	CleanName        string
	Name             string
	Score            int
	ButtonReady      bool
	CorrectQuestions []string
	LastUpdate       time.Time
	BuzzedIn         time.Time
	Websocket        *websocket.Conn
	WsClose          chan bool
}

type UpdatePlayer struct {
	ScoreDiff   *int
	ButtonReady *bool
	LastUpdate  *time.Time
	BuzzedIn    *time.Time
	Websocket   *websocket.Conn
}

type playerStore struct {
	mu          sync.RWMutex
	playerData  map[string]Player
	playerNames map[string]string
}

func cleanName(name string) string {
	return strings.TrimSpace(strings.ToLower(name))
}

// returns existing player data
func (ps *playerStore) GetPlayer(token string) (Player, bool) {
	dlog.DLog("GetPlayer()")
	ps.mu.RLock()
	player, ok := ps.playerData[token]
	ps.mu.RUnlock()
	dlog.DLog("leaving GetPlayer()")
	return player, ok
}

// creates a new player and returns the token
func (ps *playerStore) InsertPlayer(playerName string) (string, error) {
	cleanName := cleanName(playerName)
	if playerName == "" || cleanName == "" {
		return "", errors.New("no player name provided")
	}

	dlog.DLog("InsertPlayer()")
	defer dlog.DLog("leaving InsertPlayer()")
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if ps.playerNames[cleanName] != "" {
		dlog.DLog(ps.playerData[ps.playerNames[cleanName]])
		return "", errors.New("a player with this name already exists")
	}

	token := generateToken()
	for _, ok := ps.playerData[token]; ok; {
		token = generateToken()
	}

	player := Player{
		CleanName:        cleanName,
		Name:             playerName,
		Score:            0,
		ButtonReady:      true,
		CorrectQuestions: make([]string, 0),
		LastUpdate:       time.Now(),
		BuzzedIn:         time.Time{},
		Websocket:        nil,
		WsClose:          make(chan bool, 1),
	}

	ps.playerData[token] = player
	ps.playerNames[cleanName] = token

	dlog.DLog("Number of players:", len(ps.playerData))
	return token, nil
}

// updates an existing player
func (ps *playerStore) PutPlayer(token string, playerUpdates UpdatePlayer) (Player, error) {
	dlog.DLog("PutPlayer()")
	defer dlog.DLog("leaving PutPlayer()")
	ps.mu.Lock()
	defer ps.mu.Unlock()

	var player Player
	var ok bool
	// check this player exists
	if player, ok = ps.playerData[token]; !ok {
		dlog.DLog("player not found")
		return Player{}, errors.New("player not found")
	}

	// update the player attributes
	if playerUpdates.ScoreDiff != nil {
		player.Score += *playerUpdates.ScoreDiff
	}
	if playerUpdates.ButtonReady != nil {
		player.ButtonReady = *playerUpdates.ButtonReady
	}
	if playerUpdates.LastUpdate != nil {
		player.LastUpdate = *playerUpdates.LastUpdate
	}
	if playerUpdates.BuzzedIn != nil {
		player.BuzzedIn = *playerUpdates.BuzzedIn
	}
	if playerUpdates.Websocket != nil {
		dlog.DLog("updating player with websocket")
		if player.Websocket != nil {
			ps.mu.Unlock()
			// dlog.DLog("CLOSING EXISTING WEBSOCKET")
			player.WsClose <- true
			ps.mu.Lock()
			// } else {
			// dlog.DLog("NO EXISTING WEBSOCKET :)")
		}
		player.Websocket = playerUpdates.Websocket
	}

	player.LastUpdate = time.Now()
	ps.playerData[token] = player

	return player, nil
}

// updates an existing player's score to 0
func (ps *playerStore) ZeroPlayer(token string) error {
	dlog.DLog("ZeroPlayer()")
	defer dlog.DLog("leaving ZeroPlayer()")
	ps.mu.Lock()
	defer ps.mu.Unlock()

	var player Player
	var ok bool
	// check this player exists
	if player, ok = ps.playerData[token]; !ok {
		dlog.DLog("player not found")
		return errors.New("player not found")
	}

	player.Score = 0
	player.LastUpdate = time.Now()
	ps.playerData[token] = player

	return nil
}

func (ps *playerStore) NilPlayerWS(token string, ws *websocket.Conn) error {
	dlog.DLog("NilPlayerWS()")
	defer dlog.DLog("leaving NilPlayerWS()")
	ps.mu.Lock()
	defer ps.mu.Unlock()

	var player Player
	var ok bool
	// check this player exists
	if player, ok = ps.playerData[token]; !ok {
		dlog.DLog("player not found")
		return errors.New("player not found")
	}

	if player.Websocket == ws {
		player.Websocket = nil
		ps.playerData[token] = player
	}

	return nil
}

// deletes the player
func (ps *playerStore) DeletePlayer(token string) error {
	dlog.DLog("DeletePlayer()")
	defer dlog.DLog("leaving DeletePlayer()")
	ps.mu.Lock()
	defer ps.mu.Unlock()
	var player Player
	var ok bool
	if player, ok = ps.playerData[token]; !ok {
		return errors.New("player not found")
	}

	if player.Websocket != nil {
		player.WsClose <- true
	}

	dlog.DLog("deleting player from store")
	delete(ps.playerData, token)
	delete(ps.playerNames, player.CleanName)
	return nil
}

// returns a list of all players
func (ps *playerStore) AllPlayers() []*Player {
	dlog.DLog("AllPlayers()")
	var allPlayers []*Player
	ps.mu.RLock()
	for _, player := range ps.playerData {
		allPlayers = append(allPlayers, &player)
	}
	dlog.DLog("leaving AllPlayers()")
	ps.mu.RUnlock()
	return allPlayers
}

type TokenPlayer struct {
	Token  string
	Player Player
}

// returns a list of all tokens and their corresponding players
func (ps *playerStore) AllTokenPlayers() []*TokenPlayer {
	dlog.DLog("AllTokenPlayers()")
	var allPlayers []*TokenPlayer
	ps.mu.RLock()
	for token, player := range ps.playerData {
		allPlayers = append(allPlayers, &TokenPlayer{token, player})
	}
	dlog.DLog("leaving AllTokenPlayers()")
	ps.mu.RUnlock()
	return allPlayers
}

type PlayerNameToken struct {
	Name  string
	Token string
}

// returns a list of all player names and their corresponding tokens
func (ps *playerStore) AllNamesTokens() []*PlayerNameToken {
	dlog.DLog("AllNamesTokens()")
	ps.mu.RLock()
	allNamesTokens := make([]*PlayerNameToken, len(ps.playerNames))
	i := 0
	for name, token := range ps.playerNames {
		allNamesTokens[i] = &PlayerNameToken{
			Name:  name,
			Token: token,
		}
		i++
	}
	dlog.DLog("leaving AllNamesTokens()")
	ps.mu.RUnlock()
	return allNamesTokens
}

// returns the token for a given player name
func (ps *playerStore) NameToToken(name string) (string, bool) {
	dlog.DLog("NameToToken()")
	ps.mu.RLock()
	token, ok := ps.playerNames[cleanName(name)]
	dlog.DLog("leaving NameToToken()")
	ps.mu.RUnlock()
	return token, ok
}

// returns a boolean indicating if the token and name match
func (ps *playerStore) VerifyTokenName(token, name string) bool {
	dlog.DLog("VerifyTokenName()")
	defer dlog.DLog("leaving VerifyTokenName()")
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	storedToken, ok := ps.playerNames[cleanName(name)]
	return ok && storedToken == token
}

func (ps *playerStore) BuzzIn(token, name string) bool {
	dlog.DLog("BuzzIn()")
	defer dlog.DLog("leaving ResetBuzzers()")
	ps.mu.Lock()
	player, ok := ps.playerData[token]
	if !ok {
		ps.mu.Unlock()
		return false
	}
	if !player.BuzzedIn.IsZero() || player.Name != name {
		dlog.DLog("already buzzed in")
		ps.mu.Unlock()
		return false
	}
	ps.mu.Unlock()

	now := time.Now()
	f := false
	_, err := ps.PutPlayer(token, UpdatePlayer{BuzzedIn: &now, ButtonReady: &f})
	if err != nil {
		dlog.DLog("failed to buzz into player store")
		return false
	}

	return true
}

func (ps *playerStore) ResetBuzzers() {
	dlog.DLog("ResetBuzzers()")
	defer dlog.DLog("leaving ResetBuzzers()")
	ps.mu.Lock()
	defer ps.mu.Unlock()
	dlog.DLog("got lock")

	// buffer containing ready buzzer
	var buf bytes.Buffer
	handlers.RenderComponent(&buf, "buzz-button.html", handlers.Play{Ready: true})

	for token, player := range ps.playerData {
		// update vars
		player.ButtonReady = true
		player.BuzzedIn = time.Time{}
		dlog.DLog("reseting buzzer for", player.Name)

		// send message to websocket/client
		if player.Websocket != nil {
			err := player.Websocket.WriteMessage(websocket.TextMessage, buf.Bytes())
			// FIX THIS ERROR HANDLING. probably causes error with handler
			if err != nil {
				dlog.DLog("error reseting buzzer")
				player.Websocket.Close()
				player.Websocket = nil
			}
		} else {
			dlog.DLog("websocket is nil")
		}

		ps.playerData[token] = player
	}
}

func (ps *playerStore) ResetGame() {
	dlog.DLog("ResetGame()")
	defer dlog.DLog("leaving ResetGame()")
	ps.mu.Lock()
	defer ps.mu.Unlock()

	// close websocket connections
	for _, player := range ps.playerData {
		player.WsClose <- true
	}

	// reset player data and names/tokens
	ps.playerData = make(map[string]Player)
	ps.playerNames = make(map[string]string)
}

var PlayerStore playerStore = playerStore{
	mu:          sync.RWMutex{},
	playerData:  make(map[string]Player),
	playerNames: make(map[string]string),
}
