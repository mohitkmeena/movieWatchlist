# 🎬 Movie Watchlist REST API  

A simple REST API for managing a movie watchlist using **Go** and **MongoDB**. This API allows users to perform CRUD operations on movie records.  

## 🚀 Endpoints  

### 📌 Retrieve Movies  
- **GET** `/movies` - Fetch all movies.  
- **GET** `/movies/{id}` - Fetch a specific movie by its ID.  

### 🎬 Add a Movie  
- **POST** `/movies` - Add a new movie with the following JSON body:  
  ```json
  {
    "name": "Movie Title",
    "watched": false
  }
  ```  

### ✅ Update Movie Status  
- **PUT** `/movies/{id}` - Update the watched status of a movie.  

### ❌ Delete Movies  
- **DELETE** `/movies` - Delete all movies.  
- **DELETE** `/movies/{id}` - Delete a specific movie by its ID.  

## 🛠 Running the Application  

### Prerequisites  
Ensure you have **Go** installed and a **MongoDB** instance running.  

### Steps to Run  
1. Clone the repository:  
   ```sh
   git clone https://github.com/mohitkmeena/movieWatchlist.git
   ```
2. Navigate to the project directory:  
   ```sh
   cd movieWatchlist
   ```
3. Update the **MongoDB connection string** in `service.go` as per your MongoDB server URL.  
4. Install dependencies:  
   ```sh
   go get github.com/gorilla/mux
   go get go.mongodb.org/mongo-driver/mongo
   ```
5. Run the application:  
   ```sh
   go run main.go
   ```  

### 🔍 Testing the API  
- Use **Postman** or `curl` to test API endpoints.  
- Make sure your **MongoDB server** is running before executing the above command.  

---
