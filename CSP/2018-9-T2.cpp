/**
* Topic: CSP-2018-9-���
* Author: Bao Wenjie
* Email: bwj_678@qq.com
* Date: 2020/8/20
*/


//�������ƣ�	���
//ʱ�����ƣ�	1.0s
//�ڴ����ƣ�	256.0MB
//����������
//��������
//
//СH��СW������һ�����ϣ����˷ֿ���ˣ�������˵Ĺ��̿�������Ϊ��ȥ������һЩ��Ȼ��ȥ�Աߵ�һ���㳡�Ѳ�װ�ϳ������˶�Ҫ��n�ֲˣ�����Ҳ��Ҫװn�γ�������ģ�����СH��˵��n�����ཻ��ʱ���[a1, b1], [a2, b2]...[an, bn]��װ��������СW��˵��n�����ཻ��ʱ���[c1, d1], [c2, d2]...[cn, dn]��װ�������У�һ��ʱ���[s, t]��ʾ���Ǵ�ʱ��s��ʱ��t���ʱ�䣬ʱ��Ϊt - s��
//���������Ǻ����ѣ����Ƕ��ڹ㳡��װ����ʱ������죬������֪�����ǿ����Ķ೤ʱ�䡣
//
//�����ʽ
//
//����ĵ�һ�а���һ��������n����ʾʱ��ε�������
//������n��ÿ��������ai��bi������СH�ĸ���װ����ʱ��Ρ�
//������n��ÿ��������ci��di������СW�ĸ���װ����ʱ��Ρ�
//
//�����ʽ
//
//���һ�У�һ������������ʾ���˿����Ķ೤ʱ�䡣
//
//��������
//
//4
//1 3
//5 6
//9 13
//14 15
//2 4
//5 7
//10 11
//13 14
//
//�������
//
//3
//
//���ݹ�ģ��Լ��
//
//�������е�����������1 �� n �� 2000, ai < bi < ai + 1��ci < di < ci + 1, 
//�������е�i(1 �� i �� n)�У�1 �� ai, bi, ci, di �� 1000000��
#include <iostream>
#include <vector>
#include <algorithm>
using namespace::std;
struct Time
{
	int begin;
	int end;
};
typedef Time* pTime;
bool cmp (pTime& a, pTime& b)
{
	if (a->begin < b->begin)
		return true;
	else if (a->begin == b->begin)
		return a->end < b->end;
	else
		return false;
}
int main()
{
	// ����
	int n;
	scanf_s("%d", &n);
	vector<struct Time*> H(n);
	vector<struct Time*> W(n);
	for (int i(0); i < n; i++)
	{
		struct Time* time = new struct Time;
		scanf_s("%d%d", &(time->begin), &(time->end));
		H[i] = time;
	}
	for (int i(0); i < n; i++)
	{
		struct Time* time = new struct Time;
		scanf_s("%d%d", &(time->begin), &(time->end));
		W[i] = time;
	}
	// ����
	sort(H.begin(), H.end(), cmp);
	sort(W.begin(), W.end(), cmp);
	// ����
	long long sum = 0;
	int h, w;
	for (h = 0, w = 0; h < n, w < n; )
	{
		if (H[h]->begin <= W[w]->begin)
		{
			if (H[h]->end >= W[w]->end)
			{
				sum += (W[w]->end - W[w]->begin);
				w++;
			}
			else if (H[h]->end >= W[w]->begin)
			{
				sum += (H[h]->end - W[w]->begin);
				h++;
			}
			else
			{
				h++;
			}
		}
		else if(H[h]->begin > W[w]->begin)
		{
			if (W[w]->end >= H[h]->end)
			{
				sum += (H[h]->end - H[h]->begin);
				h++;
			}
			else if (W[w]->end >= H[h]->begin)
			{
				sum += (W[w]->end - H[h]->begin);
				w++;
			}
			else
			{
				w++;
			}
		}
	}
	printf("%lld\n", sum);


}