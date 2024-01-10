FROM golang:1.21
RUN apk add --no-cache --update gcc g++

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

RUN sqlite3 mound.sqlite

COPY . /build

RUN CGO_ENABLED=1 go build -o /todos

CMD ["/todos"]