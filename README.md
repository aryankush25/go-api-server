# Go API Server

This is a Go-based API server using the Gin framework. The server includes two main modules: Users and Organizations. Each module has its own CRUD APIs. The authentication is email-password-based with a two-step signup process involving OTP verification. AWS SES is used to send emails.

## Features

- **Gin Framework**: Fast, minimalist web framework for Go.
- **PostgreSQL**: Relational database for storing user and organization data.
- **AWS SES**: Service for sending emails.
- **JWT Authentication**: Secure user authentication.
- **OTP Verification**: Two-step signup process.

## Project Structure

```
go-api-server/
├── main.go
├── controllers/
│   ├── user_controller.go
│   ├── organization_controller.go
├── models/
│   ├── user.go
│   ├── organization.go
├── services/
│   ├── auth_service.go
│   ├── email_service.go
│   ├── otp_generator.go
├── repositories/
│   ├── user_repository.go
│   ├── organization_repository.go
├── middlewares/
│   ├── auth_middleware.go
├── config/
│   ├── db.go
│   ├── email.go
│   ├── env.go
├── routes/
│   ├── user_routes.go
│   ├── organization_routes.go
├── .env
├── go.mod
├── go.sum
```

## Getting Started

### Prerequisites

- Go 1.20 or higher
- PostgreSQL
- AWS account with SES configured
- Verified email address in AWS SES

### Installation

1. **Clone the repository**

   ```sh
   git clone https://github.com/aryankush25/go-api-server.git
   cd go-api-server
   ```

2. **Set up environment variables**

   Create a `.env` file in the root directory and populate it with your configuration:

   ```env
   # Database configuration
   DB_DSN=postgres://username:password@localhost:5432/dbname?sslmode=disable

   # Email configuration
   EMAIL_FROM=your_verified_email@example.com

   # AWS SES configuration
   AWS_REGION=us-east-1
   AWS_ACCESS_KEY_ID=your_access_key_id
   AWS_SECRET_ACCESS_KEY=your_secret_access_key

   # JWT Secret
   JWT_SECRET=your_secret_key

   # Server port
   PORT=:8080
   ```

3. **Download dependencies**

   ```sh
   go mod tidy
   ```

4. **Set up PostgreSQL**

   Ensure you have PostgreSQL installed and running. Create a new database for your project.

5. **Apply database migrations**

   If you have any SQL migration files, apply them to set up the database schema.

6. **Run the application**

   ```sh
   go run .
   ```

### Usage

#### Endpoints

- **User Signup Step 1**

  ```sh
  POST /users/signup-step1
  ```

  Request body:

  ```json
  {
    "email": "user@example.com",
    "password": "password"
  }
  ```

- **User Signup Step 2**

  ```sh
  POST /users/signup-step2
  ```

  Request body:

  ```json
  {
    "email": "user@example.com",
    "otp": "123456"
  }
  ```

- **Create Organization**

  ```sh
  POST /organizations/
  ```

  Request body:

  ```json
  {
    "name": "Organization Name"
  }
  ```

### Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
