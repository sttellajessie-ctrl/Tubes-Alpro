// admin 
package main 
import "fmt"

const NMAX int = 100

type Menu struct {
	Name     	string
	Category 	string
	Price    	int
	composition string
	status  	bool
}
type menuList [NMAX]Menu

func main () {
	var n int 
	//interface menu  
	fmt.Println("Who are using this cafe's catalog?")
	var user string
	fmt.Scanln(&user)
	// only admin and customer can access the menu
	if user == "admin" {
		adminMenu()
	} else if user == "customer" {
		// customerMenu()
	} else {
		fmt.Println("Invalid user")
	}
	
}

func displayMenu() {
	fmt.Println("What would you like to do?")
	fmt.Println("1. Add menu")
	fmt.Println("2. Delete menu")
	fmt.Println("3. Edit menu")
	fmt.Println("4. Show menu")
	fmt.Println("5. Exit")

	fmt.Print("Enter your choice: ")
} 

func adminMenu() {
	var list menuList
	var choose int
	fmt.Println("Welcome to the admin menu!")
	displayMenu()
	fmt.Scan(&choose)

	for choose != 5 {
		if choose == 1 {
			addMenu(&list)
		} else if choose == 2 {
			// delete menu function
		} else if choose == 3 {
			// edit menu function
		} else if choose == 4 {
			// show menu func
			var priceOrder string
			fmt.Println("please select a price order (asc/desc):")
			fmt.Scan(&priceOrder)

			if priceOrder == "asc" {
			// sort menu by price ascending

			} else if priceOrder == "desc" {
			// sort menu by price descending
				decending(list, &n)
			}

		} else if choose == 5 {
			fmt.Println("Exiting admin menu...")
		} else {
			fmt.Println("Invalid option")
		}
	}
}


func addMenu(list *menuList, n *int) {
	fmt.Println("How many menu do you want to add?")
	fmt.Scan(&n)
	for i := 1; i <= *n; i++ {
		fmt.Print("Menu name:")
		fmt.Scan(&list[i].Name)

		fmt.Print("Menu category:")
		fmt.Scan(&list[i].Category)

		fmt.Print("Menu price:")
		fmt.Scan(&list[i].Price)

		fmt.Print("Menu composition:")
		fmt.Scan(&list[i].composition)

		fmt.Print("Menu status (true/false):")
		fmt.Scan(&list[i].status)
	}
}

func decending(list menuList, n *int) {
	// sort menu by price descending
	var temp menuList
	var i, j int
	for i = 1; i < *n; i++ {
		temp = list[i]
		j = i - 1
		for j >= 0 && list[j].Price < temp.Price {
			list[j+1] = list[j]
			j = j - 1
		}
		list[j+1] = temp
	}
}