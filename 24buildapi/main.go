package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//model for course - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"-"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake Db

var courses []Course

// middleware

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Api")

	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "101", CourseName: "DBMS", CoursePrice: 100, Author: &Author{Fullname: "Khubaib Khalil", Website: "www.google.com"}})

	courses = append(courses, Course{CourseId: "102", CourseName: "MERN", CoursePrice: 100, Author: &Author{Fullname: "Mani", Website: "www.mern.com"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers

//serveHome Route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get All Courses")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	json.NewEncoder(w).Encode("No Course Found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one Course")
	w.Header().Set("Content-Type", "application/json")
	// what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Send Some data")
		return
	}
	// data = {}
	var c Course
	_ = json.NewDecoder(r.Body).Decode(&c)
	if c.IsEmpty() {
		json.NewEncoder(w).Encode("Please Send Some data")
		return
	}

	for _, course := range courses {
		if course.CourseName == c.CourseName {
			json.NewEncoder(w).Encode("Course already exists")
			return
		}
	}
	// check duplicate title
	//generate unique id , string
	//append new course
	seed := rand.New(rand.NewSource(time.Now().Unix()))

	c.CourseId = strconv.Itoa(seed.Intn(100))

	courses = append(courses, c)
	json.NewEncoder(w).Encode(c)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one Course")
	w.Header().Set("Content-Type", "application/json")

	//grab id
	params := mux.Vars(r)
	//loop , id , remove , add with my id

	for index, c := range courses {
		if c.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)

			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one Course")
	w.Header().Set("Content-Type", "application/json")

	//grab id
	params := mux.Vars(r)
	//loop , id , remove , add with my id

	for index, c := range courses {
		if c.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(c)
			break
		}
	}
}
