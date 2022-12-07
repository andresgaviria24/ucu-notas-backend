package dto

type Students struct {
	Id       int64    `json:"id"`
	Email    string   `json:"email"`
	FullName string   `json:"name"`
	Nota     []string `json:"notes"`
	Average  string   `json:"average"`
}

type CreateNote struct {
	Note string `json:"note"`
}
