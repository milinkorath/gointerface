package main

import "fmt"

func main() {

	u := User1{"Ram", "ram@gmail.com"}
	// u1 := User1{"Raju", "ranju@gmail.com"}
	add_u := &u
	// add_u1 := &u1
	u.Notify()
	// u1.Notify()
	add_u.Notify()
	// add_u1.Notify()
	u.AddNotify()
	// u1.AddNotify()
	add_u.AddNotify()

}

type User1 struct {
	Name  string
	Email string
}

func (u User1) Notify() error {
	fmt.Println(u.Name)
	fmt.Println(u.Email)
	return nil
}
func (u *User1) AddNotify() error {
	fmt.Println(u.Name)
	fmt.Println(u.Email)
	return nil
}
