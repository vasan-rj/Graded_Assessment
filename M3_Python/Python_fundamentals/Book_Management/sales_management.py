from model import BookDetails, Client, SaleRecord
from model import inventory , clients,transactions
# inventory = []
# clients = []
# transactions = []

def process_sale(buyer_name, title_of_book, qty_requested):
    try:
        requested_qty = int(qty_requested)
        for item in inventory:
            if item.title.strip().lower() == title_of_book.strip().lower():
                if item.quantity >= requested_qty:
                    item.quantity -= requested_qty
                    clients.append(transactions(buyer_name, title_of_book, requested_qty))
                    return f"Transaction completed! Stock left: {item.quantity}"
                else:
                    return f"Insufficient stock: Only {item.quantity} available."
        return "Book not found in inventory."
    except ValueError:
        return "Error: Enter a valid numerical value for quantity."

def list_sales():
    if len(transactions) == 0:
        return "Sales log is currently empty."
    return "\n".join(record.display() for record in transactions)
