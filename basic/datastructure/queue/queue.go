package queue

type Queue struct {
	Items []interface{}
}

func (queue *Queue) Enqueue(num interface{}) {
	queue.Items = append(queue.Items, num)
}

func (queue *Queue) Dequeue() interface{} {
	temp := queue.Items[0]
	queue.Items = queue.Items[1:]
	return temp
}
