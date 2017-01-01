package packages

//Country class
type Country struct {
	ID   int64  `gorm:"primary_key" json:"id"`
	Name string `gorm:"not null" json:"name"`
}
