package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	fileName := os.Args[1]

	var inputData []InputData

	inputFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	var decoder *json.Decoder
	decoder = json.NewDecoder(inputFile)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&inputData)
	if err != nil {
		switch err.(type) {
		case *json.SyntaxError:
			log.Println("Input is not JSON")
			os.Exit(1)
		case *json.UnmarshalTypeError:
			log.Println("Incorrect input data format")
			os.Exit(2)
		default:
			if IsUnknownFieldError(err) {
				log.Println("Incorrect input data format")
				os.Exit(2)
			} else {
				log.Fatal("Необработанная ошибка парсинга json файла: ", err)
			}
		}

	}
	vertices = make(map[string]int)
	if !CheckLoop(&inputData) {
		log.Println("Incorrect graph: found loop!")
		os.Exit(3)
	}
	CheckCycle(&inputData)

	inputs := make(map[uint]int)
	for _, item := range inputData {
		for _, i := range item.InputFrom {
			inputs[i] = 1
		}
	}
	var functions []string
	for i := 0; i < len(inputData); i++ {
		if _, ok := inputs[uint(i)]; !ok {
			functions = append(functions, getFunction(&inputData, uint(i)))
		}
	}
	result:= strings.Join(functions, ",")

	err = ioutil.WriteFile("output.txt", []byte(result), 0644)
	if err != nil {
		log.Fatal("ioutil.WriteFile error : ", err)
	}
	return
}

func getFunction(data *[]InputData, i uint) string {
	 res := (*data)[i].Name
	if len((*data)[i].InputFrom) > 0 {
		res += `(`
		var tmp []string
		for _, i := range (*data)[i].InputFrom {
			tmp = append(tmp, getFunction(data, i))
		}
		res += strings.Join(tmp, ",")
		res += ")"
	}

	return res
}