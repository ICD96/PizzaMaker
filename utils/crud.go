package utils

import (
	"errors"
	"fmt"
	"time"
)

func CreateOrder(id int, availablePizzas *PizzaMap) Order {
	choose := 0
	newOrder := &Order{
		id:          id,
		timeCreated: time.Now(),
		status:      "Created",
		creatorName: "Aziz",
	}
	for i := 0; i <= 10; i++ {
		fmt.Scan(&choose)
		check := choose <= 4
		if check {
			if choose == 0 {
				break
			}
			newOrder.orderedPizza = append(newOrder.orderedPizza, availablePizzas.p[choose])
			newOrder.orderCost = newOrder.orderCost + availablePizzas.p[choose].cost
			newOrder.timeToPrepare = newOrder.timeToPrepare + availablePizzas.p[choose].timeToPrepare
		} else {
			err := errors.New("нет такой позиции")
			fmt.Println(err)
			i--
		}
	}
	return *newOrder
}

func ChangeStatus(orders *OrderMap, statuses *StatusMap, orderKey int) error {
	err := validateKey(orders, orderKey)
	if err == nil {
		fmt.Println("На какой статус изменить")
		for k, v := range statuses.s {
			fmt.Println(k, v)
		}
		chooseStatus := 0
		fmt.Scan(&chooseStatus)
		if chooseStatus <= len(statuses.s) && chooseStatus >= 0 {
			if entry, ok := orders.O[orderKey]; ok {
				entry.status = statuses.s[chooseStatus]
				orders.O[orderKey] = entry
			}
		} else {
			err := errors.New("не правильный выбор")
			fmt.Println(err)
			return err
		}
	} else {
		fmt.Println(err)
		return err
	}
	return nil
}

func validateKey(orders *OrderMap, key int) error {
	if orders.O[key].status != "" {
		return nil
	}
	err := errors.New("key doesnt exist")
	return err
}

func DeleteOrder(orders *OrderMap, key int) {
	if validateKey(orders, key) != nil {
		delete(orders.O, key)
	}
}

func ListOrders(orders *OrderMap) {
	for k, _ := range orders.O {
		fmt.Print("№", k, " ")
	}
}
