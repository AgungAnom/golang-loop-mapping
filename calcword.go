package main

import (
	"fmt"
)


func main() {
var mapping map[string]int
mapping = map[string]int{}
var str string = "selamat malam"

for _,value := range str{
	fmt.Printf("%c \n",value)
	mapping[string(value)] += 1
}

fmt.Println(mapping)

}