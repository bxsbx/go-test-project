package relationship

import "fmt"

// 主题 （被观察者）

type Subject interface {
	Register(observer SObject)
	Unregister(observer SObject)
	Notify()
}

// ConcreteSubject 具体主题
type ConcreteSubject struct {
	observers []SObject
	Message   string
}

func (s *ConcreteSubject) Register(observer SObject) {
	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) Unregister(observer SObject) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) Notify() {
	for _, observer := range s.observers {
		fmt.Println(observer.BFunc(), "接收到", s.Message)

	}
}

func (s *ConcreteSubject) SendMsg(msg string) {
	s.Message = msg
	fmt.Println("发送消息", msg)
	s.Notify()
}
