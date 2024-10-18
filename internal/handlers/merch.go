package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Merch struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	Availability bool      `json:"availability"`
	CreatedAt    time.Time `json:"createdAt"`
	ModifiedAt   time.Time `json:"modifiedAt"`
}

type Response struct {
	CurPage    int     `json:"curPage"`
	TotalItems int     `json:"totalItems"`
	Merch      []Merch `json:"merch"`
}

func (app *App) MerchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		merchList, err := app.fetchMerch()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		response := Response{
			CurPage:    1,
			TotalItems: len(merchList),
			Merch:      merchList,
		}

		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Printf("JSON encoding error: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func (app *App) fetchMerch() ([]Merch, error) {
	merchList := []Merch{}
	pageSize := 10
	page := 1
	offset := (page - 1) * pageSize
	query := "SELECT id, name, price, availability, createdAt, modifiedAt FROM Merchandise ORDER BY id OFFSET @offset ROWS FETCH NEXT @pageSize ROWS ONLY;"
	rows, err := app.DB.Query(query, sql.Named("offset", offset), sql.Named("pageSize", pageSize))
	if err != nil {
		log.Printf("Database query error: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		merch := Merch{}
		err = rows.Scan(&merch.ID, &merch.Name, &merch.Price, &merch.Availability, &merch.CreatedAt, &merch.ModifiedAt)
		if err != nil {
			log.Printf("Row scanning error: %v", err)
			return nil, err
		}
		merchList = append(merchList, merch)
	}
	err = rows.Err()
	if err != nil {
		log.Printf("Row iteration error: %v", err)
		return nil, err
	}

	return merchList, nil
}
