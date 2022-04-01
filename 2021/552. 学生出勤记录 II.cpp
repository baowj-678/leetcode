/**
 * @author: Bao Wenjie
 * @link: https://leetcode-cn.com/problems/student-attendance-record-ii/
 * @date: 2021/8/20
 */
#include <iostream>

using namespace::std;

class Solution {
public:
    static constexpr int MOD = 1'000'000'007;

    int checkRecord(int n) {
        vector<vector<vector<int>>> dp(n + 1, vector<vector<int>>(2, vector<int>(3)));  // ���ȣ�A ����������β���� L ������
        dp[0][0][0] = 1;
        for (int i = 1; i <= n; i++) {
            // �� P ��β������
            for (int j = 0; j <= 1; j++) {
                for (int k = 0; k <= 2; k++) {
                    dp[i][j][0] = (dp[i][j][0] + dp[i - 1][j][k]) % MOD;
                }
            }
            // �� A ��β������
            for (int k = 0; k <= 2; k++) {
                dp[i][1][0] = (dp[i][1][0] + dp[i - 1][0][k]) % MOD;
            }
            // �� L ��β������
            for (int j = 0; j <= 1; j++) {
                for (int k = 1; k <= 2; k++) {
                    dp[i][j][k] = (dp[i][j][k] + dp[i - 1][j][k - 1]) % MOD;
                }
            }
        }
        int sum = 0;
        for (int j = 0; j <= 1; j++) {
            for (int k = 0; k <= 2; k++) {
                sum = (sum + dp[n][j][k]) % MOD;
            }
        }
        return sum;
    }
};

int main()
{
    Solution solution;
    cout << solution.checkRecord(5);
}