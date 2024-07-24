# Todo List

## Routers

### Health Check

- **Endpoint:** `GET /health`
    - **Response:** `OK`

### Create Task

- **Endpoint:** `POST /api/todo-list/tasks`
    - **Body:**
      ```json
      { 
      "title": "Купить книгу", 
      "activeAt": "2023-08-04" 
      }
      ```

### Update Task

- **Endpoint:** `PUT /api/todo-list/tasks/{ID}`
    - **Body:**
      ```json
      { 
      "title": "Купить книгу - Высоконагруженные приложения", 
      "activeAt": "2023-08-05" 
      }
      ```

### Delete Task

- **Endpoint:** `DELETE /api/todo-list/tasks/{ID}`
    - **Response:** `status: 204`

### Mark Task as Done

- **Endpoint:** `PUT /api/todo-list/tasks/{ID}/done`
    - **Response:** `status: 204`

### Get Tasks

- **Endpoint:** `GET /api/todo-list/tasks?status=active` or `GET /api/todo-list/tasks?status=done`
    - **Response:**
      ```json
      [
          {
          "id": "65f19340848f4be025160391",
          "title": "Купить книгу - Высоконагруженные приложения",
          "activeAt": "2023-08-05"
          },
          {
          "id": "75f19340848f4be025160392",
          "title": "Купить квартиру",
          "activeAt": "2023-08-05"
          },
          {
          "id": "45f19340848f4be025160394",
          "title": "Купить машину",
          "activeAt": "2023-08-05"
          }
      ]
      ```

## Models Structure

```sql
Request
{
    title: string,
    activeAt: string,
}

Task
{
    id: string,
    title: string,
    activeAt: string,
    done: bool,
}
```

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/itelman/hl-task2
   cd hl-task2
   ```
2. **Run the program:**
   ```bash
   go run main.go
   ```
3. **Check the server:**
   Open your browser and go to http://localhost:8080 to ensure the server is running properly.

**LINK: https://todo-list-hqf0.onrender.com/**
