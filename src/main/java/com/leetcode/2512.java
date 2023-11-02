package com.leetcode;

import java.util.*;
import java.util.stream.Collectors;

class Solution2512 {
    public List<Integer> topStudents(String[] positive_feedback, String[] negative_feedback, String[] report, int[] student_id, int k) {
        HashSet<String> pos = new HashSet<>(Arrays.asList(positive_feedback));
        HashSet<String> neg = new HashSet<>(Arrays.asList(negative_feedback));
        HashMap<Integer, Integer> scores = new HashMap<>();
        for (int i = 0; i < report.length; i++) {
            String[] ss = report[i].split(" ");
            int score = 0;
            for (String s: ss) {
                if (pos.contains(s)) {
                    score += 3;
                }
                if (neg.contains(s)) {
                    score -= 1;
                }
            }
            scores.put(student_id[i], score);
        }
        List<Integer> res = Arrays.stream(student_id).boxed().collect(Collectors.toList());
        Collections.sort(res, new Comparator<Integer>() {
            @Override
            public int compare(Integer o1, Integer o2) {
                Integer score1 = scores.get(o1);
                Integer score2 = scores.get(o2);
                if (Objects.equals(score1, score2)) {
                    return o1.compareTo(o2);
                }
                return -score1.compareTo(score2);
            }
        });
        return res.subList(0, k);
    }
}