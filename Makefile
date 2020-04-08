CC=`which go`
SRC=main.go move.go watch.go config.go

all:
	$(CC) run $(SRC)

build:
	$(CC) build -o auto-move $(SRC)

install:
	$(CC) install