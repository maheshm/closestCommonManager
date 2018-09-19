package employee

import "ccm/model/employee_model"

func Create(name string) *employee_model.Employee {
	employee := employee_model.Create(name)
	return employee.Save()
}

func Link(manager, employee *employee_model.Employee) {
	manager.Link_employee(employee)
}
