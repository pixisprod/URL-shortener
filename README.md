# 🔗 URL Shortener (Go + Redis + PostgreSQL)

A blazing-fast and minimal URL shortener built with **Go**, **Gin**, **PostgreSQL**, and **Redis** cache.  
Runs fully in **Docker Compose** — no manual setup required.

---

## ⚙️ Stack
- **Go** — backend & REST API (Gin)
- **PostgreSQL** — persistent storage
- **Redis** — cache layer for short links
- **Migrate** — database migrations
- **Docker Compose** — service orchestration

---

## 🚀 Run Locally

Clone the repository:

```bash
git clone https://github.com/pixisprod/URL-shortener.git
cd url-shortener
```

Start all services:

```bash
docker compose up --build
```

The app will be available at:
```
http://localhost:8080
```

---

## 📡 API Endpoints

| Method | Endpoint             | Description            |
|--------|----------------------|------------------------|
| POST   | `/api/links/cut`     | Generate a short link  |
| GET    | `/api/links/r/:hash` | Redirect to full URL   |

---

## 🧱 Project Structure

```
.
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── cache/
│   ├── config/
│   ├── controller/
│   ├── database/
│   ├── model/
│   ├── repository/
│   ├── route/
│   └── service/
├── migrations/
├── Dockerfile
├── docker-compose.yml
└── go.mod
```

---

## 📋 To-Do
- [x] TTL for links
- [ ] Healthcheck
- [ ] Metrics
- [ ] Rate limiting
- [ ] Logging
- [ ] Unit testing
- [ ] Analytics

- Account system

---

## 📄 License

MIT © 2025 — [PixisProd](https://github.com/pixisprod)
