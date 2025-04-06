# Go CRUD Application with Gin and GORM

This project is a simple CRUD (Create, Read, Update, Delete) application built using the Go programming language, the Gin web framework, and GORM for database management. It demonstrates how to work with RESTful API endpoints, perform database migrations, and interact with an SQLite database.

## Features

- Create a new post
- Read a list of posts or a specific post by ID
- Update an existing post
- Delete a post by ID
- Database migration with GORM

## Technologies Used

- Go (Golang)
- Gin Web Framework
- GORM (Object Relational Mapping)
- SQLite (Database)

## Getting Started

### Prerequisites

To run this project, you will need the following installed on your machine:

- Go 1.23 or higher
- SQLite (database file)
- `go get` to fetch the required dependencies

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/YashTripathi01/CRUD-Go.git
   cd CRUD-Go
   ```

2. **Set up environment variables:**
   Create a `.env` file in the root directory and add the following:

   ```env
   PORT=3000
   DATABASE_URL=your-database-file.db
   ```

   Replace `your-database-file.db` with the path where you'd like to store your SQLite database file.

3. **Install dependencies:**
   Run the following command to install the required Go modules:

   ```bash
   go mod tidy
   ```

4. **Run database migrations:**
   To set up the database schema, run:

   ```bash
   go run migrate.go
   ```

5. **Run the application:**
   Start the server by running:

   ```bash
   go run main.go
   ```

   The application will be running on `http://localhost:8080`.

## API Endpoints

### 1. **Create Post**

- **Endpoint:** `POST /posts`
- **Request Body:**

  ```json
  {
    "title": "Post Title",
    "body": "Post Body"
  }
  ```

- **Response:**
  ```json
  {
    "post": {
      "id": 1,
      "title": "Post Title",
      "body": "Post Body"
    }
  }
  ```

### 2. **Get All Posts**

- **Endpoint:** `GET /posts`
- **Response:**
  ```json
  {
    "posts": [
      {
        "id": 1,
        "title": "Post Title",
        "body": "Post Body"
      },
      ...
    ]
  }
  ```

### 3. **Get Post by ID**

- **Endpoint:** `GET /posts/:id`
- **Response:**
  ```json
  {
    "post": {
      "id": 1,
      "title": "Post Title",
      "body": "Post Body"
    }
  }
  ```

### 4. **Update Post**

- **Endpoint:** `PUT /posts/:id`
- **Request Body:**

  ```json
  {
    "title": "Updated Post Title",
    "body": "Updated Post Body"
  }
  ```

- **Response:**
  ```json
  {
    "post": {
      "id": 1,
      "title": "Updated Post Title",
      "body": "Updated Post Body"
    }
  }
  ```

### 5. **Delete Post**

- **Endpoint:** `DELETE /posts/:id`
- **Response:**
  ```json
  {
    "post": "Post deleted"
  }
  ```

## Code Structure

- **`main.go`**: Initializes the application, loads environment variables, connects to the database, and sets up the routes for the API.
- **`controllers/postsController.go`**: Contains the handlers for the CRUD operations related to posts.
- **`initializers/database.go`**: Responsible for connecting to the SQLite database and setting up the GORM instance.
- **`migrations/migrate.go`**: Runs the migration to create the database schema (Post model).
- **`models/post.go`**: Defines the Post model for the database.

## Contributing

Feel free to fork this repository and create pull requests if you'd like to contribute improvements or bug fixes. If you find any issues or have suggestions, feel free to open an issue in the GitHub repository.

## License

This project is open-source and available under the [MIT License](LICENSE).
