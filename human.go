package main

import (
	"fmt"
	"strings"
	"golang.org/x/exp/maps"
)

type Human struct {
	Location Room
	PutOn    []*Cloth
	Weight   int
	Take     map[int]*Thing
}

func NewUser(room Room) *Human {
	u := new(Human)
	u.Location = room
	u.Weight = 0
	u.PutOn = make([]*Cloth, 0)
	u.Take = make(map[int]*Thing)
	return u
}

func (h *Human) GoTo(room string) error {
	for _, v := range h.Location.Entrance {
		if v == room {
			for _, v := range h.Location.AppliedItems {
				if v.Where == room {
					if v.Applied {
						break
					} else {
						return fmt.Errorf("%w нужно что-то сделать с элементом %s", ErrNotAllDone, v.Name)
					}
				}
			}
			h.Location = *Locations[room]
			return nil
		}
	}
	return fmt.Errorf("%w: %s", ErrThisIsNot, room)
}

func (h Human) PrintCanGo() {
	fmt.Println("можно пойти в", strings.Join(h.Location.Entrance, ", "))
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

func (h *Human) ToPutOn(cloth string) error {
	if v, ok := h.Location.Clothes[cloth]; ok {
		h.PutOn = append(h.PutOn, v)
		h.Weight += v.Weight
		delete(h.Location.Clothes, cloth)
		return nil
	}
	return fmt.Errorf("%w: %s", ErrThisIsNot, cloth)
}

func (h *Human) ToTake(thing string) error {
	if v, ok := h.Location.Things[thing]; ok {
		if h.Weight < v.Weight {
			return fmt.Errorf("%w %s", ErrNoCapacity, thing)
		}
		h.Take[v.Id] = v
		h.Weight -= v.Weight
		delete(h.Location.Things, thing)
		return nil
	}
	return fmt.Errorf("%w: %s", ErrThisIsNot, thing)
}

func (h *Human) Apply(thing, appliedItem string) error {
	if v, ok := h.Location.AppliedItems[appliedItem]; ok {
		if v1, ok := h.Take[v.IdThing]; ok && v1.Name == thing {
			v.Applied = true
			delete(h.Take, v1.Id)
			h.Weight += v1.Weight
			return nil
		} else {
			return fmt.Errorf("%w: %s", ErrNotTake, thing)
		}
	} else {
		return fmt.Errorf("%w: %s", ErrThisIsNot, appliedItem)
	}
}