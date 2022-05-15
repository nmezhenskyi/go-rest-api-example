// Package model provides type definitions for domain models.
package model

type Wine struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Label    string `json:"label,omitempty"`
	Volume   string `json:"volume"`
	Region   string `json:"region"`
	Producer string `json:"producer"`
	Year     int    `json:"year"`
	Alcohol  string `json:"alcohol"`
	Price    string `json:"price,omitempty"`
}
