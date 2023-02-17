dev:
	go run main.go server dev :8090
build:
	make clean
	mkdir bin
	go build -o bin/creative ./main.go
clean:
	rm -rf bin
migratedev:
	go run main.go migrate dev