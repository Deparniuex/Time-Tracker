package entity

type User struct {
	ID             int64  `json:"id"`
	PassportSerie  int    `json:"passport_serie"`
	PassportNumber int    `json:"passport_number"`
	First_name     string `json:"name"`
	Surname        string `json:"surname"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"address"`
}
