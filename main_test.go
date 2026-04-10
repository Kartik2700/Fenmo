package main

import (
	"testing"
)

func TestExpenseFiltering(t *testing.T) {
	// Setup mock data
	expenses := []Expense{
		{Amount: 100, Category: "Food", Description: "Lunch"},
		{Amount: 200, Category: "Rent", Description: "April Rent"},
	}

	// Test Case 1: Filter by Food
	filter := "Food"
	var results []Expense
	for _, e := range expenses {
		if filter == "" || e.Category == filter {
			results = append(results, e)
		}
	}

	if len(results) != 1 {
		t.Errorf("Expected 1 expense, got %d", len(results))
	}

	if results[0].Category != "Food" {
		t.Errorf("Expected category 'Food', got %s", results[0].Category)
	}
}
