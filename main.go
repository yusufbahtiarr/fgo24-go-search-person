package main

import (
	"fmt"
	"strings"
)

func searchPerson(users []string, search string) {
	result:=""
	for _, person := range users {
		if strings.Contains(strings.ToLower(person), strings.ToLower(search)) {
			result = person
		}
	}
	if result != ""{
	fmt.Println(result)
	}else{
		fmt.Println("Data tidak ditemukan")
	}
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
	
	// result := searchPerson(users, "clement")
	// fmt.Println(result)
	searchPerson(users, "clement")
}
