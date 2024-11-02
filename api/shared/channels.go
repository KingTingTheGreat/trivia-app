package shared

var (
	LeaderboardChan = make(chan bool, 100)
	BuzzedInChan    = make(chan bool, 100)
	PlayerListChan  = make(chan bool, 100)
)
