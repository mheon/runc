FROM fedora:22

RUN dnf -y upgrade
RUN dnf -y install golang git mercurial

RUN mkdir /gopath

ENV GOPATH /gopath
ENV PATH $PATH:$GOPATH/bin

COPY . /gopath/src/github.com/opencontainers/runc
WORKDIR /gopath/src/github.com/opencontainers/runc

RUN go get github.com/tools/godep

CMD ["godep", "go", "build", "-o", "runc", "."]
