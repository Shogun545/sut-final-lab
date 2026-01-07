package entity

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title string `valid:"required~Title is required,length(3|100)~Title must length between 3 to 100"`
	Price float64 `valid:"required~Price is required,range(50|5000)~Price must be between 50 and 5000"`
	Code string `valid:"required~Code is required,matches(^BK+[0-9]{6}$)~Code must start with BK followed by 6 digits (0-9)"`
}