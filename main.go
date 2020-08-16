package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type task struct {
	ID int `json:ID`
	Name string `json:Name`
	Content string `json:content`
}


type allTasks []task

var tasks = allTasks {
	{
		ID: 1,
		Name: "Task One",
		Content: "Some Content",
	},
}


func indexRoute(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Welcome to my API change")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    taskID, err := strconv.Atoi(vars["id"])

	if (err != nil) {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for _, task := range tasks {
		if (task.ID == taskID) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(task)
		}
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
    taskID, err := strconv.Atoi(vars["id"])

	if (err != nil) {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for i, task := range tasks {
		if (task.ID == taskID) {
			tasks = append(tasks[:i], tasks[i + 1:]...)
			fmt.Fprintf(w, "The task %v has been remove succesfully", taskID)
		}
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
    var updateTask task
	vars := mux.Vars(r)
    taskID, err := strconv.Atoi(vars["id"])

	if (err != nil) {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)

    if (err != nil) {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	json.Unmarshal(reqBody, &updateTask)

	for i, task := range tasks {
		if (task.ID == taskID) {
			tasks = append(tasks[:i], tasks[i + 1:]...)  
			updateTask.ID = taskID
			tasks = append(tasks, updateTask)
			fmt.Fprintf(w, "The task %v has been updated succesfully", taskID)
		}
	}


}

func createTask(w http.ResponseWriter, r *http.Request) {
  var newTask task

  reqBody, err := ioutil.ReadAll(r.Body)

  if (err != nil) {
	  fmt.Fprintf(w, "Insert a Valid Task")
  } 
  
  json.Unmarshal(reqBody, &newTask)
  
  newTask.ID = len(tasks) + 1
  
  tasks = append(tasks, newTask)

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(newTask)
}

func main()  {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	log.Fatal(http.ListenAndServe(":4000", router)) 
}