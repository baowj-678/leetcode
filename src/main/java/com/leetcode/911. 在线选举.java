/**
 * @author BaoWJ
 * @link https://leetcode-cn.com/problems/online-election/
 * @date 2021/12/14
 */
package com.leetcode;


import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

class TopVotedCandidate {

    private int[] persons;
    private int[] times;
    private int[] topPerson;

    public TopVotedCandidate(int[] persons, int[] times) {
        this.persons = persons;
        this.times = times;
        this.topPerson = new int[persons.length];
        Map<Integer, Integer> map = new HashMap();
        int topVote = 0;
        int topP = -1;
        for (int i = 0; i < persons.length; i++) {
            int tmpP = persons[i];
            if (!map.containsKey(persons[i])) {
                map.put(persons[i], 0);
            }
            int tmpVote = map.get(tmpP) + 1;
            if (tmpVote >= topVote) {
                topVote = tmpVote;
                topP = tmpP;
            }
            map.put(tmpP, tmpVote);
            topPerson[i] = topP;
        }
    }

    public int q(int t) {
        int left = 0, right = times.length - 1;
        while (left < right) {
            int mid = (left + right + 1) / 2;
            if (times[mid] == t) {
                right = mid;
                break;
            } else if (times[mid] > t) {
                right = mid - 1;
            } else if (times[mid] < t) {
                left = mid;
            }
        }
        return topPerson[right];
    }
}