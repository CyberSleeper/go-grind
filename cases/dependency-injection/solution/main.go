package main

import "fmt"

type DataStore interface {
	Query(id string) string
}

// Imagine this is a heavy SQL driver
type RealPostgresDB struct{}

func (db RealPostgresDB) Query(id string) string {
	return "Real User from DB: " + id
}

type MockDB struct{}

func (db MockDB) Query(id string) string {
	return "Fake User: " + id
}

type Service struct {
	db DataStore
}

func (s Service) GetUser(id string) string {
	return s.db.Query(id)
}

func main() {
	db := MockDB{}
	svc := Service{db: db}

	fmt.Println(svc.GetUser("123"))
}
