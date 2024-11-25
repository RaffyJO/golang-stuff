<div align="center">
  <h3 align="center">Learn RESTful API Part 1 - CRUD Operations</h3>

  <p align="center">
    This project is a simple backend application that allows users to create, read, update, and delete categories. The application uses a RESTful API to handle CRUD operations.
    <br/>
    <br/>
</div>

#### 🚀 Project Overview

This project is a simple backend application that allows users to create, read, update, and delete categories. The application uses a RESTful API to handle CRUD operations. It also demonstrates how to use Go's built-in `net/http` package to handle HTTP requests and responses.

The project consists of the following files and folders:

- `main.go`: The main entry point for the application. It sets up the routes for the CRUD operations and starts the HTTP server.
- `injector.go and wire.go`: A file that defines the dependency injection container and the wire configuration for the application.
- `app`: A folder that contains configuration database and routes files.
- `controllers`: A folder that contains the controllers for the CRUD operations.
- `db/migrations`: A folder that contains the database migration files.
- `exception`: A folder that contains the exception handling files.
- `helper`: A folder that contains the helper functions.
- `middleware`: A folder that contains the middleware files with `X-API-Key` authentication.
- `models`: A folder that contains the data model category.
- `repositiory`: A folder that contains the repository files for the CRUD operations.
- `services`: A folder that contains the service files for the CRUD operations.
- `test`: A folder that contains the unit test files.

#### 🛠️ Tech Stack

The project is built using the following technologies:

- Programming Language: Go
- Database: PostgreSQL

#### 📋 API Contracts

The API contracts for this project are as follows:

- Create Category
  - Method: `POST`
  - URL: `/api/categories`
  - Request Header:
    ```json
    {
      "Content-Type": "application/json",
      "X-API-Key": "RAHASIA"
    }
    ```
  - Request Body:
    ```json
    {
      "name": "Golang"
    }
    ```
  - Response Body:
    ```json
    {
      "code": 201,
      "status": "OK",
      "data": {
        "id": 1,
        "name": "Golang",
        "created_at": "2022-01-01T00:00:00Z",
        "updated_at": "2022-01-01T00:00:00Z"
      }
    }
    ```

- Update Category
  - Method: `PUT`
  - URL: `/api/categories/:categoryId`
  - Request Header:
    ```json
    {
      "Content-Type": "application/json",
      "X-API-Key": "RAHASIA"
    }
    ```
  - Request Body:
    ```json
    {
      "name": "Golang"
    }
    ```
  - Response Body:
    ```json
    {
      "code": 200,
      "status": "OK",
      "data": {
        "id": 1,
        "name": "Golang",
        "created_at": "2022-01-01T00:00:00Z",
        "updated_at": "2022-01-01T00:00:00Z"
      }
    }
    ```

- Delete Category
  - Method: `DELETE`
  - URL: `/api/categories/:categoryId`
  - Request Header:
    ```json
    {
      "Content-Type": "application/json",
      "X-API-Key": "RAHASIA"
    }
    ```
  - Response Body:
    ```json
    {
      "code": 200,
      "status": "OK"
    }
    ```

- Find Category By Id
  - Method: `GET`
  - URL: `/api/categories/:categoryId`
  - Request Header:
    ```json
    {
      "Content-Type": "application/json",
      "X-API-Key": "RAHASIA"
    }
    ```
  - Response Body:
    ```json
    {
      "code": 200,
      "status": "OK",
      "data": {
        "id": 1,
        "name": "Golang",
        "created_at": "2022-01-01T00:00:00Z",
        "updated_at": "2022-01-01T00:00:00Z"
      }
    }
    ```

- Find All Categories
  - Method: `GET`
  - URL: `/api/categories`
  - Request Header:
    ```json
    {
      "Content-Type": "application/json",
      "X-API-Key": "RAHASIA"
    }
    ```
  - Response Body:
    ```json
    {
      "code": 200,
      "status": "OK",
      "data": [
        {
          "id": 1,
          "name": "Golang",
          "created_at": "2022-01-01T00:00:00Z",
          "updated_at": "2022-01-01T00:00:00Z"
        },
        {
          "id": 2,
          "name": "Python",
          "created_at": "2022-01-01T00:00:00Z",
          "updated_at": "2022-01-01T00:00:00Z"
        }
      ]
    }
    ```

- Error Response
  ```json
  {
    "code": 400,
    "status": "BAD REQUEST",
    "data": {
      "name": [
        "Name is required"
      ]
    }
  }
  ```

#### Database Schema

The database schema for this project is as follows:

```sql
CREATE TABLE categories (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(200) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
```

#### 📞 Contact or Ask Questions

For support or inquiries, please contact:

- Email: rafiteguh6@gmail.com
- Linkedin: https://www.linkedin.com/in/raffyjo