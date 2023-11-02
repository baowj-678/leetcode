package com.leetcode;

import javafx.util.Pair;

import java.util.*;

class StockPrice {
    HashMap<Integer, Integer> map;
    Queue<Pair<Integer, Integer>> maxQ;
    Queue<Pair<Integer, Integer>> minQ;
    int curTimestamp;

    public StockPrice() {
        this.map = new HashMap<>();
        this.minQ = new PriorityQueue<>(
                (o1, o2) -> -o2.getValue().compareTo(o1.getValue())
        );
        this.maxQ = new PriorityQueue<>(
                (o1, o2) -> o2.getValue().compareTo(o1.getValue())
        );
        this.curTimestamp = 0;
    }

    public void update(int timestamp, int price) {
        if (timestamp > this.curTimestamp) {
            this.curTimestamp = timestamp;
        }
        this.map.put(timestamp, price);
        this.maxQ.add(new Pair<>(timestamp, price));
        this.minQ.add(new Pair<>(timestamp, price));
    }

    public int current() {
        return this.map.get(this.curTimestamp);
    }

    public int maximum() {
        while(true) {
            Pair<Integer, Integer> peek = this.maxQ.peek();
            if (peek == null) {
                return -1;
            }
            if (Objects.equals(peek.getValue(), this.map.get(peek.getKey()))) {
                return peek.getValue();
            } else {
                this.maxQ.remove();
            }
        }
    }

    public int minimum() {
        while(true) {
            Pair<Integer, Integer> peek = this.minQ.peek();
            if (peek == null) {
                return -1;
            }
            if (Objects.equals(peek.getValue(), this.map.get(peek.getKey()))) {
                return peek.getValue();
            } else {
                this.minQ.remove();
            }
        }
    }

    public static void main(String[] args) {
        StockPrice obj = new StockPrice();
        obj.update(3, 9651);
        obj.update(3, 6897);
        System.out.println(obj.maximum());
    }
}
