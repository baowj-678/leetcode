/**
* Topic: CSP-2018-12-CIDR�ϲ�
* Author: Bao Wenjie
* Email: bwj_678@qq.com
* Date: 2020/8/19
*/
#include <iostream>
#include <algorithm>
#define	MAX_EDGE	100010
#define	MAX_VETEX	50010
struct Edge {
	int v;
	int u;
	int t;
};
struct Edge edge[MAX_EDGE];
int vertex[MAX_VETEX];
using namespace::std;

// ���������
bool operator < (struct Edge& a, struct Edge& b);
// ���鼯
int find(int x, int vertex[]);
int main()
{
	// ��������
	int n, m, root;
	int v, u, t;
	scanf_s("%d%d%d", &n, &m, &root);
	for (int i(0); i < m; i++)
	{
		scanf_s("%d%d%d", &(edge[i].v), &(edge[i].u), &(edge[i].t));
	}
	// ��ʼ������edge����
	memset(vertex, 0, sizeof(int) * MAX_VETEX);
	sort(edge, edge + m);
	for (int i(0); i <= n; i++)
		vertex[i] = i;
	// kruskal�㷨
	int maxValue = 0;
	for (int i(0); i < m; i++)
	{
		if (find(edge[i].v, vertex) != find(edge[i].u, vertex))
		{
			if (maxValue < edge[i].t)
				maxValue = edge[i].t;
			vertex[find(edge[i].u, vertex)] = find(edge[i].v, vertex);
		}
	}
	printf_s("%d\n", maxValue);
}

/**
* ���������
*/
bool operator < (struct Edge& a, struct Edge& b)
{
	return (a.t < b.t);
}

/**
* ���鼯
*/
int find(int x, int vertex[])
{
	int tmp = x;
	while (tmp != vertex[tmp])
	{
		tmp = vertex[tmp];
	}
	// ѹ��·��
	while (tmp != x)
	{
		int last = vertex[x];
		vertex[x] = tmp;
		x = last;
	}
	return tmp;
}