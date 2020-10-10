local_bin=~/.local/bin

todo:
	ack "// TODO" --type=go

build:
	go build -o backup main.go 

install: build
	cp -f backup $(local_bin)/backup
