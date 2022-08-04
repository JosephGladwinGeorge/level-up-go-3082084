package main

import "log"

// the number of Gophers we need to serve lunch to
const consumerCount = 300

// foodCourses represents the types of resources to pass to the consumer goroutines.
var foodCourses = []string{
	"Caprese Salad",
	"Spaghetti Carbonara",
	"Vanilla Panna Cotta",
}

// takeLunch is the consumer function for the lunch simulation
// Change the signature of this function as required
func takeLunch(name string) {
	panic("NOT IMPLEMENTED YET")
}

// serveLunch is the producer function for the lunch simulation.
// Change the signature of this function as required
func serveLunch(course string) {
	panic("NOT IMPLEMENTED YET")
}

func main() {
	log.Printf("Welcome to the GopherCon lunch! Serving %d gophers.\n", consumerCount)
}
