package main

import "fmt"

func main() {

	a := Admin{User2{"user1", "user1@gmail.com"}, "admin"}
	sendNotification1(a)
	a.User2.notify1()

}

type User2 struct {
	Name  string
	Email string
}

type Admin struct {
	User2
	level string
}
type Notifier1 interface {
	notify1() error
}

func (u User2) notify1() error {
	fmt.Println(u.Email)
	fmt.Println(u.Name)
	return nil
}

func sendNotification1(n Notifier1) {
	n.notify1()
}
