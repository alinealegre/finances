# Finances Project
## Overview
This project, is a financial management system implemented in Go. It provides functionalities related to managing debts, user authentication, and user data. The project is organized into different directories, each responsible for a specific aspect of the application.


## Technologies Used
The project utilizes the following technologies:

Go (Golang 1.19): Chosen for its simplicity, performance, and concurrency support, making it suitable for building web applications.
PostgreSQL: A powerful, open-source relational database management system used for data storage.
Docker: Utilized for containerization, allowing for easy deployment and scalability of the application.

### Documentation
The documentation for this project can be found on Postman:
[Finances API Documentation](https://www.postman.com/docking-module-geologist-57881386/workspace/finances/collection/33172072-51b20f76-dc60-42a8-bf56-2aa3439e4e3e?action=share&creator=33172072)

# How to Run
To run the project, follow these steps:

1. Ensure you have Docker installed on your system.
2. Execute the following commands in the project directory:


docker-compose up -d db      // To start the PostgreSQL database

docker-compose up -d finances // To start the application
