all:
	go get github.com/tools/godep
	godep go build -tags seccomp -o runc .

install:
	cp runc /usr/local/bin/runc
	rm runc

clean:
	rm runc
