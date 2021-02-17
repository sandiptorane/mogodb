package crud

import (
	"fmt"
	"testing"
)

func TestDatastore_CreateDB(t *testing.T) {
	datastore := &Datastore{
		Client: Connect(),
	}
	datastore.CreateDB("userrepo", "users")
}

func TestDatastore_DropDB(t *testing.T) {
	datastore := &Datastore{
		Client: Connect(),
	}
	datastore.DropDB("UserRepo")
}

func TestDatastore_FindOne(t *testing.T) {
	client := Connect()
	if client != nil {
		fmt.Println("mongodb connected successfully!")
	}
	datastore := &Datastore{
		Client: client,
	}
	datastore.FindOne("sandip123")
}

func TestDatastore_ProjectionFind(t *testing.T) {
	client := Connect()
	if client != nil {
		fmt.Println("mongodb connected successfully!")
	}
	datastore := &Datastore{
		Client: client,
	}
	datastore.ProjectionFind("sandip123")
}

func TestDatastore_FindAll(t *testing.T) {
	client := Connect()
	if client != nil {
		fmt.Println("mongodb connected successfully!")
	}
	datastore := &Datastore{
		Client: client,
	}
	datastore.FindAll()
}

func TestDatastore_ListCollection(t *testing.T) {
	client := Connect()
	if client != nil {
		fmt.Println("mongodb connected successfully!")
	}
	datastore := &Datastore{
		Client: client,
	}
	datastore.ListCollection()
}

func TestDatastore_Update(t *testing.T) {
	client := Connect()
	if client != nil {
		fmt.Println("mongodb connected successfully!")
	}
	datastore := &Datastore{
		Client: client,
	}
	datastore.UpdateAge("sandip123", 22)
}

func TestDatastore_Insert(t *testing.T) {
	client := Connect()
	if client != nil {
		fmt.Println("mongodb connected successfully!")
	}
	datastore := &Datastore{
		Client: client,
	}
	person := &Person{
		Username: "dhiraj123",
		Name:     "Dhiraj",
		Lastname: "Gurav",
		Age:      22,
	}
	datastore.Insert(person)
}

func TestDatastore_DeleteOne(t *testing.T) {
	client := Connect()
	if client != nil {
		fmt.Println("mongodb connected successfully!")
	}
	datastore := &Datastore{
		Client: client,
	}
	datastore.DeleteOne("sandip123")
}
