package dto

type EmployeeResponse struct {
	Name     string `json:"name"`
	Age      int    `json:"age" `
	JobTitle string `json:"job_title"`
	Company  string `json:"company"`
}

type EmployeeRequest struct {
	Name     string `json:"name"`
	Age      int    `json:"age" `
	JobTitle string `json:"job_title"`
	Company  string `json:"company"`
}
