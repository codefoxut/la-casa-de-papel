package main

import (
	"fmt"
	"sort"
)


func numOfRescueBoats(people []int, weightLimit int) int {
	sort.Ints(people)
	fmt.Println(people)
	var numOfBoats, leftTracker, rightTracker int
	rightTracker = len(people) - 1

	for leftTracker <= rightTracker {
		if leftTracker == rightTracker {
			numOfBoats += 1
			fmt.Printf("Boat; %d has %d only\n", numOfBoats, people[rightTracker])
			break
		}
		if people[leftTracker] + people[rightTracker] <= weightLimit {
			numOfBoats  += 1
			leftTracker += 1
			rightTracker -= 1
			fmt.Printf("Boat; %d has %d and %d\n", numOfBoats, people[leftTracker-1], people[rightTracker+1])
		} else {
			numOfBoats += 1
			rightTracker -= 1
			fmt.Printf("Boat; %d has %d only\n", numOfBoats, people[rightTracker+1])
		}
	}
	
	return numOfBoats
}

func main(){
	peoples := []int{10, 5, 13, 6, 9, 3, 5, 3, 8, 7, 8, 7, 11, 15, 9, 16, 10, 14, 4, 12, 10, }
	limit := 17
	fmt.Printf("numOfRescueBoats %d\n", numOfRescueBoats(peoples, limit))

}