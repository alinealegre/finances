Finances Project
Overview
This project, is a financial management system implemented in Go. It provides functionalities related to managing debts, user authentication, and user data. The project is organized into different directories, each responsible for a specific aspect of the application.

Directory Structure
The project directory structure is as follows:

finances/
|-- controllers/
|   |-- debts_controller.go
|   |-- login_controller.go
|   |-- score_controller.go
|   |-- user_controller.go
|
|-- database/
|   |-- migrations/
|   |   |-- migrations.go
|   |-- database.go
|   |-- Dockerfile
|
|-- models/
|   |-- debts.go
|   |-- login.go
|   |-- user.go
|
|-- server/
|   |-- middlewares/
|   |-- |-- admin_middleware.go
|   |   |-- auth_middleware.go
|   |   |-- check_cpf_middleware.go
|   |-- routes/
|   |   |-- routes.go
|   |-- server.go
|
|-- services/
|   |-- jwt_service.go
|   |-- score_service.go
|   |-- sha256_service.go
|
|-- tests/
|   |-- controllers/
|   |   |-- debts_controller_test.go
|   |-- services/ 
|   |   |-- jwt_controller_test.go

|
|-- docker-compose.yml
|-- Dockerfile
|-- main.go


Technologies Used
The project utilizes the following technologies:

Go (Golang): Chosen for its simplicity, performance, and concurrency support, making it suitable for building web applications.
PostgreSQL: A powerful, open-source relational database management system used for data storage.
Docker: Utilized for containerization, allowing for easy deployment and scalability of the application.

Dependencies
The project relies on several third-party dependencies, managed using Go Modules:

go 1.19

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.9.1
	github.com/joho/godotenv v1.5.1
	github.com/stretchr/testify v1.8.4
	gorm.io/driver/postgres v1.5.6
	gorm.io/gorm v1.25.7
)

How to Run
To run the project, follow these steps:

1. Ensure you have Docker installed on your system.
2. Execute the following commands in the project directory:

docker-compose up -d db      // To start the PostgreSQL database
docker-compose up -d finances // To start the application

The application will then be accessible according to the configurations set in the code.

Repository
The project repository can be found at https://github.com/alinealegre/finances.

For any further instructions or inquiries, please refer to the README.md file in the repository or contact the project owner.