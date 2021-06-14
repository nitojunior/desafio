package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
		log.Println("HIT \"/\"", "Status: 405")
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - OK!"))
		log.Println("HIT \"/\"", "Status: 200")
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
		log.Println("HIT \"/login\"", "Status: 405")
	} else {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Internal Server Error!"))
			log.Fatal(err)
		}

		var u User

		json.Unmarshal(body, &u)

		if u.Username == "" || u.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("403 - Bad Request"))
			log.Println("HIT \"/login\"", "Status: 403")
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"token":"2d2f3349-e9a1-4b73-bd18-b236ac1dc02a"}`))
			log.Println("HIT \"/login\"", "Status: 200")
		}

	}
}

func Content(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
		log.Println("HIT \"/content\"", "Status: 405")
	} else {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Internal Server Error!"))
			log.Fatal(err)
		}

		var t Token

		json.Unmarshal(body, &t)

		if t.Token != "2d2f3349-e9a1-4b73-bd18-b236ac1dc02a" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("403 - Bad Request"))
			log.Println("HIT \"/content\"", "Status: 403")
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("200 - OK!"))
			log.Println("HIT \"/content\"", "Status: 200")
		}

	}
}

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
		log.Println("HIT \"/healthcheck\"", "Status: 405")
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - OK!"))
		log.Println("HIT \"/healthcheck\"", "Status: 200")
	}
}

func hadleRequests() {
	log.Println("Serving HTTP")
	http.HandleFunc("/", Homepage)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/content", Content)
	http.HandleFunc("/healthcheck", Healthcheck)
	http.ListenAndServe(":8080", nil)
}

func main() {
	hadleRequests()
}
