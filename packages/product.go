package packages

//Product class
type Product struct {
	ID    int64   `gorm:"primary_key" json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	Stock int     `json:"stock"`
	Brand string  `json:"brand"`
	Image string  `json:"image"`
}
