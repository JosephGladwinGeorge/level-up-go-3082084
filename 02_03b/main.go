package main

import (
	"log"
	"fmt"
)

// the number of attendees we need to serve lunch to
const consumerCount = 300

// foodCourses represents the types of resources to pass to the consumers
var foodCourses = []string{
	"Caprese Salad",
	"Spaghetti Carbonara",
	"Vanilla Panna Cotta",
}

// takeLunch is the consumer function for the lunch simulation
// Change the signature of this function as required
func takeLunch(name string,courses []chan string,done chan<- struct{}) {
	for _,i:=range courses{
		log.Printf("%s ate %s",name,<-i)
	}
	done <- struct{}{}
}

// serveLunch is the producer function for the lunch simulation.
// Change the signature of this function as required
func serveLunch(course string,out chan<- string,done <-chan struct{}) {
	for {
		select{
		case out<-course:
		case <-done:
			return
		}
	}
}

func main() {
	log.Printf("Welcome to the conference lunch! Serving %d attendees.\n",
		consumerCount)
	
	var courses []chan string

	doneServing:=make(chan struct{})
	doneEating:= make(chan struct{})

	for _,i := range foodCourses{
		course:=make(chan string)
		courses = append(courses, course)
		go serveLunch(i,course,doneServing)
	}

	for i:=0;i< consumerCount;i++{
		name:= fmt.Sprintf("Guest %d",i)
		go takeLunch(name,courses,doneEating)
	}

	for i:=0;i<consumerCount;i++{
		<-doneEating
	}

	close(doneServing)
}
