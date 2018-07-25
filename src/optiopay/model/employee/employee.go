package employee_model

import "optiopay/datastore/memory" //should have been an interface rather


type Employee struct {
	Id int
	Name string
	Manages []*Employee
}

func Create(name string) *Employee {
	employee := &Employee{Name: name}
	return employee
}

func (employee_obj *Employee) Save() *Employee {
	var data map[string]interface{}
	data = make(map[string]interface{})

	data["name"] = employee_obj.Name //Not checking if names are duplicates
	manages_employee_id := []int{}
	for _, employee := range employee_obj.Manages {
		manages_employee_id = append(manages_employee_id, employee.Id)
	}
	data["manages_employee_id"] = manages_employee_id
	if employee_obj.Id == 0 {
		saved_data := memory.Set(data)
		employee_obj.Id = saved_data.Id
	} else {
		memory.Update(employee_obj.Id, data)
	}

	return employee_obj
}

func (manager_obj *Employee) Link_employee(employee_obj *Employee) {
	manager_obj.Manages = append(manager_obj.Manages, employee_obj)
	manager_obj.Save()
}


func Delete() {}

func Find() {}

func Find_by_name(){}
