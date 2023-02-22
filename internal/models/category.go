package models

// Category category struct
type Category struct {
	Model
	CategoryID uint        `json:"category_id"`
	Name       string      `json:"name"`
	IsParent   bool        `json:"is_parent"`
	Total      int         `json:"total" gorm:"<-:false"`
	Categories []*Category `json:"categories,omitempty" gorm:"foreignKey:CategoryID;references:ID"`
}

// TableName override table name for gorm
func (Category) TableName() string {
	return "categories"
}
