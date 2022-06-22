package main

import "fmt"

type SalaryCal interface {
	CalSalary() int
}

type perm struct {
	empId int
	basic int
	pf    int
}

type contract struct {
	empId int
	basic int
}

type freelancer struct {
	empId int
	ratePerHour int
	hours int
}

func (p perm) CalSalary() int {
	return p.basic + p.pf
}

func (c contract) CalSalary() int {
	return c.basic
}

func (f freelancer) CalSalary() int {
	return f.hours * f.ratePerHour
}

func totalExpense(emps []SalaryCal) {
	total := 0
	for _, emp := range emps {
		total += emp.CalSalary()
	}
	fmt.Println(total)
}
func main() {
	p1 := perm{
		empId: 1,
		basic: 1000,
		pf:    100,
	}
	p2 := perm{
		empId: 2,
		basic: 2000,
		pf:    200,
	}
	c1 := contract{
		empId: 3,
		basic: 3000,
	}
	f1 := freelancer{
		empId: 4,
		ratePerHour: 10,
		hours: 20,
	}
	totalExpense([]SalaryCal{p1, p2, c1,f1})
}
