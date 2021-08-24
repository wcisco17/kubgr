FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get https://github.com/wcisco17/kubgr
RUN go run github.com/prisma/prisma-client-go db push
RUN go run github.com/prisma/prisma-client-go db generate

# COPY main.go go.* /book/ /client/
# RUN CGO_ENABLED=0 go build -o ./bin/demo

# FROM scratch
# COPY --from=build /bin/demo /bin/demo
# ENTRYPOINT ["/bin/demo"]