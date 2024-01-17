package models

import (
	"time"
)

type Employees struct {
	IdModel
	EmployeesModel
}

type EmployeesModel struct {
	EmployeeID int       `json:"employee_id" gorm:"->created"`
	Name       string    `json:"name"`
	Lastname   string    `json:"lastname"`
	Birthday   time.Time `json:"birthday"`
	Age        int       `json:"age"`
	Email      string    `json:"email" validate:"required,email"`
}

type EmployeeResponse struct {
	EmployeesModel
}

type EmployeeCreateRequest struct {
	EmployeesModel
}

type EmployeeUpdateRequest struct {
	Name     string    `json:"name"`
	Lastname string    `json:"lastname"`
	Birthday time.Time `json:"birthday"`
	Email    string    `json:"email"`
}
