package main

import (
	"fmt"
	"strings"
)

func searchPerson(users []string, search string) string {
	for _, person := range users {
		if strings.Contains(strings.ToLower(person), strings.ToLower(search)) {
			return person
		}
	}
	return "Data tidak ditemukan"
}

func main(){
	users := []string{
		"Lenane Graham",
		"Ervin Howel",
		"Clementine Bauch",
		"Patricia Lebstack",
		"Mrs. Dennis Schulist",
		"Kurtis Weissnat",
		"Nicholas Runolfsdottir",
		"Glena Reichert",
		"Clementina DuBuque",
	}
	
	result := searchPerson(users, "clement")
	fmt.Println(result) // Output: Bob
}
