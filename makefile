run:
	go run main.go

build:
	go build -o bin/main main.go

test:
	go build -o bin/main main.go && ./input_test.sh