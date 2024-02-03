package model

type Icon struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"not null" `
	Category  string `json:"category"`
	ObjectKey string `json:"object_key"`
}

type IconNameResponse struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}
