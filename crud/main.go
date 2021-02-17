package crud

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type Datastore struct {
	Client *mongo.Client
}

type Person struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Username string             `bson:"username"`
	Lastname string             `bson:"lastname"`
	Age      uint               `bson:"age"`
}

// DB connection
func Connect() *mongo.Client {
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return client
}

//create database and collection
func (d *Datastore) CreateDB(dbname, collName string) {
	db := d.Client.Database(dbname)
	collection := db.Collection(collName)
	//set username field as unique
	_, err := collection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("collection %s is created in database %s:\n", collection.Name(), db.Name())
	fmt.Println(collection.Indexes().List(context.TODO()))
}

func (d *Datastore) DropDB(dbname string) {
	db := d.Client.Database(dbname)
	if err := db.Drop(context.TODO()); err != nil {
		fmt.Println(err)
	}
}

func (d *Datastore) Insert(p *Person) {
	db := d.Client.Database("userrepo")
	collection := db.Collection("users")
	result, err := collection.InsertOne(context.TODO(), p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result after insert:", result.InsertedID)
}

func (d *Datastore) FindOne(username string) {
	collection := d.Client.Database("userrepo").Collection("users")
	person := Person{}
	err := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&person)
	if err != nil {
		fmt.Println("error in findOne :", err)
		return
	}
	fmt.Println("fetched data:", person)
}

func (d *Datastore) DeleteOne(username string) {
	collection := d.Client.Database("userrepo").Collection("users")
	/*a,err := primitive.ObjectIDFromHex(	"5ffe81833d081940d16cf26e")
	if err!=nil{
		fmt.Println("cannot parse object id: ",err)
		return
	}*/
	result, err := collection.DeleteOne(context.TODO(), bson.M{"username": username})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func (d *Datastore) FindAll() {
	collection := d.Client.Database("userrepo").Collection("users")
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Println("error in findAll :", err)
		return
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	for _, result := range results {
		fmt.Println(result)
	}
}

func (d *Datastore) ProjectionFind(username string) {
	collection := d.Client.Database("userrepo").Collection("users")
	person := Person{}
	//exclude lastname while retrieving
	opts := options.FindOne().SetProjection(bson.M{"lastname": 0})
	err := collection.FindOne(context.TODO(), bson.M{"username": username}, opts).Decode(&person)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("exclude lastname:", person)
	//include lastname while retrieving
	opts = options.FindOne().SetProjection(bson.M{"lastname": 1})
	err = collection.FindOne(context.TODO(), bson.M{"username": username}, opts).Decode(&person)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(person)
}

func (d *Datastore) UpdateAge(username string, age int) {
	collection := d.Client.Database("userrepo").Collection("users")
	update := bson.M{
		"$set": bson.M{"age": age},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), bson.M{"username": username}, update)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("match count:", updateResult.MatchedCount)
}

func (d *Datastore) ListCollection() {
	db := d.Client.Database("userrepo")

	collections, err := db.ListCollections(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Println("error in findAll :", err)
		return
	}
	var results []bson.M
	if err = collections.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}

func main() {
	client := Connect()
	if client != nil {
		fmt.Println("mongodbPractice connected successfully!")
	}
	datastore := &Datastore{
		Client: client,
	}
	datastore.CreateDB("userrepo", "users")
	//datastore.Insert()
	//datastore.ListCollection()
	//client.Database("userrepo").Collection("users")
	//datastore.ProjectionFind()
	//datastore.FindOne()
	// datastore.FindAll()
	//datastore.DeleteOne()

}
