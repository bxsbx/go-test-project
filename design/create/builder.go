package create

//--------------建造者模式-------------

type Product struct {
	A int
	B string
	C bool
}

type BuilderInterface interface {
	GetProduct() Product
	CreateA()
	CreateB()
	CreateC()
}

type ABuilder struct {
	Product Product
}

func (a *ABuilder) GetProduct() Product {
	return a.Product
}

func (a *ABuilder) CreateA() {
	a.Product.A = 10
}

func (a *ABuilder) CreateB() {
	a.Product.B = "100"
}

func (a *ABuilder) CreateC() {
	a.Product.C = false
}

type BBuilder struct {
	Product Product
}

func (b *BBuilder) GetProduct() Product {
	return b.Product
}

func (b *BBuilder) CreateA() {
	b.Product.A = 20
}

func (b *BBuilder) CreateB() {
	b.Product.B = "200"
}

func (b *BBuilder) CreateC() {
	b.Product.C = true
}

type Director struct {
	Builder BuilderInterface
}

// 某个产品构建的顺序是相同的，只是组件生产不同
func (d Director) GetProductFromBuilder(builder BuilderInterface) Product {
	builder.CreateA()
	builder.CreateB()
	builder.CreateC()
	return builder.GetProduct()
}

func (d *Director) GetProductFromBuilder2() Product {
	d.Builder.CreateA()
	d.Builder.CreateB()
	d.Builder.CreateC()
	return d.Builder.GetProduct()
}
