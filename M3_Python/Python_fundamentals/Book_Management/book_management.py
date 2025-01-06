from model import BookDetails, inventory

def add_inventory_item(name, writer, cost, stock):
    try:
        cost = float(cost)
        stock = int(stock)
        if cost <= 0 or stock <= 0:
            raise ValueError("Cost and Stock must be positive values.")
        inventory.append(BookDetails(name, writer, cost, stock))
        return "Item successfully added to inventory!"
    except ValueError as error:
        return str(error)

def list_inventory():
    if not inventory:
        return "Inventory is empty."
    return "\n".join(item.show_details() for item in inventory)

def find_inventory_item(keyword):
    results = [item.show_details() for item in inventory if keyword.lower() in item.name.lower() or keyword.lower() in item.writer.lower()]
    return "\n".join(results) if results else "No matching items found."
