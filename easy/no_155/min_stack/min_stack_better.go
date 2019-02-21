package min_stack

type MinStackBetter struct {
	stackList []int
	min       int
}

/** initialize your data structure here. */
func ConstructorBetter() MinStackBetter {
	return MinStackBetter{}
}

func (this *MinStackBetter) PushBetter(x int) {
	if len(this.stackList) == 0 || (len(this.stackList) > 0 && x <= this.min) {
		this.stackList = append(this.stackList, this.min)
		this.min = x
	}
	this.stackList = append(this.stackList, x)
}

func (this *MinStackBetter) PopBetter() {
	popValue := this.stackList[len(this.stackList)-1]

	this.stackList = this.stackList[:len(this.stackList)-1]
	if popValue == this.min {
		this.min = this.stackList[len(this.stackList)-1]
		this.stackList = this.stackList[:len(this.stackList)-1]
	}

}

func (this *MinStackBetter) TopBetter() int {
	return this.stackList[len(this.stackList)-1]
}

func (this *MinStackBetter) GetMinBetter() int {
	return this.min
}
