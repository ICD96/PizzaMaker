package main

import (
	"fmt"
	"os"
	"utils/utils"
)

var choose int

func main() {
	orders := utils.NewOrderMap()
	availablePizzas := utils.NewPizzaMap()
	statuses := utils.NewStatusMap()
	idCount := 0
	listMenu(orders, availablePizzas, statuses, idCount)
}

func listMenu(orders *utils.OrderMap, availablePizzas *utils.PizzaMap, statuses *utils.StatusMap, idCount int) {
	for {
		fmt.Println("1. Сделать заказ")
		fmt.Println("2. Изменить статус заказа")
		fmt.Println("3. Удалить заказ")
		fmt.Println("4. Показать имеющиеся заказы")
		fmt.Println("0. Выход")
		fmt.Scan(&choose)
		switch choose {
		case 1:
			fmt.Println("Какую пиццу закажите? (Можно заказать не более 10)")
			fmt.Println("1. Маргарита")
			fmt.Println("2. Пепперони")
			fmt.Println("3. Сырная")
			fmt.Println("4. Мясная")
			fmt.Println("0. Закончить добавление")
			orders.O[idCount] = utils.CreateOrder(idCount, availablePizzas)
			idCount++
		case 2:
			fmt.Println("У какого заказа изменить статус")
			utils.ListOrders(orders)
			chooseOrder := 0
			fmt.Scan(&chooseOrder)
			utils.ChangeStatus(orders, statuses, chooseOrder)
		case 3:
			fmt.Println("Какой заказ удалить")
			utils.ListOrders(orders)
			fmt.Scan(&choose)
			delete(orders.O, choose)
		case 4:
			utils.ListOrders(orders)
		case 0:
			os.Exit(1)
		}
	}
}
