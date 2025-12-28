package main

type Interval struct {
	Low  int
	High int
}

type Node struct {
	Interval Interval
	Max      int
	Left     *Node
	Right    *Node
}

func main() {
	
}
