local_bin=~/.local/bin


build:
	go build -o backup main.go 

install: build
	cp -f backup $(local_bin)/backup
