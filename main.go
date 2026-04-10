package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

// Expense structure to store details about expense
type Expense struct {
	ID          int       `json:"id"`
	Amount      int64     `json:"amount"`      // Stored in cents [cite: 33]
	Category    string    `json:"category"`    // [cite: 34]
	Description string    `json:"description"` // [cite: 35]
	Date        string    `json:"date"`        // [cite: 36]
	CreatedAt   time.Time `json:"created_at"`  // [cite: 37]
}

// In-Memory Database
type DB struct {
	sync.RWMutex
	expenses []Expense
	nextID   int
}

var store = &DB{
	expenses: []Expense{},
	nextID:   1,
}

func main() {
	mux := http.NewServeMux()

	// Post endpoint to register the Expenses
	mux.HandleFunc("POST /expenses", RegisterExpense)

	log.Println("Server starting on :8080 (In-Memory Mode)...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func RegisterExpense(w http.ResponseWriter, r *http.Request) {
	var e Expense
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	store.Lock()
	defer store.Unlock()
	e.ID = store.nextID
	e.CreatedAt = time.Now()
	store.expenses = append(store.expenses, e)
	store.nextID++

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(e)
}
