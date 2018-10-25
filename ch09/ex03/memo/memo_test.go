package memo_test

import (
	"testing"

	"github.com/maxmellon/The-Program-Language-Go/ch09/ex03/memo"
	"github.com/maxmellon/The-Program-Language-Go/ch09/ex03/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Concurrent(t, m)
}

func TestSequentialCancel(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.SequentialCancel(t, m)
}

func TestConcurrentCancel(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.ConcurrentCancel(t, m)
}
