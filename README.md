# Social Media API

This project is a social media API built using Golang. It provides endpoints for managing social media posts, users, and interactions. This README will guide you through setting up the project, running tests, and understanding how to interact with the application.

## Table of Contents
1. [Setup Instructions](#setup-instructions)
2. [Running Tests](#running-tests)
3. [Application Overview](#application-overview)
4. [Interacting with the API](#interacting-with-the-api)

## Setup Instructions

To set up the project locally, follow these steps:

1. **Install Golang**: Ensure that you have Golang installed on your machine. You can download it from the [official Go website](https://golang.org/dl/).

2. **Clone the Repository**: Clone the repository to your local machine using:
    ```sh
    git clone https://github.com/VisarutJDev/social_media_api.git
    cd social_media_api
    ```

3. **Install Dependencies**: Install the required dependencies using:
    ```sh
    go mod tidy
    ```

4. **Configure Environment Variables**: Set up your environment variables. Create a `.env` file in the root directory and add necessary configurations (e.g., database connection strings, API keys).

5. **Run the Application**: Start the application using:
    ```sh
    go run main.go
    ```

## Running Tests

To ensure the application is working correctly, you can run the tests included in the project. Follow these steps:

1. **Navigate to the Project Directory**: Ensure you are in the project directory.
    ```sh
    cd social_media_api
    ```

2. **Run Tests**: Execute the following command to run all tests:
    ```sh
    go test ./...
    ```

## Application Overview

The Social Media API is designed to provide a backend service for a social media platform. It includes functionalities such as user management and post creation.

### Features:
- **User Management**: Register, authenticate, and manage user profiles.
- **Post Management**: Create, update, delete, and retrieve posts.

## API Document Swagger
By running project with 
```sh
go run main.go
```
you can access http://localhost:8080/docs/index.html to review and interact with APIs Document (Swagger)

## Interacting with the API

The API provides several endpoints to interact with the social media platform. Below are examples of how to use some of the main endpoints.

### User Registration

**Endpoint**: `POST /users/register`

**Request Body**:
```json
{
    "username": "exampleuser",
    "email": "user@example.com",
    "password": "password123"
}
```

**Response**:
```json
{
    "message": "User registered successfully",
    "userId": "12345"
}
```

### User Authentication

**Endpoint**: `POST /users/login`

**Request Body**:
```json
{
    "email": "user@example.com",
    "password": "password123"
}
```

**Response**:
```json
{
    "message": "Login successful",
    "token": "your-jwt-token"
}
```

### Create a Post

**Endpoint**: `POST /posts`

**Request Body**:
```json
{
    "title": "My First Post",
    "content": "This is the content of my first post."
}
```

**Response**:
```json
{
    "message": "Post created successfully",
    "postId": "67890"
}
```

### Like a Post

**Endpoint**: `POST /posts/{postId}/like`

**Request Body**: Empty

**Response**:
```json
{
    "message": "Post liked successfully"
}
```

### Comment on a Post

**Endpoint**: `POST /posts/{postId}/comment`

**Request Body**:
```json
{
    "comment": "This is a comment on the post."
}
```

**Response**:
```json
{
    "message": "Comment added successfully",
    "commentId": "54321"
}
```

For more detailed documentation on all available endpoints, refer to the API documentation included in the project.

---

Feel free to customize this README further based on any additional specificities or requirements of your project.
