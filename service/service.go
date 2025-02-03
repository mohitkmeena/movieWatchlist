package service


import (
	"context"
	"fmt"
	
	"log"
	"mongoapi/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb+srv://root:root@cluster0.1skef.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

// connect with mongo db
func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)
	//connect to mongodb
	
	client, err := mongo.Connect( clientOption)
	checkerr(err)
	fmt.Println("mongo connection success")

	collection=client.Database(dbName).Collection(colName)
	//collection instance is ready
	fmt.Println("collection instance is ready")
	 
}
//monfo db helpers
//insert one record
 func InsertOneMovie(movie *model.Netflix) bson.ObjectID{
	inserted,err:=collection.InsertOne(context.Background(),movie)
	checkerr(err)
	fmt.Println("inserted one movie with id ",inserted.InsertedID)
	return inserted.InsertedID.(bson.ObjectID)

 }
//update 
func UpdateOneMovie(movieId string){
    id,err:=bson.ObjectIDFromHex(movieId)
	checkerr(err)
	//shorter and clear not care about the capital  bson.M  for ordered bson.D
	filter:=bson.M{"_id":id}
	update:=bson.M{"$set":bson.M{"watched":true}}
	
	updated,err:=collection.UpdateOne(context.Background(),filter,update)
	checkerr(err)
	fmt.Println("updated ",updated.ModifiedCount)

}

//delete 
func DeleteOne(movieId string)  {
	id,err:=bson.ObjectIDFromHex(movieId)
	checkerr(err)
	//shorter and clear not care about the capital  bson.M  for ordered bson.D
	filter:=bson.M{"_id":id}
	deleted,err:=collection.DeleteOne(context.Background(),filter)
	checkerr(err)
	fmt.Println("deleted count ",deleted.DeletedCount)

}
//delete all movie
func DeleteAll()  {
	
	//shorter and clear not care about the capital  bson.M  for ordered bson.D
	filter:=bson.M{}//use this without value
	deleted,err:=collection.DeleteMany(context.Background(),filter)
	checkerr(err)
	fmt.Println("deleted count ",deleted.DeletedCount)

}
//get all
func GetAll() []bson.M  {
	cursor,err:=collection.Find(context.Background(),bson.M{})
	
	checkerr(err)
	var movies [] bson.M
	for cursor.Next(context.TODO()){
		var movie bson.M
		err:=cursor.Decode(&movie)
		checkerr(err)
		movies = append(movies, movie)

	}
	defer cursor.Close(context.Background())
    return movies
}

//get one movie
func GetOne(movieId string) bson.M{
	id,err:=bson.ObjectIDFromHex(movieId)
	checkerr(err)
	filter:=bson.M{"_id":id}
	var movie bson.M
	err=collection.FindOne(context.Background(),filter).Decode(&movie)
	checkerr(err)
	return movie
}

 //check error  
 func checkerr(err error)  {
	if err != nil {
		log.Fatal(err)
	}
 }