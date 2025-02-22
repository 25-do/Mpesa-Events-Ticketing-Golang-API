# Ticketing API

## Overview
The **Ticketing API** provides a robust system for managing event tickets, including booking, validation, and third-party access token management with **BigCache**.

## Prerequisites
- Install **Go** (latest version): [Download Go](https://go.dev/dl/)
- Install **Git**
- Setup a **PostgreSQL/MySQL** database (or your preferred DB)

## Installation
### 1. Clone the Repository
```sh
git clone https://github.com/your-username/ticketing-api.git
cd ticketing-api
```

### 2. Install Dependencies
```sh
go mod tidy
```

### 3. Configure Environment Variables
Create a `.env` file in the root directory and add:
```env
DB_HOST=your_database_host
DB_PORT=your_database_port
DB_USER=your_database_user
DB_PASSWORD=your_database_password
DB_NAME=your_database_name
CACHE_EXPIRY=59 # Access token expiry in minutes
```

### 4. Run Database Migrations (if applicable)
```sh
go run migrate.go
```

### 5. Start the Server
```sh
go run main.go
```
The API should now be running on `http://localhost:8080`

## API Endpoints
### Authentication
```go
// Authenticate user
POST /api/auth/login

// Register a new user
POST /api/auth/register
```

### Events
```go
// Get all events
GET /api/event/

// Create an event
POST /api/event/create
{
  "name": "Soul Fest",
  "description": "soul generation festival",
  "venue_id": 5,
  "organizer_id": 1,
  "start_time": "2025-02-02T07:05:27.860335+03:00",
  "end_time": "2025-02-02T07:05:27.860335+03:00",
  "is_free": true,
  "image_url": "https://cloudinary.com"
}


```



### Third-Party Access Token Management
The API internally fetches an access token from MPESA DARAJA API and stores it in **BigCache** with a 59-minute expiry.


## Deployment
For production, use Docker or a cloud platform like AWS/GCP.

### Build & Run with Docker
```sh
docker build -t ticketing-api .
docker run -p 8080:8080 ticketing-api
```

## Contributing
Feel free to submit pull requests or open issues.

## License
MIT License

