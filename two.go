package main

import "fmt"

func main() {

	u := User{"first", "last"}
	// sendNotification(u)
	sendNotification(&u)
}

type Notifier interface {
	notify() error
}

func sendNotification(n Notifier) {
	n.notify()
}

type User struct {
	firstName string
	lastName  string
}

func (u *User) notify() error {
	fmt.Println(u.firstName)
	fmt.Println(u.lastName)
	return nil
}
