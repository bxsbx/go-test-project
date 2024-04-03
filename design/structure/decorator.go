package structure

// 装饰者抽象类
type Decorator struct {
	Obj SObject
}

func (d *Decorator) AFunc() int {
	if d.Obj != nil {
		return d.Obj.AFunc()
	}
	return 0
}

func (d *Decorator) BFunc() string {
	if d.Obj != nil {
		return d.Obj.BFunc()
	}
	return ""
}

// 具体装饰者 A
type ConcreteDecoratorA struct {
	Obj SObject
}

func (d *ConcreteDecoratorA) AFunc() int {
	return d.Obj.AFunc() + 10
}

func (d *ConcreteDecoratorA) BFunc() string {
	return d.Obj.BFunc() + "DecoratorA"
}

// 具体装饰者 B
type ConcreteDecoratorB struct {
	Obj SObject
}

func (d *ConcreteDecoratorB) AFunc() int {
	return d.Obj.AFunc() + 100
}

func (d *ConcreteDecoratorB) BFunc() string {
	return d.Obj.BFunc() + "DecoratorB"
}
