//�����������飬��дһ���������������ǵĽ�����
//
//
//
//ʾ�� 1��
//
//���룺nums1 = [1, 2, 2, 1], nums2 = [2, 2]
//�����[2]
//ʾ�� 2��
//
//���룺nums1 = [4, 9, 5], nums2 = [9, 4, 9, 8, 4]
//�����[9, 4]
//
//
//˵����
//
//�������е�ÿ��Ԫ��һ����Ψһ�ġ�
//���ǿ��Բ�������������˳��
//
//��Դ�����ۣ�LeetCode��
//���ӣ�https ://leetcode-cn.com/problems/intersection-of-two-arrays
//����Ȩ������������С���ҵת������ϵ�ٷ���Ȩ������ҵת����ע��������
#include <iostream>
#include <vector>
#include <algorithm>
using namespace::std;


class Solution {
public:
    vector<int> intersection(vector<int>& nums1, vector<int>& nums2) {
        vector<int> ans;
        sort(nums1.begin(), nums1.end());
        sort(nums2.begin(), nums2.end());
        int i = 0, j = 0;
        for (i = 0, j = 0; i < nums1.size() && j < nums2.size();)
        {
            if (nums1[i] < nums2[j])
                i++;
            else if (nums1[i] > nums2[j])
                j++;
            else
            {
                if (ans.empty() || ans.back() != nums1[i])
                {
                    ans.push_back(nums1[i]);
                }
                i++;
                j++;
            }
        }
        return ans;
    }
};

int main()
{
    Solution solution = Solution();
    vector<int> a = { 1, 2, 3 };
    vector<int> b = { 2, 3, 4 };
    solution.intersection(a, b);
}