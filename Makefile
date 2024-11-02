install: 
	npm install

clean: 
	(rm -r .next && rm go-backend.exe; exit 0)

build: clean install
	npm run build

ip: 
	@current_ip=$$(curl -s ifconfig.me); \
	if grep -q "^IP=" .env.local; then \
		sed -i "s/^IP=.*/IP=\"$$current_ip\"/" .env.local; \
	else \
		echo "IP=\"$$current_ip\"" >> .env.local; \
	fi

run: build
	npm run start
