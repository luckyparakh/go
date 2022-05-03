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

//Model for course

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

//Helper -- file
func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("API Demo")

	//Seeding
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "React JS",
		CoursePrice: 299,
		Author: &Author{
			Fullname: "Rishi",
			Website:  "youtube.com",
		},
	})
	courses = append(courses, Course{
		CourseId:    "4",
		CourseName:  "Node JS",
		CoursePrice: 199,
		Author: &Author{
			Fullname: "Sam",
			Website:  "udemy.com",
		},
	})

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4000", r))
}

//Controller - file
//Serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Learning API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get a course")
	w.Header().Set("content-type", "application/json")

	//grab id from req
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found")
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a course")
	w.Header().Set("content-type", "application/json")
	if r.Body == nil {
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode("Please provide some data.")
		return
	}
	var ipCourse Course
	json.NewDecoder(r.Body).Decode(&ipCourse)
	if ipCourse.isEmpty() {
		json.NewEncoder(w).Encode("Please provide some data.")
		return
	}
	for _, course := range courses {
		if course.CourseName == ipCourse.CourseName {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Duplicate Course")
			return
		}
	}
	rand.Seed(time.Now().UnixNano())
	ipCourse.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, ipCourse)
	json.NewEncoder(w).Encode(ipCourse)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a course")
	w.Header().Set("content-type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please provide some data.")
		return
	}
	params := mux.Vars(r)
	var dummyCourse Course
	flag := false
	for index, course := range courses {
		if course.CourseId == params["id"] {
			flag = true
			courses = append(courses[:index], courses[index+1:]...)
			err := json.NewDecoder(r.Body).Decode(&dummyCourse)
			if err != nil {
				panic(err)
			}
			dummyCourse.CourseId = params["id"]
			courses = append(courses, dummyCourse)
			json.NewEncoder(w).Encode(dummyCourse)
			return
		}
	}
	if !flag {
		json.NewEncoder(w).Encode("ID not found")
	}
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a course")
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	flag := false
	for index, course := range courses {
		if course.CourseId == params["id"] {
			flag = true
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted.")
			return
		}
	}
	if !flag {
		json.NewEncoder(w).Encode("ID not found")
	}
}
