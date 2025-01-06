class BookDetails:
    def __init__(self, name, writer, cost, stock):
        self.name = name
        self.writer = writer
        self.cost = cost
        self.stock = stock

    def show_details(self):
        return f"Name: {self.name}, Writer: {self.writer}, Cost: {self.cost}, Stock: {self.stock}"


class Client:
    def __init__(self, full_name, contact_email, contact_number):
        self.full_name = full_name
        self.contact_email = contact_email
        self.contact_number = contact_number

    def show_details(self):
        return f"Full Name: {self.full_name}, Email: {self.contact_email}, Contact Number: {self.contact_number}"


class SaleRecord:
    def __init__(self, client_name, book_name, units_sold):
        self.client_name = client_name
        self.book_name = book_name
        self.units_sold = units_sold

    def show_details(self):
        return f"Client: {self.client_name}, Book: {self.book_name}, Units Sold: {self.units_sold}"


# Centralized data repositories
inventory = []
clients = []
transactions = []
