package main

import (
	"fmt"
)

func sendEmail(name string, ch <-chan string) {
	message := <-ch
	fmt.Printf("Sending email to %s: %s\n", name, message)
}

func main() {

	users := []string{"Alice", "Bob", "Charlie"}
	ch := make(chan string)
	for _, user := range users {
		go sendEmail(user, ch)
		ch <- fmt.Sprintf("Hello %s, this is a notification email!", user)
	}
	
	fmt.Printf("Main function completed")
}
