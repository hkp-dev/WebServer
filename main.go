package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runAll/models"
	"strconv"

	"github.com/gorilla/mux"
	// "runAll/services"
)

var (
	bindAddr = "0.0.0.0:8080"
	userFile = "users.json"
)

func HashPassword(pwd string) string {
	hash := sha256.New()
	hash.Write([]byte(pwd))
	return hex.EncodeToString(hash.Sum(nil))
}
func main() {
	//first: create a new router
	r := mux.NewRouter()
	r.Use()
	//second: register a handler
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintf(w, "OK"); err != nil {
			log.Printf("%s - %s - Error: %s", r.RemoteAddr, r.RequestURI, err)
			return
		}
	}).Methods(http.MethodGet)
	r.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		htmlcontent := `<!DOCTYPE html>
						<html lang="en">
						<head>
							<meta charset="UTF-8">
							<meta name="viewport" content="width=device-width, initial-scale=1.0">
							<title>Home</title>
							<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC" crossorigin="anonymous">
							<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js" integrity="sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM" crossorigin="anonymous"></script>
						</head>
						<body>
							<nav class="navbar navbar-expand-sm navbar-dark bg-danger">
								<div class="container-fluid">
									<a class="navbar-brand text-primary" href="#">WeeDigitalX</a>
									<button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#mynavbar">
										<span class="navbar-toggler-icon"></span>
									</button>
									<div class="collapse navbar-collapse" id="mynavbar">
										<ul class="navbar-nav me-auto">
											<li class="nav-item">
												<a class="nav-link" href="#">Home</a>
											</li>
											<li class="nav-item">
												<a class="nav-link" href="#"></a>
											</li>
											<li class="nav-item">
												<a class="nav-link" href="#">Contact</a>
											</li>
										</ul>
										<ul class="navbar-nav ms-auto">
											<li class="nav-item">
												<a class="btn btn-outline-light me-2" href="#">Đăng Nhập</a>
											</li>
											<li class="nav-item">
												<a class="btn btn-light" href="/register">Đăng Ký</a>
											</li>
										</ul>
									</div>
								</div>
							</nav>
						</body>
						</html>
						`

		w.Header().Set("Content-Type", "text/html")
		if _, err := fmt.Fprint(w, htmlcontent); err != nil {
			log.Printf("%s - %s - Error: %s", r.RemoteAddr, r.RequestURI, err)
			return
		}
	}).Methods(http.MethodGet)
	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			htmlcontent := `<!DOCTYPE html>
							<html lang="en">
							<head>
								<meta charset="UTF-8">
								<meta name="viewport" content="width=device-width, initial-scale=1.0">
								<title>Đăng Ký</title>
								<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC" crossorigin="anonymous">
								<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js" integrity="sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM" crossorigin="anonymous"></script>
							</head>
							<body>
								<nav class="navbar navbar-expand-sm navbar-dark bg-danger">
									<div class="container-fluid">
										<a class="navbar-brand text-primary" href="#">WeeDigitalX</a>
										<button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#mynavbar">
											<span class="navbar-toggler-icon"></span>
										</button>
										<div class="collapse navbar-collapse" id="mynavbar">
											<ul class="navbar-nav me-auto">
												<li class="nav-item">
													<a class="nav-link" href="#">Home</a>
												</li>
												<li class="nav-item">
													<a class="nav-link" href="#"></a>
												</li>
												<li class="nav-item">
													<a class="nav-link" href="#">Contact</a>
												</li>
											</ul>
											<ul class="navbar-nav ms-auto">
												<li class="nav-item">
													<a class="btn btn-outline-light me-2" href="#">Login</a>
												</li>
												<li class="nav-item">
													<a class="btn btn-light" href="/register">Register</a>
												</li>
											</ul>
										</div>
									</div>
								</nav>

								<div class="container mt-4">
									<h2 class="text-center text-primary">REGISTER</h2>
									<form id="registrationForm" method="POST" action="/register">
										<div class="mb-3">
											<label for="username" class="form-label">Username</label>
											<input type="text" class="form-control" id="username" name="username" required>
										</div>
										<div class="mb-3">
											<label for="password" class="form-label">Password</label>
											<input type="password" class="form-control" id="password" name="password" required>
										</div>
										<div class="mb-3">
											<label for="name" class="form-label">Name</label>
											<input type="text" class="form-control" id="name" name="name" required>
										</div>
										<div class="mb-3">
											<label for="age" class="form-label">Age</label>
											<input type="number" class="form-control" id="age" name="age" required>
										</div>
										<button type="submit" class="btn btn-primary">Register</button>
									</form>
								</div>
							</body>
							</html>
							`
			w.Header().Set("Content-Type", "text/html")
			if _, err := fmt.Fprint(w, htmlcontent); err != nil {
				log.Printf("%s - %s - Error: %s", r.RemoteAddr, r.RequestURI, err)
				return
			}
		} else if r.Method == http.MethodPost {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Invalid form data", http.StatusBadRequest)
				log.Printf("%s - %s - Error: %s", r.RemoteAddr, r.RequestURI, err)
				return
			}
			var users []models.User
			var user models.User
			//get value from form
			user.Username = r.FormValue("username")
			user.Name = r.FormValue("name")
			hashPwd := r.FormValue("password")
			user.Password = HashPassword(hashPwd)

			//conversion age from int to string
			ageStr := r.FormValue("age")
			age, err := strconv.Atoi(ageStr)
			if err != nil {
				http.Error(w, "Invalid age format", http.StatusBadRequest)
				log.Printf("%s - %s - Error: %s", r.RemoteAddr, r.RequestURI, err)
			}

			user.Age = age
			file, err := os.Open(userFile)
			if err != nil {
				if !os.IsNotExist(err) {
					log.Printf("%s - %s - Error: %s", r.RemoteAddr, r.RequestURI, err)
				}
			} else {
				defer file.Close()
				if err := json.NewDecoder(file).Decode(&users); err != nil {
					log.Printf("%s - %s - Error: %s", r.RemoteAddr, r.RequestURI, err)
				}
			}
			user.ID = len(users) + 1    // auto increase
			users = append(users, user) // add new user
			file, err = os.Create(userFile)
			if err != nil {
				http.Error(w, "Error creating users file", http.StatusInternalServerError)
				return
			}
			defer file.Close()

			if err := json.NewEncoder(file).Encode(users); err != nil {
				http.Error(w, "Error writing to users file", http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// htmlcontent := ``
		}
	}).Methods(http.MethodPost, http.MethodGet)
	//create file log to follow action of client
	logFile, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error openning log file: %s", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println("Server started successfully")
	serv := http.Server{
		Addr:    bindAddr,
		Handler: r,
	}
	//print the bind address
	fmt.Printf("Server listening on: %s...\n", bindAddr)
	//start server
	if err := serv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v\n", err)
	}
	// http.HandleFunc("/add_user", services.AddUser)
	// http.HandleFunc("/get_user", services.GetUser)
}
