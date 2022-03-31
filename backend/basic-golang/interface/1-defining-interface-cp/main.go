package main

import "fmt"

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk membuat interface dan mengimplementasikan interface.
// Buatlah interface Employee yang memiliki method signature GetBonus() int
// Buatlah implementasi interface Employee pada objek Manager, SeniorEngineer, dan JuniorEngineer.
// Pada objek Manager, SeniorEngineer, dan JuniorEngineer memiliki satu atribut yaitu BaseSalary.
// Untuk menghitung bonus terdapat tiga aturan sebagai berikut:
// Bonus untuk Manager adalah 3 * BaseSalary
// Bonus untuk SeniorEngineer adalah 2 * BaseSalary
// Bonus untuk JuniorEngineer adalah 3 * BaseSalary

// TODO: answer here
type Employee interface {
	FindByName(name string) (Employee, error)
	FindByOrigin(origin string) ([] Employee, error)
}
 
type Employee struct {
	Name string
	Origin string
}

type EmployeeWithMap struct {
	data map[string]Employee
}

func Employee (employees []Employee) int {
	// Hitunglah total bonus yang dikeluarkan dari list of Employee
	// TODO: answer here
	fmt.Println("FindByName from Employee")
	// implementations..
	return Employee{}, nil
}

func main() {
	// Buatlah objek konkret untuk masing-masing objek dan panggil function TotalEmployeeBonus. Print total bonus untuk semua employee.
	// TODO: answer here
	var Employee Employee

	// Employee menggunakan map
	Employee =EmployeeWithMap{
		data: map[string]Employee{},
	}
	Employee.FindByName("Manager")

	// Employee menggunakan MySQL
	EmployeeWithMySQL := 
	Employee = EmployeeWithMySQL{}


}
