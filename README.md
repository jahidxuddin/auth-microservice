# auth-microservice

## Installation Steps

1. [Install Go](#1-install-go)
2. [Clone the Project](#2-clone-the-project)
3. [Build the Project](#3-build-the-project)
4. [Generate JWT Secret](#4-generate-jwt-secret)
5. [Format the Data Source Name (DSN)](#5-format-the-data-source-name-dsn)

### 1. Install Go

If you haven't installed Go yet, follow these steps:

- Go to the [Go Downloads page](https://golang.org/dl/) and download the installer for your operating system.
- Follow the instructions to install Go.

To verify that Go is installed correctly, run the following command in your terminal:

```bash
go version
```

### 2. Clone the Project
Once Go is installed, clone the project repository to your local machine:

```bash
git clone <repository_url>
cd <repository_directory>
```
Replace ```<repository_url>``` with the URL of the repository and ```<repository_directory>``` with the name of the cloned project directory.

### 3. Build the Project
After cloning the project, navigate to the project directory and build it using the following command:

```bash
go build
```
This will generate the executable binary for your project in the current directory.

### 4. Generate JWT Secret
For the project to work, you will need a valid JWT secret key. Follow these steps to generate it:

Go to https://jwtsecret.com/generate.
Click on "Generate".
Copy the generated secret string.
The program will ask you for a JWT secret key.

### 5. Format the Data Source Name (DSN)
When configuring the database connection, the Data Source Name (DSN) must be formatted as follows:

```bash
<username>:<password>@tcp(<url>:<port>)/<db_name>
```
Replace the placeholders with your actual database credentials and details:

- ```<username>```: The username to access your database.
- ```<password>```: The password for the database user.
- ```<url>```: The host or IP address of the database server.
- ```<port>```: The port number on which the database is - running (e.g., 3306 for MySQL).
- ```<db_name>```: The name of the database you want to connect to.

Example:

```bash
user:secret@tcp(127.0.0.1:3306)/mydatabase
```

## Calling the Endpoints
Once the auth-microservice is running, you can call its endpoints to perform user authentication and related tasks.

### Available Endpoints
1. **POST /api/v1/login**
- **Description**: Authenticates a user with the provided credentials (username and password).
- **Request body**:

```json
{
  "username": "user123",
  "password": "password123"
}
```
- **Response**:

```json
{
  "message": "Login successful",
  "token": "..."
}
```

2. **POST /api/v1/register**
- **Description**: Registers a new user with the provided credentials (username, password).
- **Request body**:

```json
{
  "username": "newuser",
  "password": "newpassword123"
}
```
- **Response**:

```json
{
  "message": "User successfully registered"
}
```

## License
This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
