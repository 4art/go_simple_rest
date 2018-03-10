package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"./model"
	"./service"
)

var people []model.Person

func main() {
	people = append(people, model.Person{ID: 1, Firstname: "John", Lastname: "Smith", Address: &model.Address{City: "City X", State: "State X"}})
	people = append(people, model.Person{ID: 2, Firstname: "Maria", Lastname: "Black", Address: &model.Address{City: "City Z", State: "State Y"}})
	people = append(people, model.Person{ID: 3, Firstname: "Francis", Lastname: "Sunday"})
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	checkValid(err, w)
	var found bool = false
	for _, item := range people {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			found = true
			break
		}
	}

	if !found {
		service.ResponseNotFound(w, "id not found")
	}
}

func checkValid(err error, w http.ResponseWriter) {
	if err != nil {
		service.ResponseBadRequest(w)
	}
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person model.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = getNextFreeId(people)
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	checkValid(err, w)
	for index, item := range people {
		if item.ID == id {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

func getNextFreeId(p []model.Person) int {
	var id = 0
	for _, item := range p {
		if id <= item.ID {
			id = item.ID + 1
		}
	}
	return id
}
