# Podcast API

This project is a simple REST API for managing a list of podcasts using Go and SQLite. It allows users to store and retrieve podcast information in JSON format.

## Project Structure

```
podcast-api
├── main.go          # Entry point of the application
├── handlers         # Contains HTTP handler functions
│   └── podcast.go   # Handlers for podcast-related endpoints
├── models           # Defines the data models
│   └── podcast.go   # Podcast model and database interaction methods
├── database         # Manages database connection
│   └── db.go        # Database initialization and query execution
├── go.mod           # Module definition and dependencies
├── go.sum           # Dependency checksums
└── README.md        # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone https://github.com/yourusername/podcast-api.git
   cd podcast-api
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Run the application:**
   ```
   go run main.go
   ```

4. **Access the API:**
   The API will be available at `http://localhost:8080`.

## API Endpoints

### Create Podcast List

- **Endpoint:** `POST /podcasts`
- **Description:** Stores a list of podcasts.
- **Request Body:**
  ```json
  {
    "podcasts": [
      {
        "title": "Podcast Title",
        "description": "Podcast Description",
        "url": "http://example.com/podcast"
      }
    ]
  }
  ```

### Get Podcast List

- **Endpoint:** `GET /podcasts`
- **Description:** Retrieves the list of stored podcasts.
- **Response:**
  ```json
  {
    "podcasts": [
      {
        "id": 1,
        "title": "Podcast Title",
        "description": "Podcast Description",
        "url": "http://example.com/podcast"
      }
    ]
  }
  ```

## License

This project is licensed under the MIT License. See the LICENSE file for more details.# podcast-api
