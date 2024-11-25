<div align="center">
  <h3 align="center">Integration With External API - Music App</h3>

  <p align="center">
    This project integrates with the Spotify API, utilizes GORM for database operations, and implements unit testing to ensure code quality and reliability.
    <br/>
    <br/>
</div>

#### 🚀 Project Overview

This project integrates with the Spotify API, utilizes GORM for database operations, and implements unit testing to ensure code quality and reliability. It also demonstrates how to use Gin framework to handle HTTP requests and responses.

#### 🛠️ Tech Stack

The project is built using the following technologies:

- Programming Language: Go
- Framework: Gin, net/http
- Database: PostgreSQL
- Middleware: JWT
- ORM: GORM
- Other Tools: Docker

#### 📋 API Contracts

The API contracts for this project are as follows:

- Sign Up
  - Method: `POST`
  - URL: `/memberships/signup`
  - Request Header:
    ```json
    {
      "Content-Type": "application/json"
    }
    ```
  - Request Body:
    ```json
    {
      "email": "raffyjo@gmail.com",
      "password": "123456",
      "username": "raffy"
    }
    ```
  - Response Body:
    ```json
    {
      "status": "success",
      "message": "Sign Up Success",
      "data": null
    }
    ```

- Login
  - Method: `POST`
  - URL: `/memberships/login`
  - Request Header:
    ```json
    {
      "Content-Type": "application/json"
    }
    ```
  - Request Body:
    ```json
    {
      "email": "raffyjo@gmail.com",
      "password": "123456"
    }
    ```
  - Response Body:
    ```json
    {
      "status": "success",
      "message": "Login Success",
      "data": {
        "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiaWF0IjoxNjE5MzE0ODk5LCJleHAiOjE2MTkyMTQ5OTk5fQ.1-4-5-6-7-8-9-0-1-2-3-4-5-6-7-8-9-0",
      }
    }
    ```

- Search Tracks
  - Method: `GET`
  - URL: `/tracks/search?query=song&pageSize=10&pageIndex=1`
  - Request Header:
    ```json
    {
      "Content-Type": "application/json",
      "Authorization Type": "API Key",
      "Authorization": "<JWT>"
    }
    ```
  - Response Body:
    ```json
    {
      "status": "success",
      "message": "Successfully get search results",
      "data": {
        "limit": 2,
        "offset": 0,
        "total": 900,
        "items": [
          {
            "album_type": "album",
            "album_total_tracks": 9,
            "album_images_url": [
              "https://i.scdn.co/image/ab67616d0000b2737c52c7f7d1da8625e4376795",
              "https://i.scdn.co/image/ab67616d00001e027c52c7f7d1da8625e4376795",
              "https://i.scdn.co/image/ab67616d000048517c52c7f7d1da8625e4376795"
            ],
            "album_name": "Gajah",
            "artists_name": [
              "Tulus"
            ],
            "explicit": false,
            "id": "73DWDOjVUyJ8sAiAcySvgS",
            "name": "Sepatu",
            "is_liked": false
          },
          {
            "album_type": "single",
            "album_total_tracks": 1,
            "album_images_url": [
              "https://i.scdn.co/image/ab67616d0000b27393bdd610ba0672f6177b1564",
              "https://i.scdn.co/image/ab67616d00001e0293bdd610ba0672f6177b1564",
              "https://i.scdn.co/image/ab67616d0000485193bdd610ba0672f6177b1564"
            ],
            "album_name": "sepatu super",
            "artists_name": [
              "Rudiliem"
            ],
            "explicit": true,
            "id": "1olLNFQGdAwPlJSGlKpkLp",
            "name": "sepatu super",
            "is_liked": null
          }
        ]
      }
    }
    ```

- Like or Dislike
  - Method: `POST`
  - URL: `/tracks/track-activity`
  - Request Header:
    ```json
    {
      "Content-Type": "application/json",
      "Authorization Type": "API Key",
      "Authorization": "<JWT>"
    }
    ```
  - Request Body:
    ```json
    {
      "spotify_id": "73DWDOjVUyJ8sAiAcySvgS",
      "is_liked": true
    }
    ```
  - Response Body:
    ```json
    {
      "status": "success",
      "message": "Successfully like or dislike track",
      "data": null
    }
    ```

- Recommendation Tracks
  - Method: `GET`
  - URL: `/tracks/recommendations?trackID=73DWDOjVUyJ8sAiAcySvgS&limit=2`
  - Request Header:
    ```json
    {
      "Content-Type": "application/json",
      "Authorization Type": "API Key",
      "Authorization": "<JWT>"
    }
    ```
  - Response Body:
    ```json
    {
      "status": "success",
      "message": "Successfully get recommendations",
      "data": {
        "items": [
          {
            "album_type": "album",
            "album_total_tracks": 9,
            "album_images_url": [
              "https://i.scdn.co/image/ab67616d0000b2737c52c7f7d1da8625e4376795",
              "https://i.scdn.co/image/ab67616d00001e027c52c7f7d1da8625e4376795",
              "https://i.scdn.co/image/ab67616d000048517c52c7f7d1da8625e4376795"
            ],
            "album_name": "Gajah",
            "artists_name": [
              "Tulus"
            ],
            "explicit": false,
            "id": "73DWDOjVUyJ8sAiAcySvgS",
            "name": "Sepatu",
            "is_liked": false
          },
          {
            "album_type": "single",
            "album_total_tracks": 1,
            "album_images_url": [
              "https://i.scdn.co/image/ab67616d0000b27393bdd610ba0672f6177b1564",
              "https://i.scdn.co/image/ab67616d00001e0293bdd610ba0672f6177b1564",
              "https://i.scdn.co/image/ab67616d0000485193bdd610ba0672f6177b1564"
            ],
            "album_name": "sepatu super",
            "artists_name": [
              "Rudiliem"
            ],
            "explicit": true,
            "id": "1olLNFQGdAwPlJSGlKpkLp",
            "name": "sepatu super",
            "is_liked": null
          }
        ]
      }
    }
    ```

- Error Response
  ```json
  {
    "status": "error",
    "message": "Error Message",
    "data": null
  }
  ```

#### Database Schema

The database schema for this project is as follows:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_by VARCHAR(255) NOT NULL,
    updated_by VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE track_activities (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    spotify_id VARCHAR(255) NOT NULL,
    is_liked BOOLEAN,
    created_by VARCHAR(255) NOT NULL,
    updated_by VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
```

#### 📞 Contact or Ask Questions

For support or inquiries, please contact:

- Email: rafiteguh6@gmail.com
- Linkedin: https://www.linkedin.com/in/raffyjo