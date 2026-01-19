package main

import "fmt"

func conditional(){
	 messagesLen := 10
	 maxMessagesLen := 20

	 fmt.Println("Try to send a message of",messagesLen, "and a max length of", maxMessagesLen )

	 if messagesLen <= maxMessagesLen{
		fmt.Println("Messages sent!")
	 }else{
		fmt.Println("Messages not sent!")
	 }


}