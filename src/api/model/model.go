package model

type Product struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ImageURLs   []string `json:"image_urls"`
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
