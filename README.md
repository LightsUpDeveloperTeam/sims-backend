# Project sims-backend

This is the backend service for the SIMS project, built using Go with a modular monolith architecture. It integrates PostgreSQL as the database, Redis for caching and session management, and RabbitMQ as the message broker.

## Getting Started

These instructions will help you set up the project in a Dockerized environment for development and testing purposes.

---

## Prerequisites

Ensure you have the following installed on your system:
- Docker
- Docker Compose

---

## Makefile Commands

### Build and Run

- **Build the application inside Docker**:
  ```bash
  make docker-build
  ```
- **Run the application inside Docker**:
  ```bash
  make docker-run-app
  ```

### Containers

- **Create and start all containers (App, DB, Redis, RabbitMQ)**:
  ```bash
  make docker-run
  ```
- **Shutdown all containers**:
  ```bash
  make docker-down
  ```

### Testing

- **Run all tests inside Docker**:
  ```bash
  make docker-test
  ```
- **Run the database integration tests inside Docker**:
  ```bash
  make docker-itest
  ```

### Development

- **Live reload the application inside Docker**:
  ```bash
  make docker-watch
  ```

### Clean-up

- **Clean up binaries inside Docker**:
  ```bash
  make docker-clean
  ```

---

## Monitoring and Debugging

### RabbitMQ

RabbitMQ comes with a built-in management interface for monitoring and debugging:
- URL: [http://localhost:15672](http://localhost:15672)
- Default Username: `guest`
- Default Password: `guest`

---

## Environment Variables

The project uses the following environment variables, which are defined in a `.env` file:

```dotenv
# PostgreSQL
BLUEPRINT_DB_USERNAME=postgres
BLUEPRINT_DB_PASSWORD=password
BLUEPRINT_DB_DATABASE=mydb
BLUEPRINT_DB_PORT=5432

# RabbitMQ
RABBITMQ_DEFAULT_USER=guest
RABBITMQ_DEFAULT_PASS=guest

# Redis
REDIS_HOST=redis
REDIS_PORT=6379
```

---

## Usage

1. **Build and start the application**:
   ```bash
   make docker-run
   ```

2. **Access the API**:
  - Open your browser or use a tool like Postman to interact with the API at `http://localhost:3000`.

3. **Monitor RabbitMQ**:
  - Visit the RabbitMQ management interface at `http://localhost:15672`.

4. **Stop all services**:
   ```bash
   make docker-down
   ```

---

## Project Architecture

This project follows a modular monolith architecture, organized as follows:
- **App**: Main Go application.
- **Database**: PostgreSQL.
- **Cache/Session Manager**: Redis.
- **Message Broker**: RabbitMQ.

---

## Contributing

Contributions are welcome! Please follow the standard Git flow process:
1. Fork the repository.
2. Create a feature branch.
3. Commit your changes.
4. Submit a pull request.

---

## License

This project is licensed under the MIT License.

---

### Highlights of the Updates
1. **Docker-Focused Workflow**:
   - Updated all commands to reference the Dockerized environment.
   - Added instructions for using `make` commands for building, running, testing, and cleaning up.

2. **Redis and RabbitMQ**:
   - Included Redis and RabbitMQ usage details.
   - Added RabbitMQ monitoring instructions.

3. **Environment Variables**:
   - Highlighted the use of a `.env` file for easy configuration.

4. **Architecture Overview**:
   - Provided a brief description of the project's modular monolith architecture.

---
