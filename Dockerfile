FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on

RUN go get github.com/wcisco17/kubgr
RUN cd /build && git clone https://github.com/wcisco17/kubgr.git

RUN cd /build/kubgr/main ** go build

EXPOSE 8080

ENTRYPOINT ["/build/kubgr/main"]