package request

type Test1Params struct {
	Id string `form:"id" valid:"Required"`
}
