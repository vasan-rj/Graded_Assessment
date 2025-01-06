
from book_management import add_inventory_item, list_inventory, find_inventory_item
from customer_management import register_client, list_clients
from sales_management import process_sale, list_sales

def main_menu():
    while True:
        print("\n=== Welcome to BookMart ===")
        print("1. Manage Books")
        print("2. Manage Customers")
        print("3. Manage Sales")
        print("4. Quit Application")
        user_choice = input("Please select an option: ")

        if user_choice == "1":
            handle_books()
        elif user_choice == "2":
            handle_customers()
        elif user_choice == "3":
            handle_sales()
        elif user_choice == "4":
            print("Thank you for visiting BookMart. Goodbye!")
            break
        else:
            print("Invalid selection. Please choose again.")

def handle_books():
    while True:
        print("\n--- Book Management ---")
        print("1. Add a New Book")
        print("2. View All Books")
        print("3. Search for a Book")
        print("4. Return to Main Menu")
        book_option = input("Your choice: ")

        if book_option == "1":
            title = input("Enter Book Title: ")
            author = input("Enter Author Name: ")
            price = input("Enter Price: ")
            stock = input("Enter Quantity Available: ")
            print(add_inventory_item(title, author, price, stock))
        elif book_option == "2":
            print(list_inventory())
        elif book_option == "3":
            query = input("Search by Title or Author: ")
            print(find_inventory_item(query))
        elif book_option == "4":
            break
        else:
            print("Invalid input. Please try again.")

def handle_customers():
    while True:
        print("\n--- Customer Management ---")
        print("1. Register a Customer")
        print("2. View Customer List")
        print("3. Return to Main Menu")
        customer_option = input("Select an option: ")

        if customer_option == "1":
            full_name = input("Customer Name: ")
            email_address = input("Email Address: ")
            contact_number = input("Phone Number: ")
            print(register_client(full_name, email_address, contact_number))
        elif customer_option == "2":
            print(list_clients())
        elif customer_option == "3":
            break
        else:
            print("Invalid input. Please try again.")

def handle_sales():
    while True:
        print("\n--- Sales Management ---")
        print("1. Process a Sale")
        print("2. View Sales Records")
        print("3. Return to Main Menu")
        sales_option = input("Choose an action: ")

        if sales_option == "1":
            buyer = input("Enter Customer Name: ")
            book_name = input("Enter Book Title: ")
            quantity = input("Enter Quantity: ")
            print(process_sale(buyer, book_name, quantity))
        elif sales_option == "2":
            print(list_sales())
        elif sales_option == "3":
            break
        else:
            print("Invalid choice. Try again.")

if __name__ == "__main__":
    main_menu()
