package main

type MyStack struct {
	queue1 []int
	queue2 []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.queue1 = append(this.queue1, x)
}

func (this *MyStack) Pop() int {
	// 先将queue1中前len-1移动到queue2中，然后最后一个元素出队，即是栈顶元素，然后将queue2中元素移动回来
	for len(this.queue1) != 1 {
		this.queue2 = append(this.queue2, this.queue1[0])
		this.queue1 = this.queue1[1:]
	}
	res := this.queue1[0]
	this.queue1 = this.queue1[1:]
	for len(this.queue2) != 0 {
		this.queue1 = append(this.queue1, this.queue2[0])
		this.queue2 = this.queue2[1:]
	}
	return res
}

func (this *MyStack) Top() int {
	for len(this.queue1) != 1 {
		this.queue2 = append(this.queue2, this.queue1[0])
		this.queue1 = this.queue1[1:]
	}
	res := this.queue1[0]
	this.queue1 = this.queue1[1:]
	this.queue2 = append(this.queue2, res)
	for len(this.queue2) != 0 {
		this.queue1 = append(this.queue1, this.queue2[0])
		this.queue2 = this.queue2[1:]
	}
	return res
}

func (this *MyStack) Empty() bool {
	if len(this.queue1) == 0 && len(this.queue2) == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
