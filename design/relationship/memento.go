package relationship

// 备忘录
type Memento struct {
	Context string
	Obj     AStruct
}

// 发起者
type Originator struct {
	Obj AStruct
}

func (o *Originator) Write(a int, b string) {
	o.Obj.A = a
	o.Obj.B = b
}

func (o *Originator) GetObj() AStruct {
	return o.Obj
}

func (o *Originator) SaveObjToMemento() Memento {
	return Memento{Context: o.Obj.B, Obj: o.Obj}
}

func (o *Originator) RestoreFromMemento(memento Memento) {
	o.Obj = memento.Obj
}

// Caretaker 管理者，负责保存备忘录
type Caretaker struct {
	mementos []Memento
}

func (c *Caretaker) AddMemento(memento Memento) {
	c.mementos = append(c.mementos, memento)
}

func (c *Caretaker) GetMemento(index int) Memento {
	if index >= 0 && index < len(c.mementos) {
		return c.mementos[index]
	}
	return Memento{}
}

func (c *Caretaker) PopMemento() Memento {
	if len(c.mementos) > 0 {
		memento := c.mementos[len(c.mementos)-1]
		c.mementos = c.mementos[:len(c.mementos)-1]
		return memento
	}
	return Memento{}
}
