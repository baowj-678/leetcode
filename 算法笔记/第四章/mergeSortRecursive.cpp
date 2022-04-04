/**
* Description: �鲢���򣨵ݹ�ʵ�֣�
* Author: Bao Wenjie
* Date: 2021/1/17
*/


#include <vector>

using namespace::std;
/**
* ������A��[L1, R1],[L2, R2]����ϲ�Ϊ��������
*/
void merge(vector<int> &A, int L1, int R1, int L2, int R2)
{
	int i = L1, j = L2, idx = 0;
	vector<int> tmp(R2 - L1 + 10);
	while (i <= R1 && j <= R2)
	{
		if (A[i] <= A[j])
			tmp[idx++] = A[i++];
		else
			tmp[idx++] = A[j++];
	}
	while (i <= R1) tmp[idx++] = A[i++];
	while (j <= R2) tmp[idx++] = A[j++];
	for (int i = 0; i < idx; i++)
	{
		A[L1 + i] = tmp[i];
	}
}

/**
* ������[left, right]����
*/
void mergeSort(vector<int>& A, int left, int right)
{
	if (left < right)
	{
		int mid = (left + right) / 2;
		mergeSort(A, left, mid);
		mergeSort(A, mid + 1, right);
		merge(A, left, mid, mid + 1, right);
	}
}
