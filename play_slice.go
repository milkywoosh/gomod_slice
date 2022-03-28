package gomod_slice

import "fmt"

func Traverse_slice(data []string) {
	for _, value := range data {
		fmt.Println(value)
	}
}
