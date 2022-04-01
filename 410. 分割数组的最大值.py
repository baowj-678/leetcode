# 给定一个非负整数数组和一个整数 m，你需要将这个数组分成 m 个非空的连续子数组。设计一个算法使得这 m 个子数组各自和的最大值最小。

# 注意:
# 数组长度 n 满足以下条件:

# 1 ≤ n ≤ 1000
# 1 ≤ m ≤ min(50, n)
# 示例:

# 输入:
# nums = [7,2,5,10,8]
# m = 2

# 输出:
# 18

# 解释:
# 一共有四种方法将nums分割为2个子数组。
# 其中最好的方式是将其分为[7,2,5] 和 [10,8]，
# 因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。

# 来源：力扣（LeetCode）
# 链接：https://leetcode-cn.com/problems/split-array-largest-sum
# 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


class Solution:
    def splitArray(self, nums, m: int) -> int:
        low = max(nums)
        high = sum(nums)
        while(low < high):
            mid = (low + high) // 2
            cnt = 0
            sub_sum = 0
            for num in nums:
                sub_sum += num
                if(sub_sum > mid):
                    cnt += 1
                    sub_sum = num
            if(sub_sum > 0):
                cnt += 1
            if(cnt > m):
                low = mid + 1
            else:
                high = mid
        return low


if __name__ == '__main__':
    solution = Solution()
    nums = [7, 8, 8, 1]
    print(solution.splitArray(nums, 3))
