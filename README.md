# Train-Booking



Clone the project

```bash
  git clone https://github.com/roshanikhairnar/Train-Booking.git
```

Go to the project directory

```bash
  cd Train-Booking
```

## Run Program Locally

Open Terminal and Run
``` bash
go run server/cmd/main.go
```

Open another Terminal
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
### 5. ModifySeat

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
### 3. GetUsersBySection

- **Endpoint**: `/v1/users/section/{section}`
- **Method**: `GET`
```
curl --location 'http://localhost:8080/v1/users/section/A'
```

### 4. RemoveUser

- **Endpoint**: `/v1/user/{userId}`
- **Method**: `DELETE`
```
curl --location --request DELETE 'http://localhost:8080/v1/user/user-127'
```


