package create

//key(Factory)-key(object)

// --------------抽象工厂-------------------
type SFactory interface {
	//适用于不同的工厂能有不同的对象
	CreateSObject(key string) SObject
	CreateLObject(key string) LObject

	//适用于不同工厂获取所有对象各一个
	CreateSObjectStruct() SObject
	CreateLObjectStruct() LObject
}

//---------------工厂A-----------------

type AFactory struct{}

// 如果需要每次生成新的用如下方法
func (f AFactory) CreateSObject(key string) SObject {
	switch key {
	case "A":
		return &AStruct{}
	case "B":
		return &BStruct{}
	}
	return nil
}

func (f AFactory) CreateLObject(key string) LObject {
	switch key {
	case "C":
		return &CStruct{}
	}
	return nil
}

// ----------------通过方法返回对应的对象-------------

func (f AFactory) CreateSObjectStruct() SObject {
	return &AStruct{}
}

func (f AFactory) CreateLObjectStruct() LObject {
	return &CStruct{}
}

//---------------工厂B-----------------

type BFactory struct{}

func (f BFactory) CreateSObject(key string) SObject {
	switch key {
	case "B":
		return &BStruct{}
	}
	return nil
}

// 如果需要每次生成新的用如下方法
func (f BFactory) CreateLObject(key string) LObject {
	switch key {
	case "C":
		return &CStruct{}
	case "D":
		return &DStruct{}
	}
	return nil
}

// ----------------通过方法返回对应的对象-------------

func (f BFactory) CreateSObjectStruct() SObject {
	return &BStruct{}
}

func (f BFactory) CreateLObjectStruct() LObject {
	return &DStruct{}
}

//-------------工厂的工厂-----------------

type AbstractFactory struct{}

// 或者使用map存放
func (a AbstractFactory) GetFactory(key string) SFactory {
	switch key {
	case "A":
		return &AFactory{}
	case "B":
		return &BFactory{}
	}
	return nil
}

func (a AbstractFactory) CreatAFactory() SFactory {
	return &AFactory{}
}

func (a AbstractFactory) CreatBFactory() SFactory {
	return &BFactory{}
}
