# MiniPay 💳

MiniPay is a lightweight digital wallet and payment management system built with Go. It provides secure user authentication, wallet management, balance deposits, peer-to-peer transfers, and transaction tracking through a RESTful API.

The project demonstrates modern backend development practices including JWT-based authentication, layered architecture, database persistence, and secure transaction handling.

---

## Features

### User Authentication

* User registration
* User login
* Password hashing using BCrypt
* JWT token generation and validation
* Protected API routes

### Wallet Management

* Create a personal wallet
* View wallet details
* Deposit funds into wallet
* Retrieve current wallet balance

### Transactions

* Transfer money between wallets
* Transaction recording and storage
* View transaction history
* Balance validation before transfers

### Security

* Password encryption with BCrypt
* JWT authentication middleware
* Protected routes
* Database-backed persistence

---

## Tech Stack

### Backend

* Go (Golang)
* Gin Web Framework

### Database

* MySQL
* GORM ORM

### Authentication

* JWT (JSON Web Tokens)
* BCrypt Password Hashing

### Frontend

* HTML
* JavaScript

---

## Project Architecture

```text
minipay/
│
├── cmd/
│   └── main.go
│
├── config/
│   ├── config.go
│   ├── database.go
│   └── jwt.go
│
├── handlers/
│   ├── auth.go
│   ├── wallet.go
│   └── transaction.go
│
├── middleware/
│   └── auth.go
│
├── models/
│   ├── user.go
│   ├── wallet.go
│   └── transaction.go
│
├── repositories/
│   ├── wallet_repository.go
│   └── transaction_repository.go
│
├── services/
│   ├── auth_service.go
│   ├── wallet_service.go
│   └── transaction_service.go
│
├── ui/
│   └── index.html
│
├── go.mod
└── README.md
```

---

## System Design

MiniPay follows a layered architecture:

### Handler Layer

Handles incoming HTTP requests and API responses.

### Service Layer

Contains business logic such as:

* Wallet operations
* Money transfers
* Authentication workflows

### Repository Layer

Responsible for database interactions.

### Database Layer

Stores:

* Users
* Wallets
* Transactions

---

## Database Models

### User

```go
type User struct {
    Name     string
    Email    string
    Password string
}
```

### Wallet

```go
type Wallet struct {
    UserID  uint
    Balance float64
}
```

### Transaction

```go
type Transaction struct {
    SenderWalletID   uint
    ReceiverWalletID uint
    Amount           float64
    Status           string
}
```

---

## API Endpoints

### Public Routes

#### Register User

```http
POST /register
```

Request:

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

---

#### Login User

```http
POST /login
```

Request:

```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

Response:

```json
{
  "message": "Login successful",
  "token": "<jwt-token>"
}
```

---

### Protected Routes

All routes below require:

```http
Authorization: Bearer <jwt-token>
```

---

#### Get Current User

```http
GET /me
```

---

#### Create Wallet

```http
POST /wallet/create
```

---

#### Get Wallet Information

```http
GET /wallet
```

---

#### Deposit Funds

```http
POST /wallet/deposit
```

Example:

```json
{
  "amount": 1000
}
```

---

#### Transfer Funds

```http
POST /transfer
```

Example:

```json
{
  "receiver_wallet_id": 2,
  "amount": 250
}
```

---

#### Transaction History

```http
GET /transactions
```

---

## Installation

### Clone Repository

```bash
git clone https://github.com/SulakshanCGhimire/minipay.git
cd minipay
```

### Install Dependencies

```bash
go mod tidy
```

### Configure Environment Variables

Create a `.env` file:

```env
DB_USER=root
DB_PASSWORD=password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=minipay

JWT_SECRET=your_secret_key
```

### Run MySQL

Ensure MySQL is running and the configured database exists.

### Start Server

```bash
go run cmd/main.go
```

Server runs at:

```text
http://localhost:8080
```

---

## Testing the API

### Health Check

```http
GET /ping
```

Response:

```json
{
  "message": "pong"
}
```

---

## Future Improvements

* Transaction reversal support
* Account verification
* Admin dashboard
* Real-time notifications
* Mobile application
* QR-based payments
* Multi-currency wallets
* Payment gateway integration
* Fraud detection module
* Transaction analytics dashboard

---

## Learning Outcomes

This project demonstrates:

* REST API Development
* Backend Engineering
* JWT Authentication
* Password Security
* Database Design
* Layered Architecture
* Repository Pattern
* Financial Transaction Processing

---

## Author

**Sulakshan Chandra Ghimire**

Computer Engineering Student

GitHub: https://github.com/SulakshanCGhimire

---

## License

This project is licensed under the MIT License.
