package types

type User struct {
	FirstName     string `json:"first_name" form:"first_name"`
	LastName      string `json:"last_name" form:"last_name"`
	Position      string `json:"position" form:"position"`
	Email         string `json:"email" form:"email"`
	Company       string `json:"company" form:"company"`
	License       string `json:"license_no" form:"license_no"`
	Country       string `json:"country" form:"country"`
	SignaturePath string `json:"signature" form:"signature"`
	EmployeeID    string `json:"employee_id" form:"employee_id"`
}
