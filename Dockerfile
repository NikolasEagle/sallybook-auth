FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY funcs ./funcs

COPY structs ./structs

COPY main.go ./

RUN go build -o ./bin/app .

EXPOSE 8001

CMD [ "./bin/app" ]