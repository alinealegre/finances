# Finances Project
## Overview
This project, is a financial management system implemented in Go. It provides functionalities related to managing debts, user authentication, and user data. The project is organized into different directories, each responsible for a specific aspect of the application.


## Technologies Used
The project utilizes the following technologies:

Go (Golang 1.19): Chosen for its simplicity, performance, and concurrency support, making it suitable for building web applications.
PostgreSQL: A powerful, open-source relational database management system used for data storage.
Docker: Utilized for containerization, allowing for easy deployment and scalability of the application.

# How to Run
To run the project, follow these steps:

1. Ensure you have Docker installed on your system.
2. Execute the following commands in the project directory:


docker-compose up -d db      // To start the PostgreSQL database

docker-compose up -d finances // To start the application
