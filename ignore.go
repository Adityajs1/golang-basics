package main 
import "fmt"

func ignore(){
   first, last := getNames()
   fmt.Println("Welcome to textio ",first,last)
}

func getNames()(string, string){
	return "Aditya", "Barney"
}