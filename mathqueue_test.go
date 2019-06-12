package mathqueue_test

import (
	"math"
	"testing"

	"github.com/arendtio/mathqueue"
)

func check(t *testing.T, target int, result int) {
	if target != result {
		t.Fatal("Result:", result, "does not match target:", target)
	}
}

func TestMedian(t *testing.T) {
	mq := mathqueue.NewMathQueue()
	mq.Enqueue(10)
	check(t, 10, mq.Median())
	mq.Enqueue(50)
	check(t, 30, mq.Median())
	mq.Enqueue(20)
	check(t, 20, mq.Median())
	mq.Enqueue(80)
	check(t, 35, mq.Median())
	mq.Dequeue() // 10
	check(t, 50, mq.Median())
	mq.Dequeue() // 50
	check(t, 50, mq.Median())
	mq.Dequeue() // 20
	check(t, 80, mq.Median())
}

func TestPercentile(t *testing.T) {
	mq := mathqueue.NewMathQueue()
	// from 0 to 100 (101 elements)
	for i := 0; i < 101; i++ {
		mq.Enqueue(i)
	}
	check(t, 99, mq.Percentile(99))
}

func TestQuantile(t *testing.T) {
	mq := mathqueue.NewMathQueue()
	mq.Enqueue(10)
	check(t, 10, mq.Quantile(1, 3))
	mq.Enqueue(30)
	check(t, 10, mq.Quantile(1, 3))
	mq.Enqueue(80)
	check(t, 20, mq.Quantile(1, 3))
	mq.Enqueue(40)
	check(t, 30, mq.Quantile(1, 3))
}

func TestMax(t *testing.T) {
	mq := mathqueue.NewMathQueue()
	mq.Enqueue(30)
	check(t, 30, mq.Max())
	mq.Enqueue(-40)
	check(t, 30, mq.Max())
	mq.Enqueue(80)
	check(t, 80, mq.Max())
	mq.Enqueue(50)
	check(t, 80, mq.Max())
	mq.Dequeue() // 30
	check(t, 80, mq.Max())
	mq.Dequeue() // -40
	check(t, 80, mq.Max())
	mq.Dequeue() // 80
	check(t, 50, mq.Max())
}

func TestMin(t *testing.T) {
	mq := mathqueue.NewMathQueue()
	mq.Enqueue(30)
	check(t, 30, mq.Min())
	mq.Enqueue(10)
	check(t, 10, mq.Min())
	mq.Enqueue(-40)
	check(t, -40, mq.Min())
	mq.Enqueue(80)
	check(t, -40, mq.Min())
	mq.Dequeue()
	check(t, -40, mq.Min())
	mq.Dequeue()
	check(t, -40, mq.Min())
	mq.Dequeue()
	check(t, 80, mq.Min())
}

func TestLimits(t *testing.T) {
	mq := mathqueue.NewMathQueue()
	mq.Enqueue(math.MaxInt64)
	mq.Enqueue(1)
	if mq.Sum() != -9223372036854775808 {
		t.Fatal("Limits sohuld be tested more thoroughly")
	}
}
