//��Ŀ����
//����һ���ַ����ж��Ƿ��ǻ��Ĵ��������Ĵ�����һ�������ͷ�����һ�����ַ��������硰level�����ߡ�noon���ȵȾ��ǻ��Ĵ���
//
//����
//һ���ַ��������Ȳ�����255��
//
//���
//����ǻ��Ĵ��������YES�������������NO����
//
//�������� Copy
//12321
//������� Copy
//YES

#include <cstdio>
#include <cstring>


int main()
{
	char str[260];
	scanf_s("%s", str);
	int length = strlen(str);
	for (int i(0), j(length - 1); i <= j; i++, j--)
	{
		if (str[i] != str[j])
		{
			printf("NO\n");
			break;
		}
	}
	printf("YES\n");
	return 0;
}