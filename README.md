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
└── go.mod
```

### Directory Breakdown

- **cmd/**: Entry points for the client and server applications.
  - **client/**: Contains the main entry point for the client application.
  - **server/**: Contains the main entry point for the server application.

- **internal/**: Core business logic of the application.
  - **client/**: Manages RPC calls from the client.
  - **server/db/**: Handles database connection management.
  - **server/rpc/**: Contains RPC server setup and method definitions.
  - **server/repository/**: Manages database operations related to the user.
  - **shared/**: Contains models and types shared across the application.

## Getting Started

### Prerequisites

- Go 1.20 or later
- PostgreSQL
- Any RPC client for testing

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/balasl342/go-simple-rpc-calls.git
   ```

2. **Install dependencies:**

   Ensure that you are in the project root directory and run:

   ```bash
   go mod tidy
   ```

3. **Set up environment variables:**

   Configure your environment variables for database connection and other settings as needed.

   Example ```.env``` file:

   ```
   user=username
   password=password
   dbname=mydb
   host=localhost
   port=5432
   ```

### Running the Application

1. **Start the RPC Server:**

   Navigate to the server directory and run the server:

   ```bash
   cd cmd/server
   go run main.go
   ```

   The server will start and listen for RPC calls on port 1234.

2. **Run the RPC Client:**

   In a new terminal window, navigate to the client directory and run the client:

   ```bash
   cd cmd/client
   go run main.go
   ```

   The client will make an RPC call to the server and print the result.

### Database Integration (Optional)

If you want to integrate a database (PostgreSQL) into your RPC server, ensure that your PostgreSQL service is running and the environment variables are correctly set.

### Making RPC Calls

The application includes sample RPC methods, ```CreateProducts and GetProductsByID```. You can test this method using the provided client, or by using another RPC client of your choice.

### Extending the Application

- **Adding New RPC Methods:** To add new methods, create new functions in ```internal/server/rpc/methods.go``` and register them in ```internal/server/rpc/server.go```.
  
- **Database Operations:** To add or modify database operations, update the ```internal/server/repository/``` package.

### License

This project is licensed under the MIT License. See the LICENSE file for details.
