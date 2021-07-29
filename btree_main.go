package main

import "github.com/LalithSrinivas/golang-basics/BTree"

// This is an example program using all the functions of Btree

func main() {
	err := BTree.Logger.Sync()
	if err != nil {
		return
	} // flushes buffer, if any
	var rootNode = &BTree.Node{Value: 10}
	BTree.ChangeLogLevel("Info")
	println("Log level set to Info")
	rootNode.Insert(15)
	rootNode.Insert(12)
	println("Changed log level to Debug")
	BTree.ChangeLogLevel("Debug")
	rootNode.Insert(9)
	rootNode.Inorder()
	println()
	println()
	rootNode.Insert(16)
	rootNode.Insert(0)
	BTree.ChangeLogLevel("Warn")
	println("Changed log level to Warn")
	rootNode.Insert(15)
	rootNode.Insert(13)
	rootNode.Insert(12)
	rootNode.Inorder()
}
