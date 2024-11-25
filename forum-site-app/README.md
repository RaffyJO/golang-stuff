<div align="center">
  <h3 align="center">Learn RESTful API Part 2 - Forum Site App</h3>

  <p align="center">
    This project is a simple forum application that allows users to create, read forum posts, create comments and like or dislike posts.
    <br/>
    <br/>
</div>

#### 🚀 Project Overview

This project is a simple forum application that allows users to create, read forum posts, create comments and like or dislike posts. It also demonstrates how to use Gin framework to handle HTTP requests and responses.

#### 🛠️ Tech Stack

The project is built using the following technologies:

- Programming Language: Go
- database: MySQL
- Framework: Gin
- Middleware: JWT
- Other Tools: Docker

#### 📋 API Contracts

The API contracts for this project are as follows:

- Sign Up
  - Method: `POST`
  - URL: `/memberships/sign-up`
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
      "message": "Sign Up Success"
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
        "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiaWF0IjoxNjE5MzE0ODk5LCJleHAiOjE2MTkyMTQ5OTk5fQ.1-4-5-6-7-8-9-0-1-2-3-4-5-6-7-8-9-0"
      }
    }
    ```

- Refresh Token
  - Method: `POST`
  - URL: `/memberships/refresh-token`
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
      "refreshtoken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiaWF0IjoxNjE5MzE0ODk5LCJleHAiOjE2MTkyMTQ5OTk5fQ.1-4-5-6-7-8-9-0-1-2-3-4-5-6-7-8-9-0"
    }
    ```
  - Response Body:
    ```json
    {
      "status": "success",
      "message": "Refresh Token Success",
      "data": {
        "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiaWF0IjoxNjE5MzE0ODk5LCJleHAiOjE2MTkyMTQ5OTk5fQ.1-4-5-6-7-8-9-0-1-2-3-4-5-6-7-8-9-0",
      }
    }
    ```

- Create Post
  - Method: `POST`
  - URL: `/posts/create`
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
      "PostTitle": "Hello World",
      "PostContent": "This is a sample post",
      "PostHastags": ["hastag1", "hastag2"]
    }
    ```
  - Response Body:
    ```json
    {
      "status": "success",
      "message": "Create Post Success",
    }
    ```

- Create Comment
  - Method: `POST`
  - URL: `/posts/comment/:postID`
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
      "commentContent": "This is a sample comment"
    }
    ```
  - Response Body:
    ```json
    {
      "status": "success",
      "message": "Create Comment Success",
    }
    ```

- Upsert User Activity
  - Method: `PUT`
  - URL: `/posts/user-activity/:postID`
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
      "isLiked": true
    }
    ```
  - Response Body:
    ```json
    {
      "status": "success",
      "message": "Upsert User Activity Success",
    }
    ```

- Get All Post
  - Method: `GET`
  - URL: `/posts?pageSize=10&pageIndex=1`
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
      "message": "Get All Post Success",
      "data": [
        {
          "id": 1,
          "userID": 1,
          "username": "raffy",
          "postTitle": "Hello World",
          "postContent": "This is a sample post",
          "postHastags": [
            "hastag1",
            "hastag2"
          ],
          "isLiked": false
        }
      ],
      "Pagination": {
        "limit": 10,
        "offset": 0
      }
    }
    ```

- Get Post By Id
  - Method: `GET`
  - URL: `/posts/:postID`
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
      "message": "Get Post By Id Success",
      "data": {
        "postDetail": {
          "id": 1,
          "userID": 1,
          "username": "raffy",
          "postTitle": "Hello World",
          "postContent": "This is a sample post",
          "postHastags": [
            "hastag1",
            "hastag2"
          ],
          "isLiked": false
        },
        "likeCount": 3,
        "comments": [
          {
            "id": 1,
            "userID": 1,
            "username": "raffy",
            "commentContent": "This is a sample comment"
          },
          {
            "id": 2,
            "userID": 1,
            "username": "raffy",
            "commentContent": "This is a sample comment"
          },
        ]
      }
    }
    ```

- Error Response
  ```json
  {
    "status": "error",
    "message": "Error Message",
  }
  ```

#### Database Schema

The database schema for this project is as follows:

```sql
CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) AUTO_INCREMENT PRIMARY KEY,
  `username` varchar(255) NOT NULL UNIQUE,
  `email` varchar(255) NOT NULL,
  `password` varchar(512) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_by` LONGTEXT NOT NULL,
  `updated_by` LONGTEXT NOT NULL
)

CREATE TABLE `posts` (
  `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `user_id` INT NOT NULL,
  `post_title` varchar(255) NOT NULL,
  `post_content` LONGTEXT NOT NULL,
  `post_hastags` LONGTEXT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_by` LONGTEXT NOT NULL,
  `updated_by` LONGTEXT NOT NULL
  CONSTRAINT `fk_user_id_posts` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
)

CREATE TABLE IF NOT EXISTS `comments` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `post_id` INT NOT NULL,
  `user_id` INT NOT NULL,
  `comment_content` LONGTEXT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_by` LONGTEXT NOT NULL,
  `updated_by` LONGTEXT NOT NULL,
  CONSTRAINT `fk_user_id_comments` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_post_id_comments` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`)
)

CREATE TABLE IF NOT EXISTS `user_activities` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `post_id` INT NOT NULL,
  `user_id` INT NOT NULL,
  `is_liked` BOOLEAN NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_by` LONGTEXT NOT NULL,
  `updated_by` LONGTEXT NOT NULL,
  CONSTRAINT `fk_user_id_user_activities` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_post_id_user_activities` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`)
)

CREATE TABLE IF NOT EXISTS `refresh_tokens` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `user_id` INT NOT NULL,
  `refresh_token` TEXT NOT NULL,
  `expired_at` TIMESTAMP NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_by` LONGTEXT NOT NULL,
  `updated_by` LONGTEXT NOT NULL,
  CONSTRAINT `fk_user_id_user_refresh_tokens` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
)
```

#### 📞 Contact or Ask Questions

For support or inquiries, please contact:

- Email: rafiteguh6@gmail.com
- Linkedin: https://www.linkedin.com/in/raffyjo