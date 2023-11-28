package models
import "gorm.io/gorm"
type Advertisement struct {
    ID          uint    `gorm:"primary_key; autoIncrement" json:"id"`
    Title       *string `json:"title"`
    Description *string `json:"description"`
    Category    *string `json:"category"`
    Price       *float64 `json:"price"`
}
func MigrateAdv(db *gorm.DB) error{
	err:=db.AutoMigrate(&Advertisement{})
	return err
}