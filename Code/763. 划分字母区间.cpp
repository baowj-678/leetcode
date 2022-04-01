//�ַ��� S ��Сд��ĸ��ɡ�����Ҫ������ַ�������Ϊ�����ܶ��Ƭ�Σ�ͬһ����ĸֻ����������е�һ��Ƭ�Ρ�����һ����ʾÿ���ַ���Ƭ�εĳ��ȵ��б�
//
//
//
//ʾ�� 1��
//
//���룺S = "ababcbacadefegdehijhklij"
//�����[9, 7, 8]
//���ͣ�
//���ֽ��Ϊ "ababcbaca", "defegde", "hijhklij"��
//ÿ����ĸ��������һ��Ƭ���С�
//�� "ababcbacadefegde", "hijhklij" �Ļ����Ǵ���ģ���Ϊ���ֵ�Ƭ�������١�
//
//
//��ʾ��
//
//S�ĳ�����[1, 500]֮�䡣
//Sֻ����Сд��ĸ 'a' �� 'z' ��
//
//��Դ�����ۣ�LeetCode��
//���ӣ�https ://leetcode-cn.com/problems/partition-labels
//����Ȩ������������С���ҵת������ϵ�ٷ���Ȩ������ҵת����ע��������
#include <vector>
#include <iostream>

using namespace::std;

class Solution {
public:
    vector<int> partitionLabels(string S) {
        int End[26];
        for (int i = 0; i < S.length(); i++)
        {
            End[S[i] - 'a'] = i;
        }
        vector<int> ans;
        int start = 0, end = 0;
        for (int i(0); i < S.length(); i++)
        {
            int end = max(end, End[S[i] - 'a']);
            if (end == i)
            {
                ans.push_back(end - start + 1);
                start = end + 1;
            }
        }
        return ans;
    }
};

int main()
{
    string a = "ababcbacadefegdehijhklij";
    Solution b = Solution();
    vector<int> ans = b.partitionLabels(a);
    for (int i(0); i < ans.size(); i++)
    {
        printf("%d ", ans[i]);
    }
}