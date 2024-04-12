package relationship

type InterpreterContext struct {
	TempMap map[string]int
}

type Expression interface {
	interpret(context InterpreterContext) int
}

type TerminalExpressionA struct {
}

func (t *TerminalExpressionA) interpret(context InterpreterContext) int {
	return context.TempMap["TA"]
}

type TerminalExpressionB struct {
}

func (t *TerminalExpressionB) interpret(context InterpreterContext) int {
	return context.TempMap["TB"]
}

type NonTerminalExpressionA struct {
	CurA Expression
	CurB Expression
}

func (n *NonTerminalExpressionA) interpret(context InterpreterContext) int {
	return n.CurA.interpret(context) - n.CurB.interpret(context)
}

type NonTerminalExpressionB struct {
	CurA Expression
	CurB Expression
}

func (n *NonTerminalExpressionB) interpret(context InterpreterContext) int {
	return n.CurA.interpret(context) + n.CurB.interpret(context)
}
