FROM golang:alpine

LABEL org.opencontainers.image.authors="Jannik Bach"

ENV GIN_MODE=release
ENV APP_HOST=0.0.0.0

WORKDIR /go/app

COPY ../backend/ /go/app/

# Run the two commands below to install git and dependencies for the project. 
RUN apk update && apk add --no-cache git
RUN go get ./

RUN go build -o backend .

EXPOSE $PORT

ENTRYPOINT ["./backend"]
