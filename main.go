package main

import (
	"fmt"
	"strings"

	"golang.org/x/exp/maps"
)

type Room struct {
	Name         string
	Entrance     []string
	Clothes      map[string]*Cloth
	Things       map[string]*Thing
	AppliedItems map[string]*AppliedItem
}

type Human struct {
	Location Room
	PutOn    []*Cloth
	Weight   int
	Take     map[int]*Thing
}

type Cloth struct {
	Name   string
	Weight int
}

type Thing struct {
	Id     int
	Name   string
	Where  string
	Weight int
}

type AppliedItem struct {
	Name    string
	IdThing int
	Applied bool
	Where   string
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
	r.Clothes = make(map[string]*Cloth)
	r.Things = make(map[string]*Thing)
	r.AppliedItems = make(map[string]*AppliedItem)
	return r
}

func NewUser(room Room) *Human {
	u := new(Human)
	u.Location = room
	u.Weight = 0
	u.PutOn = make([]*Cloth, 0)
	u.Take = make(map[int]*Thing)
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

func NewAppliedItem(name string, idThing int, where string) *AppliedItem {
	a := new(AppliedItem)
	a.Name = name
	a.IdThing = idThing
	a.Applied = false
	a.Where = where
	return a
}

func (r *Rooms) AddRoom(name string, entrance ...string) {
	room := NewRoom(name, entrance)
	(*r)[name] = room
}

func (h *Human) GoTo(room string) bool {
	for _, v := range h.Location.Entrance {
		if v == room {
			for _, v := range h.Location.AppliedItems {
				if v.Where == room {
					if v.Applied {
						break
					} else {
						fmt.Print("не все сделано. ")
						return false
					}
				}
			}
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
		r.Clothes[cloth] = NewCloth(cloth, weight)
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

func (r *Room) AddThing(things ...*Thing) {
	for _, thing := range things {
		(*thing).Id = len(r.Things)
		r.Things[thing.Name] = thing

	}
}

func (h *Human) ToTake(thing string) bool {
	if v, ok := h.Location.Things[thing]; ok {
		if h.Weight < v.Weight {
			fmt.Println("некуда класть")
			return false
		}
		h.Take[v.Id] = v
		h.Weight -= v.Weight
		delete(h.Location.Things, thing)
		return true
	}
	return false
}

func (h *Human) Apply(thing, appliedItem string) bool {
	if v, ok := h.Location.AppliedItems[appliedItem]; ok {
		if v1, ok := h.Take[v.IdThing]; ok && v1.Name == thing{
			v.Applied = true
			delete(h.Take, v1.Id)
			h.Weight += v1.Weight
			return true
		} else {
			fmt.Println("вы не взяли этот предмет: ", thing)
			return false
		}
	} else {
		fmt.Println("тут нет элемента", appliedItem)
		return false
	}
}

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
			if User.GoTo(parametr1) {
				User.PrintCanGo()
			} else {
				fmt.Println("нет такого")
			}
		case "осмотреться":
			User.LookAround()
		case "надеть":
			fmt.Scan(&parametr1)
			if User.ToPutOn(parametr1) {
				fmt.Println("Вы надели:", parametr1)
			} else {
				fmt.Println("нет такого")
			}
		case "взять":
			fmt.Scan(&parametr1)
			if User.ToTake(parametr1) {
				fmt.Println("Вы взяли:", parametr1)
			} else {
				fmt.Println("нет такого")
			}
		case "применить":
			fmt.Scan(&parametr1)
			fmt.Scan(&parametr2)
			if User.Apply(parametr1, parametr2) {
				fmt.Println("получилось")
			}
		default:
			fmt.Println("неизвестная команда")
		}
	}
}
