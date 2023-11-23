package com.leetcode;

class Solution1334 {
    public int findTheCity(int n, int[][] edges, int distanceThreshold) {
        int[][] dist = new int[n][n];
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                dist[i][j] = -1;
            }
            dist[i][i] = 0;
        }
        for (int[] edge: edges) {
            dist[edge[0]][edge[1]] = edge[2];
            dist[edge[1]][edge[0]] = edge[2];
        }

        for (int k = 0; k < n; k++) {
            for (int start = 0; start < n; start++) {
                for (int end = start+1; end < n; end++) {
                    if (dist[start][k] != -1 && dist[k][end] != -1) {
                        if (dist[start][end] == -1) {
                            dist[start][end] =  dist[start][k] + dist[k][end];
                        } else {
                            dist[start][end] = Math.min(dist[start][end], dist[start][k] + dist[k][end]);
                        }
                        dist[end][start] = dist[start][end];
                    }
                }
            }
        }

        int res = 0, min = n;
        for (int i = 0; i < n; i++) {
            int cnt = 0;
            for (int j = 0; j < n; j++) {
                if (dist[i][j] != -1 && dist[i][j] <= distanceThreshold) {
                    cnt++;
                }
            }
            if (cnt <= min) {
                min = cnt;
                res = i;
            }
        }
        return res;
    }


    public static void main(String[] args) {
        Solution1334 so = new Solution1334();
        int[][] next = new int[][]{{0,1,2},{0,4,8},{1,2,3},{1,4,2},{2,3,1},{3,4,1}};
        System.out.println(so.findTheCity(5, next, 2));
    }
}