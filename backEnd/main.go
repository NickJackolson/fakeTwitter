package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// User is a model
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Response given back to frontend after login
type Response struct {
	Username string
	Token    string
}

// Article is a model for articles which will be used in website
// it contains ID, title, content, post_time and user_info
type Article struct {
	Id      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Ptime   time.Time `json:"ptime"`
	Author  string    `json:"author"`
}

// Articles as slice initialized here
var articles []Article

var dbName string = "./data.db"

// Users as slice initialized here
var users []User

func connectToDB(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	return db
}

func homePage(w http.ResponseWriter, r *http.Request) {
	//var message = "Home Page of Mete App"
	_ = json.NewEncoder(w).Encode("Home")
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var serverUser User
	var dbUser User
	var resp Response
	_ = json.NewDecoder(r.Body).Decode(&serverUser)
	db := connectToDB(dbName)
	defer db.Close()

	if serverUser.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	rows, err := db.Query("SELECT * FROM user WHERE email = ?", serverUser.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Db error")
		return
	}
	for rows.Next() {
		rows.Scan(&dbUser.Id, &dbUser.Username, &dbUser.Password, &dbUser.Email)
	}
	rows.Close()
	if dbUser.Username == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("E-mail is incorrect")
		return
	}
	if dbUser.Password != serverUser.Password {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Password is incorrect")
		return
	}
	w.WriteHeader(http.StatusOK)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":  dbUser.Username,
		"email": dbUser.Email,
	})
	tokenString, err := token.SignedString([]byte("secretKey"))
	resp.Username = dbUser.Username
	resp.Token = tokenString
	_ = json.NewEncoder(w).Encode(resp)
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	db := connectToDB(dbName)
	defer db.Close()

	if user.Username == "" {
		json.NewEncoder(w).Encode("Need A Username")
		return
	}

	statement, err := db.Prepare("INSERT INTO user (username, password, email) VALUES (?, ?, ?)")
	if err != nil {
		json.NewEncoder(w).Encode("db error")
		return
	}
	_, err = statement.Exec(user.Username, user.Password, user.Email)
	if err != nil {
		json.NewEncoder(w).Encode("Check Credientals")
		return
	}
	json.NewEncoder(w).Encode("Success")
}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var articleList []Article
	var instance Article
	// reqToken := r.Header.Get("token")
	// if reqToken == "" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	_ = json.NewEncoder(w).Encode("Unauthorized access!")
	// 	return
	// }
	// user := r.Header.Get("user")
	// token, err := jwt.Parse(reqToken, "secretKey")
	// if err != nil {
	// 	panic(err)
	// }
	// claims := token.Claims.(jwt.MapClaims)
	// for key, value := range claims {
	// 	if value != user {
	// 		w.WriteHeader(http.StatusBadRequest)
	// 		_ = json.NewEncoder(w).Encode("Unauthorized access!")
	// 		return
	// 	}
	// }
	db := connectToDB(dbName)
	defer db.Close()
	rows, err := db.Query("SELECT id,title,content,ptime,author FROM article ORDER BY ptime DESC")
	if err != nil {
		fmt.Println("error occured while getting article")
	}
	for rows.Next() {
		rows.Scan(&instance.Id, &instance.Title, &instance.Content, &instance.Ptime, &instance.Author)
		articleList = append(articleList, instance)
	}

	rows.Close()
	json.NewEncoder(w).Encode(articleList)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get any params

	var instance Article
	instance.Id, _ = strconv.Atoi(params["id"])

	db := connectToDB(dbName)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM article WHERE id= ?", params["id"])
	if err != nil {
		fmt.Fprintf(w, "article couldn't be found in db")
	}
	for rows.Next() {
		rows.Scan(&instance.Id, &instance.Title, &instance.Content, &instance.Ptime, &instance.Author)
		json.NewEncoder(w).Encode(&instance)
	}
	rows.Close()
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article)

	params := mux.Vars(r) // get any params
	article.Id, _ = strconv.Atoi(params["id"])

	db := connectToDB(dbName)
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM article WHERE id= ?")
	if err != nil {
		fmt.Fprintf(w, "article couldn't be found in db")
	}
	statement.Exec(params["id"])

	statement, err = db.Prepare("INSERT INTO article (title, content, author) VALUES (?, ?, ?)")
	defer statement.Close()
	if err != nil {
		fmt.Fprintf(w, "Error inserting into db")
		return
	}
	statement.Exec(article.Title, article.Content, article.Author)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get any params
	db := connectToDB(dbName)
	defer db.Close()
	fmt.Printf("%q\n", params["id"])
	statement, err := db.Prepare("DELETE FROM article WHERE id= ?")
	if err != nil {
		fmt.Fprintf(w, "article couldn't be found in db")
	}
	statement.Exec(params["id"])
}

func postArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article)

	db := connectToDB(dbName)
	defer db.Close()

	if article.Content == "" {
		var message = "Error! No content given for new post"
		fmt.Fprintf(w, message)
		return
	}

	statement, err := db.Prepare("INSERT INTO article (title, content, author) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = statement.Exec(article.Title, article.Content, article.Author)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func createEmptyDB() {
	db := connectToDB(dbName)
	defer db.Close()

	// check if database exists else create and fill!
	_, err := db.Query("SELECT * FROM article")
	if err != nil {
		fmt.Println("Creating DB")
		statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY, username TEXT, password TEXT, email TEXT, UNIQUE(username,email))")
		if err != nil {
			fmt.Println("Error at user creation")
			return
		}
		statement.Exec()

		statement, err = db.Prepare("CREATE TABLE article (id INTEGER PRIMARY KEY, title NVARCHAR(50), content NVARCHAR(50), ptime TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL, author NVARCHAR(50))")

		if err != nil {
			fmt.Println("Error2 at article creation")
			log.Fatal(err)
			return
		}
		statement.Exec()

		return
	}
	fmt.Println("DB Already Exists")
}

func handleRequests() {

	myRouter := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-CSRF-Token", "X-Requested-With", "Content-Type", "Authorization", "Accept", "Accept-Encoding", "Content-Length"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/login", loginPage).Methods("POST")
	myRouter.HandleFunc("/register", registerUser).Methods("POST")
	myRouter.HandleFunc("/articles", getAllArticles).Methods("GET")
	myRouter.HandleFunc("/articles", postArticle).Methods("POST")
	myRouter.HandleFunc("/articles/{id}", getArticle).Methods("GET")
	myRouter.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(headers, methods, origins)(myRouter)))
}

func main() {
	//send data to db if there is none
	createEmptyDB()
	//start API
	handleRequests()
}
