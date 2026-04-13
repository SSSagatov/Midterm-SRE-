# Report Guide for the Teacher

## Step 1. Containerization
Explain that your project is split into:
- frontend service (Nginx + HTML/CSS/JS)
- backend service (Go REST API)
- database service (PostgreSQL)
- observability services (Prometheus, Grafana, Node Exporter)

Mention that all services are started through Docker Compose.

## Step 2. SLI, SLO, Error Budget

### SLI 1 — Availability
Formula:
Good requests / Total requests

Prometheus idea:
```promql
1 - (sum(rate(http_requests_total{status=~"5.."}[5m])) / sum(rate(http_requests_total[5m])))
```

SLO:
- 99.5% monthly availability

Monthly Error Budget:
- 30 days = 43,200 minutes
- 0.5% = 216 minutes of allowed failed availability

### SLI 2 — Latency
Formula:
Requests under 500 ms / Total requests

Possible PromQL:
```promql
sum(rate(http_request_duration_seconds_bucket{le="0.5"}[5m])) / sum(rate(http_request_duration_seconds_count[5m]))
```

SLO:
- 95% of requests should complete in under 500 ms

## Step 3. Monitoring and Dashboard
Golden Signals included:
- latency
- traffic
- errors
- host resource usage

## Step 4. Alerting Validation
Two alerts:
- Warning: HighLatencyWarning
- Critical: HighErrorRateCritical

Manual validation:
- `/debug/slow?seconds=6`
- `/debug/fail`

Add screenshots of FIRING alerts in Prometheus.
