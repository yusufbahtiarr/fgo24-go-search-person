package utils

import (
	"fmt"
	"strings"
)

func SearchPerson(users []string, search string) {
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
