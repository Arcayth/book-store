# Book Store API

This is a simple Book Store API built with Golang, Gin, and MongoDB.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)


## Features

- CRUD operations for managing books.
- MongoDB integration for data storage.
- RESTful API using the Gin framework.
- Air for hot-reloading during development.

## Getting Started

### Prerequisites

Make sure you have the following installed on your machine:

- [Golang](https://golang.org/)
- [Air](https://github.com/cosmtrek/air) (for development)

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/haythamchanouni/book-store.git
    ```

2. Install the dependencies:
    ```bash
    cd book-store-api
    go mod tidy
    ```

## Usage

1. Create a `.env` file in the root directory and add the following environment variables:

    ```bash
    MONGODB_URI='<your-mongodb-uri>'
    MONGO_DB_NAME='<your-mongodb-database-name>'
    ```

2. Start the API server:
    ```bash
    go run main.go
    ```

    or, for development with hot-reloading:
    ```bash
    air init
    air
    ```

3. Access the API at `http://localhost:8080`.

## API Endpoints

| Method | Endpoint | Description |
| --- | --- | --- |
| GET | `/books` | Get all books |
| GET | `/books/:id` | Get a book by ID |
| POST | `/books` | Create a new book |
| PUT | `/books/:id` | Update a book by ID |
| DELETE | `/books/:id` | Delete a book by ID |


