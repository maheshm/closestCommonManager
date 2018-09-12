package employee_model

import "testing"

func TestCreate(t *testing.T) {
	employee := Create("Claire")
	if employee.Name != "Claire" {
		t.Errorf("Create was incorrect, got: %s, want %s", employee.Name, "Claire")
	}
}

func TestSave(t *testing.T) {
	employee := Create("Claire")
	employee_obj := employee.Save()
	if employee_obj.Id != 1 {
		t.Errorf("Save failed. id got: %d, want: %d", employee_obj.Id, 1)
	}

	employee2 := Create("Claire 2")
	employee_obj = employee2.Save()
	if employee_obj.Id != 2 {
		t.Errorf("Save failed. id got: %d, want: %d", employee_obj.Id, 2)
	}

}

func TestLink_employee(t *testing.T) {
	employee := Create("Claire")
	employee_obj := employee.Save()
	employee2 := Create("Claire 2")
	employee2_obj := employee2.Save()

	employee_obj.Link_employee(employee2_obj)

	if len(employee_obj.Manages) != 1 {
		t.Errorf("Link failed. lenght got: %d, want: %d", len(employee_obj.Manages), 1)
	}

}
