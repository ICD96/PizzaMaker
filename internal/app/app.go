package app

import (
	"controller/internal/controller"
	"controller/internal/entity"
	"controller/internal/usecase/repo"
	"fmt"
	"os"
)

var choose int

func Run() {
	orders := entity.NewOrderMap()
	availablePizzas := entity.NewPizzaMap()
	statuses := entity.NewStatusMap()
	idCount := 1
	listMenu(orders, availablePizzas, statuses, idCount)
}

func listMenu(orders *entity.OrderMap, availablePizzas *entity.PizzaMap, statuses *entity.StatusMap, idCount int) {
	for {
		fmt.Println("1. Сделать заказ")
		fmt.Println("2. Изменить статус заказа")
		fmt.Println("3. Удалить заказ")
		fmt.Println("4. Показать имеющиеся заказы")
		fmt.Println("5. Reader")
		fmt.Println("6. Save")
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
			orders.O[idCount] = controller.CreateOrder(idCount, availablePizzas)
			idCount++
		case 2:
			fmt.Println("У какого заказа изменить статус")
			controller.ListOrders(orders)
			chooseOrder := 0
			fmt.Scan(&chooseOrder)
			controller.ChangeStatus(orders, statuses, chooseOrder)
		case 3:
			fmt.Println("Какой заказ удалить")
			controller.ListOrders(orders)
			fmt.Scan(&choose)
			delete(orders.O, choose)
		case 4:
			controller.ListOrders(orders)
		case 5:
			repo.Read()
		case 6:
			repo.Save(orders)
		case 0:
			os.Exit(1)
		}
	}
}
