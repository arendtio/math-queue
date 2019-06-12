package mathqueuebasic_test

import (
	"math"
	"testing"

	mathqueuebasic "github.com/arendtio/mathqueue/basic"
)

func TestLength(t *testing.T) {
	mq := mathqueuebasic.NewMathQueueBasic()
	mq.Enqueue(42)
	if mq.Length() != 1 {
		t.Fatal("Length does not match expected Length:", mq.Length(), "vs", 1, "(expected)")
	}
	mq.Dequeue()
	if mq.Length() != 0 {
		t.Fatal("Length does not match expected Length:", mq.Length(), "vs", 0, "(expected)")
	}
}

func TestSum(t *testing.T) {
	mq := mathqueuebasic.NewMathQueueBasic()
	mq.Enqueue(1)
	mq.Enqueue(3)
	mq.Enqueue(5)
	if mq.Sum() != 9 {
		t.Fatal("Sum", mq.Sum(), "does not match expectation 9")
	}
	mq.Enqueue(-4)
	if mq.Sum() != 5 {
		t.Fatal("Sum", mq.Sum(), "can't handle negative numbers")
	}
	mq.Dequeue()
	if mq.Sum() != 4 {
		t.Fatal("Dequeue does not update the Sum:", mq.Sum(), "expected 4")
	}
}

func TestAvg(t *testing.T) {
	mq := mathqueuebasic.NewMathQueueBasic()
	mq.Enqueue(1) // [1]
	if mq.Avg() != 1 {
		t.Fatal("Avg failed on step 1:", mq.Avg())
	}
	mq.Enqueue(3) // [1, 3]
	if mq.Avg() != 2 {
		t.Fatal("Avg failed on step 2:", mq.Avg())
	}
	mq.Enqueue(5) // [1, 3, 5]
	if mq.Avg() != 3 {
		t.Fatal("Avg failed on step 3:", mq.Avg())
	}
	mq.Enqueue(-2) // [1, 3, 5, -2]
	if mq.Avg() != 1 {
		t.Fatal("Avg failed on step 4:", mq.Avg())
	}
	mq.Dequeue() // [3, 5, -2]
	if mq.Avg() != 2 {
		t.Fatal("Avg failed on step 5:", mq.Avg())
	}
	mq.Dequeue() // [5, -2]
	if mq.Avg() != 1 {
		t.Fatal("Avg failed on step 6:", mq.Avg())
	}
}

func TestLimits(t *testing.T) {
	mq := mathqueuebasic.NewMathQueueBasic()
	mq.Enqueue(math.MaxInt64)
	mq.Enqueue(1)
	if mq.Sum() != -9223372036854775808 {
		t.Fatal("Limits sohuld be tested more thoroughly")
	}
}
