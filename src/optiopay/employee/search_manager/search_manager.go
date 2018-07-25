package search_manager

//yet to be implemented
import (
	"optiopay/model/employee_model"
)

var ceo_of_company *employee_model.Employee

func Setceo(ceo *employee_model.Employee) {
	ceo_of_company = ceo
}

func Search_common_manager (employee1, employee2 *employee_model.Employee) *employee_model.Employee {
	manager := employee_model.Search_common_manager(ceo_of_company, employee1, employee2)
	return manager
}
