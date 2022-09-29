// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Drink struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Flavour string `json:"flavour"`
	Price   *int   `json:"price"`
	Type    string `json:"type"`
	ML      int    `json:"mL"`
}

type NewDrink struct {
	ID      *string `json:"id"`
	Name    string  `json:"name"`
	Flavour string  `json:"flavour"`
	ML      int     `json:"mL"`
	Type    string  `json:"type"`
	Price   *int    `json:"price"`
}
