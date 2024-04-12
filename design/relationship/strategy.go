package relationship

type StrategyEnv struct {
	Obj SObject
}

func (s *StrategyEnv) ExecFunc1(x1, x2 int) int {
	return s.Obj.AFunc(x1, x2)
}
