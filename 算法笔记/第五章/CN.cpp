/**
* Description: ���������
* Author: Bao Wenjie
* Date: 2021/1/21
*/


/****************************** ��������� ********************************************/
/**
* ����һ��ֱ�Ӽ���
*/
long long C_1(long long n, long long m)
{
	long long ans = 1;
	for (long long i = 1; i <= n; i++)
		ans *= i;
	for (long long i = 1; i <= m; i++)
		ans /= i;
	for (long long i = 1; i <= n - m; i++)
		ans /= i;
	return ans;
}

/**
* ��������ͨ�����ƹ�ʽ����
*/
long long C_2(long long n, long long m)
{
	if (m == 0 || m == n)
		return 1;
	return C_2(n - 1, m) + C_2(n - 1, m - 1);
}

/**
* ������ + ��¼��ʷֵ
*/
long long res[100][100] = { 0 };
long long C_2_dp(long long n, long long m)
{
	if (m == 0 || m == n)
		return 1;
	if (res[n][m] != 0)
		return res[n][m];
	return res[n][m] = C_2_dp(n - 1, m) + C_2_dp(n - 1, m - 1);
}


/**
* ��������������η���
*/
long long C_3(long long n, long long m)
{
	long long ans = 1;
	for (long long i = 1; i <= m; i++)
	{
		ans = ans * (n - m + i) / i;
	}
	return ans;
}



/********************************** �����������ģ ***************************************/
int Lucas(const int n, const int m, const int p)
{
	if (m == 0)
		return 1;
	return C_1(n % p, m % p) * Lucas(n / p, m / p, p) % p;
}