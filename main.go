package main

import "fmt"

const NMAX int = 100

type Menu struct {
	Name            string
	Category        string
	Price           int
	Ingredients     [10]string
	IngredientCount int
	Status          bool
	Stock           int
}
type menuList [NMAX]Menu

type Order struct {
	Name     string
	Price    int
	Quantity int
}
type orderList [NMAX]Order

func main() {
	var catalog menuList
	var menuCount int
	var exit bool
	var user string
	exit = false
	menuCount = 0

	for !exit {
		fmt.Println("Welcome to our cafe!")
		fmt.Println("login as (admin/customer/exit):")
		fmt.Scan(&user)
		// only admin and customer can access the menu
		if user == "admin" {
			fmt.Println(`                
     /\_/\                                                /\_/\
    ( o.o )                                              ( o.o )
  _▄▄█████▄▄____________________________________________▄▄█████▄▄_
            ███╗   ███╗███████╗███╗   ██╗██╗   ██╗
            ████╗ ████║██╔════╝████╗  ██║██║   ██║
            ██╔████╔██║█████╗  ██╔██╗ ██║██║   ██║
            ██║╚██╔╝██║██╔══╝  ██║╚██╗██║██║   ██║
            ██║ ╚═╝ ██║███████╗██║ ╚████║╚██████╔╝
            ╚═╝     ╚═╝╚══════╝╚═╝  ╚═══╝ ╚═════╝
 ═══════════════════════  (=^·^=)  ═════════════════════════════					
			`)
			adminMenu(&catalog, &menuCount)
		} else if user == "customer" {
			fmt.Println(`                
     /\_/\                                                            /\_/\
    ( o.o )                                                          ( o.o )
  _▄▄█████▄▄________________________________________________________▄▄█████▄▄_
         ██╗    ██╗███████╗██╗     ██████╗ ██████╗ ███╗   ███╗███████╗
         ██║    ██║██╔════╝██║    ██╔════╝██╔═══██╗████╗ ████║██╔════╝
         ██║ █╗ ██║█████╗  ██║    ██║     ██║   ██║██╔████╔██║█████╗  
         ██║███╗██║██╔══╝  ██║    ██║     ██║   ██║██║╚██╔╝██║██╔══╝  
        ╚███╔███╔╝███████╗███████╗╚██████╗╚██████╔╝██║ ╚═╝ ██║███████╗
         ╚══╝╚══╝ ╚══════╝╚══════╝ ╚═════╝ ╚═════╝ ╚═╝     ╚═╝╚══════╝
 ════════════════════════════  (=^·^=)  ═══════════════════════════════════						
			`)
			customerMenu(&catalog, menuCount)
		} else if user == "exit" {
			exit = true
			fmt.Println("Goodbye!", `\ (^_^)`)
		} else {
			fmt.Println("Invalid user. Please type exactly 'admin', 'customer', or 'exit'.")
		}
	}
}

func adminMenu(catalog *menuList, menuCount *int) {
	var exit bool = false
	for !exit {
		var choose int

		fmt.Println("\n--- Admin Menu ---")

		fmt.Println("1. Add menu")
		fmt.Println("2. Delete menu")
		fmt.Println("3. Edit menu")
		fmt.Println("4. Show menu")
		fmt.Println("5. Show statistics")
		fmt.Println("6. Exit")
		fmt.Scan(&choose)
		if choose == 1 {
			addMenu(&*catalog, &*menuCount)
		} else if choose == 2 {
			// delete menu function
			deleteMenu(&*catalog, &*menuCount)
		} else if choose == 3 {
			// edit menu function
			showMenu(*catalog, *menuCount)
			editMenu(&*catalog, &*menuCount)
		} else if choose == 4 {
			// show menu function
			showMenu(*catalog, *menuCount)
		} else if choose == 5 {
			statisticPerCatergory(*catalog, *menuCount)
		} else if choose == 6 {
			exit = true
			fmt.Println("Goodbye!", `\ (^_^)`)
		} else {
			fmt.Println("Invalid option")
		}
	}
}

func addMenu(catalog *menuList, menuCount *int) {
	var n, i, j int
	fmt.Println("How many menu do you want to add?")
	fmt.Scan(&n)
	for i = 0; i < n; i++ {

		if *menuCount >= NMAX {
			fmt.Println("Menu list is full. Cannot add more menu.")
			return
		}

		var newMenu Menu

		fmt.Print("Menu name (use _ for spaces): ")
		fmt.Scan(&newMenu.Name)

		fmt.Print("Menu category (Food/Drink): ")
		fmt.Scan(&newMenu.Category)

		fmt.Print("Menu price: ")
		fmt.Scan(&newMenu.Price)

		fmt.Print("How many ingredients are in the composition? (Max 10): ")
		fmt.Scan(&newMenu.IngredientCount)
		if newMenu.IngredientCount > 10 {
			fmt.Println("Limited to 10 ingredients.")
			newMenu.IngredientCount = 10
		}
		if newMenu.IngredientCount > 0 {
			fmt.Println("Enter each ingredient: ")
			for j = 0; j < newMenu.IngredientCount; j++ {
				fmt.Printf("Ingredient %d (use _ for spaces): ", j+1)
				fmt.Scan(&newMenu.Ingredients[j])
			}
		}
		newMenu.Status = false
		fmt.Print("Menu stock: ")
		fmt.Scan(&newMenu.Stock)
		if newMenu.Stock > 0 {
			newMenu.Status = true
		}
		(*catalog)[*menuCount] = newMenu
		*menuCount++
	}
	fmt.Println("Menus added successfully!")
}

func deleteMenu(catalog *menuList, menuCount *int) {
	var idx, i int
	if *menuCount == 0 {
		fmt.Println("The menu was empty, nothing to delete")
		return
	}
	showMenu(*catalog, *menuCount)
	fmt.Print("Enter the name of menu to delete: ")
	var name string
	fmt.Scan(&name)
	idx = SearchMenu(*catalog, *menuCount, name)
	if idx == -1 {
		fmt.Println("Menu not found")
		return
	}
	for i = idx; i < *menuCount-1; i++ {
		(*catalog)[i] = (*catalog)[i+1]
	}
	*menuCount = *menuCount - 1
	fmt.Printf("Menu '%s' deleted succesfully\n", name)
}

func showMenu(catalog menuList, menuCount int) {
	if menuCount == 0 {
		fmt.Println("The menu is currently empty.")
		return
	}
	fmt.Println("\n--- Current Menu Catalog ---")
	var i int
	for i = 0; i < menuCount; i++ {
		var menu Menu = catalog[i]
		fmt.Printf("%d. [%s] %s - Rp%d (Available: %t)\n", i+1, menu.Category, menu.Name, menu.Price, menu.Status)

		fmt.Print("    Composition: ")
		if menu.IngredientCount == 0 {
			fmt.Print(" None")
		} else {
			var j int
			for j = 0; j < menu.IngredientCount; j++ {
				fmt.Print(menu.Ingredients[j])
				if j < menu.IngredientCount-1 {
					fmt.Print(", ")
				}
			}
		}
		fmt.Println()
	}
}

func editMenu(catalog *menuList, menuCount *int) {
	var price, i, choose, stock int
	var enter bool
	var MenuName string
	var found bool
	enter = false
	for !enter {
		fmt.Println("Choose Menu to edit:")
		fmt.Scan(&MenuName)
		found = false
		for i = 0; i < *menuCount && !found; i++ {
			if MenuName == (*catalog)[i].Name {
				found = true
				fmt.Println("What do you want to edit?")
				fmt.Println("1. Price")
				fmt.Println("2. Stock")
				fmt.Scan(&choose)
				if choose == 1 {
					fmt.Println("Enter new price:")
					fmt.Scan(&price)
					(*catalog)[i].Price = price
					fmt.Println("The price of menu", (*catalog)[i].Name, "has been updated to Rp", price)
				} else if choose == 2 {
					fmt.Println("Enter new stock:")
					fmt.Scan(&stock)
					(*catalog)[i].Stock = stock
					if (*catalog)[i].Stock > 0 {
						(*catalog)[i].Status = true
					} else {
						(*catalog)[i].Status = false
					}
					fmt.Println("The stock of menu", (*catalog)[i].Name, "has been updated to", stock)
				} else {
					fmt.Println("Invalid option. No changes made.")
				}
			}
		}
		if !found {
			fmt.Println("Menu not found. Please enter a valid menu name.")
		} else {
			enter = true
		}
	}
}

// customer menu
func customerMenu(catalog *menuList, menuCount int) {
	fmt.Println("welcome to our cafe!")
	if menuCount == 0 {
		fmt.Println("Sorry, the menu is currently empty, please come back later.")
		return
	}
	fmt.Println("Here is our menu:")
	fmt.Println("categories: Drink / Food ")
	fmt.Print("Please choose a category: ")
	var category string
	fmt.Scan(&category)
	fmt.Print("Please select a price order (1: Low to High, 2: High to Low): ")
	var priceOrder int
	fmt.Scan(&priceOrder)

	var sorted menuList
	sorted = *catalog
	if priceOrder == 1 {
		sorted = sortbyAsc(sorted, menuCount)
	} else {
		sorted = sortbyDesc(sorted, menuCount)
	}
	fmt.Printf("Here is our %s menu:\n", category)
	showMenuByCategory(sorted, menuCount, category)

	var currentOrders orderList
	var orderCount int = 0
	var exit bool

	exit = false
	for !exit {
		fmt.Println("what would you like to do?")
		fmt.Println("1. Show menu")
		fmt.Println("2. Add order")
		fmt.Println("3. delete order")
		fmt.Println("4. edit order")
		fmt.Println("5. show bill")
		fmt.Println("6. exit")
		var choose int
		fmt.Scan(&choose)

		if choose == 1 {
			//show menu func
			fmt.Println("Here is our menu:")
			fmt.Println("categories: Drink / Food ")
			fmt.Print("Please choose a category: ")
			var category string
			fmt.Scan(&category)
			fmt.Print("Please select a price order (1: Low to High, 2: High to Low): ")
			var priceOrder int
			fmt.Scan(&priceOrder)

			var sorted menuList
			sorted = *catalog
			if priceOrder == 1 {
				sorted = sortbyAsc(sorted, menuCount)
			} else {
				sorted = sortbyDesc(sorted, menuCount)
			}
			fmt.Printf("Here is our %s menu:\n", category)
			showMenuByCategory(sorted, menuCount, category)

		} else if choose == 2 {
			//add order func
			addOrder(&currentOrders, &orderCount, &*catalog, menuCount)
		} else if choose == 3 {
			//delete order func
			deleteOrder(&currentOrders, &orderCount, &*catalog, menuCount)
		} else if choose == 4 {
			//edit order
			editOrder(&currentOrders, orderCount)
		} else if choose == 5 {
			//show bill for customer
			showBill(currentOrders, orderCount)
		} else if choose == 6 {
			exit = true
			fmt.Println("Thank you for visiting our cafe, see you next time!")
		} else {
			fmt.Println("Invalid option")
		}
	}
}

func addOrder(cart *orderList, orderCount *int, catalog *menuList, menuCount int) {
	var orderName string
	var idxMenu int
	var quantity int

	if *orderCount >= NMAX {
		fmt.Println("Order list is full. Cannot add more orders.")
		return
	}

	fmt.Print("Menu name: ")
	fmt.Scan(&orderName)

	idxMenu = SearchMenu(*catalog, menuCount, orderName)
	if idxMenu == -1 || !(*catalog)[idxMenu].Status {
		fmt.Println("Menu not found or unavailable.")
		return
	}

	fmt.Print("Quantity: ")
	fmt.Scan(&quantity)

	if quantity <= 0 {
		fmt.Println("Invalid quantity.")
		return
	}
	if quantity > (*catalog)[idxMenu].Stock {
		fmt.Printf("Not enough stock. Stock remaining: %d\n", (*catalog)[idxMenu].Stock)
		return
	}

	(*cart)[*orderCount].Name = (*catalog)[idxMenu].Name
	(*cart)[*orderCount].Price = (*catalog)[idxMenu].Price
	(*cart)[*orderCount].Quantity = quantity

	(*catalog)[idxMenu].Stock -= quantity
	(*orderCount)++
	*cart = sortCartByName(*cart, *orderCount)

	fmt.Println("Order added successfully!")
	if (*catalog)[idxMenu].Stock == 0 {
		(*catalog)[idxMenu].Status = false
		fmt.Printf("'%s' is now out of stock.\n", (*catalog)[idxMenu].Name)
	}
}

func showMenuByCategory(catalog menuList, menuCount int, category string) {
	var found bool
	found = false
	var count, i int
	count = 0
	for i = 0; i < menuCount; i++ {
		if catalog[i].Category == category {
			count = count + 1
			var menu Menu = catalog[i]
			fmt.Printf("%d. %s - Rp%d, Stock: %d, (Available: %t)\n", count, menu.Name, menu.Price, menu.Stock, menu.Status)
			found = true
		}
	}
	if !found {
		fmt.Printf("Sorry, we don't have any %s menu at the moment.\n", category)
	}
}

func showBill(cart orderList, orderCount int) {
	var i, totalAmount, totalItem int
	totalAmount = 0
	totalItem = 0

	if orderCount == 0 {
		fmt.Println("No orders have been placed yet.")
		return
	}

	fmt.Println("=========================================")
	fmt.Println("              CAFE RECEIPT               ")
	fmt.Println("=========================================")
	fmt.Printf("  %-20s %6s %8s\n", "Item", "Qty", "Price")
	fmt.Println("-----------------------------------------")

	for i = 0; i < orderCount; i++ {
		var subtotal int
		subtotal = cart[i].Price * cart[i].Quantity
		fmt.Printf("  %-20s %6d %8d\n", cart[i].Name, cart[i].Quantity, subtotal)
		fmt.Printf("  @ Rp%d/item\n", cart[i].Price)
		totalAmount += subtotal
		totalItem += cart[i].Quantity
	}

	fmt.Println("-----------------------------------------")
	fmt.Printf("  %-20s %6d\n", "Total Items", totalItem)
	fmt.Println("-----------------------------------------")
	fmt.Printf("  %-20s %6s %8d\n", "TOTAL", "Rp", totalAmount)
	fmt.Println("=========================================")
	fmt.Println("       Thank you for your order!         ")
	fmt.Println("      Please come back again soon!       ")
	fmt.Println("-----------------------------------------")
	fmt.Println("        ❤️  with love by Aulia           ")
	fmt.Println("              & Jessie  ❤️               ")
	fmt.Println("=========================================")
}

func editOrder(cart *orderList, orderCount int) {
	var quantity, i int
	var enter, found bool
	var MenuName string
	enter = false
	for !enter {
		found = false
		fmt.Println("Choose Menu to edit:")
		fmt.Scan(&MenuName)
		for i = 0; i < orderCount && !found; i++ {
			if MenuName == (cart)[i].Name {
				found = true
				fmt.Println("Enter new quantity:")
				fmt.Scan(&quantity)
				(cart)[i].Quantity = quantity
				fmt.Println("The quantity of menu", (cart)[i].Name, "has been updated to", quantity)
			}
		}
		if !found {
			fmt.Println("Menu not found. Please enter a valid menu name.")
		} else {
			enter = true
		}
	}
}

func deleteOrder(cart *orderList, orderCount *int, catalog *menuList, menuCount int) {
	var idx, i, menuIdx int
	if *orderCount == 0 {
		fmt.Println("The order list is empty, nothing to delete")
		return
	}
	showOrder(*cart, *orderCount)
	fmt.Print("Enter the name of menu to delete: ")
	var name string
	fmt.Scan(&name)
	idx = SearchOrderBiner(*cart, *orderCount, name)
	if idx == -1 {
		fmt.Println("Menu not found")
		return
	}

	// restore stock back to catalog
	menuIdx = SearchMenu(*catalog, menuCount, (*cart)[idx].Name)
	if menuIdx != -1 {
		(*catalog)[menuIdx].Stock += (*cart)[idx].Quantity
		if (*catalog)[menuIdx].Stock > 0 {
			(*catalog)[menuIdx].Status = true
		}
	}

	for i = idx; i < *orderCount-1; i++ {
		(*cart)[i] = (*cart)[i+1]
	}
	*orderCount = *orderCount - 1
	fmt.Printf("Menu '%s' deleted succesfully and stock restored\n", name)
}

// show order menu

func showOrder(cart orderList, orderCount int) {
	if orderCount == 0 {
		fmt.Println("The order list is currently empty.")
		return
	}
	fmt.Println("\n--- Current Order List ---")
	var i int
	for i = 0; i < orderCount; i++ {
		var order Order = cart[i]
		fmt.Printf("%d. %s - Rp%d x%d\n", i+1, order.Name, order.Price, order.Quantity)
	}
}

func SearchOrderBiner(cart orderList, orderCount int, name string) int {
	var left, right, mid, idx int
	left = 0
	right = orderCount - 1
	idx = -1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if cart[mid].Name == name {
			idx = mid
		} else if cart[mid].Name < name {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return idx
}

//helper function
func SearchMenu(catalog menuList, menuCount int, name string) int {
	var i, idx int
	idx = -1
	for i = 0; i < menuCount && idx == -1; i++ {
		if catalog[i].Name == name {
			idx = i
		}
	}
	return idx
}

func sortbyAsc(list menuList, n int) menuList {
	var i, j, minIdx int
	for i = 0; i < n-1; i++ {
		minIdx = i
		for j = i + 1; j < n; j++ {
			if list[j].Price < list[minIdx].Price {
				minIdx = j
			}
		}
		var temp Menu = list[minIdx]
		list[minIdx] = list[i]
		list[i] = temp
	}
	return list
}

func sortbyDesc(list menuList, n int) menuList {
	var i, j int
	var keyMenu Menu
	for i = 1; i < n; i++ {
		keyMenu = list[i]
		j = i - 1
		for j >= 0 && list[j].Price < keyMenu.Price {
			list[j+1] = list[j]
			j = j - 1
		}
		list[j+1] = keyMenu
	}
	return list
}

func sortCartByName(list orderList, n int) orderList {
	var i, j int
	var keyOrder Order
	for i = 1; i < n; i++ {
		keyOrder = list[i]
		j = i - 1
		for j >= 0 && list[j].Name > keyOrder.Name {
			list[j+1] = list[j]
			j = j - 1
		}
		list[j+1] = keyOrder
	}
	return list
}

func statisticPerCatergory(catalog menuList, menuCount int) {
	var drinkCount, i, foodCount, totalPrice int
	var averagePrice float64
	drinkCount = 0
	foodCount = 0
	totalPrice = 0

	if menuCount == 0 {
		fmt.Println("No menus available for statistics.")
		return
	}

	for i = 0; i < menuCount; i++ {
		if catalog[i].Category == "Drink" {
			drinkCount = drinkCount + 1
		} else if catalog[i].Category == "Food" {
			foodCount = foodCount + 1
		}
		totalPrice = totalPrice + catalog[i].Price
	}
	averagePrice = float64(totalPrice) / float64(menuCount)

	fmt.Printf("Statistics by Category:\n")
	fmt.Printf("Drinks: %d\n", drinkCount)
	for i = 0; i < menuCount; i++ {
		if catalog[i].Category == "Drink" {
			fmt.Printf(" - %s: Rp%d\n", catalog[i].Name, catalog[i].Price)
		}
	}
	fmt.Printf("Food: %d\n", foodCount)
	for i = 0; i < menuCount; i++ {
		if catalog[i].Category == "Food" {
			fmt.Printf(" - %s: Rp%d\n", catalog[i].Name, catalog[i].Price)
		}
	}
	fmt.Printf("Total Items: %d\n", menuCount)
	fmt.Printf("Average Price: Rp%.0f\n", averagePrice)
}
