package models

type Todo struct {
	ID          uint   `json:"id" gorm:"primary_key;auto_increment;not_null"`
	Title       string `json:"title" gorm:"not_null"`
	Description string `json:"description" gorm:"not_null"`
	Status      bool   `json:"status" gorm:"not_null"`
}
