package types

type User struct {
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	Company    string `json:"company"`
	Position   string `json:"position"`
	EmployeeID string `json:"employee_id"`
	Email      string `json:"email"`
	Country    string `json:"country"`
	License    string `json:"license"`
}
