FROM golang as build

WORKDIR /app

COPY go.mod /app
COPY go.sum /app

#RUN apk --no-cache add git gcc \
#    && go mod download
RUN go mod download

COPY . /app

RUN go build -o /exporter

# build a small statically linked applet
RUN GOOS=linux GOARCH=amd64 go build -o /exporter -ldflags="-s -w"
# lets make is smaller: https://blog.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/
#RUN upx --brute /exporter

FROM debian:stretch
COPY --from=build /exporter /
ENTRYPOINT ["/exporter"]