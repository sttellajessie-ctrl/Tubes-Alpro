# README

## Group Identity
 **Group Name:** Study Syndicate

 **Group Theme:** Digital Café Menu Catalogue Application.

 **Members:**
  * Jessie Sttella (NIM: 103012540001)
  * Dania Aulia (NIM: 103012540024)

## Description
This application is designed to manage and process menu data for both food and beverages in a cafe. It serves as an efficient data management and inventory monitoring tool tailored specifically for cafe operators. It helps cafe operators keep track of menu availability, pricing, and stock levels in real time.


## Specification
The application is split into two distinct user interfaces based on roles, each possessing specific data manipulation privileges:

### 1. Admin Interface
Grants the user to Add, Edit, Delete, and View global items within the master cafe menu catalogue.

### 2. Customer Interface
Grants the user purchasing privileges to manage their temporary shopping cart. Customers can:
* **Add** items to their cart
* **Modify** order quantities
* **Delete** ordered items
* **View** their current bill

### System and Data Rules
* **Menu Attributes:** Menu name, Category, Price, Ingredients (with a maximum of 10 ingredients), IngredientCount, Status (indicating whether the item is currently in stock), and Stock.
* **Maximum Capacity:** The system can store a maximum of 100 menu items ($NMAX = 100$).
* **Customer Features:** Customers can view menu records filtered dynamically by a specific category and sorted by price in either ascending (low-to-high) or descending (high-to-low) order.
* **Automatic Inventory Management:** 
  * Whenever a customer successfully adds an item to their order cart, the system automatically decrements that item's inventory in the master catalogue based on the selected quantity.
  * If a menu item's stock reaches zero (“0”), the system instantly updates its Status flag to `false`.

## How The Project Works

### 1. Adding an Item
* **Admin:** Allows the Admin to type in new foods or drinks to sell. The admin inputs the item's Name, Category, Price, Ingredients, and Stock. The app automatically checks if the overall menu is full (max 100 items) and if the recipe has too many ingredients (max 10). If everything looks good and the stock is above 0, the app turns the item's status to **Available (true)**.
* **Customer:** Allows the Customer to add items based on what has been inputted by the Admin.

### 2. Changing an Item (Update)
* Allows the Admin to update details for menu price and stock.
* Allows the Customer to update the quantity of their ordered items.
* The system uses **Selection Sort** to find the location of the item.

### 3. Deleting an Item
* **Admin:** Allows the Admin to remove an item from the Menu. It uses *Selection Sort* to find the item’s location by its name.
* **Customer:** Allows the Customer to remove an item from the Cart but prevents them from removing the item from the main Menu.
* **List Mechanism:** To avoid leaving an ugly, empty gap in the middle of the list, the app **shifts every item after it one step to the left**, effectively drawing over the deleted item.

### 4. Showing the Menu
Displays the items so users can see what to manage (Admin) or buy (Customer):
* **Admin View:** Prints a giant list showing everything, including ingredients, exact stock counts, and whether it's available or not.
* **Customer View:** Keeps things clean. It asks the customer to choose a category (like just "Food" or just "Drink") so they aren't overwhelmed.

### 5. Sorting the Price
Rearranges the menu based on price so customers can shop smart:
* **Low to High (Ascending):** The app uses a **Selection Sort** algorithm. It scans the menu, finds the absolute cheapest item, puts it at the top, and repeats until sorted.
* **High to Low (Descending):** The app uses an **Insertion Sort** algorithm. It takes items one by one and inserts them into their correct position from most expensive to cheapest.


## Task Distribution

### Member 1: Jessie Sttella
* **Project Architecture & Flow:**
  * Played a primary role in planning, brainstorming, and designing the project's foundational logic.
  * Prepared the initial project concept, feature planning, and program flow.
  * Implemented the main layout and structural parts of the Admin Menu.
  * Implemented the Exit Menu function.
* **Data Management & Operations:**
  * Developed the Edit Data feature to allow updates to the catalogue.
  * Developed the Delete Data feature for removing catalogue entries.
* **Search Algorithms & Transaction System:**
  * Implemented the Binary Search algorithm for efficient data retrieval.
  * Developed the Add Order functionality to handle customer selections.
  * Developed the Show Bill system to generate and present billing details.
* Collaborated and worked closely with Member 2 throughout the entire discussion, integration, and development process.

### Member 2: Dania Aulia
* **Project Contribution & Integration:**
  * Contributed actively to the core development and integration phase of the project.
  * Implemented vital parts of the Admin Menu interface.
* **Catalogue Input & Display System:**
  * Developed the Add Menu feature within the Admin Menu.
  * Developed the Show Menu feature to view general catalogue items.
  * Developed the Show by Category feature for filtered data viewing.
* **Algorithm & Analytics Development:**
  * Developed the complete Sorting feature, implementing both Insertion Sort and Selection Sort mechanisms.
  * Programmed both Ascending and Descending sorting capabilities for data organization.
  * Developed the Show Order feature to track active selections.
  * Developed the Statistic per Category feature to provide analytical data insights.
* **Testing & Refinement:**
  * Participated heavily in system testing, debug discussions, and programmatic refinements to ensure total project cohesion.


> **Side Note:**
> Both members have maintained an equal and highly collaborative workflow, successfully splitting the core algorithms (Sorting vs. Searching) and backend management (CRUD operations vs. User Order/Statistics tracking) to reach the current 80% completion mark.