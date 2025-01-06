from flask import Flask, jsonify, request
import sqlite3
import os

application = Flask(__name__)

DB_FILE = 'library_catalog.db'

# Function to initialize the database
def setup_database():
    if not os.path.exists(DB_FILE):
        conn = sqlite3.connect(DB_FILE)
        cur = conn.cursor()
        cur.execute('''
            CREATE TABLE library (
                book_id INTEGER PRIMARY KEY AUTOINCREMENT,
                book_title TEXT NOT NULL,
                author_name TEXT NOT NULL,
                year_published INTEGER NOT NULL,
                book_genre TEXT NOT NULL
            )
        ''')
        
        books_to_insert = [
            ("The Great Escape", "Ruby Stark", 1925, "Fiction"),
            ("Silent Spring", "Rachel Carson", 1962, "Non-Fiction"),
            ("Dune", "Frank Herbert", 1965, "Science Fiction"),
            ("The Da Vinci Code", "Dan Brown", 2003, "Thriller"),
        ]

        cur.executemany('''
            INSERT INTO library (book_title, author_name, year_published, book_genre)
            VALUES (?, ?, ?, ?)
        ''', books_to_insert)
        conn.commit()
        conn.close()
        print("Database has been set up.")

# Function to establish a database connection
def connect_to_db():
    conn = sqlite3.connect(DB_FILE)
    conn.row_factory = sqlite3.Row
    return conn

# Endpoint to add a new book
@application.route('/library', methods=['POST'])
def create_book():
    details = request.get_json()
    required_keys = ['book_title', 'author_name', 'year_published', 'book_genre']

    if not all(key in details for key in required_keys):
        return jsonify({"error": "Incomplete information", "details": "All fields are mandatory."}), 400

    try:
        conn = connect_to_db()
        cur = conn.cursor()
        cur.execute('''
            INSERT INTO library (book_title, author_name, year_published, book_genre)
            VALUES (?, ?, ?, ?)
        ''', (details['book_title'], details['author_name'], details['year_published'], details['book_genre']))
        conn.commit()
        new_id = cur.lastrowid
        conn.close()
        return jsonify({"message": "Book added successfully", "book_id": new_id}), 201
    except Exception as error:
        return jsonify({"error": "Database issue", "details": str(error)}), 500

# Endpoint to fetch books, with optional filters
@application.route('/library', methods=['GET'])
def fetch_books():
    genre_filter = request.args.get('book_genre')
    author_filter = request.args.get('author_name')
    query = 'SELECT * FROM library WHERE 1=1'
    params = []

    if genre_filter:
        query += ' AND book_genre = ?'
        params.append(genre_filter)
    if author_filter:
        query += ' AND author_name = ?'
        params.append(author_filter)

    conn = connect_to_db()
    result_set = conn.execute(query, params).fetchall()
    conn.close()
    return jsonify([dict(record) for record in result_set]), 200

# Endpoint to fetch a single book by ID
@application.route('/library/<int:book_id>', methods=['GET'])
def fetch_book_by_id(book_id):
    conn = connect_to_db()
    record = conn.execute('SELECT * FROM library WHERE book_id = ?', (book_id,)).fetchone()
    conn.close()
    if record is None:
        return jsonify({"error": "Not Found", "details": "No book matches the provided ID."}), 404
    return jsonify(dict(record)), 200

# Endpoint to modify book details
@application.route('/library/<int:book_id>', methods=['PUT'])
def modify_book(book_id):
    updates = request.get_json()
    conn = connect_to_db()
    existing_book = conn.execute('SELECT * FROM library WHERE book_id = ?', (book_id,)).fetchone()

    if existing_book is None:
        return jsonify({"error": "Not Found", "details": "No book matches the provided ID."}), 404

    conn.execute('''
        UPDATE library
        SET book_title = ?, author_name = ?, year_published = ?, book_genre = ?
        WHERE book_id = ?
    ''', (updates.get('book_title', existing_book['book_title']),
          updates.get('author_name', existing_book['author_name']),
          updates.get('year_published', existing_book['year_published']),
          updates.get('book_genre', existing_book['book_genre']),
          book_id))
    conn.commit()
    conn.close()
    return jsonify({"message": "Book details updated successfully."}), 200

# Endpoint to delete a book by ID
@application.route('/library/<int:book_id>', methods=['DELETE'])
def remove_book(book_id):
    conn = connect_to_db()
    existing_book = conn.execute('SELECT * FROM library WHERE book_id = ?', (book_id,)).fetchone()

    if existing_book is None:
        return jsonify({"error": "Not Found", "details": "No book matches the provided ID."}), 404

    conn.execute('DELETE FROM library WHERE book_id = ?', (book_id,))
    conn.commit()
    conn.close()
    return jsonify({"message": "Book removed successfully."}), 200

if __name__ == "__main__":
    setup_database()  # Initialize the database during the first run
    application.run(debug=True)
