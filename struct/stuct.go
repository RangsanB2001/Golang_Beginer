package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {
	// employee1 := employee{
	// 	employeeID:   "101",
	// 	employeeName: "Tuu",
	// 	phone:        "002818218",
	// }
	// fmt.Println("employee1 =", employee1)

	employeeList := []employee{}
	employee1 := employee{
		employeeID:   "101",
		employeeName: "tuu",
		phone:        "0292922202",
	}

	employeeList = append(employeeList, employee1)

	// employeeList[0] = employee{
	// 	employeeID:   "101",
	// 	employeeName: "tuu",
	// 	phone:        "0292922202",
	// }
	// employeeList[1] = employee{
	// 	employeeID:   "102",
	// 	employeeName: "tuuPoom",
	// 	phone:        "0292239202",
	// }
	// employeeList[2] = employee{
	// 	employeeID:   "103",
	// 	employeeName: "tuuKuy",
	// 	phone:        "02929223202",
	// }
	fmt.Println("employee = ", employeeList)
}
