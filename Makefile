gen:
	cd games && \
	go run github.com/99designs/gqlgen@v0.17.45
	cd players && \
	go run github.com/99designs/gqlgen@v0.17.45
	cd leaderboards && \
	go run github.com/99designs/gqlgen@v0.17.45
	cd monolith && \
	go run github.com/99designs/gqlgen@v0.17.45

up: down
	go run ./games &
	go run ./players &
	go run ./leaderboards &
	cd router && \
	  npm run start

mono: down
	go run ./monolith &

down:
	-sudo kill $(shell sudo lsof -ti:8080)
	-sudo kill $(shell sudo lsof -ti:8081)
	-sudo kill $(shell sudo lsof -ti:8082)
	-sudo kill $(shell sudo lsof -ti:8083)
