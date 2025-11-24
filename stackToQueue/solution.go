package stackToQueue

import "fmt"

type MyQueue struct {
	in  []int
	out []int
}

func Constructor() MyQueue {
	return MyQueue{
		in:  []int{},
		out: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) Pop() int {
	this.Move()
	element := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return element
}

func (this *MyQueue) Peek() int {
	this.Move()
	if len(this.out)==0{
		return 0
	}
	return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {
	if len(this.out) == 0 && len(this.in) == 0 {
		return true
	}
	return false
}

func (this *MyQueue) Move() {
	if len(this.out) == 0 {
		for len(this.in) > 0 {
			element := this.in[len(this.in)-1]
			this.in = this.in[:len(this.in)-1]
			this.out = append(this.out, element)
		}
	}
}

func TestCode() {
	queue := Constructor()

	queue.Push(10)
	queue.Push(17)
	queue.Push(22)

	fmt.Println(queue.in)

	pop1 := queue.Pop()
	fmt.Println(pop1)

	peek1 := queue.Peek()
	fmt.Println(peek1)

	empty1 := queue.Empty()
	fmt.Println(empty1)

	//****************************

	fmt.Println(queue.out)

	pop2 := queue.Pop()
	fmt.Println(pop2)

	peek2 := queue.Peek()
	fmt.Println(peek2)

	empty2 := queue.Empty()
	fmt.Println(empty2)

	//*****************************

	fmt.Println(queue.out)

	pop3 := queue.Pop()
	fmt.Println(pop3)

	peek3 := queue.Peek()
	fmt.Println(peek3)

	empty3 := queue.Empty()
	fmt.Println(empty3)
}
