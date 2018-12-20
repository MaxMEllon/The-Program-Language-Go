package cyclic

import "testing"

func TestIsCyclic(t *testing.T) {
	type Cycle struct {
		Value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}
	if !IsCyclic(c) {
		t.Errorf("cyclic のはず")
	}

	var nc Cycle
	nc = Cycle{42, &Cycle{32, nil}}
	if IsCyclic(nc) {
		t.Errorf("cyclic ではないはず")
	}
}
