FROM golang:latest

RUN apt-get update
RUN apt-get upgrade -y

ENV GIN_MODE=release

RUN mkdir /app
COPY . /usr/local/go/src/github.com/guillermodoghel/minesweeper-API/

CMD go run /usr/local/go/src/github.com/guillermodoghel/minesweeper-API/cmd

EXPOSE 8080
