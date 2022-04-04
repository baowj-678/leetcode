/*
����һ������ nums����дһ������������ 0 �ƶ��������ĩβ��ͬʱ���ַ���Ԫ�ص����˳��

ʾ��:

����: [0,1,0,3,12]
���: [1,3,12,0,0]
˵��:

������ԭ�����ϲ��������ܿ�����������顣
�������ٲ���������
*/

#include <iostream>
#include <vector>

using namespace::std;

void moveZeroes(vector<int>& nums) 
{
	int point2zero = 0;
	int point2nums = 0;
	for (; point2nums < nums.size(); point2nums++)
	{
		if (nums[point2nums] != 0)
		{
			if (point2nums == point2zero)
			{
				point2zero++;
				continue;
			}
			if (nums[point2zero] == 0)
			{
				nums[point2zero] = nums[point2nums];
				nums[point2nums] = 0;
				point2zero++;
			}
			else
			{
				while (point2zero < point2nums && nums[point2zero] == 0)
					point2zero++;
				nums[point2zero] = nums[point2nums];
				nums[point2nums] = 0;
				point2zero++;
			}
		}
	}
}

void moveZeroesBetter(vector<int>& nums)
{
	int point2ok = 0;
	for (int i(0); i < nums.size(); i++)
	{
		if (nums[i] > 0)
		{
			int temp = nums[i];
			nums[i] = 0;
			nums[point2ok] = temp;
			point2ok++;
		}
	}
}

int main()
{
	int n;
	cin >> n;
	int x;
	vector<int> nums;
	for (int i(0); i < n; i++)
	{
		cin >> x;
		nums.push_back(x);
	}
	moveZeroesBetter(nums);
	for (int i(0); i < n; i++)
	{
		cout << nums[i] << " ";
	}
}

//accepted