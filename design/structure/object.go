package structure

//---------------------具体对象-------------------

type SObject interface {
	AFunc() int
	BFunc() string
}

//---------------对象A-----------------

type AStruct struct {
	A int
	B string
}

func (a *AStruct) AFunc() int {
	return a.A * 100
}

func (a *AStruct) BFunc() string {
	return a.B
}

//--------------对象B-----------------

type BStruct struct {
	A int
	B string
}

func (b *BStruct) AFunc() int {
	return b.A * 10
}

func (b *BStruct) BFunc() string {
	return b.B
}
