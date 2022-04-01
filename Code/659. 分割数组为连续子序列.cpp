//����һ��������������������� num�����ܰ����ظ����֣������㽫���Ƿָ��һ�����������У�����ÿ�������ж���������������ҳ�������Ϊ 3 ��
//
//���������������ָ�򷵻� true �����򣬷��� false ��
//
//
//
//ʾ�� 1��
//
//���� : [1, 2, 3, 3, 4, 5]
//��� : True
//���� :
//����Էָ��������������������:
//1, 2, 3
//3, 4, 5
//
//
//ʾ�� 2��
//
//���� : [1, 2, 3, 3, 4, 4, 5, 5]
//��� : True
//���� :
//����Էָ��������������������:
//1, 2, 3, 4, 5
//3, 4, 5
//
//
//ʾ�� 3��
//
//���� : [1, 2, 3, 4, 4, 5]
//��� : False
//
//
//��ʾ��
//
//��������鳤�ȷ�ΧΪ[1, 10000]
//
//��Դ�����ۣ�LeetCode��
//���ӣ�https ://leetcode-cn.com/problems/split-array-into-consecutive-subsequences
//����Ȩ������������С���ҵת������ϵ�ٷ���Ȩ������ҵת����ע��������
#include <iostream>
#include <vector>
#include <unordered_map>

using namespace::std;


class Solution {
public:
    bool isPossible(vector<int>& nums) {
        unordered_map<int, int> countMap;
        unordered_map<int, int> endMap;
        for (auto& x : nums) {
            int count = countMap[x] + 1;
            countMap[x] = count;
        }
        for (auto& x : nums) {
            int count = countMap[x];
            if (count > 0) {
                int prevEndCount = endMap[x - 1];
                if (prevEndCount > 0) {
                    countMap[x] = count - 1;
                    endMap[x - 1] = prevEndCount - 1;
                    endMap[x] = endMap[x] + 1;
                }
                else {
                    int count1 = countMap[x + 1];
                    int count2 = countMap[x + 2];
                    if (count1 > 0 && count2 > 0) {
                        countMap[x] = count - 1;
                        countMap[x + 1] = count1 - 1;
                        countMap[x + 2] = count2 - 1;
                        endMap[x + 2] = endMap[x + 2] + 1;
                    }
                    else {
                        return false;
                    }
                }
            }
        }
        return true;
    }
};


int main()
{
    vector<int> a = {1,2,2,3,4,5,6,7,8};
    Solution so;
    cout << so.isPossible(a);
}