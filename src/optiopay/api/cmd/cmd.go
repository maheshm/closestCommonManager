package cmd

import (
	"fmt"
	"optiopay/employee/create"
	"optiopay/employee/search_manager"
)

func Init() {
	fmt.Println("Init")

	claire := employee.Create("Claire")
	john := employee.Create("John")
	sarah := employee.Create("Sarah")

	four := employee.Create("4")
	five := employee.Create("5")

	six := employee.Create("6")
	seven := employee.Create("7")

	employee.Link(claire, john)
	employee.Link(claire, sarah)

	employee.Link(john, four)
	employee.Link(john, five)

	employee.Link(sarah, six)
	employee.Link(sarah, seven)

	search_manager.Setceo(claire)

	four_five := search_manager.Search_common_manager(four, five)
	six_john := search_manager.Search_common_manager(six, john)

	fmt.Printf("Four, Five - %s \nSix, John - %s\n", four_five.Name, six_john.Name)
}
