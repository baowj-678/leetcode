/**
* Author: Bao Wenjie
* Description: ��С������
* Date: 2021/1/18
*/

#include <cstdio>

int gcd(int a, int b)
{
	if (b == 0) return a;
	return gcd(b, a % b);
}

int lcm(int a, int b)
{
	return a * b / gcd(a, b);
}