package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

func (Customer) TableName() string {
	return "customer"
}
