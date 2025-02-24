# 🚀 **Step-by-Step Guide to Run the Project**

## **1️⃣ Install Required Dependencies**
Ensure you have the following installed on your system:

- [Docker](https://www.docker.com/products/docker-desktop/) (latest version)
- [Docker Compose](https://docs.docker.com/compose/install/)
- OpenSSL (for self-signed SSL certificate)

---

## **2️⃣ Clone the Repository**
If you haven’t already, clone your repository:

```bash
git clone https://github.com/lightsupdeveloperteam/sims-backend.git
cd sims-backend
```

---

## **3️⃣ Setup `.env` Configuration**
Create a `.env` file at the project root:

```bash
touch .env
```

Then, add the following environment variables:

```ini
# Application Config
PORT=3000
APP_ENV=local

# Database Configuration
BLUEPRINT_DB_HOST=db
BLUEPRINT_DB_PORT=5432
BLUEPRINT_DB_DATABASE=blueprint
BLUEPRINT_DB_USERNAME=melkey
BLUEPRINT_DB_PASSWORD=password1234
BLUEPRINT_DB_SCHEMA=public

# Redis Configuration
REDIS_HOST=redis
REDIS_PORT=6379

# RabbitMQ Configuration
RABBITMQ_DEFAULT_USER=guest
RABBITMQ_DEFAULT_PASS=guest

# JWT Secret Key
JWT_SECRET_KEY=my_super_secret_key

# Traefik Configuration
FQDN_DOMAIN=api.simsbe.test
```

> **🔹 Note:**  
> - Ensure that `FQDN_DOMAIN` is set to `api.simsbe.test` as used in Traefik.
> - The database and Redis use the **container names** (`db`, `redis`) as their hosts since they are within the same Docker network.

---

## **4️⃣ Generate SSL Certificates**
Since Traefik requires a valid SSL certificate, generate a self-signed certificate:

```bash
mkdir certs
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout certs/api.simsbe.test.key \
  -out certs/api.simsbe.test.crt \
  -subj "/CN=api.simsbe.test/O=Local Testing"
```

This will create:
- `certs/api.simsbe.test.key` (private key)
- `certs/api.simsbe.test.crt` (certificate)

---

## **5️⃣ Update Hosts File**
You need to map `api.simsbe.test` to `127.0.0.1` for local testing.

Edit your **hosts file**:

- **Linux/macOS:**  
  ```bash
  sudo nano /etc/hosts
  ```
- **Windows:**  
  Edit `C:\Windows\System32\drivers\etc\hosts` with Notepad (Run as Administrator).

Add this line:

```
127.0.0.1  api.simsbe.test
```

Save and exit.

---

## **6️⃣ Build & Run Docker Containers**
Now, start all services using **Docker Compose**:

```bash
docker-compose up --build
```

This will:
- Build and run the **SIMS backend API**.
- Start PostgreSQL, Redis, RabbitMQ.
- Start **Traefik as a reverse proxy with HTTPS**.

---

## **7️⃣ Verify That Services Are Running**
Check running containers:

```bash
docker ps
```

You should see something like this:

```bash
CONTAINER ID   IMAGE                  PORTS                                NAMES
xxxxxxxxxxxx   traefik:latest         0.0.0.0:80->80/tcp, 0.0.0.0:443->443/tcp  traefik
xxxxxxxxxxxx   sims-backend           3000/tcp                             sims-backend
xxxxxxxxxxxx   postgres:latest        0.0.0.0:5432->5432/tcp               sims-postgres
xxxxxxxxxxxx   redis:latest           0.0.0.0:6379->6379/tcp               sims-redis
xxxxxxxxxxxx   rabbitmq:3-management  0.0.0.0:5672->5672/tcp, 0.0.0.0:15672->15672/tcp sims-rabbitmq
```

---

## **8️⃣ Test the API**
Now, test if your API is running by visiting:

🔹 **Secure API Endpoint:**  
👉 [https://api.simsbe.test](https://api.simsbe.test) (your browser might warn about the self-signed SSL)

Or use **cURL**:

```bash
curl -k https://api.simsbe.test
```

🔹 **RabbitMQ Management UI:**  
👉 [http://localhost:15672](http://localhost:15672)  
Login with:
- **Username:** `guest`
- **Password:** `guest`

🔹 **Traefik Dashboard:**  
👉 [http://localhost:8080](http://localhost:8080) (only for debugging)

---

## **9️⃣ Stop the Project**
To stop all services:

```bash
docker-compose down
```

This will gracefully stop and remove the containers.

---

## **🔄 Optional: Auto-Restart on Code Changes**
For live reloading, install **Air** (Go hot-reload tool):

```bash
go install github.com/cosmtrek/air@latest
```

Then, start the API inside the container in development mode:

```bash
docker exec -it sims-backend air
```

This will reload the Go app whenever code changes are detected.

---

# 🎯 **Project Summary**
| **Component** | **Description** | **Access URL** |
|--------------|----------------|---------------|
| **Go API** | Main backend service | [https://api.simsbe.test](https://api.simsbe.test) |
| **PostgreSQL** | Database | `postgres://melkey:password1234@db:5432/blueprint` |
| **Redis** | Caching & session store | `redis://redis:6379` |
| **RabbitMQ** | Message broker | `amqp://guest:guest@rabbitmq:5672/` |
| **RabbitMQ UI** | Management dashboard | [http://localhost:15672](http://localhost:15672) |
| **Traefik** | Reverse proxy & SSL termination | [http://localhost:8080](http://localhost:8080) |

---

# ✅ **Congratulations!** 🎉  
You have successfully set up a **fully containerized development environment** for your **SIMS Backend** project using Docker, Traefik, PostgreSQL, Redis, and RabbitMQ!

🚀 **Now you’re ready to start building your application!** 🚀
