package stack

import "sync"

//type Item interface{}
//string改Item

type StringStack struct {
	items []string
	mutex sync.Mutex
}

//将元素压入栈顶
func (stack *StringStack) Push(item string) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock() //在性能方面，使用defer比在return语句之前解锁堆栈代价更大。

	stack.items = append(stack.items, item)
}

//移除并返回栈顶元素
func (stack *StringStack) Pop() string {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	if len(stack.items) == 0 {
		//return nil
		return ""
	}

	lastItem := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]

	return lastItem
}

//检查堆栈是否为空
func (stack *StringStack) IsEmpty() bool {
	stack.mutex.Lock() //这里加锁是因为len函数是不安全的。
	defer stack.mutex.Unlock()

	return len(stack.items) == 0
}

//清空堆栈
func (stack *StringStack) Clear() {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	stack.items = nil
}

//生成并返回原堆栈的副本
func (stack *StringStack) Dump() []string {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	var copiedStack = make([]string, len(stack.items))
	copy(copiedStack, stack.items)

	return copiedStack
}

//返回栈顶元素
func (stack *StringStack) Top() string {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	if len(stack.items) == 0 {
		//return nil
		return ""
	}

	return stack.items[len(stack.items)-1]
}
