package models

//Users will store the details of user
type Users struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
