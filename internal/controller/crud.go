package controller

import (
	"controller/internal/entity"
	"errors"
	"fmt"
	"time"
)

func CreateOrder(id int, availablePizzas *entity.PizzaMap) entity.Order {
	choose := 0
	newOrder := &entity.Order{
		Id:          id,
		TimeCreated: time.Now(),
		Status:      "Created",
		CreatorName: "Aziz",
	}
	for i := 0; i <= 10; i++ {
		fmt.Scan(&choose)
		check := choose <= 4
		if check {
			if choose == 0 {
				break
			}
			newOrder.OrderedPizza = append(newOrder.OrderedPizza, choose)
			newOrder.OrderCost = newOrder.OrderCost + availablePizzas.P[choose].Cost
			newOrder.TimeToPrepare = newOrder.TimeToPrepare + availablePizzas.P[choose].TimeToPrepare
		} else {
			err := errors.New("нет такой позиции")
			fmt.Println(err)
			i--
		}
	}
	return *newOrder
}

func ChangeStatus(orders *entity.OrderMap, statuses *entity.StatusMap, orderKey int) error {
	err := validateKey(orders, orderKey)
	if err == nil {
		fmt.Println("На какой статус изменить")
		for k, v := range statuses.S {
			fmt.Println(k, v)
		}
		chooseStatus := 0
		fmt.Scan(&chooseStatus)
		if chooseStatus <= len(statuses.S) && chooseStatus >= 0 {
			if entry, ok := orders.O[orderKey]; ok {
				entry.Status = statuses.S[chooseStatus]
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

func validateKey(orders *entity.OrderMap, key int) error {
	if orders.O[key].Status != "" {
		return nil
	}
	err := errors.New("key doesnt exist")
	return err
}

func DeleteOrder(orders *entity.OrderMap, key int) {
	if validateKey(orders, key) != nil {
		delete(orders.O, key)
	}
}

func ListOrders(orders *entity.OrderMap) {
	for k := range orders.O {
		fmt.Print("№", k, " ")
	}
	fmt.Println()
}
