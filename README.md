## Overview

The Agnos Backend project is a project created for interviews at Agnos Health. The project uses Go with the Gin framework, Nginx as a reverse proxy, and PostgreSQL for the database. The project will be built using Docker.

## Setup

### Prerequisites

- Docker
- Docker Compose

### Running the Application

1. **Clone the Repository**

    ```sh
    git clone https://github.com/Jackietanakit/tanakit_thangjantaraprapab_agnos_backend.git
    cd agnosbackend
    ```

2. **Build and Start Containers**

    ```sh
    docker-compose up --build
    ```

    This command will build the Docker images for the Go application and start the services defined in `docker-compose.yml`, including Nginx, Go, and PostgreSQL.

3. **Access the Application**

    Once the containers are running, you can access the API at `http://localhost:8080`. Nginx will proxy requests to the Go service.

## API Endpoints

- GET /: Returns "Hello, World!" message
- POST /api/strong_password_steps: Accepts a JSON payload with an `init_password` field and returns the minimus steps required to make the password strong.

### Example Request
    ```sh
    curl -X POST http://localhost:8080/api/strong_password_steps \
    -H "Content-Type: application/json" \
    -d '{"init_password": "password123"}'
    ```
### Example Response
    ```sh
    {
    "num_of_steps": 1
    }
    ```

## Running Tests

- The unit test checks weather or not the logic of the minimum steps is correct.

To run the unit tests for the Go application:

    ```sh
    docker-compose run golang_service go test ./...
    ```


