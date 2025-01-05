# Trivia

A live buzzer system

## Setup

### Requirements

- Install [Go version 1.23 or higher](https://go.dev/doc/install)

## How to run

- Run the **run.sh** file
  - You will then be prompted to set your password and IP address
- You can also manually intialize and run
    1. Manually create and configure a file called *.env* or use the command ```go run init/init.go```
         - The file should include values of the form ```IP="..."``` and ```PASSWORD="..."```
    4. Build the project using the command ```go build -o trivia-app.exe```
    5. Finally, run the output executable called **trivia-app.exe**

## Pages 

### Host

- Visit _HostIP_:8080/control to access controls
  - Enter the configured password
  - Select a player's name and the number of points you'd like to give them
    - A positive value will add to their score while a negative value will subtract from it. Negative scores are possible.
  - Reset Buzzers will clear the current buzzed in list/state but will not affect anything else
  - Removing a player permanently deletes all of their information
  - Reset Game will permanently delete all players and all of their information
- Visit and display the _HostIP_:8080/join page to display a QR codes for players to easily access the website

### Players

1. Visit _HostIP_:8080 while on the same network or scan the QR code on the /join page
2. Enter your name and advance to the next page
3. You will now see a buzzer, clicking it will buzz you in

### Additional Pages

#### Leaderboard

- See the players ranked by score

#### Buzzed-In

- See the players as they buzz in
- This page should play a buzzer sound when players buzz in

#### Game

- You can see both the leaderboard and the buzzed in information on this page
