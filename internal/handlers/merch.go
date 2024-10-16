package handlers

import (
	"encoding/json"
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
		merchList := []Merch{}

		rows, err := app.DB.Query("SELECT * FROM Merchandise")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			merch := Merch{}
			err = rows.Scan(&merch.ID, &merch.Name, &merch.Price, &merch.Availability, &merch.CreatedAt, &merch.ModifiedAt)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			merchList = append(merchList, merch)
		}

		response := Response{
			CurPage:    1,
			TotalItems: len(merchList),
			Merch:      merchList,
		}

		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}
