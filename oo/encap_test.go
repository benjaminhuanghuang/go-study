package oostudy

import (
	"fmt"
	"testing"
)

type Employee struct {
	ID   string
	Name string
	Age  int
}

func (e Employee) ToString() string {
	return fmt.Sprintf("ID: %s/Name:%s/Age%d", e.ID, e.Name, e.Age)
}

func (e *Employee) ToString() string {
	return fmt.Sprintf("ID: %s/Name:%s/Age%d", e.ID, e.Name, e.Age)
}

func TestCreateObject(t *testing.T) {
	e := Employee{"0", "Bob", 20}
}
