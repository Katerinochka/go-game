package main

import (
	"fmt"
	"strings"
)

type Room struct {
	Name     string
	Entrance []string
	Clothes  map[string][]Cloth
	Things   map[string][]Thing
}

type Human struct {
	Location Room
	PutOn    []Cloth
}

type Cloth struct {
	Name   string
	Things map[string]Thing
}

type Thing struct {
	Name string
}

type Rooms map[string]*Room

var Locations Rooms
var User Human

func NewLocation() *Rooms {
	locations := make(Rooms)
	return &locations
}

func NewRoom(name string, entrance []string) *Room {
	r := new(Room)
	r.Name = name
	r.Entrance = entrance
	r.Clothes = make(map[string][]Cloth)
	r.Things = make(map[string][]Thing)
	return r
}

func NewUser(room Room) *Human {
	u := new(Human)
	u.Location = room
	return u
}

func NewCloth(name string) *Cloth {
	c := new(Cloth)
	c.Name = name
	return c
}

func NewThing(name string) *Thing {
	t := new(Thing)
	t.Name = name
	return t
}

func (r *Rooms) AddRoom(name string, entrance ...string) {
	room := NewRoom(name, entrance)
	(*r)[name] = room
}

func (h *Human) GoTo(room string) bool {
	for _, v := range h.Location.Entrance {
		if v == room {
			h.Location = *Locations[room]
			return true
		}
	}
	return false
}

func (h Human) PrintCanGo() {
	fmt.Println("можно пойти в", strings.Join(h.Location.Entrance, ", "))
}

func (r *Room) AddClothes(where string, clothes ...string) {
	for _, cloth := range clothes {
		r.Clothes[where] = append(r.Clothes[where], *NewCloth(cloth))
	}
}

func (h Human) LookAround() {
	fmt.Print("ты находишься в ", h.Location.Name, ". ")
	for k, v := range h.Location.Clothes {
		if len(v) > 0 {
			fmt.Print(k, ": ", v)
		}
	}
	h.PrintCanGo()
}

func (h *Human) ToPutOn(cloth string) bool {
	for k, v := range h.Location.Clothes {
		for k1, v1 := range v {
			if v1.Name == cloth {
				h.PutOn = append(h.PutOn, v1)
				h.Location.Clothes[k] = append(h.Location.Clothes[k][:k1], h.Location.Clothes[k][k1+1:]...)
				return true
			}
		}
	}
	return false
}

func (r *Room) AddThing(where string, things ...string) {
	for _, thing := range things {
		r.Things[where] = append(r.Things[where], *NewThing(thing))
	}
}

func (h *Human) Take(thing string) {
	// if len(h.PutOn) == 0 {
	// 	fmt.Println("некуда положить")
	// 	return
	// }
	// for k, v := range h.Location.Things {
	// 	for k1, v1 := range v {
	// 		if v1.Name == thing {
	// 			h.PutOn[0].Things[thing] = *NewThing(thing)
	// 			h.Location.Things
	// 		}
	// 	}
	// }
}

func init() {
	Locations = *NewLocation()
	Locations.AddRoom("кухня", "коридор")
	Locations.AddRoom("коридор", "кухня", "комната", "улица")
	Locations.AddRoom("комната", "коридор")
	Locations["комната"].AddClothes("на стуле", "рюкзак")
	Locations["комната"].AddThing("на столе", "конспекты")
	Locations.AddRoom("улица", "коридор")
	User = *NewUser(*Locations["кухня"])
}

func main() {
	var command, parametr string

	for {
		fmt.Scan(&command)
		switch command {
		case "идти":
			fmt.Scan(&parametr)
			if User.GoTo(parametr) {
				User.PrintCanGo()
			} else {
				fmt.Println("нет такого")
			}
		case "осмотреться":
			User.LookAround()
		case "надеть":
			fmt.Scan(&parametr)
			if User.ToPutOn(parametr) {
				fmt.Println("Вы надели:", parametr)
			} else {
				fmt.Println("нет такого")
			}
		default:
			fmt.Println("неизвестная команда")
		}
	}
}
