package models

// UnitType unit type
type UnitType uint

const (
	// UnitTypeUnknown unit type unknown
	UnitTypeUnknown UnitType = iota
	// Gram unit type gram
	Gram
	// Kilogram unit type kilogram
	Kilogram
	// Piece unit type piece or unit
	Piece
	// Pair unit type pair
	Pair
	// Set unit type set
	Set
	// Litre unit type litre
	Litre
	// Pound unit type pound
	Pound
	// CubicCentimetre unit type cubic centimetre
	CubicCentimetre
	// Carton unit type cartan
	Carton
)

// Type type
func (u UnitType) Type() *Type {
	name := "-"
	switch u {
	case Gram:
		name = "กรัม"

	case Kilogram:
		name = "กิโลกรัม"

	case Piece:
		name = "ชิ้น"

	case Pair:
		name = "คู่"

	case Set:
		name = "เซต"

	case Litre:
		name = "ลิตร"

	case Pound:
		name = "ปอนด์"

	case CubicCentimetre:
		name = "เซนติเมตร"

	case Carton:
		name = "กล่อง"

	}

	return &Type{
		Name:  name,
		Value: int(u),
	}
}

// GetName get name
func (p Product) GetName() string {
	if p.ProductMasterName != "" {
		return p.ProductMasterName + " " + p.Name
	}

	return p.Name
}

// Product product structure
type Product struct {
	Model
	Name                  string           `json:"name"`
	TotalSell             float64          `json:"total_sell"`
	SKU                   string           `json:"sku" gorm:"column:sku"`
	Quantity              float64          `json:"quantity"`
	ProductMasterName     string           `json:"product_master_name,omitempty" gorm:"<-:false" copier:"-"`
	ProductMasterImageIDs Int64Array       `json:"product_master_image_ids,omitempty" gorm:"type:bigint[];<-:false" copier:"-"`
	Unit                  UnitType         `json:"unit"`
	Price                 float64          `json:"price"`
	AvailableStock        float64          `json:"available_stock"`
	IsPreOrder            bool             `json:"is_pre_order"`
	IsVariant             bool             `json:"is_variant"`
	ProductID             uint             `json:"product_id"`
	Weight                float64          `json:"weight,omitempty"`
	Width                 float64          `json:"width,omitempty"`
	Height                float64          `json:"height,omitempty"`
	Length                float64          `json:"length,omitempty"`
	CategoryID            uint             `json:"category_id,omitempty"`
	Description           string           `json:"description"`
	CoverImageIDs         Int64Array       `json:"cover_image_ids,omitempty" gorm:"type:bigint[]"`
	CreatedByUserID       uint             `json:"created_by_user_id"`
	UpdatedByUserID       *uint            `json:"updated_by_user_id,omitempty"`
	DeletedByUserID       *uint            `json:"deleted_by_user_id,omitempty"`
	IsPublish             bool             `json:"is_publish"`
	CoverImages           []*File          `json:"cover_images,omitempty" gorm:"-"`
	Cagtegory             *Category        `json:"category,omitempty" gorm:"foreignKey:CategoryID;references:ID"`
	ProductVariants       []*Product       `json:"product_variants,omitempty" gorm:"foreignKey:ProductID;references:ID"`
	Variants              []*Variant       `json:"variants,omitempty" gorm:"foreignKey:ProductID;references:ID"`
	VariantOptions        []*ReturnVariant `json:"variant_options,omitempty" gorm:"-"`
}

// ReturnVariant return variant struct
type ReturnVariant struct {
	OptionName  string `json:"option_name"`
	VariantName string `json:"variant_name"`
}

// Variant variant struct
type Variant struct {
	Model
	OptionName      string `json:"option_name"`
	VariantName     string `json:"variant_name"`
	ProductID       uint   `json:"product_id"`
	CreatedByUserID uint   `json:"created_by_user_id"`
	UpdatedByUserID *uint  `json:"updated_by_user_id,omitempty"`
	DeletedByUserID *uint  `json:"deleted_by_user_id,omitempty"`
}

// TableName override table name
func (Variant) TableName() string {
	return "variants"
}

// TableName override table name
func (Product) TableName() string {
	return "products"
}
