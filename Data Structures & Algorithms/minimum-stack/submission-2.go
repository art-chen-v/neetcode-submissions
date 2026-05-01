type Node struct {
	val int
	curMin int
}

type MinStack struct {
	data []Node
}

func Constructor() MinStack {
	return MinStack{
		data: make([]Node, 0),
	}
}

func (this *MinStack) Push(val int) {
	if len(this.data) > 0 && this.GetMin() < val {
		this.data = append(this.data, Node{val: val, curMin: this.GetMin()})
	} else {
		this.data = append(this.data, Node{val: val, curMin: val})
	}
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1].val
}

func (this *MinStack) GetMin() int {
	return this.data[len(this.data)-1].curMin
}
