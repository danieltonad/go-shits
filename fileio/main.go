package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func readFile(path string) {
	databyte, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fmt.Println("File contetnt: \n", string(databyte))
}

func writeToCsv(path string, data []User) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for _, user := range data {
		_, err := file.WriteString(fmt.Sprintf("%s,%d,%s\n", user.Name, user.Age, user.Address))
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Data written to file successfully")
}

func main() {
	// fmt.Println("File IO")
	// content := "Bull shit"

	// file, err := os.Create("file.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// io.WriteString(file, content)

	// os.WriteFile("file.txt", []byte(content), 0644)
	users := []User{}
	users = append(users, User{Name: "Damiel", Age: 25, Address: "40, lorem, ipsum"})
	users = append(users, User{Name: "Bleje", Age: 20, Address: "40, lorem, ipsum"})
	users = append(users, User{Name: "Bleje", Age: 20, Address: "40, lorem, ipsum"})
	users = append(users, User{Name: "Bleje", Age: 20, Address: "40, lorem, ipsum"})
	users = append(users, User{Name: "Bleje", Age: 20, Address: "40, lorem, ipsum"})
	users = append(users, User{Name: "Bleje", Age: 20, Address: "40, lorem, ipsum"})
	users = append(users, User{Name: "Bleje", Age: 20, Address: "40, lorem, ipsum"})

	writeToCsv("csv_test_file.csv", users)
}
