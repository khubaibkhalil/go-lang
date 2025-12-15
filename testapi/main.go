package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string `json:"id"`
	CourseName  string `json:"cname"`
	CoursePrice int    `json:"-"`

	Author *Author `json:"author"`
}

type Author struct {
	FullName string
	Website  string
}

// var courses Course = Course{CourseId: "1001", CourseName: "DBMS", CoursePrice: 1, Author: &Author{FullName: "KK", Website: "facebook.com"}}

var courses []Course = []Course{
	Course{CourseId: "1001", CourseName: "DBMS", CoursePrice: 1, Author: &Author{FullName: "KK", Website: "facebook.com"}},
	Course{CourseId: "1001", CourseName: "DBMS", CoursePrice: 1, Author: &Author{FullName: "KK", Website: "facebook.com"}}}

func main() {
	// fmt.Printf("%+v", courses)
	// fmt.Print(*courses[0].Author)

	r := mux.NewRouter()

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getALlCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourseByid).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))

}

// func serveHome() {

// }

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello world</>"))
}

func getALlCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourseByid(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, id := range courses {
		if id.CourseId == params["id"] {
			json.NewEncoder(w).Encode(courses[i])
		}
	}

	return
}
