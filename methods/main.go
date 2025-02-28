package main

import "fmt"

type User struct {
	Name    string
	Email   string
	Address string
}

func (u User) GetFullName() string {
	return u.Name
}

func main() {
	user := User{Name: "John Doe", Email: "", Address: "Nairobi"}
	fmt.Print(user.GetFullName())
}
