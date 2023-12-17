💡 # GoFr Bookstore The GoFr Bookstore is a comprehensive book management system built on the GoFr framework,
incorporating seamless integration with Redis and MySQL databases. This README provides step-by-step instructions 
for setting up and running the application.

To get started with GoFr use the following command:
go get gofr.dev
The latest version of go in your system should be installed.

API Requests made using postman

1.GET ALL Request
Get all the Books ![Screenshot](C:\Users\Fayez Khan\go\bookstore\gofr\GET ALL.png)

2.CREATE Request
Add books to table
![Screenshot](C:\Users\Fayez Khan\go\bookstore\gofr\Create.png)

3.GET BY ID Request
Fetch book by BOOK ID
![Screenshot](C:\Users\Fayez Khan\go\bookstore\gofr\GET BY ID.png)

4.UPDATE Request
UPDATE the book's Author
![Screenshot](C:\Users\Fayez Khan\go\bookstore\gofr\UPDATE.png)

5.🗑️ DELETE Request
Deletes an existing book record
![Screenshot](C:\Users\Fayez Khan\go\bookstore\gofr\DELETE.png)

UML Diagrams

Project Structure
![Screenshot](C:\Users\Fayez Khan\go\bookstore\gofr\project structure.png)

ROUTES
![Routes](C:\Users\Fayez Khan\go\bookstore\gofr\routes.png)

Database 🛢️
The project utilizes a MySQL database to store information about book. The MySQL database is set up as a POSTMAN container for easy deployment and management.

Database Schema 🗃️
Table: books
Columns: id INT AUTO_INCREMENT PRIMARY KEY, name  VARCHAR(255) NOT NULL, author  VARCHAR(255), publication  VARCHAR(255) 
![Screenshot](C:\Users\Fayez Khan\go\bookstore\gofr\db schema.png)

Database Setup 🐳
POSTMAN Container ![Screenshot](C:\Users\Fayez Khan\go\bookstore\gofr\database setup.png)

To run the MySQL database locally as a postman container, use the following command:

mysql xmysql -h localhost u-Fayez-2403 -p Fayez@1724 -d bookstore

3306:3306 -d mysql:8.0.30
After this create a new connection in MYSQL workbench and make a projects table with the above defined schema to successfully setup the database.

⚡️ RUN

Now simply run the following command in your terminal to run the application:
go run cmd/main.go

cmd/main.go
![image](C:\Users\Fayez Khan\go\bookstore\gofr\main.go.png)

pkg/config/app.go
![image](C:\Users\Fayez Khan\go\bookstore\gofr\app.go.png)

pkg/controllers/book-controller.go
![image](C:\Users\Fayez Khan\go\bookstore\gofr\book-controller-go.png)

pkg/models/book.go
![image](C:\Users\Fayez Khan\go\bookstore\gofr\book.go.png)

pkg/routes/bookstore-routes.go
![image](C:\Users\Fayez Khan\go\bookstore\gofr\bookstore-routes.go.png)

pkg/utils/utils.go
![image](C:\Users\Fayez Khan\go\bookstore\gofr\utils.go.png)
