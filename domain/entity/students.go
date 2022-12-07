package entity

type Students struct {
	Id       int64  `gorm:"PRIMARY_KEY;NOT NULL;TYPE:bigint;COLUMN:id" json:"id"`
	Email    string `gorm:"NOT NULL;COLUMN:email" json:"email"`
	FullName string `gorm:"NOT NULL;COLUMN:full_name" json:"full_name"`
	Nota     string `gorm:"NOT NULL;COLUMN:nota" json:"note"`
}
