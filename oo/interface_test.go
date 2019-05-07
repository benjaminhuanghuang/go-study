package oostudy

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}
type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "Go-Hello world"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
