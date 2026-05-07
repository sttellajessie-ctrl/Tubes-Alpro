package main

import "fmt"

const NMAX int = 100

type Menu struct {
	Name     string
	Category string
	Price    float64
}

func main() {
	//interface menu
	var n int
	fmt.Println("Welcome to our cafe's catalog")
	fmt.Println("1. Add menu")
	fmt.Scan(&n)
	if n == 1 {
		addMenu()
	}
}

func addMenu() {
	var menu Menu
	fmt.Print("Menu name:")
	fmt.Scanln(&menu.Name)
	fmt.Print("Menu category:")
	fmt.Scanln(&menu.Category)
	fmt.Print("Menu price:")
	fmt.Scanln(&menu.Price)
}
