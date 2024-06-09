package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// model for Course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Name    string `json:"authorname"`
	Website string `json:"authorwebsite"`
}

// middleware/helper to check if courseid and coursename are not empty
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

// fake db
var courses []Course

func main() {
	fmt.Println("API")
	r := mux.NewRouter()

	//seeding (feeding sample db data for testing)
	courses = append(courses, Course{CourseId: "4", CourseName: "Golang", CoursePrice: 599, Author: &Author{Name: "Amit Mishra", Website: "go.dev"}})
	courses = append(courses, Course{CourseId: "1", CourseName: "Python", CoursePrice: 399, Author: &Author{Name: "Christopher Nolan", Website: "dream@inception.com"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hey, this is an API by amitthisside</h1>"))
}

// get courses controller
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// get one course
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One course")
	w.Header().Set("Content-Type", "application/json")

	//get parameters from request
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	//if course not found
	json.NewEncoder(w).Encode("Could not find a course with the given Course ID")
}

// add one course controller
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add One Course")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	//if passed json is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	//if json is {}
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	course.CourseId = strconv.Itoa(rand.Intn(1000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

// update one course
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One Course")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	//grab course id from request
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	//id not found
	json.NewEncoder(w).Encode("Course ID not found")
}

// delete one course
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One Course")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(&course)
			courses = append(courses[:index], courses[index+1:]...)
			return
		}
	}
	json.NewEncoder(w).Encode("Couldn't find data with the given Course ID")
}
