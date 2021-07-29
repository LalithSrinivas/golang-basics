package BTree

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var encoderCfg = zap.NewProductionEncoderConfig()
var atomicLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
var Logger = zap.New(zapcore.NewCore(
	zapcore.NewJSONEncoder(encoderCfg),
	zapcore.Lock(os.Stdout),
	atomicLevel,
))
var sugar = Logger.Sugar()

type Node struct {
	Value int
	right *Node
	left  *Node
}

func (t *Node) Insert(val int) {
	if t == nil {
		sugar.Warnf("Null root node")
		return
	}

	if val > t.Value {
		if t.right == nil {
			sugar.Debugf("Right")
			sugar.Infof("Added node with value %d successfully!", val)
			t.right = &Node{Value: val}
		} else {
			sugar.Debugf("Right")
			t.right.Insert(val)
		}
	} else {
		if t.left == nil {
			sugar.Debugf("Left")
			sugar.Infof("Added node with value %d successfully!", val)
			t.left = &Node{Value: val}
		} else {
			sugar.Infof("Added node with value %d successfully!", val)
			t.left.Insert(val)
		}
	}
	return
}

func (t *Node) Inorder() {
	if t == nil {
		return
	}
	t.left.Inorder()
	println(t.Value)
	t.right.Inorder()
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
