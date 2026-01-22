package main

import "fmt"

// 1️⃣ Interface
type Sender interface {
	Send() string
}

// 2️⃣ Base struct
type User struct {
	name string
}

// 3️⃣ Nested struct
type Message struct {
	from User
	text string
}

// 4️⃣ Method implementing interface
func (m Message) Send() string {
	return m.from.name + ": " + m.text
}

// 5️⃣ Function using interface
func teste(s Sender) {
	fmt.Println(s.Send())
}

// 6️⃣ main function
func inter() {
	teste(Message{
		from: User{name: "Aditya"},
		text: "Learning Go interfaces!",
	})
}
