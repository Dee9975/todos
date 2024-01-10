FROM golang:1.21

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . /build

RUN go build -o /todos

CMD ["/todos"]