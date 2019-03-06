package queue

type Queue struct {
	Items []int
}

func (queue *Queue) Enqueue(num int) {
	queue.Items = append(queue.Items, num)
}

func (queue *Queue) Dequeue() int {
	temp := queue.Items[0]
	queue.Items = queue.Items[1:]
	return temp
}
