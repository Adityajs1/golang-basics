package main

import "fmt"

func formatString(){
	const name = "Aditya Tiwari"
    const openRate = 30.5

    msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, openRate)
    fmt.Println(msg)
}