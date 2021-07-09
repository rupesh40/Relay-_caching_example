FROM golang:latest
WORKDIR /cache_example
COPY go.mod .
RUN go mod download
COPY . .
RUN go build 
CMD [ "./cache_example" ] 
