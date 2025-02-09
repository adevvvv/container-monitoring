FROM golang:1.22.2-alpine

WORKDIR /app

COPY ../go.mod . 
COPY ../go.sum . 

COPY backend/ ./backend

WORKDIR /app/backend

RUN go mod tidy
RUN go mod download
RUN go build -o backend .

EXPOSE 8080

CMD ["./backend"]
