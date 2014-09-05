package main

type user struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}
