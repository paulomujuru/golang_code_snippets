package main

import "fmt"

type Set map[string]struct{}

func main() {
	s := make(Set)
	//add items
	s["Item1"] = struct{}{}
	s["Item2"] = struct{}{}
	//get and print items
	fmt.Println(getSetValues(s))
}

func getSetValues(s Set) []string {
	var retVal []string
	for k, _ := range s {
		retVal = append(retVal, k)
	}
	return retVal
}

//  the way channels work in golang is that when they are two threads to communicate with each other thread one will wait till  the other thread recieves the instruction then they will excute the instruction
