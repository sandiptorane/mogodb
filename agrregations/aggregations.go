package main

import "mongodb/crud"

func main() {
	client :=crud.Connect()
	datastore := &crud.Datastore{
		Client: client,
	}
	datastore.GroupByAge()
	datastore.Match()
}
