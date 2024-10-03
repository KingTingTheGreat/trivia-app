# Trivia

A live buzzer to play trivia on your local network

## Setup

### Backend

- Install [Go version 1.23 or higher](https://go.dev/doc/install)
- Create a **.env.local** file with the path go-backend/.env and add a password; PASSWORD="yourpasswordhere"

### Frontend

- Install [Node.js](https://nodejs.org/en/download/)
- Edit the file nextjs-frontend/ip.ts to include your device's IP address on your local network

## How to run

- Use the command ```npm run build``` to build the program
- Use the command ```npm run start``` to run it

### Players
-   [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
-   [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

- Visit _HostIP_:3000 while on the same network
- Enter your name and advance to the next page
- You will now see a buzzer, clicking it will notify the host

### Host

- Visit _HostIP_:3000/host to view the players who have buzzed in sorted in chronological order, as well as the time they buzzed in down to the millisecond
- This page will also display the players ranked by score, score seniority is used as a tie-breaker when scores are equal
- Go to _HostIP_:3000/control to access controls
  - Enter the password you put in your **.env.local** file
  - Enter a player's name and the number of points you'd like to give them
    - A positive value will add to their score while a negative value will subtract from it. Negative scores are possible.
  - Reset Buzzers will clear the current buzzed in list/state but will not affect anything else

### Additional Pages

#### Leaderboard

- Used to only view the players ranked by score

#### Stats

- View players ranked by score as well as the questions each player has received points for

#### Buzzed-In

- See the players as they buzz in for the current question
- This page should play a buzzer sound when players buzz in
