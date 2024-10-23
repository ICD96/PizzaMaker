package main

import (
	"fmt"
	"utils/utils"
)

func main() {
	orders := utils.NewOrderMap()
	availablePizzas := utils.NewPizzaMap()
	statuses := utils.NewStatusMap()
	listMenu(orders, availablePizzas, statuses)
}

func listMenu(orders *utils.OrderMap, availablePizzas *utils.PizzaMap, statuses *utils.StatusMap) {
	for {
		fmt.Println("1. Сделать заказ")
		fmt.Println("2. Изменить статус заказа")
		fmt.Println("3. Удалить заказ")
		fmt.Println("4. Показать имеющиеся заказы")
		fmt.Println("0. Выход")
		utils.Menu(orders, availablePizzas, statuses)
	}
}
