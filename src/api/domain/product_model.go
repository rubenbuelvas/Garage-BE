package domain

const (
	StatusAvailable  = "available"
	StatusOutOfStock = "out_of_stock"

	// TODO : Get this from a config file or database
	CategoryElectronics    = "electronics"
	CategoryClothing       = "clothing"
	CategoryHomeAppliances = "home_appliances"
	CategoryBooks          = "books"
	CategoryPlants         = "plants"
	CategoryAutomotive     = "automotive"
)

// TODO: Load this when starting the application
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
