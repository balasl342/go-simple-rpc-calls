# Go RPC Application

This is a simple RPC (Remote Procedure Call) application built in Go. The application is divided into client and server components, each with their own responsibilities. The server exposes an RPC method, while the client calls this method remotely.

## Project Structure

The project is organized as follows:

```
rpc-application/
│
├── cmd/
│   ├── client/
│   │   └── main.go
│   └── server/
│       └── main.go
│
├── internal/
│   ├── client/
│   │   └── router.go
│   ├── server/
│   │   ├── db/
│   │   │   └── connection.go
│   │   ├── rpc/
│   │   │   ├── methods.go
│   │   │   └── server.go
│   │   └── repository/
│   │       └── user_repository.go
│   └── shared/
│       └── models.go
│
├── pkg/
│   └── config/
│       └── config.go
└── go.mod
```
