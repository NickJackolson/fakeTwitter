# FakeTwitter
A fake twitter app. Runs Vue on frontend and Golang on backend. It has almost no styling and code may sometimes differ from best-practice methods.
### Login Page
![Image of LoginPage](https://github.com/NickJackolson/fakeTwitter/blob/master/pictures/fakeTwitter1.png)
### Register Page
![Image of RegisterPage](https://github.com/NickJackolson/fakeTwitter/blob/master/pictures/fakeTwitter2.png)
### Home Page
![Image of RegisterPage](https://github.com/NickJackolson/fakeTwitter/blob/master/pictures/fakeTwitter3.png)


## How to run
### Starting with GO backend
GO [compiler](https://golang.org "compiler") must be installed and all backend files should be in your systems GOPATH which usually is located at `C:/Users/"YourUsername"/go/`
Make sure your GO PATH directory structure looks like this:
```
 ├───bin
 ├───pkg
 └───src
 └───github.com
     └───"YourGithubUsername"
          └───BackendFolder  (Should include main.go and data.db)
```

Before starting up the backend make sure you have installed:
"jwt-go"," handlers", "mux" and "go-sqlite3".
Install them by typing:
```
go get github.com/dgrijalva/jwt-go
go get github.com/gorilla/handlers
go get github.com/gorilla/mux
go get github.com/mattn/go-sqlite3

```
After all installed use ```go run main.go``` in BackendFolder to start the backend at port :8081.
### Vue JS FrontEnd
[NodeJs](https://nodejs.org/en/ "NodeJs") should be present in your system. Also you need Vue in order to run the frontend.

`npm install -g @vue/cli`

Before starting the frontend, following dependencies must be installed: vue-router and axios.
```
npm install vue-router
npm install axios
```
After install is completed use `npm run serve` from root of the frontend.

## How to run On Docker
### BackEnd
Change to backend directory.
First build the docker image by using docker command: 
```docker build -t go-app . ```
then use:
```docker run -it -p 8081:8081``` 
Go api will be ready at 8081

### FrontEnd
Change to frontend directory.
First build the frontend's docker image by typing
``` docker build -t dockerized-vuejs ```

then use 
```docker run -it -p 8080:8080``` 
in frontEnd directory. While it runs the app will be ready at localhost:8080