/**
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

若队列为空，pop_front 和 max_value 需要返回 -1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

type MaxQueue struct {
	size          int
	queue         []int
	maxValueQueue []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if this.size > 0 {
		return this.maxValueQueue[0]
	} else {
		return -1
	}
}

func (this *MaxQueue) Push_back(value int) {
	if this.size == 0 {
		this.size = 1
		this.queue = []int{value}
		this.maxValueQueue = []int{value}
	} else {
		this.queue = append(this.queue, value)
		this.size++

		if value > this.maxValueQueue[0] {
			this.maxValueQueue = []int{value}
		} else {
			i := len(this.maxValueQueue) - 1
			for i >= 0 {
				if this.maxValueQueue[i] < value {
					i--
				} else {
					break
				}
			}

			this.maxValueQueue = this.maxValueQueue[:i+1]
			this.maxValueQueue = append(this.maxValueQueue, value)
		}
	}
}

func (this *MaxQueue) Pop_front() int {
	if this.size == 0 {
		return -1
	}

	front := this.queue[0]
	this.queue = this.queue[1:]
	this.size--

	if front == this.maxValueQueue[0] {
		this.maxValueQueue = this.maxValueQueue[1:]
	}
	return front
}
