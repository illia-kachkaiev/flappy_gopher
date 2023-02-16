GAME_FILE_NAME=flappy_gopher

build:
	go build -o $(GAME_FILE_NAME)

run:
	./$(GAME_FILE_NAME)


test:
	go test