package main

import "fmt"

type Item struct {
	id    int
	name  string
	price int
}

type ItemNotFoundErrror struct {
	Id int
}

func (i *ItemNotFoundErrror) Error() string {
	return fmt.Sprintf("Id : %d Not Found", i.Id)
}

var items = []Item{
	{1, "MacBook", 11000000},
	{2, "Watch", 300000},
	{3, "Samsung S24 FE", 10000000},
}

func getItemById(id int) (Item, error) {
	for _, item := range items{
		if item.id == id {
			return item, nil
		}
	}
	return	Item{}, &ItemNotFoundErrror{Id: id}
}

func main() {
	result, err := getItemById(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}