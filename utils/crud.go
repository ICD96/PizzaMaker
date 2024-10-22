package utils

import (
	"errors"
	"fmt"
	"time"
)

func listMenu() {
	fmt.Println("1. Сделать заказ")
	fmt.Println("2. Изменить статус заказа")
	fmt.Println("3. Удалить заказ")
	fmt.Println("0. Выход")
	menu()
}
func menu() {
	choose := 0
	idCount := 0
	var pizzaToOrder []int
	orders := NewOrderMap()
	availablePizzas := NewPizzaMap()
	//statuses := NewStatusMap()
	for {
		listMenu()
		fmt.Scan(choose)
		switch choose {
		case 1:
			fmt.Println("Какую пиццу закажите? (Можно заказать не более 10)")
			fmt.Println("1. Маргарита")
			fmt.Println("2. Пепперони")
			fmt.Println("3. Сырная")
			fmt.Println("4. Мясная")
			fmt.Println("0. Закончить добавление")
			for i := 0; i <= 10; i++ {
				fmt.Scan(choose)
				if choose <= 4 && choose != 0 {
					pizzaToOrder = append(pizzaToOrder, choose)
				} else {
					err := errors.New("такой позиции нет")
					fmt.Println(err)
				}
				if choose == 0 {
					break
				}
			}
			orders.o[idCount] = CreateOrder(idCount, pizzaToOrder, availablePizzas)
		case 2:
			fmt.Println("У какого заказа изменить статус")
			chooseOrder := 0
			fmt.Scan(chooseOrder)
			if orders.o[chooseOrder].status != "" {
				fmt.Println("На какой статус изменить")
				fmt.Println("1. Ожидание")
				fmt.Println("2. В процессе")
				fmt.Println("3. Готов")
				fmt.Println("4. Изменен")
				chooseStatus := 0
				fmt.Scan(chooseStatus)
				if choose <= 4 {

				}
			}
		}
	}
}
func CreateOrder(id int, pizzas []int, availablePizzas *pizzaMap) Order {
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
