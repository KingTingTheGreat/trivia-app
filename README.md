# Trivia

A live buzzer system

## Setup

### Backend

- Install [Go version 1.23 or higher](https://go.dev/doc/install)

### Frontend

- Install [Node.js](https://nodejs.org/en/download/)

## How to run

- Use the command ```make run```
  - You will then be prompted to set your password and IP address
- You can also manually intialize and run
    1. Manually configre a file called *.env.local* or use the command ```go run init/init.go```
    2. Install frontend dependencies using the command ```npm install```
    3. Build the project using the command ```npm run build```
    4. Finally, run it using the command ```npm run start```

## Pages 

### Host

- Visit _HostIP_:3000/control to access controls
  - Enter the configured password
  - Select a player's name and the number of points you'd like to give them
    - A positive value will add to their score while a negative value will subtract from it. Negative scores are possible.
  - Reset Buzzers will clear the current buzzed in list/state but will not affect anything else
  - Removing a player permanently deletes all of their information
- Visit and display the _HostIP:3000_/join page to display a QR codes for players to easily access the website

### Players

- Visit _HostIP_:3000 while on the same network or scan the QR code on the /join page
- Enter your name and advance to the next page
- You will now see a buzzer, clicking it will buzz you in

### Additional Pages

#### Leaderboard

- Used to only view the players ranked by score

#### Buzzed-In

- See the players as they buzz in for the current question
- This page should play a buzzer sound when players buzz in

#### Game

- You can see both the leaderboard and the buzzed in information on this page
