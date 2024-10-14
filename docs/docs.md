## Documentation

### API Endpoints

#### GET `/person`
- **Description**: Retrieve all persons.
- **Response**: 
  ```json
  [
    {
      "id": "uuid",
      "name": "John Doe",
      "age": 30,
      "hobbies": ["reading", "swimming"]
    },
    ...
  ]
  ```

#### GET `/person/${personId}`
- **Description**: Retrieve a person by `personId`.
- **Response**: 
  ```json
  {
    "id": "uuid",
    "name": "John Doe",
    "age": 30,
    "hobbies": ["reading", "swimming"]
  }
  ```

#### POST `/person`
- **Description**: Create a new person.
- **Request Body**:
  ```json
  {
    "name": "John Doe",
    "age": 30,
    "hobbies": ["reading", "swimming"]
  }
  ```
- **Response**: 
  ```json
  {
    "id": "uuid",
    "name": "John Doe",
    "age": 30,
    "hobbies": ["reading", "swimming"]
  }
  ```

#### PUT `/person/${personId}`
- **Description**: Update an existing person.
- **Request Body**:
  ```json
  {
    "name": "John Doe",
    "age": 31,
    "hobbies": ["reading", "swimming", "cycling"]
  }
  ```
- **Response**: 
  ```json
  {
    "id": "uuid",
    "name": "John Doe",
    "age": 31,
    "hobbies": ["reading", "swimming", "cycling"]
  }
  ```

#### DELETE `/person/${personId}`
- **Description**: Delete a person by `personId`.
- **Response**: 
  ```json
  {
    "message": "Person deleted successfully"
  }
  ```

### Error Handling

- **404 Not Found**: For non-existing endpoints.
  ```json
  {
    "error": "Resource not found"
  }
  ```

- **500 Internal Server Error**: For internal server errors.
  ```json
  {
    "error": "Internal server error"
  }
  ```

### Cross-Site Resource Sharing (CORS)

To ensure the API is accessible by frontend apps hosted on a different domain, configure CORS in your Go application. Here is an example using the `github.com/rs/cors` package:

```go
package main

import (
    "github.com/rs/cors"
    "net/http"
)

func main() {
    mux := http.NewServeMux()
    // Define your handlers here

    handler := cors.Default().Handler(mux)
    http.ListenAndServe(":8080", handler)
}
```

This configuration allows all origins by default. You can customize the CORS settings as needed.
