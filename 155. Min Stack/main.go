type MinStack struct {
	data   []int
	sorted []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.sorted) == 0 || val <= this.GetMin() {
		this.sorted = append(this.sorted, val)
	}
	this.data = append(this.data, val)
}

func (this *MinStack) Pop() {
	if this.Top() == this.GetMin() {
		this.sorted = this.sorted[:len(this.sorted)-1]
	}
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.sorted[len(this.sorted)-1]
}
