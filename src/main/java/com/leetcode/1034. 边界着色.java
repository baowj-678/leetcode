/**
 * @author Bao WJ
 * @date 2021/12/7
 * @link https://leetcode-cn.com/problems/coloring-a-border/
 */
package com.leetcode;

import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;
import java.util.ArrayList;
import java.util.LinkedList;
import java.util.Queue;

class Solution1034 {
    public int[][] colorBorder(int[][] grid, int row, int col, int color) {
        Queue<Integer[]> Q = new LinkedList<>();
        int colorTmp = color;
        color = -2;
        int preColor = grid[row][col];
        grid[row][col] = color;
        Q.offer(new Integer[] {row, col});
        while (!Q.isEmpty()) {
            Integer[] tmp =  Q.poll();
            row = tmp[0];
            col = tmp[1];
            int num = 0;
            if (row > 0 && (grid[row-1][col] == preColor || grid[row-1][col] == color || grid[row-1][col] == -1)) {
                num++;
                if (grid[row-1][col] == preColor) {
                    grid[row-1][col] = color;
                    Q.offer(new Integer[] {row-1, col});
                }
            }
            if (col > 0 && (grid[row][col-1] == preColor || grid[row][col-1] == color || grid[row][col-1] == -1)) {
                num++;
                if (grid[row][col-1] == preColor) {
                    grid[row][col-1] = color;
                    Q.offer(new Integer[] {row, col-1});
                }

            }
            if (row + 1 < grid.length && (grid[row+1][col] == preColor || grid[row+1][col] == color || grid[row+1][col] == -1)) {
                num++;
                if (grid[row+1][col] == preColor) {
                    grid[row+1][col] = color;
                    Q.offer(new Integer[] {row+1, col});
                }

            }
            if (col + 1 < grid[0].length && (grid[row][col+1] == preColor || grid[row][col+1] == color || grid[row][col+1] == -1)) {
                num++;
                if (grid[row][col+1] == preColor) {
                    grid[row][col+1] = color;
                    Q.offer(new Integer[] {row, col+1});
                }

            }
            if (num == 4) {
                grid[row][col] = -1;
            }
        }
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[0].length; j++) {
                if (grid[i][j] == -1) {
                    grid[i][j] = preColor;
                } else if (grid[i][j] == color) {
                    grid[i][j] = colorTmp;
                }
            }
        }
        return grid;
    }

//    public static void main(String[] args) throws InstantiationException, IllegalAccessException, NoSuchMethodException, InvocationTargetException {
//        int[][] grid = new int[][] {{1,1},{1,2}};
//        Class<Solution> solution = Solution.class;
//        Solution instance = solution.getDeclaredConstructor(null).newInstance();
//        Method colorBorder = solution.getMethod("colorBorder", int[][].class, int.class, int.class, int.class);
//        colorBorder.invoke(instance, grid, 0, 0, 3);
//    }
}