package packages

//City class
type City struct {
	ID        int64  `gorm:"primary_key" json:"id"`
	CountryID int64  `gorm:"not null" json:"country_id"`
	Name      string `gorm:"not null" json:"name"`
}
