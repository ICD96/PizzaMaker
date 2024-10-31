package entity

import "time"

type Order struct {
	Id            int
	TimeCreated   time.Time
	CreatorName   string
	Status        string
	OrderCost     int
	TimeToPrepare time.Duration
	OrderedPizza  []int
}
type pizza struct {
	Id            int
	Name          string
	Cost          int
	TimeToPrepare time.Duration
}
type PizzaMap struct {
	P map[int]pizza
}
type OrderMap struct {
	O map[int]Order
}
type StatusMap struct {
	S map[int]string
}

func NewStatusMap() *StatusMap {
	return &StatusMap{map[int]string{1: "Ожидание", 2: "В процессе", 3: "Готов", 4: "Отменен"}}
}
func NewOrderMap() *OrderMap {
	return &OrderMap{map[int]Order{}}
}

func NewPizzaMap() *PizzaMap {
	return &PizzaMap{map[int]pizza{
		1: {1, "Маргарита", 10, 5 * time.Minute},
		2: {2, "Пеперони", 15, 7 * time.Minute},
		3: {3, "Сырная", 18, 6 * time.Minute},
		4: {4, "Мясная", 20, 8 * time.Minute},
	}}
}
