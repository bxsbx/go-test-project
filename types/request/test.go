package request

type Test1Params struct {
	Id int `form:"id" valid:"Required"`
}
