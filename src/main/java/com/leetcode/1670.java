package com.leetcode;

import jdk.nashorn.internal.ir.SplitReturn;

import java.util.LinkedList;

class FrontMiddleBackQueue {
    // first-----head------last  first-------tail--------last
    LinkedList<Integer> head;
    LinkedList<Integer> tail;

    public FrontMiddleBackQueue() {
        this.head = new LinkedList<>();
        this.tail = new LinkedList<>();
    }

    public void pushFront(int val) {
        if (this.head.size() > this.tail.size()) {
            this.tail.addFirst(this.head.removeLast());
        }
        this.head.addFirst(val);
    }

    public void pushMiddle(int val) {
        if (this.head.size() > this.tail.size()) {
            this.tail.addFirst(this.head.removeLast());
        }
        this.head.addLast(val);
    }

    public void pushBack(int val) {
        this.tail.addLast(val);
        if (this.head.size() < this.tail.size()) {
            this.head.addLast(this.tail.removeFirst());
        }
    }

    public int popFront() {
        if (this.head.size() == 0) {
            return -1;
        }
        int res = this.head.removeFirst();
        if (this.head.size() < this.tail.size()) {
            this.head.addLast(this.tail.removeFirst());
        }
        return res;
    }

    public int popMiddle() {
        if (this.head.size() == 0) {
            return -1;
        }
        int res = this.head.removeLast();
        if (this.head.size() < this.tail.size()) {
            this.head.addLast(this.tail.removeFirst());
        }
        return res;
    }

    public int popBack() {
        if (this.tail.size() == 0) {
            if (this.head.size() == 0) {
                return -1;
            } else {
                return this.head.removeLast();
            }
        }
        int res = this.tail.removeLast();
        while(this.head.size() - this.tail.size() > 1) {
            this.tail.addFirst(this.head.removeLast());
        }
        return res;
    }
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * FrontMiddleBackQueue obj = new FrontMiddleBackQueue();
 * obj.pushFront(val);
 * obj.pushMiddle(val);
 * obj.pushBack(val);
 * int param_4 = obj.popFront();
 * int param_5 = obj.popMiddle();
 * int param_6 = obj.popBack();
 */