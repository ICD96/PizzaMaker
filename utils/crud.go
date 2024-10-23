package utils

import (
	"errors"
	"fmt"
	"os"
	"time"
)

var idCount int

func Menu(orders *OrderMap, availablePizzas *PizzaMap, statuses *StatusMap) {
	choose := 0
	var pizzaToOrder []int
	chooseOrder := 0
	fmt.Scan(&choose)
	switch choose {
	case 1:
		fmt.Println("Какую пиццу закажите? (Можно заказать не более 10)")
		fmt.Println("1. Маргарита")
		fmt.Println("2. Пепперони")
		fmt.Println("3. Сырная")
		fmt.Println("4. Мясная")
		fmt.Println("0. Закончить добавление")
		for i := 0; i <= 10; i++ {
			fmt.Scan(&choose)
			check := choose <= 4
			if check {
				if choose == 0 {
					break
				}
				pizzaToOrder = append(pizzaToOrder, choose)
			} else {
				err := errors.New("нет такой позиции")
				fmt.Println(err)
				i--
			}
		}
		orders.o[idCount] = CreateOrder(idCount, pizzaToOrder, availablePizzas)
		idCount++
	case 2:
		fmt.Println("У какого заказа изменить статус")
		listOrders(orders)
		fmt.Scan(&chooseOrder)
		if orders.o[chooseOrder].status != "" {
			fmt.Println("На какой статус изменить")
			for k, v := range statuses.s {
				fmt.Println(k, v)
			}
			chooseStatus := 0
			fmt.Scan(&chooseStatus)
			if choose <= 4 && choose >= 0 {
				if entry, ok := orders.o[chooseOrder]; ok { // Не очень понял почему именно так нужно делать, надо потом разобраться
					entry.status = statuses.s[chooseStatus]
					orders.o[chooseOrder] = entry
				}
			} else {
				err := errors.New("не правильный выбор")
				fmt.Println(err)
			}
		} else {
			fmt.Println("такого заказа нет")
		}
	case 3:
		fmt.Println("Какой заказ удалить")
		listOrders(orders)
		fmt.Scan(&chooseOrder)
		delete(orders.o, chooseOrder)
	case 0:
		os.Exit(1)
	}
}

func CreateOrder(id int, pizzas []int, availablePizzas *PizzaMap) Order {
	newOrder := &Order{
		id:          id,
		timeCreated: time.Now(),
		status:      "Created",
		creatorName: "Aziz",
	}
	for i := range pizzas {
		newOrder.orderedPizza = append(newOrder.orderedPizza, availablePizzas.p[i])
		newOrder.orderCost = newOrder.orderCost + availablePizzas.p[i].cost
		newOrder.timeToPrepare = newOrder.timeToPrepare + availablePizzas.p[i].timeToPrepare
	}
	return *newOrder
}

func listOrders(orders *OrderMap) {
	for k, _ := range orders.o {
		fmt.Print("№", k, " ")
	}
}
