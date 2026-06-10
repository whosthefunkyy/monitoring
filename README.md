```mermaid
graph LR
    %% Стилизация блоков
    classDef collector fill:#E6522C,stroke:#fff,stroke-width:2px,color:#fff;
    classDef prom fill:#E6522C,stroke:#fff,stroke-width:2px,color:#fff;
    classDef loki fill:#F1A325,stroke:#fff,stroke-width:2px,color:#fff;
    classDef grafana fill:#F26122,stroke:#fff,stroke-width:2px,color:#fff;
    classDef alert fill:#DF2626,stroke:#fff,stroke-width:2px,color:#fff;
    classDef tg fill:#0088cc,stroke:#fff,stroke-width:2px,color:#fff;

    %% Секция Метрик (Prometheus)
    subgraph Metrics_Pipeline [Metrics Pipeline]
        Node[Node Exporter]:::collector -->|Scrapes| Prom[Prometheus Server]:::prom
        Cadvisor[cAdvisor]:::collector -->|Scrapes| Prom
        GoAPI[Go API Metrics]:::collector -->|Scrapes| Prom
    end

    %% Секция Логов (Loki)
    subgraph Logs_Pipeline [Logs Pipeline]
        Logs[Application Logs] -->|Collects| Tail[Promtail]:::loki
        Tail -->|Pushes| Loki[Grafana Loki]:::loki
    end

    %% Визуализация и Алертинг
    Prom -->|1. Fires Alerts| AM[Alertmanager]:::alert
    AM -->|2. Sends Notification| TG[Telegram Bot]:::tg

    %% Grafana как единый UI
    Grafana[Grafana Dashboards]:::grafana -.->|Queries Metrics| Prom
    Grafana -.->|Queries Logs| Loki

    %% Настройка связей для Grafana
    style Grafana fill:#F26122,stroke:#1F1F22,stroke-width:2px,color:#fff;
```

# Monitoring Stack with Prometheus, Grafana, Loki & Alertmanager

## Overview

Hi, I'm Artem (whosthefunky).
This project is a hands-on monitoring and observability stack built with Docker Compose.
The goal of the project was to practice real-world monitoring, logging, alerting, and infrastructure troubleshooting workflows commonly used in DevOps and Cloud environments.

The stack includes:

* Metrics collection with Prometheus
* Visualization with Grafana
* Container monitoring with cAdvisor
* Host monitoring with Node Exporter
* Centralized logging with Loki + Promtail
* Alerting with Alertmanager
* Telegram alert notifications
* A small Go API application generating logs and metrics traffic

---

# Tech Stack

* Docker & Docker Compose
* Prometheus
* Grafana
* Loki
* Promtail
* Alertmanager
* Node Exporter
* cAdvisor
* Golang

---

# Architecture

Metrics flow:

Node Exporter / cAdvisor / Go API
→ Prometheus
→ Grafana dashboards
→ Alertmanager
→ Telegram notifications

Logs flow:

Docker logs
→ Promtail
→ Loki
→ Grafana Explore

---

# Dashboards

The project contains two dashboards:

## 1. Infrastructure Dashboard

Includes:

* CPU usage
* Memory usage
* Disk usage
* Network activity
* Host metrics
* Application logs panel
* Alert list

<img width="1621" height="999" alt="new-dash" src="https://github.com/user-attachments/assets/51ff76ee-96bb-4f29-838e-c99dc0cf20b3" />



---

## 2. Containers Dashboard

Includes:

* Container CPU usage
* Container memory usage
* Network traffic
* Docker container monitoring with cAdvisor
<img width="1622" height="1078" alt="CONTAINERS" src="https://github.com/user-attachments/assets/ab79d2b8-8534-434d-8df2-2a904989cb3a" />

---

# Logging

Logs are collected from Docker containers using Promtail and stored in Loki.

Grafana Explore is used for:

* searching logs
* filtering by containers
* troubleshooting application issues

Example log query:

```logql
{compose_service="go-api"} |= "GET"
```

<img width="1621" height="1035" alt="query" src="https://github.com/user-attachments/assets/3bf02bce-3355-442e-94b5-52e0662668f0" />


---

# Alerting

Alertmanager is configured with Telegram notifications.

Implemented alerts:

* ServiceDown
* HighCPU
* HighMemory
* ContainerHighMemory

Example:

* stopping a container triggers a Prometheus alert
* Alertmanager sends a Telegram notification

<img width="1920" height="878" alt="prometheus-alert" src="https://github.com/user-attachments/assets/b9dc33a7-f4d5-4596-9145-8d3b9bb02e9b" />

<img width="1300" height="1061" alt="telegram-alert" src="https://github.com/user-attachments/assets/2ed0a915-cf9e-4afa-ad2b-579b9119f555" />


---

# How to Run
<img width="672" height="303" alt="tree" src="https://github.com/user-attachments/assets/932cd227-6ff3-490d-810e-3ec78ec07d79" />

```bash
docker compose up -d
```

Grafana:
http://localhost:3000

Prometheus:
http://localhost:9090

Alertmanager:
http://localhost:9093

---


