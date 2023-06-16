# Personal Finances Tracker

The Personal Finances Tracker is a Go-based web application that allows users to track their income, expenses, and manage their personal finances. This project utilizes GORM as the ORM (Object Relational Mapping) library, Gorilla Mux as the HTTP router and MySQL as the database.

## Features

- Create, Read, Update, and Delete (CRUD) operations for income and expenses records
- User authentication and authorization for secure access to the application
- Filtering, pagination, and search functionality for listing income and expenses records
- Integration with a MySQL database to store and retrieve data

## Dependencies

- GORM: GORM is a powerful ORM library for Go, which simplifies database operations and provides an easy-to-use API for interacting with databases. You can find more information about GORM [here](https://gorm.io/).

- Gorilla Mux: Gorilla Mux is a popular HTTP router and URL matcher for Go. It provides flexible routing capabilities for handling HTTP requests and building RESTful APIs. You can find more information about Gorilla Mux [here](https://github.com/gorilla/mux).

- MySQL: MySQL is a widely-used open-source relational database management system. It provides robust features for data storage and retrieval. You can find more information about MySQL [here](https://www.mysql.com/).

## Getting Started

### Prerequisites

- Go: Make sure you have Go installed on your system. You can download and install it from the official Go website [here](https://golang.org/dl/).

- MySQL: Install and configure MySQL on your system. You can download MySQL Community Edition from the official MySQL website [here](https://dev.mysql.com/downloads/installer/).

### Installation

1. Clone the repository:

    ```bash
    git clone git@github.com:KaungKhantNaing16/Go-PersonalFinanceTracker.git

2. Change to the project repository:

    ```bash
    cd Go-PersonalFinanceTracker

3. Install the project dependencies:

    ```bash
    go mod download

4. Run the project

    ```bash
    go run main.go

Open your web browser and visit http://localhost:8000 to access the personal finances tracker application.


