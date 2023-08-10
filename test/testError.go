package main

import (
	"StandardProject/common/errorz"
	"errors"
)

func A() error {
	//return errorz.CodeMsgError(100, "122", errors.New("vsva"))
	//return errors.Unwrap(errors.New("vsv"))
	//return errors.New("3434")
	//return errors.Wrap(errors.New("scsca"), "gen sql")
	//a := 12
	//fmt.Println(a)
	return errorz.CodeMsgError(1001, "csfewefwef", errors.New("csc"))
}

func B() error {
	err := A()
	//kk := "casvvq"
	//fmt.Printf(kk)
	//fmt.Printf("%p\n", err)
	//return errorz.CodeMsgError(1000, "vewweve111v", err)
	//return errorz.CodeMsgError(1000, "vewweve111v", errors.New("csc"))
	return err
}

func C() error {
	err := B()
	//fmt.Printf("%p\n", err)
	return errorz.CodeMsgError(1006, "nynyy", err)
}

type Info struct {
	Name  string
	MIcs  string
	Icscw int
}

func main() {
	//list := []string{}
	//
	//fmt.Println(len(list))
	//list = append(list, "2323")
	//list = append(list, "verberb")
	//list = append(list, "verberb")
	//list = append(list, "verberb")
	//list = append(list, "verberb")
	//fmt.Println(list)
	//fmt.Println(cap(list))
	//
	////list[5] = "2323"
	//
	//stack := errorz.NewStack[string](8)
	//stack.Push("cascav")
	//stack.Push("berberb")
	//fmt.Println(stack.Size())
	//fmt.Println(stack.Pull())
	//fmt.Println(stack.Top())
	//stack.Pull()
	//fmt.Println(stack.Top())
	//fmt.Println(stack.Pull())
	//fmt.Println(stack.Size())
	//
	//stack.CutCapSize()
	//
	//err := C()
	//fmt.Printf("%p\n", err)
	////fmt.Println(err.Error())
	//stack := errorz.GetErrorStackList(err)
	//fmt.Println(stack)
	////fmt.Println(Info{stack, "vewevw", 23})
	//marshal, _ := json.Marshal(stack)
	//fmt.Println(string(marshal))

}
