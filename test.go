package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phonenumber"`
	IsActive    bool      `json:"isactive"`
}

func createJsonFile(validUsers []User) {
	// Convert to JSON
	json_data, err := json.Marshal(validUsers)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//print json data
	//fmt.Println(string(json_data))

	//create json file
	json_file, err := os.Create("user_data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer json_file.Close()

	json_file.Write(json_data)
	json_file.Close()

}

func readCSV() []User {
	csv_file, err := os.Open("user_data.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csv_file.Close()

	r := csv.NewReader(csv_file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var user User
	var users []User
	type ID uuid.UUID
	for _, rec := range records {
		user.ID, _ = uuid.FromString(rec[0])
		user.Name = rec[1]
		user.Email = rec[2]
		user.PhoneNumber = rec[3]
		user.IsActive, _ = strconv.ParseBool(rec[4])
		if check_validity(user.Name, user.PhoneNumber, user.Email) {
			users = append(users, user)
		} else {
			continue
		}
	}
	//For checking duplicate ID
	visited := make(map[ID]bool)
	var validUsers []User
	for _, values := range users {
		if !visited[ID(values.ID)] {
			visited[ID(values.ID)] = true

		} else {
			id := uuid.NewV4()
			log.Warn("Duplicate User ID found..Assigning new User ID")
			values.ID = id
		}
		validUsers = append(validUsers, values)
	}
	return validUsers
}

func main() {
	var validUsers []User = readCSV()
	createJsonFile(validUsers)

}
