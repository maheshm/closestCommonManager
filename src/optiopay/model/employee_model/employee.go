package employee_model

import (
	"optiopay/datastore/memory" //should have been an interface rather
	"optiopay/lib/lca"
	"strconv"
)

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

func build_data(data *Employee) *lca.Data {
	//this should be elsewhere
	employee_node := &lca.Data{Name: data.Name}
	for _, employee := range data.Manages {
		employee_child_node := build_data(employee)
		employee_node.Children = append(employee_node.Children, employee_child_node)
	}
	return employee_node
}

func Search_common_manager (ceo, employee1, employee2 *Employee) *Employee {
	data := build_data(ceo)
	lca.Init(data)
	common_manager := lca.Search(employee1.Name, employee2.Name)

	common_manager_obj := Find_by_name(common_manager)
	return common_manager_obj
}

func Delete() {}

func Find() {}

func build_employee_obj(data *memory.Entry) *Employee {
	employee_obj := &Employee{Id: data.Id, Name: data.Data["name"].(string)}
	for _, id := range data.Data["manages_employee_id"].([]int) {
		me_entry := memory.Get("Id", strconv.Itoa(id))
		me_obj := build_employee_obj(me_entry)
		employee_obj.Manages = append(employee_obj.Manages, me_obj)
	}
	return employee_obj
}

func Find_by_name(name string) *Employee{
	employee := memory.Get("name", name)
	employee_obj := build_employee_obj(employee)
	return employee_obj
}
