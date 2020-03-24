FROM golang:alpine3.11

WORKDIR /app

COPY . .

RUN cd cmd/http && go build 

ENTRYPOINT /app/cmd/http/http