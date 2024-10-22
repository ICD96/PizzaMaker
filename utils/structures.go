package utils

import "time"

type Order struct {
	id            int
	timeCreated   time.Time
	creatorName   string
	status        string
	orderCost     float64
	orderedPizza  []pizza
	timeToPrepare time.Duration
}
type pizza struct {
	id            int
	name          string
	cost          float64
	timeToPrepare time.Duration
}
type pizzaMap struct {
	p map[int]pizza
}
type orderMap struct {
	o map[int]Order
}
type statusMap struct {
	s map[int]string
}

func NewStatusMap() *statusMap {
	return &statusMap{map[int]string{1: "Ожидание", 2: "В процессе", 3: "Готов", 4: "Отменен"}}
}
func NewOrderMap() orderMap {
	return orderMap{map[int]Order{}}
}

func NewPizzaMap() *pizzaMap {
	return &pizzaMap{map[int]pizza{
		1: {1, "Маргарита", 10, 5 * time.Minute},
		2: {2, "Пеперони", 15, 7 * time.Minute},
		3: {3, "Сырная", 18, 6 * time.Minute},
		4: {4, "Мясная", 20, 8 * time.Minute},
	}}
}
