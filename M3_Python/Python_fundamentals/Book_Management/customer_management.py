from model import Client, clients

def register_client(full_name, contact_email, contact_number):
    if not full_name or not contact_email or not contact_number:
        return "Error: All fields (full name, email, contact number) are mandatory."
    clients.append(Client(full_name, contact_email, contact_number))
    return "Client registered successfully!"

def list_clients():
    if not clients:
        return "No clients registered."
    return "\n".join(client.show_details() for client in clients)
