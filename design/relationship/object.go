package relationship

//---------------------具体对象-------------------

type SObject interface {
	AFunc(x1, x2 int) int
	BFunc() string
}

//---------------对象A-----------------

type AStruct struct {
	A int
	B string
}

func (a *AStruct) AFunc(x1, x2 int) int {
	return x1 + x2
}

func (a *AStruct) BFunc() string {
	return "A"
}

//--------------对象B-----------------

type BStruct struct{}

func (b *BStruct) AFunc(x1, x2 int) int {
	return x1 - x2
}

func (b *BStruct) BFunc() string {
	return "B"
}
