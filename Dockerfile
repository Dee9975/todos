FROM golang:1.21

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /todos

EXPOSE 8080
ENV HOSTNAME 0.0.0.0
ENV PORT 8080

CMD ["/todos"]