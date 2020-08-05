package tree

import "fmt"

// Print prints a Tree. Main use is to aid in debugging.
func (t *Tree) Print() {
	if t.Root == nil {
		fmt.Println("<nil>")
	}
	t.Root.print()
}

func (n *Node) print() {
	q := newQueue()
	q.push(n)
//	fmt.Printf("---Q--- %T %s\n", q, q)   		// *tree.queue	
//	fmt.Printf("--- Print --- %T %s\n", n, n)	// *tree.Node

//	ss := *n.Elem	//n.Elem è un puntatore
/*	scan := (*n.Elem).m
	for k, v := range scan {
		fmt.Printf("--Elem-- %d %s--\n", k, v) 
		for  _, rr := range v {
			fmt.Printf("--RR %s--\n",rr)
		}
	}
*/
	nodesInCurrentLevel := 1
	nodesInNextLevel := 0

	for i:=0; !q.empty(); i++ {
		do := q.pop()
/*
		if do != nil   {
			fmt.Printf("\n---Do---%s-- ",  do.Elem.Name())	// *tree.Node	

			if do.Left != nil {
				fmt.Printf("L**%s*L ",  do.Left.Elem.Name() )	// *tree.Node	
			}
			if do.Right != nil {
				fmt.Printf("R*%s*R\n",  do.Right.Elem.Name() )	// *tree.Node	
			}
		}
*/
		nodesInCurrentLevel--

		if do != nil {
			fmt.Print(i, " ", do.Elem.Name(), " ")
			q.push(do.Left)
			q.push(do.Right)
			nodesInNextLevel += 2
		}
		if nodesInCurrentLevel == 0 {
			fmt.Println()
		}
		nodesInCurrentLevel = nodesInNextLevel
		nodesInNextLevel = 0
	}
	fmt.Println()
}

type queue []*Node

// newQueue returns a new queue.
func newQueue() queue {
	q := queue([]*Node{})
	return q
}

// push pushes n to the end of the queue.
func (q *queue) push(n *Node) {
	*q = append(*q, n)
}

// pop pops the first element off the queue.
func (q *queue) pop() *Node {
	n := (*q)[0]
	*q = (*q)[1:]
	return n
}

// empty returns true when the queue contains zero nodes.
func (q *queue) empty() bool {
	return len(*q) == 0
}
