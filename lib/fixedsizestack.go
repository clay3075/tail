package FixedSizeStack

import (
	"container/list"
	"fmt"
)

type FixedSizeStack struct {
	stack   *list.List
	maxSize int
}

func New(size int) *FixedSizeStack {
	stack := FixedSizeStack{}
	stack.stack = list.New()
	stack.maxSize = size
	return &stack
}

func (FS *FixedSizeStack) Push(item interface{}) {
	if (FS.stack.Len() >= FS.maxSize) {
		FS.stack.Remove(FS.stack.Front())
	}
	FS.stack.PushBack(item)
}


func (FS *FixedSizeStack) Print() {
	for e := FS.stack.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}