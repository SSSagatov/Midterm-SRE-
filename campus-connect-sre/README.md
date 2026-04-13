# Campus Connect SRE Project

Production-ready midterm project with:
- **Frontend:** HTML + CSS + JavaScript + Nginx
- **Backend:** Go
- **Database:** PostgreSQL
- **Observability:** Prometheus + Grafana + Node Exporter
- **Alerting:** Warning + Critical Prometheus alerts

## 1) What the app does
Campus Connect is a small discussion board for students.
Users can:
- view recent posts
- publish a new post
- like existing posts
- check service health

This gives you meaningful business logic for SLI/SLO:
- availability of API requests
- latency of API requests
- success of post creation

## 2) Project structure

```text
campus-connect-sre/
├── backend/
├── frontend/
├── postgres/
├── prometheus/
├── grafana/
├── docker-compose.yml
└── .env.example
```

## 3) Start the project

```bash
cp .env.example .env
docker compose up -d --build
```

## 4) URLs
- Frontend: http://localhost
- Backend API: http://localhost:8080
- Health: http://localhost:8080/healthz
- Metrics: http://localhost:8080/metrics
- Prometheus: http://localhost:9090
- Grafana: http://localhost:3000

## 5) Demo API calls

Create post:
```bash
curl -X POST http://localhost:8080/api/posts   -H "Content-Type: application/json"   -d '{"author":"Aruzhan","title":"Need help with Docker","content":"Who can explain healthchecks?"}'
```

List posts:
```bash
curl http://localhost:8080/api/posts
```

Like a post:
```bash
curl -X POST http://localhost:8080/api/posts/1/like
```

## 6) Trigger alerts manually

### Warning alert (high latency)
```bash
curl "http://localhost:8080/debug/slow?seconds=6"
```

### Critical alert (force failures)
```bash
for i in {1..20}; do curl -s http://localhost:8080/debug/fail; done
```

## 7) Suggested SLI / SLO

### SLI 1: Availability
Percentage of successful HTTP requests:
- good events = requests not returning 5xx
- total events = all HTTP requests

**SLO:** 99.5% monthly availability

**Monthly error budget:**  
30 days = 43,200 minutes  
0.5% of 43,200 = **216 minutes**

### SLI 2: Latency
Percentage of API requests completed in less than 500 ms.

**SLO:** 95% of requests under 500 ms

## 8) For the report
Take screenshots of:
- website running
- docker compose services
- Prometheus targets
- Grafana dashboard
- alert in FIRING state

## 9) Notes
- The frontend is a separate service on Nginx.
- The backend is a separate Go API service.
- The database is a separate PostgreSQL service.
- This structure matches the teacher requirement more formally than a single monolith.
