# 🔗 URL Shortener

A scalable URL shortener service built with **Go**, using **Redis for caching** and **MySQL for persistence**.  
The application is containerized with **Docker** and can be deployed locally using **Kubernetes**. 🚀

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
url-shortner
│
├── cmd/
│   └── main.go
│
├── internal/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   └── model/
│
├── pkg/
│   ├── generator/
│
├── k8s/
│
├── Dockerfile
└── README.md
```

---

## 🐳 Running Locally (Docker)

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