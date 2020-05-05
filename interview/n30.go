/*
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。



示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

type MinStack struct {
	arr    []int
	minArr []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.arr = append(this.arr, x)

	minArrayLength := len(this.minArr)
	if minArrayLength == 0 || x <= this.minArr[minArrayLength-1] {
		this.minArr = append(this.minArr, x)
	}
}

func (this *MinStack) Pop() {
	e := this.arr[len(this.arr)-1]
	if e == this.minArr[len(this.minArr)-1] {
		this.minArr = this.minArr[:len(this.minArr)-1]
	}

	this.arr = this.arr[:len(this.arr)-1]
}

func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1]
}

func (this *MinStack) Min() int {
	return this.minArr[len(this.minArr)-1]
}
