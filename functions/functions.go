package main

import (
	"fmt"
	"math/rand"
)

func returnRandom(min, max int)(int){
	return rand.Intn(max - min) + min
}


func main(){
	fmt.Println(returnRandom(2, 15))
}