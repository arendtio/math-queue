package mathqueuebasic

import (
	"math"
)

type Interface interface {
	Enqueue(int)
	Dequeue() int
	Front() int
	Back() int
	Length() int
	Sum() int
	Avg() int
}

type MathQueueBasic struct {
	list []int
	sum  int
}

func NewMathQueueBasic() *MathQueueBasic {
	return &MathQueueBasic{}
}

func (mq *MathQueueBasic) Enqueue(n int) {
	mq.list = append(mq.list, n)
	mq.sum += n
}

func (mq *MathQueueBasic) Dequeue() int {
	n := mq.list[0]
	mq.list = mq.list[1:]
	mq.sum -= n
	return n
}

func (mq MathQueueBasic) Front() int {
	return mq.list[0]
}

func (mq MathQueueBasic) Back() int {
	return mq.list[len(mq.list)-1]
}

func (mq MathQueueBasic) Length() int {
	return len(mq.list)
}

func (mq MathQueueBasic) Sum() int {
	return mq.sum
}

func (mq MathQueueBasic) Avg() int {
	return int(math.Round(float64(mq.sum) / float64(len(mq.list))))
}
