package main

import "fmt"

type EmployeeRow struct {
	ID        int
	Name      string
	Position  string
	Salary    int
	ManagerID int
}
type EmployeeDB []EmployeeRow

func NewEmployeeDB() *EmployeeDB {
	return &EmployeeDB{}
}

func (db *EmployeeDB) Where(id int) *EmployeeRow {
	for i := 0; i < len(*db); i++ {
		if (*db)[i].ID == id {
			return &(*db)[i]
		}
	}
	return nil
}

func (db *EmployeeDB) Insert(name string, position string, salary int, managerID int) {
	(*db) = append(*db, EmployeeRow{
		ID:        len(*db) + 1,
		Name:      name,
		Position:  position,
		Salary:    salary,
		ManagerID: managerID,
	})
}

func (db *EmployeeDB) Update(id int, name string, position string, salary int, managerID int) {
	// TODO: answer here
	for idx, row := range *db {
		if row.ID == id {
			emp := (*db)[idx]
			emp.Name = name
			emp.Position = position
			emp.Salary = salary
			emp.ManagerID = managerID
			return
		}
	}

}

func (db *EmployeeDB) Delete(id int) {
	// TODO: answer here
	for idx, row := range *db {
		if row.ID == id {
			(*db) = append((*db)[:idx], (*db)[idx+1:]...)
			return
		}
	}
}
