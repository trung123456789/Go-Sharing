package structdemo

// User struct
type User struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
	Age  *int    `json:"age"`
}

// UserInfo struct
type UserInfo struct {
	UserID uint `gorm:"primary_key"`
	Name   string
	Age    int
}
