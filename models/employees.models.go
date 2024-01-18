package models

import (
	"time"
)

type Employees struct {
	IdModel
	EmployeesModel
}

type EmployeesModel struct {
	EmployeeID string    `json:"employee_id" gorm:"<-:created"`
	Name       string    `json:"name"`
	Lastname   string    `json:"lastname"`
	Birthday   time.Time `json:"birthday"`
	Age        int       `json:"age"`
	Email      string    `json:"email"`
	Tel        string    `json:"tel"`
}

type EmployeeResponse struct {
	ID int `json:"id"`
	EmployeesModel
}

type EmployeeUpdatedResponse struct {
	ID        int       `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
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
	Tel      string    `json:"tel"`
}

type EmployeeGeneration struct {
	ID         int    `json:"id"`
	Age        int    `json:"age"`
	Generation string `json:"generation"`
	Name       string `json:"name"`
	Lastname   string `json:"lastname"`
	Email      string `json:"email"`
}

type EmployeeGenerationResponse struct {
	EmployeeGenration []EmployeeGeneration `json:"employee_generation"`
	GenZTotal         int                  `json:"gen_z_total"`
	GenYTotal         int                  `json:"gen_y_total"`
	GenXTotal         int                  `json:"gen_x_total"`
	BabyBoomerTotal   int                  `json:"baby_boomer_total"`
	GITotal           int                  `json:"gi_generation_total"`
}
