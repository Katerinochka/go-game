package main

import (
	"errors"
	"fmt"
)

var Locations Rooms
var User Human

var ErrThisIsNot = errors.New("здесь нет такого")
var ErrNotAllDone = errors.New("здесь остались незаконченные дела.")
var ErrNoCapacity = errors.New("некуда положить")
var ErrNotTake = errors.New("вы не взяли этот предмет")

func init() {
	Locations = *NewLocation()
	Locations.AddRoom("кухня", "коридор")
	Locations.AddRoom("коридор", "кухня", "комната", "улица")
	Locations["коридор"].AppliedItems["дверь"] = NewAppliedItem("дверь", 1, "улица")
	Locations.AddRoom("комната", "коридор")
	Locations["комната"].AddClothes(map[string]int{
		"рюкзак": 10,
	})
	Locations["комната"].AddThing(
		NewThing("конспекты", 1),
		NewThing("ключи", 1))
	Locations.AddRoom("улица", "коридор")
	User = *NewUser(*Locations["кухня"])
}

func main() {
	var command, parametr1, parametr2 string

	for {
		fmt.Scan(&command)
		switch command {
		case "идти":
			fmt.Scan(&parametr1)
			if err := User.GoTo(parametr1); err != nil {
				fmt.Println(err)
			} else {
				User.PrintCanGo()
			}
		case "осмотреться":
			User.LookAround()
		case "надеть":
			fmt.Scan(&parametr1)
			if err := User.ToPutOn(parametr1); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Вы надели:", parametr1)
			}
		case "взять":
			fmt.Scan(&parametr1)
			if err := User.ToTake(parametr1); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Вы взяли:", parametr1)
			}
		case "применить":
			fmt.Scan(&parametr1)
			fmt.Scan(&parametr2)
			if err := User.Apply(parametr1, parametr2); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("получилось")
			}
		default:
			fmt.Println("неизвестная команда")
		}
	}
}
