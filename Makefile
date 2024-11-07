install: 
	npm install

clean: 
	(rm -r .next && (rm trivia-app || rm trivia-app.exe); exit 0)

build: clean install
	npm run build

initialize: 
	go run init/init.go

run: initialize build
	npm run start
