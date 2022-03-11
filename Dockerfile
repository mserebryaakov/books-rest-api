FROM golang

WORKDIR /app

COPY . /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

RUN go build ./cmd/apiserver

CMD ["./apiserver"]

EXPOSE 8080