FROM golang:1.22rc2-bookworm

WORKDIR /usr/backend/

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY ./ ./

RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]
