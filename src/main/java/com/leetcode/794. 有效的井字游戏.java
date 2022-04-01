/**
 * @author BaoWJ
 * @date 2021/12/9
 * @link https://leetcode-cn.com/problems/valid-tic-tac-toe-state/
 */
package com.leetcode;

class Solution794 {
    public boolean validTicTacToe(String[] board) {
        int xNum = 0, oNum = 0;
        char[][] charBoard = new char[3][3];
        for (int i = 0; i < board.length; i++) {
            String str = board[i];
            for (int j = 0; j < str.length(); j++) {
                char c = str.charAt(j);
                charBoard[i][j] = c;
                if (c == 'X') {
                    xNum++;
                } else if (c == 'O') {
                    oNum++;
                }
            }
        }
        if (xNum - oNum > 1 || oNum - xNum > 0) {
            return false;
        }

        if (xNum > 0) {
            xNum = 1;
        }
        if (oNum  > 0) {
            oNum = 1;
        }
        if (isVictory(charBoard, ' ')) {
            return false;
        }

        for (int i = 0; i < charBoard.length; i++) {
            for (int j = 0; j < charBoard[i].length; j++) {
                if (charBoard[i][j] != ' ') {
                    char tmp = charBoard[i][j];
                    charBoard[i][j] = ' ';
                    if (!isVictory(charBoard, tmp)) {
                        if (tmp == 'X') {
                            xNum = 0;
                        } else if (tmp == 'O') {
                            oNum = 0;
                        }
                    }
                    charBoard[i][j] = tmp;
                }
            }
        }
        if (xNum == 0 && oNum == 0) {
            return true;
        }
        return false;
    }

    boolean isVictory(char[][] board, char flag) {
        int xNum = 0, oNum = 0;
        for (int i = 0; i < board[0].length; i++) {
            if (board[1][i] == board[0][i] &&
                    board[2][i] == board[0][i]) {
                if (board[0][i] == 'X') {
                    xNum++;
                } else if (board[0][i] == 'O') {
                    oNum++;
                }
            }
        }
        for (int i = 0; i < board.length; i++) {
            if (board[i][0] == board[i][1] &&
                    board[i][0] == board[i][2]) {
                if (board[i][0] == 'X') {
                    xNum++;
                } else if (board[i][0] == 'O') {
                    oNum++;
                }
            }
        }
        if (board[0][0] == board[1][1] &&
                board[0][0] == board[2][2]) {
            if (board[1][1] == 'X') {
                xNum++;
            } else if (board[1][1] == 'O') {
                oNum++;
            }
        }
        if (board[0][2] == board[1][1] &&
                board[0][2] == board[2][0]) {
            if (board[1][1] == 'X') {
                xNum++;
            } else if (board[1][1] == 'O') {
                oNum++;
            }
        }
        if (flag == 'X') {
            return xNum > 0;
        } else if (flag == 'O') {
            return oNum > 0;
        }
        return (xNum > 0) && (oNum > 0);
    }
}