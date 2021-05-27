package rest_todo

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"` // binding: "required" это для пакета gin (обязательное поле)
}
