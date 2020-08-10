package structs

// Restaurant - object structure for representing restaurants
type Restaurant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type YelpResponse struct {
	Businesses []Restaurant `json:"businesses"`
	Total      int          `json:"total"`
}
