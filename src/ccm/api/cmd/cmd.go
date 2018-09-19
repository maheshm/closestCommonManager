package cmd

import (
	"fmt"
	"ccm/employee/create"
	"ccm/model/employee_model"
	"ccm/employee/search_manager"
	"github.com/manifoldco/promptui"
)

func ask(question string) string {
	prompt := promptui.Prompt{ Label: question, Validate: nil }

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

func get_employees_of_manager(manager *employee_model.Employee) {
	question := "Enter names of people who directly report to "+ manager.Name + " "
	for ans := ask(question); ans != "q"; ans = ask(question) {
		emp := employee.Create(ans)
		employee.Link(manager, emp)
	}

	for _, emp := range manager.Manages {
		get_employees_of_manager(emp)
	}
}

func search(ceo *employee_model.Employee) {
	cont := true
	for ; cont ; {
		emp1 := ask("Enter the name of first employee ")
		emp2 := ask("Enter the name of second employee ")
		search_manager.Setceo(ceo)

		emp1_obj := employee_model.Find_by_name(emp1)
		emp2_obj := employee_model.Find_by_name(emp2)
		manager := search_manager.Search_common_manager(emp1_obj, emp2_obj)

		fmt.Printf("The closest common Manager of %v and %v is %v\n\n", emp1, emp2, manager.Name)

		wish_to_continue := ask("Do you wish to continue? (y/n) ")
		if wish_to_continue == "n" {
			cont = false
		}
	}
}

func Init() {
	ceo := employee.Create(ask("Enter the name of CEO "))

	fmt.Println("Enter `q` to skip adding more")
	get_employees_of_manager(ceo)

	search(ceo)
}
