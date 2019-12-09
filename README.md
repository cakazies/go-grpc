# Golang with GRPC Client and Server

this repo will show you about interaction Golang client server with GRPC, this point is client request and get data from Server with GRPC

## Requirement

- install dependencies go get
  - `github.com/joho/godotenv`
  - `google.golang.org/grpc`
- First time you can see in `.env` client and server must be same
- Second is `server/proto/student.proto` and `client/proto/student.proto` both file must be same
  > Default both of this file already same

## How to Run

### Running Server First

`go run server/main.go`

### Running Client

`go run client/main.go`
