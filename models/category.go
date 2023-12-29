package models

type Category struct {
    CategoryID   int    `json:"category_id"`
    CategoryName string `json:"category_name"`
    CategoryImage string `json:"category_image"`
}
type CategoryIDWithName struct{
	CategoryID   int    `json:"category_id"`
    CategoryName string `json:"category_name"`
}