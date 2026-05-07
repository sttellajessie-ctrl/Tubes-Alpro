package main

import "fmt"

const NMAX int = 100

type Menu struct {
	Name     string
	Category string
	Price    float64
}
type menuList [NMAX]Menu

func main() {
	//interface menu
	fmt.Println("Who are using this cafe's catalog?")
	var user string
	fmt.Scanln(&user)

	// only admin and customer can access the menu
	if user == "admin" {
		adminMenu()
	} else if user == "customer" {
		customerMenu()
	} else {
		fmt.Println("Invalid user")
	}
}

func adminMenu() {
	var list menuList
	var n int
	fmt.Println("What would you like to do?")
	fmt.Println("1. Add menu")

	fmt.Scan(&n)
	if n == 1 {
		addMenu(&list)
	}
}

func addMenu(list *menuList) {
	var n int
	fmt.Println("How many menu do you want to add?")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Print("Menu name:")
		fmt.Scan(&list[i].Name)

		fmt.Print("Menu category:")
		fmt.Scan(&list[i].Category)

		fmt.Print("Menu price:")
		fmt.Scan(&list[i].Price)
	}
}
