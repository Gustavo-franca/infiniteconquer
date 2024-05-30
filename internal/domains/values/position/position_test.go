package position_test

import (
	"infiniteconquer/internal/domains/values/position"
	"testing"
)


func TestPosition(t *testing.T){
	p := position.New(1,2)
	x, y := p.Get()

	if x != 1 {
		t.Error("position x must be 1")
	}

	if y != 2 {
		t.Error("position y must be 2")
	}


}
