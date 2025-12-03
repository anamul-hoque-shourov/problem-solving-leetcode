package leetcode_232

import "fmt"

type myQueue struct {
	in  []int
	out []int
}

func constructor() myQueue {
	return myQueue{
		in:  []int{},
		out: []int{},
	}
}

func (m *myQueue) push(x int) {
	m.in = append(m.in, x)
}

func (m *myQueue) pop() int {
	m.move()
	element := m.out[len(m.out)-1]
	m.out = m.out[:len(m.out)-1]
	return element
}

func (m *myQueue) peek() int {
	m.move()
	if len(m.out) == 0 {
		return 0
	}
	return m.out[len(m.out)-1]
}

func (m *myQueue) empty() bool {
	if len(m.out) == 0 && len(m.in) == 0 {
		return true
	}
	return false
}

func (m *myQueue) move() {
	if len(m.out) == 0 {
		for len(m.in) > 0 {
			element := m.in[len(m.in)-1]
			m.in = m.in[:len(m.in)-1]
			m.out = append(m.out, element)
		}
	}
}

func TestCode() {
	queue := constructor()

	queue.push(10)
	queue.push(17)
	queue.push(22)

	fmt.Println(queue.in)

	pop1 := queue.pop()
	fmt.Println(pop1)

	peek1 := queue.peek()
	fmt.Println(peek1)

	empty1 := queue.empty()
	fmt.Println(empty1)

	//****************************

	fmt.Println(queue.out)

	pop2 := queue.pop()
	fmt.Println(pop2)

	peek2 := queue.peek()
	fmt.Println(peek2)

	empty2 := queue.empty()
	fmt.Println(empty2)

	//*****************************

	fmt.Println(queue.out)

	pop3 := queue.pop()
	fmt.Println(pop3)

	peek3 := queue.peek()
	fmt.Println(peek3)

	empty3 := queue.empty()
	fmt.Println(empty3)
}
