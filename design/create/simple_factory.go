package create

// --------------------简单工厂（常用）---------------
type SimpleFactory struct{}

// 也可以通过存放在map中 map[string]SObject 这map即为工厂（适用于可以存全局object（单例））

// 如果需要每次生成新的用如下方法（业务上常用）
func (f SimpleFactory) CreateObject(key string) SObject {
	switch key {
	case "A":
		return &AStruct{}
	case "B":
		return &BStruct{}
	}
	return nil
}

// ----------------通过方法返回对应的对象（架构上常用）-------------

func (f SimpleFactory) CreateAStruct() SObject {
	return &AStruct{}
}

func (f SimpleFactory) CreateBStruct() SObject {
	return &BStruct{}
}
