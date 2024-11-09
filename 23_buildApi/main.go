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

// Model for Course -file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper -file
func (c *Course) IsEmpty() bool {
	//return c.CourseName == "" && c.CourseId == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to the build API in golang!!!")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "Let's go with golang", CoursePrice: 190, Author: &Author{Fullname: "John Doe", Website: "johndoe.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Let's go with c#", CoursePrice: 200, Author: &Author{Fullname: "Matthew Jeng", Website: "matthewj.com"}})

	//routes
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/api/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/api/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/api/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/api/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/api/courses/{id}", deleteOneCourse).Methods("DELETE")

	//listening to a PORT
	log.Fatal(http.ListenAndServe(":8000", r))

}

// controllers -file
// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Golang</h1>"))
}

// get
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)
}

// get
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")

	w.Header().Set("Content-Type", "application/json")

	// get courseID from url
	params := mux.Vars(r)
	fmt.Println("request params", params)
	fmt.Printf("Type of params is: %T \n", params)
	//loop through courses, find matching id and return response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("course not found with the id: " + params["id"])
	return
}

// create
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create course")

	w.Header().Set("Content-Type", "application/json")

	//what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	//what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course) // decode
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	//generate unique id & convert it to string

	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}

// update
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")

	w.Header().Set("Content-Type", "application/json")

	//get courseID from url
	params := mux.Vars(r)
	fmt.Println("request params", params)
	fmt.Printf("Type of params is: %T \n", params)

	//loop through the courses , find matching id, remove it & then add the data with Id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course) // decode the data from the request body
			if course.IsEmpty() {
				json.NewEncoder(w).Encode("Please send some data")
				return
			}
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)

			fmt.Println("course", course)
			return
		}
	}

	json.NewEncoder(w).Encode("course not found with the id: " + params["id"])

}

// delete
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")

	w.Header().Set("Content-Type", "application/json")

	//get courseID from url
	params := mux.Vars(r)
	fmt.Println("request params", params)
	fmt.Printf("Type of params is: %T \n", params)

	//loop through the courses , find matching id, remove it
	for index, course := range courses {
		if course.CourseId == params["id"] {

			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("course deleted successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("course not found with the id: " + params["id"])
}
