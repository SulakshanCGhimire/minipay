# MiniPay Concepts Guide

This document explains the major concepts, principles, and software engineering practices used in the MiniPay project. It is intended for students who want to understand not only *what* was built, but *why* it was built that way.

---

# 1. What is MiniPay?

MiniPay is a digital wallet application that allows users to:

* Register and log in
* Create a wallet
* Deposit funds
* Transfer money to other users
* View transaction history

Although simplified, the project mimics the architecture used in real-world payment systems such as eSewa, Khalti, PayPal, and other digital wallets.

---

# 2. REST API Architecture

MiniPay is built as a REST API.

REST (Representational State Transfer) is an architectural style used for communication between clients and servers using HTTP.

Example:

```http
POST /login
GET /wallet
POST /transfer
```

Each endpoint represents a resource or action.

Benefits:

* Easy to understand
* Platform independent
* Used by modern web and mobile applications
* Scalable

---

# 3. Layered Architecture

The project follows a layered architecture.

```text
Client
   ↓
Handlers
   ↓
Services
   ↓
Repositories
   ↓
Database
```

Each layer has a specific responsibility.

---

## Handler Layer

Handles:

* HTTP requests
* Request validation
* Response generation

Example:

```text
POST /login
```

The handler receives user credentials and forwards them to the service layer.

Responsibilities:

* Read request data
* Call services
* Return JSON responses

---

## Service Layer

Contains business logic.

Examples:

* Validate login credentials
* Check wallet balance
* Process transfers

This is where application rules live.

Example:

```text
Cannot transfer more money than available balance.
```

This rule belongs to the service layer.

Benefits:

* Clean code
* Easier testing
* Better maintainability

---

## Repository Layer

Responsible for database communication.

Examples:

* Create user
* Find wallet
* Save transaction

Benefits:

* Database code is isolated
* Easy to switch databases later
* Cleaner architecture

---

# 4. MVC vs Layered Architecture

Many beginners learn MVC:

```text
Model
View
Controller
```

MiniPay uses a more scalable architecture:

```text
Handler
Service
Repository
Model
```

This approach is common in professional backend systems.

Advantages:

* Better separation of concerns
* Easier maintenance
* Easier testing
* More scalable

---

# 5. Authentication

Authentication answers:

> "Who is the user?"

MiniPay uses:

* Email
* Password
* JWT Token

Authentication flow:

```text
Register
    ↓
Login
    ↓
Generate JWT
    ↓
Send Token
    ↓
Access Protected Routes
```

---

# 6. Password Hashing

Passwords should never be stored directly.

Bad:

```text
password123
```

Good:

```text
$2a$10$7Qx...
```

MiniPay uses BCrypt.

Process:

```text
Password
    ↓
BCrypt Hash
    ↓
Store Hash
```

Benefits:

* Passwords cannot be recovered easily
* Industry standard approach
* Increased security

---

# 7. JWT (JSON Web Token)

JWT is used for stateless authentication.

After login:

```text
User → Login
      ↓
Server verifies credentials
      ↓
Server generates JWT
      ↓
User stores JWT
      ↓
JWT sent with future requests
```

Example:

```http
Authorization: Bearer <token>
```

Benefits:

* No session storage required
* Fast authentication
* Widely used in APIs

---

# 8. Middleware

Middleware sits between request and response.

Example:

```text
Request
   ↓
Middleware
   ↓
Handler
```

MiniPay uses authentication middleware.

Responsibilities:

* Verify JWT
* Reject invalid users
* Attach user information to request context

Benefits:

* Reusable logic
* Cleaner handlers
* Better security

---

# 9. Database Design

MiniPay uses relational database concepts.

Main entities:

```text
User
Wallet
Transaction
```

Relationships:

```text
User
  |
  | 1:1
  |
Wallet

Wallet
  |
  | 1:N
  |
Transaction
```

---

# 10. ORM (Object Relational Mapping)

MiniPay uses GORM.

Without ORM:

```sql
SELECT * FROM users;
```

With ORM:

```go
db.Find(&users)
```

Benefits:

* Less SQL code
* Faster development
* Database abstraction

---

# 11. Repository Pattern

Repository Pattern separates business logic from database logic.

Without Repository Pattern:

```text
Service
  └── SQL Queries
```

With Repository Pattern:

```text
Service
    ↓
Repository
    ↓
Database
```

Benefits:

* Better organization
* Easier testing
* Cleaner code

---

# 12. Transaction Processing

The core feature of MiniPay is money transfer.

Transfer process:

```text
Sender Wallet
      ↓
Validate Balance
      ↓
Deduct Amount
      ↓
Credit Receiver
      ↓
Record Transaction
```

Important rule:

```text
Money cannot be created or destroyed.
```

This is known as maintaining consistency.

---

# 13. Data Integrity

A payment system must preserve correct balances.

Example:

User A = 1000

Transfer 200

Result:

```text
User A = 800
User B = +200
```

The total money remains consistent.

Data integrity ensures:

* No duplicate transfers
* No missing balances
* Correct transaction records

---

# 14. API Design Principles

Good APIs should be:

### Predictable

```http
GET /wallet
POST /wallet/deposit
POST /transfer
```

### Consistent

Always use similar naming conventions.

### Stateless

Every request contains everything needed.

Benefits:

* Easier client integration
* Better scalability

---

# 15. JSON Communication

MiniPay communicates using JSON.

Example:

Request:

```json
{
  "amount": 500
}
```

Response:

```json
{
  "message": "Deposit successful"
}
```

Why JSON?

* Lightweight
* Human readable
* Industry standard

---

# 16. Error Handling

Good systems must handle failures.

Examples:

* Invalid login
* Wallet not found
* Insufficient balance
* Invalid token

Example response:

```json
{
  "error": "Insufficient balance"
}
```

Benefits:

* Better user experience
* Easier debugging

---

# 17. Environment Variables

Sensitive information should not be hardcoded.

Bad:

```go
password := "root123"
```

Good:

```env
DB_PASSWORD=root123
JWT_SECRET=mysecret
```

Benefits:

* Improved security
* Easier deployment
* Better configuration management

---

# 18. Why Payment Systems Are Challenging

Payment systems require:

* Security
* Consistency
* Reliability
* Authentication
* Auditability

Even a small wallet application introduces many real-world software engineering challenges.

MiniPay is therefore much more than a CRUD application—it introduces concepts used in fintech systems.

---

# 19. What Students Learn From This Project

By building MiniPay, students practice:

### Backend Development

* Go
* Gin Framework
* REST APIs

### Database Management

* MySQL
* GORM
* Data Modeling

### Security

* JWT
* Password Hashing
* Authentication Middleware

### Software Engineering

* Layered Architecture
* Repository Pattern
* Separation of Concerns

### Financial Systems

* Wallet Design
* Transaction Processing
* Data Integrity

---

# Final Note

MiniPay was built not only to create a functional wallet system but also to explore important backend engineering concepts such as authentication, database management, layered architecture, transaction processing, and secure API design. Understanding these concepts is more valuable than memorizing code because they apply to almost every modern software system.
