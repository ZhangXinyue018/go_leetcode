package min_stack

type MinStack struct {
	stackList    []int
	minStackList []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.stackList = append(this.stackList, x)
	if len(this.minStackList) == 0 {
		this.minStackList = append(this.minStackList, x)
	} else if x <= this.minStackList[len(this.minStackList)-1] {
		this.minStackList = append(this.minStackList, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.stackList) == 0 {
		return
	}
	popValue := this.stackList[len(this.stackList)-1]
	this.stackList = this.stackList[:len(this.stackList)-1]
	if popValue == this.minStackList[len(this.minStackList)-1] {
		this.minStackList = this.minStackList[:len(this.minStackList)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stackList[len(this.stackList)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStackList[len(this.minStackList)-1]
}
