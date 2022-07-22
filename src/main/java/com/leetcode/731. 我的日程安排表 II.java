package com.leetcode;


import java.util.*;

class MyCalendarTwo {
    private TreeMap<Integer, Integer> map;

    public MyCalendarTwo() {
        this.map = new TreeMap<>();
    }

    public boolean book(int start, int end) {
        int booked = 0;
        this.add(start, 1);
        this.add(end, -1);
        for (Map.Entry<Integer, Integer> entry : this.map.entrySet()) {
            booked += entry.getValue();
            if (booked >= 3) {
                this.add(start, -1);
                this.add(end, 1);
                return false;
            }
        }
        return true;
    }

    private void add(int key, int value) {
        Integer integer = this.map.get(key);
        if (integer == null) {
            this.map.put(key, value);
        } else {
            this.map.put(key, integer + value);
        }
    }

    public static void main(String[] args) {
//        MyCalendarTwo myCalendarTwo = new MyCalendarTwo();
//        System.out.println(myCalendarTwo.book(10, 20));
//        System.out.println(myCalendarTwo.book(50, 60));
//        System.out.println(myCalendarTwo.book(10, 40));
//        System.out.println(myCalendarTwo.book(5, 15));
//        System.out.println(myCalendarTwo.book(5, 10));
//        System.out.println(myCalendarTwo.book(25, 55));
//    }

        TreeMap<Integer, Integer> map = new TreeMap<>();
        map.put(1, 2);
        map.put(3, 2);
        map.put(9, 7);
        map.put(0, 3);
        map.put(100, 3);
        for (Map.Entry<Integer, Integer> entry : map.entrySet()) {
            System.out.println(entry.getKey() + ": " + entry.getValue());
        }
    }
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * MyCalendarTwo obj = new MyCalendarTwo();
 * boolean param_1 = obj.book(start,end);
 */

