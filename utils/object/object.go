package object

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func AddSalary(e *Employee, add int) {
	e.Salary = add + e.Salary
}

func TestObject() {
	corey := Employee{
		ID:        1,
		Name:      "Corey",
		Address:   "123 fake st.",
		DoB:       time.Date(1980, 10, 1, 0, 0, 0, 0, time.UTC),
		Position:  "dev",
		Salary:    10000,
		ManagerID: 0,
	}
	AddSalary(&corey, 1000)
	fmt.Println(corey)

	corey.Salary = corey.Salary + 5000
	fmt.Println(corey)
}
