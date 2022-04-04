﻿/**
 * @author: Bao Wenjie
 * @link: https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
 * @date: 2021/9/4
 */

/* 

用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

示例 1：

输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]
示例 2：

输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]
提示：

1 <= values <= 10000
最多会对 appendTail、deleteHead 进行 10000 次调用

*/

#include <iostream>
#include <stack>

using namespace std;

class CQueue {
public:
    stack<int> output_stack;
    stack<int> input_stack;
    CQueue() {
    	
    }

    void appendTail(int value) {
        input_stack.push(value);
    }

    int deleteHead() {
    	if(output_stack.empty())
    	{
            if (input_stack.empty())
                return -1;
            while (!input_stack.empty())
            {
                output_stack.push(input_stack.top());
                input_stack.pop();
            }
    	}
        int ans = output_stack.top();
        output_stack.pop();
        return ans;
    }
};

int main()
{
    CQueue* cq = new CQueue();
    cq->appendTail(3);
    cq->deleteHead();
    cq->deleteHead();
}
/**
 * Your CQueue object will be instantiated and called as such:
 * CQueue* obj = new CQueue();
 * obj->appendTail(value);
 * int param_2 = obj->deleteHead();
 */