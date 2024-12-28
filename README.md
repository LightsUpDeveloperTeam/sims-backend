# Project sims-backend

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## Makefile Commands

### Build and Run

- **Run build make command with tests**:
  ```bash
  make all
  ```
- **Build the application**:
  ```bash
  make build
  ```
- **Run the application**:
  ```bash
  make run
  ```

### Containers

- **Create containers for the application, database, Redis, and RabbitMQ**:
  ```bash
  make docker-run
  ```
- **Shutdown all containers**:
  ```bash
  make docker-down
  ```

### Testing

- **Run the database integration tests**:
  ```bash
  make itest
  ```
- **Run the full test suite**:
  ```bash
  make test
  ```

### Development

- **Live reload the application**:
  ```bash
  make watch
  ```

### Clean-up

- **Clean up binaries from the last build**:
  ```bash
  make clean
  ```

### Additional Information

- **Monitor RabbitMQ in the browser (GUI)**:
  - URL: [http://localhost:15672](http://localhost:15672)
  - Default Username: `guest`
  - Default Password: `guest`