## 🚀 Setup

Follow the steps below to get the project up and running:

---

### 1. Install Docker & Docker Compose
Make sure you have Docker and Docker Compose installed on your system.  
👉 Follow the official guide: [Docker Installation](https://docs.docker.com/get-started/get-docker/)

---

### 2. Create `docker-compose.yml`
Download the provided example configuration:

```bash
# Using curl
curl -L -o docker-compose.yml https://raw.githubusercontent.com/rafinhacuri/logs-docker/main/docker-compose.yml

# Or using wget
wget -O docker-compose.yml https://raw.githubusercontent.com/rafinhacuri/logs-docker/main/docker-compose.yml
```

> 💡 Alternatively, you can copy the contents directly from the [example file](https://github.com/rafinhacuri/logs-docker/blob/main/docker-compose.yml).

---

### 3. Create `.env` file
Set up your environment variables by downloading the example `.env`:

```bash
# Using curl
curl -L -o .env https://raw.githubusercontent.com/rafinhacuri/logs-docker/main/.env.example

# Or using wget
wget -O .env https://raw.githubusercontent.com/rafinhacuri/logs-docker/main/.env.example
```

> ⚙️ Edit the `.env` file to match your environment (database, ports, etc.).

---

### 4. Start the Services
Pull and start all containers:

```bash
docker compose pull
docker compose up -d --force-recreate
```

---

### 5. Verify Installation
Check that everything is running correctly:

```bash
docker compose ps
```

If all services show as `Up`, your setup is complete ✅

---

## 📝 License

This project is licensed under the [MIT License](https://github.com/rafinhacuri/logs-docker/blob/main/LICENSE).  
Copyright © 2025 [Rafael Curi Leonardo](https://github.com/rafinhacuri)