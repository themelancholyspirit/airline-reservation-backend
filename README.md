# Airline Reservation System

This is a Golang-based airline reservation system API that allows users to register, login, search for flights, book flights, and manage their bookings. The backend is built using Gin, GORM, and PostgreSQL for handling database operations.

## API Reference

#### Register a new user

```http
POST /register
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `username` | `string` | **Required**. Your username |
| `password` | `string` | **Required**. Your password |

#### Login to get a JWT token

```http
POST /login
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `username` | `string` | **Required**. Your username |
| `password` | `string` | **Required**. Your password |

#### Get a list of available flights

```http
GET /flights
```

Authentication: Bearer Token (JWT) required in the Authorization header

#### Create a new flight booking

```http
POST /bookings
```

Authentication: Bearer Token (JWT) required in the Authorization header

Example curl command:

```bash
curl -X POST http://localhost:8080/bookings \
-H "Authorization: Bearer YOUR_JWT_TOKEN" \
-H "Content-Type: application/json" \
-d '{
  "flight_id": 1,
  "seat_number": "A1",
  "booking_time": "2024-10-10T08:57:25.462Z",
  "status": "Confirmed",
  "departure_city": "JFK",
  "arrival_city": "LAX",
  "departure_time": "2024-10-11T12:42:41Z",
  "arrival_time": "2024-10-11T16:42:41Z"
}'
```

#### View a specific booking

```http
GET /bookings/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of booking to fetch |

Authentication: Bearer Token (JWT) required in the Authorization header

#### View all bookings for a user

```http
GET /bookings/user/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of user to fetch bookings for |

Authentication: Bearer Token (JWT) required in the Authorization header

