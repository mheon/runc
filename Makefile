all: container

container:
	docker build -t test/runc .
	docker run -t -i --name runc test/runc
	docker cp runc:/gopath/src/github.com/opencontainers/runc/runc .
	docker rm -f runc

install:
	cp runc /usr/local/bin/runc
	rm runc

clean:
	rm runc
