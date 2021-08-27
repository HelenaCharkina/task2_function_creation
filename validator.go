package main

import (
	"log"
	"os"
	"strings"
)

func IsUnknownFieldError(err error) bool {
	if strings.Contains(err.Error(), "unknown field") {
		return true
	}
	return false
}

func CheckLoop(data *[]InputData) bool {
	for i, item := range *data {
		for _, vertex := range item.InputFrom {
			if vertex == uint(i) {
				return false
			}
		}
	}
	return true
}
func CheckCycle(data *[]InputData) {

	for _, item := range *data {
		vertices[item.Name] = 0
	}

	for i := 0; i < len(*data); i++ {
		dfs(data, i)
	}
}
func dfs(arr *[]InputData, i int) {
	vertices[(*arr)[i].Name] = 1

	for _, j := range (*arr)[i].InputFrom {
		if vertices[(*arr)[j].Name] == 0 {
			dfs(arr, int(j))
		} else if vertices[(*arr)[j].Name] == 1 {
			log.Println("Incorrect graph: found cycle!")
			os.Exit(4)
		}
	}
	vertices[(*arr)[i].Name] = 2
}
