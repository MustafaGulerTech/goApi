# Go API with API Key Authentication

This is a simple Go project that demonstrates how to create a basic REST API protected by an API key. The API serves a list of users from a local JSON file (`db.json`).

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

To run this project, you need to have [Go](https://golang.org/doc/install) installed on your system.

### Running the Server

1.  Clone the repository or download the source code.
2.  Open your terminal and navigate to the project directory.
3.  Run the following command to start the API server:

    ```bash
    go run main.go
    ```

4.  The server will start on `http://localhost:8080`.

## API Reference

The API provides a single endpoint to retrieve a list of users. Access to this endpoint is restricted and requires a valid API key.

### Get All Users

Retrieves a list of all users from the database.

*   **URL:** `/users`
*   **Method:** `GET`
*   **Headers:**
    *   `X-API-KEY`: `your-secret-api-key` (This is the required API key)

---

### Usage Examples

You can use `curl` or any other API client to test the endpoint.

#### Example 1: Successful Request

This example shows a successful request with the correct API key.

```bash
curl -X GET http://localhost:8080/users \
-H "X-API-KEY: your-secret-api-key"
```

**Expected Response (200 OK):**

```json
[
  {
    "id": 1,
    "first_name": "John",
    "last_name": "Doe",
    "birth_date": "1990-05-15"
  },
  {
    "id": 2,
    "first_name": "Jane",
    "last_name": "Smith",
    "birth_date": "1988-11-22"
  },
  {
    "id": 3,
    "first_name": "Peter",
    "last_name": "Jones",
    "birth_date": "1995-01-30"
  }
]
```

#### Example 2: Unauthorized Request

This example shows a request with an invalid or missing API key.

```bash
curl -X GET http://localhost:8080/users \
-H "X-API-KEY: wrong-key"
```

**Expected Response (401 Unauthorized):**

```
Unauthorized
