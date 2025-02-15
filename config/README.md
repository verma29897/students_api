# Students API

A simple RESTful API for managing students, built with Go, Gin, and PostgreSQL. This API allows you to perform CRUD operations (Create, Read, Update, Delete) on student records.

## Features
- CRUD operations for students
- PostgreSQL database integration
- Configurable environment and server settings
- Graceful shutdown of the server
- Debug and Release mode support based on environment configuration

## Technologies Used
- **Go**: Programming language used for the backend.
- **Gin**: Web framework for Go.
- **PostgreSQL**: Database for storing student records.
- **Cleanenv**: For managing configuration files.
- **Godotenv**: For loading environment variables from a `.env` file.

## Setup and Installation

### Prerequisites

1. Go (1.18 or higher)
2. PostgreSQL

### 1. Clone the repository

```bash
git clone https://github.com/verma29897/students-api.git
cd students-api
go mod tidy


## Configuration in local.yaml file

### HTTP Server

- **Hostname**: The hostname on which the server will run.
  - Example: `localhost`
  
- **Address (Port)**: The port number where the server will listen for incoming requests.
  - Example: `9090`

### Example Configuration:

The HTTP server is configured with the following settings:

```yaml
ENV=development
http_server:
  Hostname: "localhost"
  Addr: "9090"

## .env Configuration

The `.env` file is used to store environment variables for your application. This file should be placed at the root of your project directory.

### Example `.env` File:

```env


# Database Configuration
GIN_MODE="release" # or "debug"
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_HOST=localhost
DB_PORT=5432
DB_NAME=your_db_name






# Students API

## Overview

This is a simple API to manage student records. You can perform CRUD operations (Create, Read, Update, Delete) on student data using the following endpoints.

## API Endpoints

### 1. `GET /api/v1/students/`
- **Description**: Retrieve a list of all students.
- **Response**:
  - **200 OK**: A JSON array of all students.

  Example response:
  ```json
  [
    {
      "id": 1,
      "first_name": "John",
      "last_name": "Doe",
      "email": "john.doe@example.com",
      "phone": "123-456-7890"
    },
    {
      "id": 2,
      "first_name": "Jane",
      "last_name": "Doe",
      "email": "jane.doe@example.com",
      "phone": "987-654-3210"
    }
  ]


### Explanation:

- **GET `/api/v1/students/`** retrieves all students.
- **GET `/api/v1/students/:id`** retrieves a student by their `id`.
- **POST `/api/v1/students/`** allows you to create a new student.
- **PUT `/api/v1/students/:id`** updates a student with the provided `id`.
- **DELETE `/api/v1/students/:id`** deletes a student by their `id`.


