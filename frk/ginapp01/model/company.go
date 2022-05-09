package model

type Company struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

// https://www.cnblogs.com/haima/p/12849729.html
func (Company) TableName() string {
	return "company"
}
