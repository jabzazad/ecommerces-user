package models

// LoginType login channel
type LoginType uint

const (
	// LoginTypeUnknown login channel
	LoginTypeUnknown LoginType = iota
	// LoginTypeNormal login normal
	LoginTypeNormal
	// LoginTypeGoogle login channel
	LoginTypeGoogle
	// LoginTypeFacebook login channel
	LoginTypeFacebook
)

// UserRole user role
type UserRole uint

const (
	// UnknownRole unknown role
	UnknownRole UserRole = iota
	// RoleCustomer user role customer
	RoleCustomer
	// RoleShopEmployee role shop employee
	RoleSeller UserRole = 5
	// RoleUser user role user
	RoleAdmin UserRole = 10
)

// User user model
type User struct {
	Model
	ImageURL  string `json:"image_url"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// TableName override table name
func (User) TableName() string {
	return "users"
}
