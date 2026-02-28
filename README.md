# 🔗 URL Shortener (Go + MySQL + Docker)

A production-ready URL Shortener built with:

-   Go (Golang)
-   MySQL
-   Docker
-   REST APIs

------------------------------------------------------------------------

## 📌 Features

-   Create short URLs
-   Redirect to original URL
-   MySQL persistent storage
-   Auto database migration
-   Dockerized setup
-   Environment variable configuration

------------------------------------------------------------------------

## 🏗 Architecture

Client → API (Go App) → MySQL Database

------------------------------------------------------------------------

## 📁 Project Structure

. 
├── main.go 
├── handlers/ 
├── models/ 
├── config/ 
├── Dockerfile 
├── docker-compose.yml 
├── README.md

------------------------------------------------------------------------

# 🚀 Getting Started

## 1️⃣ Run Locally (Without Docker)

### Create Database

CREATE DATABASE shortner;

### Set Environment Variables

export DB_USER=root\
export DB_PASS=root\
export DB_HOST=localhost:3306\
export DB_NAME=shortner

### Run Application

go run main.go

App runs on:

http://localhost:8080

------------------------------------------------------------------------

# 🐳 Run with Docker

## Build Image

docker build -t url-shortner .

## Run MySQL Container

docker run -d\
--name mysql\
-e MYSQL_ROOT_PASSWORD=root\
-e MYSQL_DATABASE=shortner\
-p 3306:3306\
mysql:8

## Run App Container

docker run -d\
--name app\
--link mysql:mysql\
-p 8080:8080\
-e DB_USER=root\
-e DB_PASS=root\
-e DB_HOST=mysql:3306\
-e DB_NAME=shortner\
url-shortner

Application will be available at:

http://localhost:8080

------------------------------------------------------------------------

# 🔌 API Endpoints

## 1️⃣ Create Short URL

POST /shorten

Request Body:

{ "url": "https://google.com" }

Response:

{ "short_url": "http://localhost:8080/abc123" }

------------------------------------------------------------------------

## 2️⃣ Redirect

GET /{shortCode}

Example:

http://localhost:8080/abc123

Redirects to the original URL.

------------------------------------------------------------------------

# ⚙️ Environment Variables

DB_USER - Database username\
DB_PASS - Database password\
DB_HOST - Database host\
DB_NAME - Database name

------------------------------------------------------------------------

# 🧠 Future Improvements

-   Kubernetes deployment
-   Persistent storage setup
-   Secrets management
-   Rate Limiting
-   Analytics tracking
-   Redis caching
-   JWT authentication

------------------------------------------------------------------------

# 👨‍💻 Author

Mayur Wadekar

------------------------------------------------------------------------

# 📜 License

MIT License
