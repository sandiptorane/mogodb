package crud

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func (d *Datastore)GroupByAge() {
	collection := d.Client.Database("userrepo").Collection("users")
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$age"},
			{"numTimes", bson.D{
				{"$sum", 1},
			}},
		}},
	}
	cursor, err := collection.Aggregate(context.TODO(), mongo.Pipeline{groupStage})
	if err != nil {
		log.Fatal(err)
	}

	//list all
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Printf("age %v appears %v times\n", result["_id"], result["numTimes"])
	}


}

func (d *Datastore)Match(){
	log.Println("inside match")
	collection := d.Client.Database("userrepo").Collection("users")
	matchstage:=bson.D{
		{"$match", bson.D{
			{"age", bson.D{{"$gte",21}}},
		}},
	}
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$age"},
			{"total", bson.D{
				{"$sum", "$age"},
			}},
		}},
	}
	cursor, err := collection.Aggregate(context.TODO(), mongo.Pipeline{matchstage,groupStage})
	if err != nil {
		log.Fatal(err)
	}

	//list all
	var results []bson.M
	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Printf("sum of age %v : %v \n", result["_id"], result["total"])
	}
	log.Println("returning from match")
}
