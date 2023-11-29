package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	connString := "user=postgres dbname=restAPI sslmode=disable"

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL1 := `
    CREATE TABLE IF NOT EXISTS entity1 (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255)
        -- Другие поля, связанные с Entity1
    )
`

	createTableSQL2 := `
    CREATE TABLE IF NOT EXISTS entity2 (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255)
        -- Другие поля, связанные с Entity2
    )
`

	createTableSQL3 := `
    CREATE TABLE IF NOT EXISTS entity3 (
        id SERIAL PRIMARY KEY,
        description VARCHAR(255)
        -- Другие поля, связанные с Entity3
    )
`

	_, err = db.Exec(createTableSQL1)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(createTableSQL2)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(createTableSQL3)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Таблицы Entity1, Entity2 и Entity3 созданы успешно.")

	http.HandleFunc("/entity1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			entities, err := entity1Repo.GetAll()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Отправляем JSON-ответ
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(entities); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

	http.HandleFunc("/entity2", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			entities, err := entity2Repo.GetAll()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Отправляем JSON-ответ
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(entities); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/entity3", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			entities, err := entity3Repo.GetAll()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(entities); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

}
