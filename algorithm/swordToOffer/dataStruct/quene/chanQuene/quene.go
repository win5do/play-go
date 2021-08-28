package chanQuene

type Quene struct {
	c chan int
}

func NewQuene() *Quene {
	return &Quene{
		c: make(chan int, 1<<10), // 2^10
	}
}

func (q *Quene) Push(x int) {
	q.c <- x
}

func (q *Quene) Pop() int {
	return <-q.c
}

func (q *Quene) Size() int {
	return len(q.c)
}
