FROM golang:1.16-alpine

RUN apk add --no-cache git
RUN mkdir /fraazo
ADD . /fraazo

# Set the Current Working Directory inside the container
WORKDIR /fraazo

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o main  .


# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["/fraazo/main"]