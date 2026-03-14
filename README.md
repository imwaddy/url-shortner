# 🔗 URL Shortener

A scalable URL shortener service built with **Go**, using **Redis for caching** and **MySQL for persistence**.  
The ## 🐳 Running Locally (Docker)

### Using Docker

Build the image:

```bash
docker build -t url-shortener .
```

Run the container:

```bash
docker run -p 8080:8080 \
  -e DB_USER=root \
  -e DB_PASS=password \
  -e DB_HOST=mysql:3306 \
  -e DB_NAME=urlshortener \
  url-shortener
```

### Using Docker Compose (Recommended)

```bash
docker-compose up -d
```s containerized with **Docker** and can be deployed locally using **Kubernetes**. 🚀

---

## 🏗 Architecture

```
Client
   |
   v
Go API (Gin)
   |
   +---- Redis (Cache)
   |
   +---- MySQL (Persistent Storage)
```

### 🔄 Flow

1. User sends a URL to shorten. 🌐
2. Service generates a short code. 🔑
3. URL is stored in **MySQL**. 🗄
4. Frequently accessed URLs are cached in **Redis**. ⚡
5. When the short URL is requested:
   - Check Redis first. ⚡
   - If not present → fetch from MySQL → update Redis. 🔁

---

## 🛠 Tech Stack

- **Go** 🐹
- **Gin Web Framework** 🌿
- **MySQL** 🗄
- **Redis** ⚡
- **Docker** 🐳
- **Kubernetes** ☸️

---

## 📁 Project Structure

```
url-shortener/
│
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
│
├── handler/
│   └── url.go                   # HTTP request handlers
│
├── service/
│   └── url.go                   # Business logic layer
│
├── repository/
│   └── url.go                   # Data access layer
│
├── config/
│   └── config.go                # Configuration management
│
├── pkg/                         # Shared utilities
│   ├── database/
│   │   └── mysql.go            # MySQL connection
│   ├── cache/
│   │   └── redis.go            # Redis client
│   ├── logger/
│   │   └── logger.go           # Logging utility
│   └── shortener/
│       └── generator.go        # Short code generator
│
├── k8s/                         # Kubernetes manifests
│   ├── app/
│   ├── mysql/
│   └── redis/
│
├── Dockerfile
├── go.mod
└── README.md
```

### 🏛️ Architecture Principles

- **Clean separation of concerns** - handler → service → repository
- **Simplified structure** - removed unnecessary `internal/` nesting
- **Reusable utilities** - shared packages in `pkg/`
- **Easy navigation** - flat structure for better developer experience

---

## � Getting Started

### Prerequisites

- Go 1.24+ 🐹
- Docker & Docker Compose 🐳
- kubectl (for Kubernetes deployment) ☸️
- MySQL 8.0+ 🗄
- Redis 6.0+ ⚡

### Environment Variables

Create a `.env` file or set these environment variables:

```bash
PORT=8080
DB_USER=root
DB_PASS=password
DB_HOST=localhost:3306
DB_NAME=urlshortener
REDIS_ADDR=localhost:6379
```

### Local Development

1. **Install dependencies:**
   ```bash
   go mod download
   ```

2. **Run MySQL and Redis:**
   ```bash
   docker-compose up -d mysql redis
   ```

3. **Run the application:**
   ```bash
   go run cmd/server/main.go
   ```

The service will be available at `http://localhost:8080`

---

## �🐳 Running Locally (Docker)

Build the image:

```bash
docker build -t url-shortner .
```

Run the container:

```bash
docker run -p 8080:8080 url-shortner
```

---

## ☸️ Running with Kubernetes

Apply the resources:

```bash
kubectl apply -f k8s/mysql.yaml
kubectl apply -f k8s/redis.yaml
kubectl apply -f k8s/url-shortner.yaml
```

Check running resources:

```bash
kubectl get pods
kubectl get svc
```

Example output:

```
NAME           TYPE        PORT(S)
mysql          ClusterIP   3306
redis          ClusterIP   6379
url-shortner   NodePort    80:31639
```

---

## 🌍 Access the Service

Use the NodePort exposed by Kubernetes:

```bash
curl http://localhost:31639/<short-code>
```

Example:

```bash
curl http://localhost:31639/YwNMMv
```

---

## 📡 API

### ✂️ Create Short URL

```
POST /shorten
```

Body:

```json
{
  "url": "https://example.com"
}
```

Response:

```json
{
  "short": "abc123"
}
```

---

### 🔁 Redirect to Original URL

```
GET /{short_code}
```

Example:

```
GET "http://localhost:31639/abc123"
```

Response: **HTTP 302 Redirect** ↪️

---

## ✨ Features

- URL shortening 🔗
- Redirect using short code ↪️
- Redis caching layer ⚡
- MySQL persistent storage 🗄
- Docker containerization 🐳
- Kubernetes deployment ☸️

---

## 🚀 Future Improvements

- Analytics for URL clicks 📊
- Rate limiting 🚦
- Expiration for short URLs ⏳
- Custom aliases 🏷
- Distributed ID generation ⚙️

---

## 👨‍💻 Author

**Mayur Wadekar**

GitHub:  
https://github.com/imwaddy

---

# 📜 License

MIT License