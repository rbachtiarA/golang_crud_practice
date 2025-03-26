package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie
var latestID int

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Set-Cookie", "id=asudelete")

	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			json.NewEncoder(w).Encode(map[string]any{
				"status": "success delete movie",
				"item":   item,
			})
			break
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)
	latestID++
	movie.ID = strconv.FormatInt(int64(latestID), 10)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(map[string]any{
		"status": "create movie success",
		"movie":  movie,
	})
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	if err := r.ParseForm(); err != nil {
		json.NewEncoder(w).Encode(map[string]any{
			"status": "error form",
			"error":  err,
		})
		return
	}

	if err := r.ParseForm(); err != nil {
		json.NewEncoder(w).Encode(map[string]any{
			"status": "error form",
		})
		return
	}
	for index, item := range movies {
		if item.ID == params["id"] {
			if title := r.FormValue("title"); title != "" {
				movies[index].Title = title
			}
			if isbn := r.FormValue("isbn"); isbn != "" {
				movies[index].Isbn = isbn
			}
			if firstname := r.FormValue("firstname"); firstname != "" {
				movies[index].Director.Firstname = firstname
			}
			if lastname := r.FormValue("lastname"); lastname != "" {
				movies[index].Director.Lastname = lastname
			}

			json.NewEncoder(w).Encode(map[string]any{
				"status": "success update movie",
				"movie":  movies[index],
			})
			return
		}
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("./templates/home/*html"))
	if err := tmpl.ExecuteTemplate(w, "home.html", movies); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies,
		Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}},
		Movie{ID: "2", Isbn: "434343", Title: "Movie Two", Director: &Director{Firstname: "Foo", Lastname: "Bar"}},
		Movie{ID: "3", Isbn: "282828", Title: "Movie Three", Director: &Director{Firstname: "Cloe", Lastname: "Won"}},
	)
	latestID = 3
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server on port 8080 \n")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server error")
	}
}
