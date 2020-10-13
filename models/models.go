package models

//Users will store the details of user
type Users struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	Name  string `json:"fullname"`
	Email string `json:"email"`
}
