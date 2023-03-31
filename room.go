package main

type Room struct {
	Name         string
	Entrance     []string
	Clothes      map[string]*Cloth
	Things       map[string]*Thing
	AppliedItems map[string]*AppliedItem
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

func (r *Room) AddClothes(clothes map[string]int) {
	for cloth, weight := range clothes {
		r.Clothes[cloth] = NewCloth(cloth, weight)
	}
}

func (r *Room) AddThing(things ...*Thing) {
	for _, thing := range things {
		(*thing).Id = len(r.Things)
		r.Things[thing.Name] = thing

	}
}

