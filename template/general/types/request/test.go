package request

type Test1Params struct {
	Id int `form:"id" valid:"Required"`
}

type TestFuncReq struct {
	Id    int     `path:"id"`                    // id主键
	Name  string  `form:"name" valid:"Required"` // 名称
	IsOk  bool    `json:"is_ok"`
	Money float64 `query:"money" valid:"Required"`
}

type TestFunc2Req struct {
	Id1    int     `form:"id1"`                     // id主键1
	Name1  string  `form:"name1" valid:"Required"`  // 名称1
	IsOk1  bool    `form:"is_ok1"`                  //是否ok
	Money1 float64 `form:"money1" valid:"Required"` //钱
}
