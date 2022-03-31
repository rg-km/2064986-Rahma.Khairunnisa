package main

import (
	"time"
)

type Item struct {
	Title    string
	Deadline time.Time
}

type Todos struct {
	items []Item
}

func (todos *Todos) Add(item Item) {
	todos.items = append(todos.items, item)
}

func (todos *Todos) GetAll() []Item {
	return todos.items // TODO: replace this
}

func (todos *Todos) GetUpcoming() []Item {
	var itemUpComing []Item
	for _, v := range todos.items {
		if time.Now().Add(+time.Hour) == v.Deadline {
			itemUpComing = append(itemUpComing, v)
		}
	}
	return itemUpComing
}

func NewItem(title string, deadline time.Time) Item {
	return Item{title, deadline}
}

func NewTodos() Todos {
	return Todos{}
}
