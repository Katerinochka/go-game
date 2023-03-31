package main

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

func NewLocation() *Rooms {
	locations := make(Rooms)
	return &locations
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