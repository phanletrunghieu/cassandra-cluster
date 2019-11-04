package main

import (
	"cassandra-go/config"
	"cassandra-go/models"
	"encoding/json"

	"github.com/gocql/gocql"
	"github.com/k0kubun/pp"
)

func main() {
	var err error
	var query string

	session := config.Session
	defer session.Close()

	// create user
	user := models.User{
		FirstName: "Hiếu",
		LastName:  "Phan",
		Age:       22,
		City:      "Los Angeles",
		Email:     "hieutrunglephan@gmail.com",
	}

	// create
	gocqlUuid := gocql.TimeUUID()
	query = `INSERT INTO users (id, firstname, lastname, email, city, age) VALUES (?, ?, ?, ?, ?, ?)`
	err = session.Query(
		query,
		gocqlUuid,
		user.FirstName,
		user.LastName,
		user.Email,
		user.City,
		user.Age).Exec()
	if err != nil {
		panic(err)
	}
	user.ID = gocqlUuid

	pp.Println(user)

	// select
	query = "SELECT * FROM users"
	iterable := session.Query(query).Iter()
	userList := []models.User{}
	m := map[string]interface{}{}
	for iterable.MapScan(m) {
		user := models.User{}
		data, _ := json.Marshal(m)
		json.Unmarshal(data, &user)
		userList = append(userList, user)
		m = map[string]interface{}{}
	}
	pp.Println(userList)
}
