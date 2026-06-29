# SpotSync

A parking reservation management system built with Go, Echo framework, and PostgreSQL.

## Features

- User authentication with JWT
- Role-based access control (Admin, Driver)
- Parking zone management
- Reservation booking system
- Real-time availability tracking

## Tech Stack

- **Go 1.26.3**
- **Echo v5** - Web framework
- **GORM** - ORM for PostgreSQL
- **JWT** - Authentication
- **Docker** - Containerization

## Getting Started

### Prerequisites

- Go 1.26.3 or higher
- PostgreSQL database
- Docker (optional)

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd SpotSync
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables:
Create a `.env` file in the root directory:
```env
PORT=8080
JWT_SECRET=your-secret-key
DSN=host=localhost user=postgres password=postgres dbname=spotsync port=5432 sslmode=disable
```

4. Run the application:
```bash
go run cmd/main.go
```

The server will start on `http://localhost:8080`

### Docker Setup

Build and run with Docker:
```bash
docker build -t spotsync .
docker run -p 8080:8080 --env-file .env spotsync
```

## API Endpoints

### Authentication

#### Register User
- **Endpoint:** `POST /api/v1/auth/register`
- **Description:** Register a new user with driver role
- **Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123",
  "role": "driver"
}
```
- **Response (201 Created):**
```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "role": "driver",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### Login
- **Endpoint:** `POST /api/v1/auth/login`
- **Description:** Authenticate user and receive JWT token
- **Request Body:**
```json
{
  "email": "john@example.com",
  "password": "password123"
}
```
- **Response (200 OK):**
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "role": "driver"
    }
  }
}
```

### Parking Zones

#### Get All Parking Zones
- **Endpoint:** `GET /api/v1/zones`
- **Description:** Retrieve all parking zones with availability
- **Authentication:** Not required
- **Response (200 OK):**
```json
{
  "success": true,
  "message": "Parking zones retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Downtown Parking",
      "type": "covered",
      "total_capacity": 100,
      "available_spots": 45,
      "price_per_hour": 5.50,
      "created_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

#### Get Parking Zone by ID
- **Endpoint:** `GET /api/v1/zones/:id`
- **Description:** Retrieve a specific parking zone by ID
- **Authentication:** Not required
- **Response (200 OK):**
```json
{
  "success": true,
  "message": "Parking zone retrieved successfully",
  "data": {
    "id": 1,
    "name": "Downtown Parking",
    "type": "covered",
    "total_capacity": 100,
    "available_spots": 45,
    "price_per_hour": 5.50,
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

#### Create Parking Zone
- **Endpoint:** `POST /api/v1/zones`
- **Description:** Create a new parking zone (Admin only)
- **Authentication:** Required (Admin role)
- **Headers:** `Authorization: Bearer <token>`
- **Request Body:**
```json
{
  "name": "Mall Parking",
  "type": "open",
  "total_capacity": 200,
  "price_per_hour": 3.00
}
```
- **Response (200 OK):**
```json
{
  "success": true,
  "message": "Parking zone created successfully",
  "data": {
    "id": 2,
    "name": "Mall Parking",
    "type": "open",
    "total_capacity": 200,
    "price_per_hour": 3.00,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### Reservations

#### Create Reservation
- **Endpoint:** `POST /api/v1/reservations`
- **Description:** Create a new parking reservation
- **Authentication:** Required (Admin or Driver role)
- **Headers:** `Authorization: Bearer <token>`
- **Request Body:**
```json
{
  "zone_id": 1,
  "license_plate": "ABC-1234"
}
```
- **Response (201 Created):**
```json
{
  "success": true,
  "message": "Reservation created successfully",
  "data": {
    "id": 1,
    "user_id": 1,
    "zone_id": 1,
    "license_plate": "ABC-1234",
    "status": "active",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### Get My Reservations
- **Endpoint:** `GET /api/v1/reservations/my-reservations`
- **Description:** Retrieve all reservations for the authenticated user
- **Authentication:** Required (Admin or Driver role)
- **Headers:** `Authorization: Bearer <token>`
- **Response (200 OK):**
```json
{
  "success": true,
  "message": "Reservations retrieved successfully",
  "data": [
    {
      "id": 1,
      "license_plate": "ABC-1234",
      "status": "active",
      "zone": {
        "id": 1,
        "name": "Downtown Parking",
        "type": "covered"
      },
      "created_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

#### Get All Reservations
- **Endpoint:** `GET /api/v1/reservations`
- **Description:** Retrieve all reservations (Admin only)
- **Authentication:** Required (Admin role)
- **Headers:** `Authorization: Bearer <token>`
- **Response (200 OK):**
```json
{
  "success": true,
  "message": "Reservations retrieved successfully",
  "data": [
    {
      "id": 1,
      "license_plate": "ABC-1234",
      "status": "active",
      "zone": {
        "id": 1,
        "name": "Downtown Parking",
        "type": "covered"
      },
      "created_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

#### Delete Reservation
- **Endpoint:** `DELETE /api/v1/reservations/:id`
- **Description:** Cancel a reservation
- **Authentication:** Required (Admin or Driver role)
- **Headers:** `Authorization: Bearer <token>`
- **Response (200 OK):**
```json
{
  "success": true,
  "message": "Reservation cancelled successfully"
}
```

## Error Responses

All endpoints return error responses in the following format:

```json
{
  "success": false,
  "message": "Error message",
  "errors": "Detailed error information"
}
```

Common HTTP status codes:
- `400 Bad Request` - Invalid request payload or validation failed
- `401 Unauthorized` - Missing or invalid authentication
- `403 Forbidden` - Insufficient permissions
- `500 Internal Server Error` - Server-side error

## Project Structure

```
SpotSync/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── config/              # Configuration management
│   ├── domain/              # Business logic
│   │   ├── parking_zones/   # Parking zone domain
│   │   ├── reservations/    # Reservation domain
│   │   └── users/           # User domain
│   ├── httpResponse/       # HTTP response utilities
│   ├── middlewares/         # Custom middlewares
│   ├── server/              # HTTP server setup
│   └── utils/               # Utility functions
├── .env                     # Environment variables
├── Dockerfile               # Docker configuration
├── go.mod                   # Go module file
└── go.sum                   # Go dependencies
```

## License

MIT License
