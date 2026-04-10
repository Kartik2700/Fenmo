package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sort"
	"sync"
	"time"
)

// Expense structure to store details about expense
type Expense struct {
	ID          int       `json:"id"`
	Amount      int64     `json:"amount"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Date        string    `json:"date"`
	CreatedAt   time.Time `json:"created_at"`
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
	// This serves index.html
	mux.Handle("/", http.FileServer(http.Dir("./")))

	// Post endpoint to register the Expenses
	mux.HandleFunc("POST /expenses", RegisterExpense)

	// Get endpoint to fetch Expenses
	mux.HandleFunc("GET /expenses", GetExpenses)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Fallback for local development
	}

	log.Printf("Server starting on : %s (In-Memory Mode)...\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}

// RegisterExpense to store the expense
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

// GetExpenses To Retrieve Expenses
func GetExpenses(w http.ResponseWriter, r *http.Request) {
	categoryFilter := r.URL.Query().Get("category")
	sortOrder := r.URL.Query().Get("sort")

	store.RLock()
	defer store.RUnlock()
	// Create a copy to avoid modifying the original during sorting/filtering
	results := make([]Expense, 0)
	for _, e := range store.expenses {
		if categoryFilter == "" || e.Category == categoryFilter {
			results = append(results, e)
		}
	}

	if sortOrder == "date_desc" {
		sort.Slice(results, func(i, j int) bool {
			return results[i].Date > results[j].Date
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
