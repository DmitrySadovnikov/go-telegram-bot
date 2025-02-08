# Telegram Bot Project

This project is a Telegram bot built with Go. It uses PostgreSQL as the database and Docker for containerization.

## Prerequisites

- Docker
- Docker Compose

## Getting Started

### Clone the Repository

```sh
git clone https://github.com/DmitrySadovnikov/go-telegram-bot.git
cd go-telegram-bot
```

### Build and Run the Project

1. **Build and start the services using Docker Compose:**

    ```sh
    docker compose up --build
    ```

2. **The application will be available at `http://localhost:8080`.**

### Environment Variables

The following environment variables are used in the project:

- `PORT`: The port on which the application runs (default: 8080).
- `DATABASE_URL`: The URL for connecting to the PostgreSQL database.

### API Endpoints

#### Messages Endpoint

- **URL:** `/api/messages`
- **Method:** `POST`
- **Description:** This endpoint processes incoming Telegram messages.

**Request:**

```sh
curl -X POST http://localhost:8080/messages -H "Content-Type: application/json" -d '{
  "update_id": 394352518,
  "message": {
    "message_id": 5,
    "from": {
      "id": 123456789,
      "is_bot": false,
      "first_name": "Dmitry",
      "last_name": "Sadovnikov",
      "username": "DmitrySadovnikov",
      "language_code": "en-RU"
    },
    "chat": {
      "id": 123456789,
      "first_name": "Dmitry",
      "last_name": "Sadovnikov",
      "username": "DmitrySadovnikov",
      "type": "private"
    },
    "date": 1527367962,
    "text": "hi"
  }
}'
```

**Response:**

```json
{
  "message": "hi"
}
```

## Development

### Running Locally

To run the project locally without Docker, ensure you have Go installed and follow these steps:

1. **Install dependencies:**

    ```sh
    go mod download
    ```

2. **Run the application:**

    ```sh
    go run main.go
    ```

### Linting

This project uses `gometalinter` for linting. To run the linter:

```sh
gometalinter --config=.gometalinter.json ./...
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
