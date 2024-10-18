package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	NextPage   int     `json:"nextPage"`
	TotalPages int     `json:"totalPages"`
	TotalItems int     `json:"totalItems"`
	Merch      []Merch `json:"merch"`
}

func (app *App) MerchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pageParam, err := validatePageParam(r.URL.Query().Get("page"), 1)
	if err != nil {
		http.Error(w, "Invalid page parameter. Please provide a positive integer.", http.StatusInternalServerError)
		return
	}
	pageSizeParam, err := validatePageParam(r.URL.Query().Get("pageSize"), 10)
	if err != nil {
		http.Error(w, "Invalid pageSize parameter. Please provide a positive integer.", http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case http.MethodGet:
		merchList, err := app.fetchMerch(pageParam, pageSizeParam)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		totalPages, totalItems, err := app.fetchMetaData(pageSizeParam)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		response := Response{
			CurPage:    pageParam,
			NextPage:   pageParam + 1,
			TotalPages: totalPages,
			TotalItems: totalItems,
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

func (app *App) fetchMerch(page int, pageSize int) ([]Merch, error) {
	merchList := []Merch{}
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

func (app *App) fetchMetaData(pageSize int) (int, int, error) { // returns totalPages and totalItems
	totalItemsQuery := "SELECT COUNT(*) FROM Merchandise"
	rows, err := app.DB.Query(totalItemsQuery)
	if err != nil {
		log.Printf("Database query error: %v", err)
		return 0, 0, err
	}
	defer rows.Close()

	var totalItems int
	if rows.Next() {
		err := rows.Scan(&totalItems)
		if err != nil {
			log.Printf("Row scanning error: %v", err)
			return 0, 0, err
		}
	}
	totalPages := (totalItems / pageSize) + 1
	return totalPages, totalItems, nil
}

func validatePageParam(param string, defaultValue int) (int, error) {
	if param == "" {
		return defaultValue, nil
	}

	page, err := strconv.Atoi(param)
	if err != nil {
		return 0, fmt.Errorf("PageParam Error: %v", err)
	}

	return page, nil
}
