# Train-Booking
This is a gRPC API implementation in Golang for a train ticket booking system, allowing users to book tickets, view ticket details, manage users, and modify seat allocations.

## Prerequisites

Before setting up and running the project, ensure you have the following prerequisites installed on your system:

### 1. Golang

- **Version**: Go 1.18 or later is recommended.
- **Installation**:
  - For macOS: `brew install go`
  - For Windows: Download and install from [Go official website](https://golang.org/dl/)
  - For Linux: Install using package manager (e.g., `sudo apt-get install golang`)

### 2. Protocol Buffers (protoc)

- **Version**: Protocol Buffers v3.14.0 or later is required.
- **Installation**:
  - macOS: `brew install protobuf`
  - Windows/Linux: Download the precompiled binaries from [Protocol Buffers releases](https://github.com/protocolbuffers/protobuf/releases)
  
### 3. gRPC for Go

- **Installation**:
  - Run `go get google.golang.org/grpc` to install the gRPC package.
  - Additionally, install the Protobuf plugin for Go using:
    ```bash
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    ```
  - Ensure the `protoc-gen-go` and `protoc-gen-go-grpc` binaries are available in your `PATH`.

### 4. gRPC-Gateway 

- If you plan to expose your gRPC services as RESTful APIs using gRPC-Gateway:
  - Install the gRPC-Gateway package using:
    ```bash
    go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
    go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
    ```
  - Ensure the `protoc-gen-grpc-gateway` and `protoc-gen-openapiv2` binaries are available in your `PATH`.



## Clone the project

```bash
  git clone https://github.com/roshanikhairnar/Train-Booking.git
```

Go to the project directory

```bash
  cd Train-Booking
```

## Run Program Locally

Open Terminal and run server
``` bash
go run server/cmd/main.go
```

Open another Terminal and run client
```bash
go run client/cmd/main.go
```
Run following curl commands

### 1. SubmitPurchase

- **Endpoint**: `/v1/Purchase`
- **Method**: `POST`

```
curl --location 'http://localhost:8080/v1/Purchase' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user": {"id": "user-197", "first_name": "kk", "last_name": "kk", "email": "kk@example.com"},
    "from": "London",
    "to": "France"
}'
```
### 2. GetTicketDetails

- **Endpoint**: `/v1/ticket/{userId}`
- **Method**: `GET`
```
curl --location 'http://localhost:8080/v1/ticket/user-197'
```
### 3. ModifySeat

- **Endpoint**: `/v1/user/{userId}/seat`
- **Method**: `PUT`
```
curl --location --request PUT 'http://localhost:8080/v1/user/user-197/seat' \
--header 'Content-Type: application/json' \
--data '{
    "userId":"user-197",
    "newSeatNumber": "1"
}'
```
### 4. GetUsersBySection

- **Endpoint**: `/v1/users/section/{section}`
- **Method**: `GET`
```
curl --location 'http://localhost:8080/v1/users/section/A'
```

### 5. RemoveUser

- **Endpoint**: `/v1/user/{userId}`
- **Method**: `DELETE`
```
curl --location --request DELETE 'http://localhost:8080/v1/user/user-127'
```


