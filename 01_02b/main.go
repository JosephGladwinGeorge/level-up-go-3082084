package main

import (
	"log"
	"strings"
	"time"
)

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	n:=strings.Split(msg, " ")
	for _,i := range(n){
		var t []string
		for p,k:= range(i){
			j:= strings.Repeat(string(k),p+1)
			t= append(t,j)
		}
		print(strings.Join(t,""))

	}

}

func main() {
	msg := "Time to learn about Go strings!"
	slowDown(msg)
}