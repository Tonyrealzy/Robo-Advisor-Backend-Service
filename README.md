# Robo-Advisor AI & Backend Services

## Overview

The robo-advisor provides personalized investment recommendations and portfolio management services to bank customers. Utilizing artificial intelligence and machine learning algorithms to analyze customer's risk tolerance, investment goals, and financial preferences. The robo-advisor then generates customized investment portfolios tailored to each customer's unique needs, optimizing returns while minimizing risk.

- **Backend Service (Go)**: A supporting backend that manages file uploads, metadata, and presigned file access. It's hosted on [https://robo-advisor-backend-service.onrender.com/api](https://robo-advisor-backend-service.onrender.com/api) with Swagger documentation on [https://robo-advisor-backend-service.onrender.com/api/swagger/index.html](https://robo-advisor-backend-service.onrender.com/api/swagger/index.html).

## 🐳 Docker Setup

We use Docker Compose to run both services together.

**Backend Service (Go + Gin)**

## 🚀 Running Locally with Docker

Make sure Docker is installed on your machine. Then run:

```
docker-compose up --build
```

This will start:

1. Go backend service on [http://localhost:8080](http://localhost:8080)
2. Swagger UI on [http://localhost:8000/api/swagger/index.html](http://localhost:5000/api/swagger/index.html)

```
Sample JSON payload:
{
"age": 30,
"location": "US",
"investmentKnowledge": "moderate",
"investmentPurpose": "retirement",
"investmentHorizon": 10,
"riskTolerance": "medium",
"amount": 50000,
"currency": "USD"
}
```
