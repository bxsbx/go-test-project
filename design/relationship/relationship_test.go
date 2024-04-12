package relationship

import (
	"fmt"
	"strconv"
	"testing"
)

func TestStrategy(t *testing.T) {
	env1 := StrategyEnv{Obj: &AStruct{}}
	fmt.Println(env1.ExecFunc1(10, 12))

	env2 := StrategyEnv{Obj: &BStruct{}}
	fmt.Println(env2.ExecFunc1(10, 12))
}

func TestTemplate(t *testing.T) {
	t1 := Template{Obj: &AStruct{}}
	t1.TemplateMethod1()

	t2 := Template{Obj: &BStruct{}}
	t2.TemplateMethod1()
}

func TestObserver(t *testing.T) {
	a := &AStruct{}
	b := &BStruct{}
	subject := &ConcreteSubject{}
	subject.Register(a)
	subject.Register(b)

	subject.SendMsg("vasva")

	subject.Unregister(a)

	subject.SendMsg("奢侈啊")
}

func TestIterator(t *testing.T) {
	collectionA := &ACollection{}
	for i := 0; i < 100; i++ {
		collectionA.Add(AStruct{B: strconv.Itoa(i)})
	}
	iteratorA := collectionA.GetIterator()
	for iteratorA.HasNext() {
		next := iteratorA.GetNext()
		fmt.Println(next)
	}

	collectionB := &BCollection{Map: make(map[string]AStruct)}
	for i := 0; i < 100; i++ {
		collectionB.Add(AStruct{B: strconv.Itoa(i)})
	}
	iteratorB := collectionB.GetIterator()
	for iteratorB.HasNext() {
		next := iteratorB.GetNext()
		fmt.Println(next)
	}

}

func TestResponsibility(t *testing.T) {
	handleA := &AHandle{Name: "A"}
	handleB := &BHandle{Name: "B", Obj: &AStruct{B: "B"}}
	handleC := &CHandle{Name: "C", Obj: &AStruct{B: "C"}}

	//handleA.SetNext(handleB)
	//handleB.SetNext(handleC)

	handleA.SetNext(handleC)
	handleC.SetNext(handleB)

	handleA.Handle()

}

func TestCommand(t *testing.T) {
	aStruct := AStruct{A: 10, B: "A"}
	aCommand := &ACommand{Obj: aStruct}
	invoker := Invoker{Command: aCommand}
	invoker.ExecuteCommand()
	invoker.CancelCommand()

	bCommand := &BCommand{Obj: &aStruct}
	invoker = Invoker{Command: bCommand}
	invoker.ExecuteCommand()
	invoker.CancelCommand()
}

func TestState(t *testing.T) {
	context := &Context{}
	context.AState = &AState{Val: 1, Context: context}
	context.BState = &BState{Val: 2, Context: context}
	context.CState = &CState{Val: 3, Context: context}

	context.CurState = context.AState

	context.A()
	context.C()
	context.B()

}

func TestMemento(t *testing.T) {
	originator := Originator{}
	caretaker := Caretaker{}

	originator.Write(1, "A")
	caretaker.AddMemento(originator.SaveObjToMemento())
	fmt.Println(originator.GetObj())

	originator.Write(2, "B")
	caretaker.AddMemento(originator.SaveObjToMemento())
	fmt.Println(originator.GetObj())

	originator.Write(3, "C")
	//caretaker.AddMemento(originator.SaveObjToMemento())
	fmt.Println(originator.GetObj())

	originator.RestoreFromMemento(caretaker.PopMemento())
	fmt.Println(originator.GetObj())

	originator.RestoreFromMemento(caretaker.GetMemento(0))
	fmt.Println(originator.GetObj())

}

func TestVisitor(t *testing.T) {
	objectStructure := NewObjectStructure()
	objectStructure.AddElement(&ElementA{B: "A"})
	objectStructure.AddElement(&ElementB{A: "B"})

	objectStructure.AcceptElement(&Visitor1{A: "AV"})
	objectStructure.AcceptElement(&Visitor2{})

}

func TestMediator(t *testing.T) {
	m := MediatorA{ObjMap: make(map[string]Colleague)}
	a := ColleagueA{Name: "A", Mediator: &m}
	b := ColleagueA{Name: "B", Mediator: &m}
	m.Add(a.Name, &a)
	m.Add(b.Name, &b)
	a.Send(b.Name, "cjsasocj")
	b.Send(a.Name, "vewevwev")
}

func TestInterpret(t *testing.T) {
	context := InterpreterContext{TempMap: make(map[string]int)}
	context.TempMap["TA"] = 10
	context.TempMap["TB"] = 20

	terminalExpressionA := &TerminalExpressionA{}
	terminalExpressionB := &TerminalExpressionB{}
	//20-10
	nonTerminalExpressionA := &NonTerminalExpressionA{CurA: terminalExpressionB, CurB: terminalExpressionA}
	//(20-10)+20
	nonTerminalExpressionB := &NonTerminalExpressionB{nonTerminalExpressionA, terminalExpressionB}
	//(20-10)-((20-10)+20)
	expressionA := NonTerminalExpressionA{nonTerminalExpressionA, nonTerminalExpressionB}

	interpret := expressionA.interpret(context)
	fmt.Println(interpret)

}
