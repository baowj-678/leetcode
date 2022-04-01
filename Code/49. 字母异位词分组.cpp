//����һ���ַ������飬����ĸ��λ�������һ����ĸ��λ��ָ��ĸ��ͬ�������в�ͬ���ַ�����
//
//ʾ�� :
//
//����: ["eat", "tea", "tan", "ate", "nat", "bat"]
//��� :
//	[
//		["ate", "eat", "tea"],
//		["nat", "tan"],
//		["bat"]
//	]
//˵����
//
//���������ΪСд��ĸ��
//�����Ǵ������˳��
//
//��Դ�����ۣ�LeetCode��
//���ӣ�https ://leetcode-cn.com/problems/group-anagrams
//����Ȩ������������С���ҵת������ϵ�ٷ���Ȩ������ҵת����ע��������

#include <iostream>
#include <string>
#include <vector>
#include <unordered_map>
#include <algorithm>

using namespace::std;

class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        unordered_map<string, vector<string>> map;
        vector<vector<string>> ans;
        for (string str : strs)
        {
            string s = str;
            sort(s.begin(), s.end());
            if (map.find(s) != map.end())
            {
                map[s].push_back(str);
            }
            else
            {
                vector<string> tmp = { str };
                map.insert(pair<string, vector<string>>(s, tmp));
            }
        }
        for (auto pa : map)
        {
            ans.push_back(pa.second);
        }
        return ans;
    }
};

int main()
{
    vector<int> tmp = { 1 };
    vector<int>& a = tmp;
    
    vector<vector<int>> A;
    A.push_back(a);
    a.push_back(9);

}