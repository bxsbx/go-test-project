package request

type Test1Params struct {
	Id int `form:"id" valid:"Required"`
}

type TestFuncReq struct {
	Id    int     `query:"id" valid:"Required"`   // id主键
	Name  string  `query:"name" valid:"Required"` // 名称
	IsOk  bool    `query:"is_ok"`
	Money float64 `query:"money"`
}
