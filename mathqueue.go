package mathqueue

import (
	"math"

	basic "github.com/arendtio/mathqueue/basic"
	sortedIntList "github.com/arendtio/sortedintlist"
)

type Interface interface {
	basic.Interface
	Median() int
	Percentile(int) int
	Quantile(int, int) int
	Min() int
	Max() int
}

type MathQueue struct {
	*basic.MathQueueBasic
	sorted *sortedIntList.List
}

func NewMathQueue() *MathQueue {
	return &MathQueue{
		MathQueueBasic: basic.NewMathQueueBasic(),
		sorted:         sortedIntList.NewSortedIntList(),
	}
}

func (mq *MathQueue) Enqueue(n int) {
	mq.MathQueueBasic.Enqueue(n)
	mq.sorted.Insert(n)
}

func (mq *MathQueue) Dequeue() int {
	n := mq.MathQueueBasic.Dequeue()
	mq.sorted.Remove(n)
	return n
}

func (mq MathQueue) Median() int {
	return mq.Quantile(1, 2)
}

func (mq MathQueue) Percentile(x int) int {
	return mq.Quantile(x, 100)
}

func (mq MathQueue) Quantile(x int, q int) int {
	l := float64(mq.sorted.Length())
	f := float64(x) / float64(q)
	pos := l * f
	index := int(pos) - 1
	//fmt.Println("Quantile Length:", l, "Factor:", f, "pos:", pos, "index:", index)
	if float64(int(pos)) == pos { // pos is an integer
		return int(math.Round(float64(mq.sorted.At(index)+mq.sorted.At(index+1)) / 2.0))
	} else {
		return mq.sorted.At(index + 1)
	}
}

func (mq MathQueue) Min() int {
	return mq.sorted.At(0)
}

func (mq MathQueue) Max() int {
	return mq.sorted.At(mq.sorted.Length() - 1)
}
