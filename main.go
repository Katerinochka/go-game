package main

import (
	"fmt"
	"strings"

	"golang.org/x/exp/maps"
)

type Room struct {
	Name     string
	Entrance []string
	Clothes  map[string]Cloth
	Things   map[string]Thing
}

type Human struct {
	Location Room
	PutOn    []Cloth
	Weight   int
	Take     []Thing
}

type Cloth struct {
	Name   string
	Weight int
}

type Thing struct {
	Name   string
	Where  string
	Weight int
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
	r.Clothes = make(map[string]Cloth)
	r.Things = make(map[string]Thing)
	return r
}

func NewUser(room Room) *Human {
	u := new(Human)
	u.Location = room
	u.Weight = 0
	return u
}

func NewCloth(name string, weight int) *Cloth {
	c := new(Cloth)
	c.Name = name
	c.Weight = weight
	return c
}

func NewThing(name string, weight int) *Thing {
	t := new(Thing)
	t.Name = name
	t.Weight = weight
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

func (r *Room) AddClothes(clothes map[string]int) {
	for cloth, weight := range clothes {
		r.Clothes[cloth] = *NewCloth(cloth, weight)
	}
}

func (h Human) LookAround() {
	fmt.Print("ты находишься в ", h.Location.Name, ". ")
	if len(h.Location.Clothes) == 0 && len(h.Location.Things) == 0 {
		fmt.Print("ничего интересного. ")
	} else {
		keys := maps.Keys(h.Location.Clothes)
		keys = append(keys, maps.Keys(h.Location.Things)...)
		fmt.Print("здесь есть: ", strings.Join(keys, ", "), ". ")
	}
	h.PrintCanGo()
}

func (h *Human) ToPutOn(cloth string) bool {
	if v, ok := h.Location.Clothes[cloth]; ok {
		h.PutOn = append(h.PutOn, v)
		h.Weight += v.Weight
		delete(h.Location.Clothes, cloth)
		return true
	}
	return false
}

func (r *Room) AddThing(things map[string]int) {
	for thing, weight := range things {
		r.Things[thing] = *NewThing(thing, weight)
	}
}

func (h *Human) ToTake(thing string) bool {
	if v, ok := h.Location.Things[thing]; ok {
		if h.Weight < v.Weight {
			fmt.Println("некуда класть")
			return false
		}
		h.Take = append(h.Take, v)
		h.Weight -= v.Weight
		delete(h.Location.Things, thing)
		return true
	}
	return false
}

func init() {
	Locations = *NewLocation()
	Locations.AddRoom("кухня", "коридор")
	Locations.AddRoom("коридор", "кухня", "комната", "улица")
	Locations.AddRoom("комната", "коридор")
	Locations["комната"].AddClothes(map[string]int{
		"рюкзак": 10,
	})
	Locations["комната"].AddThing(map[string]int{
		"конспекты": 1,
	})
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
		case "взять":
			fmt.Scan(&parametr)
			if User.ToTake(parametr) {
				fmt.Println("Вы взяли:", parametr)
			} else {
				fmt.Println("нет такого")
			}
		default:
			fmt.Println("неизвестная команда")
		}
	}
}
