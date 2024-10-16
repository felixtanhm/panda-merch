package handlers

import (
	"encoding/json"
	"net/http"
)

type Merch struct {
	CurPage    int    `json:"curPage"`
	TotalItems int    `json:"totalItems"`
	Message    string `json:"message"`
}

func MerchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		response := Merch{
			CurPage:    1,
			TotalItems: 0,
			Message:    "Hello",
		}
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}
