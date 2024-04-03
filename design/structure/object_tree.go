package structure

import "fmt"

type Component interface {
	Operation() string
}

// 叶子节点
type Leaf struct {
	Name string
}

func (l *Leaf) Operation() string {
	return l.Name
}

// 树枝节点
type Composite struct {
	Name     string
	Children []Component
}

func (c *Composite) Operation() string {
	result := fmt.Sprintf("%s\n", c.Name)
	for _, child := range c.Children {
		result += "   " + child.Operation() + "\n"
	}
	return result
}

func (c *Composite) Add(child Component) {
	c.Children = append(c.Children, child)
}

func (c *Composite) Remove(index int) {
	if index >= 0 && index < len(c.Children) {
		c.Children = append(c.Children[:index], c.Children[index+1:]...)
	}
}

func (c *Composite) GetChild(index int) Component {
	if index >= 0 && index < len(c.Children) {
		return c.Children[index]
	}
	return nil
}
