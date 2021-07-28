package main
import "./LinkedList"

func main() {
	head := &LinkedList.Node{Value: 6}
	head.InsertNode(10)
	head.InsertNode(20)
	head.InsertNode(2)
	head.InsertNode(5)
	head.PrintList()
	head2 := &LinkedList.Node{Value: 1}
	head2.InsertNode(2)
	head2.InsertNode(3)
	head2.InsertNode(4)
	head2.PrintList()
	head.Merge(head2)
	head.PrintList()
}