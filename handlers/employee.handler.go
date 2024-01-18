package handlers

import (
	"fmt"
	"strconv"
	"strings"
	"workshop/go-fiber/database"
	"workshop/go-fiber/models"
	"workshop/go-fiber/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func CreateEmployee(c *fiber.Ctx) error {
	log.Info("IN CREATE EMPS")
	db := database.DBConn
	var employee models.Employees
	var employeeResponse models.EmployeeResponse

	if err := c.BodyParser(&employee); err != nil {
		return response.BuidErrorResponse(c, 503, err.Error(), nil)
	}

	db.Create(&employee).Scan(&employeeResponse)
	return response.BuidSuccessResponse(c, 201, "Successfully Create Employees", employeeResponse)
}

func UpdateEmployee(c *fiber.Ctx) error {
	log.Info("IN UPDATE EMPS")
	db := database.DBConn
	var employeeUpdate models.Employees
	var employeeResponse models.EmployeeUpdatedResponse

	id := c.Params("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return response.BuidErrorResponse(c, 503, err.Error(), nil)
	}

	if err := c.BodyParser(&employeeUpdate); err != nil {
		return response.BuidErrorResponse(c, 503, err.Error(), employeeUpdate)
	}

	result := db.Where("id = ?", idInt).Updates(&employeeUpdate).Scan(&employeeResponse)

	if result.RowsAffected == 0 {
		return response.BuidErrorResponse(c, 404, fmt.Sprintf("Employee ID %d not found", idInt), nil)
	}

	return response.BuidSuccessResponse(c, 200,
		fmt.Sprintf("Employee ID %d is updated successfully", idInt), employeeResponse)
}

func GetEmployee(c *fiber.Ctx) error {
	log.Info("IN GET EMPS")
	db := database.DBConn
	var employeeResponse []models.EmployeeResponse

	result := db.Model(&models.Employees{}).Find(&employeeResponse)
	if result.RowsAffected == 0 {
		return response.BuidSuccessResponse(c, 404, "Employee data not found", employeeResponse)
	}

	return response.BuidSuccessResponse(c, 200, "Success", employeeResponse)
}

func GetEmployeeById(c *fiber.Ctx) error {
	log.Info("IN GET EMP ID")
	db := database.DBConn
	var employeeResponse models.EmployeeResponse

	id := c.Params("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return response.BuidErrorResponse(c, 503, err.Error(), nil)
	}

	result := db.Where("id = ?", idInt).Model(&models.Employees{}).Find(&employeeResponse)
	if result.RowsAffected == 0 {
		return response.BuidErrorResponse(c, 404, fmt.Sprintf("Employee ID %d is not found", idInt), nil)
	}

	return response.BuidSuccessResponse(c, 200, "Success", employeeResponse)
}

func RemoveEmployeeById(c *fiber.Ctx) error {
	log.Info("IN REMOVE EMP")
	db := database.DBConn
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return response.BuidErrorResponse(c, 503, err.Error(), nil)
	}

	result := db.Delete(idInt)
	if result.RowsAffected == 0 {
		return response.BuidErrorResponse(c, 404, fmt.Sprintf("Employee ID %d is not found", idInt), nil)
	}

	return response.BuidSuccessResponse(c, 200, fmt.Sprintf("Employee ID %d is deleted successfully", idInt), nil)
}

func SearchEmployee(c *fiber.Ctx) error {
	log.Info("IN SEARCFH EMPS")
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var employeeResponse []models.EmployeeResponse
	result := db.Model(&models.Employees{}).
		Where("name LIKE '" + "%" + search + "%'" +
			"OR lastname LIKE " + "'%" + search + "%'" + " OR employee_id LIKE" + "'%" + search + "%'").Find(&employeeResponse)
	if result.RowsAffected == 0 {
		return response.BuidErrorResponse(c, 404, "No data found", nil)
	}

	return response.BuidSuccessResponse(c, 200, "Success", employeeResponse)
}

func GetEmployeeGeneration(c *fiber.Ctx) error {
	db := database.DBConn
	var employeeResponse []models.EmployeeResponse

	result := db.Model(&models.Employees{}).Find(&employeeResponse)

	if result.RowsAffected == 0 {
		return response.BuidErrorResponse(c, 404, "No data found", employeeResponse)
	}

	giGenTotal := 0
	babyBoomerTotal := 0
	genX := 0
	genY := 0
	genZ := 0

	var datas []models.EmployeeGeneration

	for _, item := range employeeResponse {
		var employee models.EmployeeGeneration
		generation := ""
		if item.Age > 75 {
			generation = "G.I. Generation"
			giGenTotal += 1
		} else if item.Age >= 57 && item.Age <= 75 {
			generation = "Baby Boomer"
			babyBoomerTotal += 1
		} else if item.Age >= 42 && item.Age <= 56 {
			generation = "GenX"
			genX += 1
		} else if item.Age >= 24 && item.Age <= 41 {
			generation = "GenY"
			genY += 1
		} else {
			generation = "GenZ"
			genZ += 1
		}

		employee = models.EmployeeGeneration{
			ID:         item.ID,
			Name:       item.Name,
			Lastname:   item.Lastname,
			Age:        item.Age,
			Email:      item.Email,
			Generation: generation,
		}

		datas = append(datas, employee)
	}

	var apiResp models.EmployeeGenerationResponse

	apiResp = models.EmployeeGenerationResponse{
		EmployeeGenration: datas,
		GenZTotal:         genZ,
		GenYTotal:         genY,
		GenXTotal:         genX,
		BabyBoomerTotal:   babyBoomerTotal,
		GITotal:           giGenTotal,
	}

	return response.BuidSuccessResponse(c, 200, "success", apiResp)
}
