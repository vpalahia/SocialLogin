# Oauth 

A web page to login from different social applications(github, linkedin)

Follow the bellow steps to run and test the application
Pre-Requisite: 
There must be a Postgres server running locally on port 5432 and a database should exist with name "users" along with username as "postgres" and password set to "Welcome@123" otherwise db/init.go file can be configured according to the existing database before running the server. 


1. Clone the respoitory.  
2. Run "go run main.go" command at the root of the repository.
3. In the browser go to "http://localhost:8084/login" url.
4. Click on any one (LinkedIn or Github) icon.
5. Enter the credentials for the same and login.
6. If logged in through github, Hello and your name will be printed.
7. If logged in through linkedin, Hello and you email Id will be printed.
8. If user logs in successfully, the details will be saved in database.
9. Details of all users can be fetched through postman.
10. Postman collection file has been placed in repository.
