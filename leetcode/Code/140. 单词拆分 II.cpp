﻿//给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。
//
//说明：
//
//分隔时可以重复使用字典中的单词。
//你可以假设字典中没有重复的单词。
//示例 1：
//
//输入 :
//s = "catsanddog"
//wordDict = ["cat", "cats", "and", "sand", "dog"]
//输出 :
//	[
//		"cats and dog",
//		"cat sand dog"
//	]
//示例 2：
//
//输入 :
//s = "pineapplepenapple"
//wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
//输出 :
//	[
//		"pine apple pen apple",
//		"pineapple pen apple",
//		"pine applepen apple"
//	]
//解释 : 注意你可以重复使用字典中的单词。
//示例 3：
//
//输入 :
//s = "catsandog"
//wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出 :
//	[]
//
//来源：力扣（LeetCode）
//链接：https ://leetcode-cn.com/problems/word-break-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

#include <vector>
#include <string>
#include <iostream>
#include <unordered_set>

using namespace::std;

class Solution {
public:
    vector<string> ans;
    vector<string> wordBreak(string s, vector<string>& wordDict) {
        unordered_set<string> set;
        unordered_set<int> length;
        for (auto itor : wordDict)
        {
            set.insert(itor);
            length.insert(itor.length());
        }
        wordback(s, "", set, length);
        return this->ans;
    }

    void wordback(string s, string ss, unordered_set<string> set, unordered_set<int> length)
    {
        if (s.length() == 0)
        {
            ans.push_back(ss.substr(1));
            return;
        }
        for (auto len : length)
        {
            if (s.size() < len)
                continue;
            string word = s.substr(0, len);
            if (set.find(word) != set.end())
            {
                string ss_ = ss + " " + word;
                wordback(s.substr(len), ss_, set, length);
            }
        }
    }
};


int main()
{
    Solution solution = Solution();
    string s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
    vector<string> wordDict = { "a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa" };
    vector<string>ans = solution.wordBreak(s, wordDict);
    for (auto seq : ans)
        cout << seq << endl;
}