FROM golang:1.22.2-alpine

WORKDIR /app

COPY ../go.mod . 
COPY ../go.sum . 

COPY pinger/ ./pinger

WORKDIR /app/pinger

RUN apk add --no-cache docker-cli
RUN go mod tidy  
RUN go mod download
RUN go build -o pinger .

CMD ["./pinger"]
