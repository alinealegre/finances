# Finances Project
## Overview
This project, is a financial management system implemented in Go. It provides functionalities related to managing debts, user authentication, and user data. The project is organized into different directories, each responsible for a specific aspect of the application.


## Technologies Used
The project utilizes the following technologies:

Go (Golang): Chosen for its simplicity, performance, and concurrency support, making it suitable for building web applications.
PostgreSQL: A powerful, open-source relational database management system used for data storage.
Docker: Utilized for containerization, allowing for easy deployment and scalability of the application.

## Dependencies
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

# How to Run
To run the project, follow these steps:

1. Ensure you have Docker installed on your system.
2. Execute the following commands in the project directory:

docker-compose up -d db      // To start the PostgreSQL database
docker-compose up -d finances // To start the application