//����һ�������������е������ڵ��ż���ڵ�ֱ�����һ����ע�⣬����������ڵ��ż���ڵ�ָ���ǽڵ��ŵ���ż�ԣ������ǽڵ��ֵ����ż�ԡ�
//
//�볢��ʹ��ԭ���㷨��ɡ�����㷨�Ŀռ临�Ӷ�ӦΪ O(1)��ʱ�临�Ӷ�ӦΪ O(nodes)��nodes Ϊ�ڵ�������
//
//ʾ�� 1:
//
//����: 1->2->3->4->5->NULL
//��� : 1->3->5->2->4->NULL
//ʾ�� 2 :
//
//	���� : 2->1->3->5->6->4->7->NULL
//	��� : 2->3->6->7->1->5->4->NULL
//	˵�� :
//
//Ӧ�����������ڵ��ż���ڵ�����˳��
//����ĵ�һ���ڵ���Ϊ�����ڵ㣬�ڶ����ڵ���Ϊż���ڵ㣬�Դ����ơ�
//
//��Դ�����ۣ�LeetCode��
//���ӣ�https ://leetcode-cn.com/problems/odd-even-linked-list
//����Ȩ������������С���ҵת������ϵ�ٷ���Ȩ������ҵת����ע��������


//Definition for singly-linked list.
#include <iostream>

using namespace::std;

struct ListNode {
     int val;
     ListNode *next;
     ListNode() : val(0), next(nullptr) {}
     ListNode(int x) : val(x), next(nullptr) {}
     ListNode(int x, ListNode *next) : val(x), next(next) {}
};
 
class Solution {
public:
    ListNode* oddEvenList(ListNode* head) {
        ListNode* even_head = NULL, *even_tail = NULL, * odd_tail = NULL, *odd_head = NULL;
        bool is_even = true;
        while (head)
        {
            if (is_even)
            {
                if (even_head == NULL)
                {
                    even_head = head;
                    even_tail = head;
                }
                else
                {
                    even_tail->next = head;
                    even_tail = head;
                }
                head = head->next;
                even_tail->next = NULL;
                is_even = false;
            }
            else
            {
                if (odd_head == NULL)
                {
                    odd_head = odd_tail = head;
                }
                else
                {
                    odd_tail->next = head;
                    odd_tail = head;
                }
                head = head->next;
                odd_tail->next = NULL;
                is_even = true;
            }
        }
        if (even_tail == NULL)
            return even_head;
        even_tail->next = odd_head;
        return even_head;
    }
};