package four

import (
	"fmt"
	"math/rand"
)

var customerMap map[int]*Customer = make(map[int]*Customer)

type Customer struct {
	name   string
	birth  int
	income int
}

func insertCustomer(name string, birth int, income int) *Customer {
	customer := &Customer{name, birth, income}

	id := rand.Int()
	customerMap[id] = customer

	return customer
}

func getCustomer(id int) *Customer {
	return customerMap[id]
}

func deleteCustomer(id int) {
	delete(customerMap, id)
}

func showCustomerByBirth() {
	for id, customer := range customerMap {
		if customer.birth >= 1980 && customer.birth <= 2000 {
			fmt.Printf("%d ---> %v\n", id, &customer)
		}
	}
}

func showCustomersByAverageIncome() {
	sum := 0
	for _, customer := range customerMap {
		sum += customer.income
	}

	averageIncome := sum / len(customerMap)

	for id, customer := range customerMap {
		if customer.income >= averageIncome {
			fmt.Printf("%d ---> %v\n", id, &customer)
		}
	}
}
