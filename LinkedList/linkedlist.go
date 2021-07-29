package LinkedList

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var encoderCfg = zap.NewProductionEncoderConfig()
var atomicLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
var logger = zap.New(zapcore.NewCore(
	zapcore.NewJSONEncoder(encoderCfg),
	zapcore.Lock(os.Stdout),
	atomicLevel,
))
var sugar = logger.Sugar()

type Node struct {
	Value int
	next  *Node
}

func (rootNode *Node) InsertNode(value int) {
	if rootNode == nil {
		sugar.Warnf("Root node is null!")
		return
	} else if rootNode.next == nil {
		rootNode.next = &Node{Value: value}
		sugar.Infof("Added rootNode with value: %v", value)
	} else {
		sugar.Debugf("Moving to next rootNode")
		rootNode.next.InsertNode(value)
	}
	return
}

func (rootNode *Node) PrintList() {
	if rootNode == nil {
		sugar.Warnf("Root node is null!")
		return
	}
	println(rootNode.Value)
	rootNode.next.PrintList()
}

func (rootNode *Node) GetElement(value int) *Node {
	if rootNode == nil {
		sugar.Warnf("Root node is null!")
		return nil
	}
	if rootNode.Value == value {
		return rootNode
	} else {
		return rootNode.next.GetElement(value)
	}
}

func ChangeLogLevel(newLevel string) {
	switch newLevel {
	case "Debug":
		atomicLevel.SetLevel(zap.DebugLevel)
	case "Warn":
		atomicLevel.SetLevel(zap.WarnLevel)
	default:
		atomicLevel.SetLevel(zap.InfoLevel)
	}
}

func (root *Node) Merge(root2 *Node) {
	pointer := root
	for pointer.next != nil {
		pointer = pointer.next
	}
	pointer.next = root2
	return
}
