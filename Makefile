local_bin=~/.local/bin


build:
	go build -o backup main.go 

install: build
	cp backup $(local_bin)/backup
