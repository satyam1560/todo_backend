# Todo App

A simple RESTful API service for managing todos, built with Go.

---

## 📁 Project Structure
```
project-root/
├── main.go   
├── internal/                
│   ├── database/
│   |       ├── generated/
│   │       |    ├── db.go      
|   |       |    ├── models.go
|   |       |    └── todo.sql.go
│   |       ├── query/
|   |       |    └── todo_query.sql
│   |       ├── schema/
│   │       |   └── todo_schema.sql  
│   │       └── db_init.go 
|   |
│   └── handlers/
│       └── todo_handlers.go        
|               
├── api/                   
│   ├── router/
│   │   └── router.go      
│   └── middleware/
│       └── middleware.go  
|
├── go.mod                   
├── go.sum                  
├── sqlc.yaml                
├── .env                  
├── .air.toml                  
├── .gitignore    
├── tests/       
├── docs/     
└── README.md               

```


## 🌐 Base URL
http://localhost:8080


---

## 📌 API Endpoints

### Create a Todo
- **URL**: `/api/todos`
- **Method**: `POST`
- **Description**: Create a new todo item.

---

### Get All Todos
- **URL**: `/api/todos`
- **Method**: `GET`
- **Description**: Retrieve all todo items.

---

### Get a Todo by ID
- **URL**: `/api/todos/{id}`
- **Method**: `GET`
- **Description**: Retrieve a specific todo item by its ID.

---

### Update a Todo by ID
- **URL**: `/api/todos/{id}`
- **Method**: `PUT`
- **Description**: Update a specific todo item by its ID.

---

### Delete a Todo by ID
- **URL**: `/api/todos/{id}`
- **Method**: `DELETE`
- **Description**: Delete a specific todo item by its ID.

---

to run 
air