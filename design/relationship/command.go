package relationship

import "fmt"

type Command interface {
	Execute()
	Cancel()
}

type ACommand struct {
	Obj AStruct //接收者A
}

func (a *ACommand) Execute() {
	a.Obj.AFunc(10, 2)
	fmt.Println("执行了ACommand")
}

func (a *ACommand) Cancel() {
	a.Obj.BFunc()
	fmt.Println("取消了ACommand")
}

type BCommand struct {
	Obj SObject //接收者B
}

func (b *BCommand) Execute() {
	b.Obj.AFunc(10, 2)
	fmt.Println("执行了BCommand")
}

func (b *BCommand) Cancel() {
	b.Obj.BFunc()
	fmt.Println("取消了BCommand")
}

// 调用者
type Invoker struct {
	Command Command
}

func (i *Invoker) ExecuteCommand() {
	i.Command.Execute()
}

func (i *Invoker) CancelCommand() {
	i.Command.Cancel()
}
