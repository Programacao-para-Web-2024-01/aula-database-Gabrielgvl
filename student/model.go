package student

type Student struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
