FROM golang:alpine

LABEL org.opencontainers.image.authors="Jannik Bach"

WORKDIR /go/app

COPY ./ /go/app/

# Run the two commands below to install git and dependencies for the project. 
RUN apk update && apk add --no-cache git
RUN go get ./

RUN go build -o backend .

EXPOSE $PORT

ENTRYPOINT ["./backend"]
