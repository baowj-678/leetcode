//��Ŀ����
//����һ����n��Ȼ������n����ֵ������ͬ��������һ��ֵx��������ֵ����������е��±꣨��0��ʼ������������������� - 1����
//
//����
//���������ж��飬����n(1 <= n <= 200)����������n������Ȼ������x��
//
//���
//����ÿ������, ����������
//
//�������� Copy
//4
//1 2 3 4
//3
//������� Copy
//2

//link http://codeup.cn/problem.php?cid=100000576&pid=1

#include <cstdio>

const int maxn = 210;
int nums[maxn];
int main()
{
	int n, x;
	for (int i(0); i < n; i++)
	{
		scanf("%d", nums + i);
	}
	scanf("%d", &x);
	for (int i(0); i < n; i++)
		if (nums[i] == x)
			return i;
	return -1;
}