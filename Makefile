install: 
	npm install

clean: 
	(rm -r .next && rm go-backend.exe; exit 0)

build: clean install
	npm run build

run: build
	npm run start
