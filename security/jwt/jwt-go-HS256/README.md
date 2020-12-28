Based on the https://www.sohamkamani.com/golang/2019-01-01-jwt-authentication/#creating-the-http-server


Running our application
To run this application, build and run the Go binary:

go build
./jwt-go-example
Now, using any HTTP client with support for cookies (like Postman, or your web browser) make a sign-in request with the appropriate credentials:

POST http://localhost:8000/signin

{"username":"user1","password":"password1"}
You can now try hitting the welcome route from the same client to get the welcome message:

GET http://localhost:8000/welcome
Hit the refresh route, and then inspect the clients cookies to see the new value of the token cookie:

POST http://localhost:8000/refresh
You can find the working source code for this example here.