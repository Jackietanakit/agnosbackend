
## Setup

### Prerequisites

- Docker
- Docker Compose

### Running the Application

1. **Clone the Repository**

    ```sh
    git clone <repository-url>
    cd agnosbackend
    ```

2. **Build and Start Containers**

    ```sh
    docker-compose up --build
    ```

    This command will build the Docker images for the Go application and start the services defined in `docker-compose.yml`, including Nginx, Go, and PostgreSQL.

3. **Access the Application**

    Once the containers are running, you can access the API at `http://localhost:8080`. Nginx will proxy requests to the Go service.

### Running Tests

To run the unit tests for the Go application, use the following command:

```sh
docker-compose run golang_service go test ./...
