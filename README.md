# 🐳 Logs-Docker Setup Guide

Welcome to **Logs-Docker** — a lightweight environment to get your logging stack up and running quickly with Docker.  

---

## ✅ Quick Start Checklist

### 🔧 Step 1 — Install Prerequisites
- [ ] Install **Docker**  
- [ ] Install **Docker Compose**  
👉 Official guide: [Get Docker](https://docs.docker.com/get-started/get-docker/)

---

### 📦 Step 2 — Get the `docker-compose.yaml`
Choose one of the options below to download the configuration file:

<details>
<summary>🔽 Using curl</summary>

```bash
curl -L -o docker-compose.yaml https://raw.githubusercontent.com/rafinhacuri/logs-docker/main/docker-compose.yaml
```
</details>

<details>
<summary>🔽 Using wget</summary>

```bash
wget -O docker-compose.yaml https://raw.githubusercontent.com/rafinhacuri/logs-docker/main/docker-compose.yaml
```
</details>

Alternatively, copy it directly from the [example file](https://github.com/rafinhacuri/logs-docker/blob/main/docker-compose.yaml).

---

### 📝 Step 3 — Configure Environment
Download and prepare your `.env` file:

<details>
<summary>🔽 Using curl</summary>

```bash
curl -L -o .env https://raw.githubusercontent.com/rafinhacuri/logs-docker/main/.env.example
```
</details>

<details>
<summary>🔽 Using wget</summary>

```bash
wget -O .env https://raw.githubusercontent.com/rafinhacuri/logs-docker/main/.env.example
```
</details>

Then **edit the `.env`** with your database, ports, and other settings.

---

### 🚀 Step 4 — Launch Services
Run the following commands:

```bash
docker compose pull
docker compose up -d --force-recreate
```

---

### 🔍 Step 5 — Verify Installation
Check running containers:

```bash
docker compose ps
```

If all services show as `Up`, you’re ready! 🎉

---

## 📜 License

> Licensed under the [MIT License](https://github.com/rafinhacuri/logs-docker/blob/main/LICENSE)  
> © 2025 [Rafael Curi Leonardo](https://github.com/rafinhacuri)  

![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)