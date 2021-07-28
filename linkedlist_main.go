package main
import "./LinkedList"

func main() {
	head := &LinkedList.Node{Value: 6}
	head.InsertNode(10)
	head.InsertNode(20)
	head.InsertNode(2)
	head.InsertNode(5)
	head.PrintList()
}