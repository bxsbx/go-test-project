package request

type Test1Params struct {
	Id int `form:"id" valid:"Required"`
}

type TestFuncReq struct {
	Id    int     `path:"id"`                    // id主键
	Name  string  `form:"name" valid:"Required"` // 名称
	IsOk  bool    `json:"is_ok"`
	Money float64 `form:"money"`
}
