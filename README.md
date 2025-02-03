# Movie Watchlist REST API

This project is a simple REST API for managing a movie watchlist. It allows users to perform CRUD operations on movie records stored in a MongoDB database.

## Endpoints

- **GET** `/movies`  
  Retrieve all movies.

- **GET** `/movies/{id}`  
  Retrieve a specific movie by its ID.

- **POST** `/movies`  
  Add a new movie with the following JSON body:
{
"name": "Movie Title",
"watched": false
}
- **PUT** `/movies/{id}`  
Update the watched status of a movie by its ID. The request body should include:

- **DELETE** `/movies`  
Delete all movies.

- **DELETE** `/movies/{id}`  
Delete a specific movie by its ID.

## Running the Application

To run this application, ensure you have Go installed along with MongoDB set up. You can clone this repository and run it using the following commands:

1. Clone the repository:
git clone https://github.com/mohitkmeena/movieWatchlist.git

2. Navigate to the project directory:
cd movieWatchlist
3.Make sure you changed MongoDB connectionString in service.go as per your mongodb server url .
3. Run the application:
go run main.go

Make sure your MongoDB server is running before executing the above command. 

## Additional Information

- You can use tools like Postman or curl to test the API endpoints.
- Ensure you have installed necessary dependencies using:
go get github.com/gorilla/mux
go get go.mongodb.org/mongo-driver/mongo
