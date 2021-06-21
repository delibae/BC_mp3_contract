package main

import (
    "encoding/json"
    // "fmt"
    // "io/ioutil"
    "os"
	"J2D.com/BC"
	"log"
	"fmt"
	
)

func main() {
	
	var config []BC.Block
	file, err := os.Open("./BCM.json")
	defer file.Close()
	if err != nil { 
		log.Fatal(err) 
	} 
	decoder := json.NewDecoder(file) 
	err = decoder.Decode(&config) 
	if err != nil { 
		log.Fatal(err) 
	} 
	fmt.Println(config)
}