#include <stdio.h>
#include <vector>
#include <algorithm>
using namespace::std;

int jump(vector<int>& nums)
{
    int ans = 0;
    int start = 0;
    int end = 1;
    while (end < nums.size())
    {
        int maxPos = 0;
        for (int i = start; i < end; i++)
        {
            // ��������Զ�ľ���
            maxPos = max(maxPos, i + nums[i]);
        }
        start = end;      // ��һ�������㷶Χ��ʼ�ĸ���
        end = maxPos + 1; // ��һ�������㷶Χ�����ĸ���
        ans++;            // ��Ծ����
    }
    return ans;
}

int main()
{

}