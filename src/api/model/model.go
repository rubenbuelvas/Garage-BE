package model

const (
	StatusAvailable  = "available"
	StatusOutOfStock = "out_of_stock"

	CategoryElectronics    = "electronics"
	CategoryClothing       = "clothing"
	CategoryHomeAppliances = "home_appliances"
	CategoryBooks          = "books"
	CategoryPlants         = "plants"
	CategoryAutomotive     = "automotive"
)

var Categories = []string{
	CategoryElectronics,
	CategoryClothing,
	CategoryHomeAppliances,
	CategoryBooks,
	CategoryPlants,
	CategoryAutomotive,
}

var Statuses = []string{
	StatusAvailable,
	StatusOutOfStock,
}

type Product struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ImageURLs   []string `json:"image_urls"`
	Category    string   `json:"category"`
	Status      string   `json:"status"`
	Price       float64  `json:"price"`
	Quantity    int      `json:"quantity"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
