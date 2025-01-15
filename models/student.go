package models

type Student struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"size:100;not null" json:"name"`
	Age   int    `gorm:"not null" json:"age"`
	Grade string `gorm:"size:10;not null" json:"grade"`
}
