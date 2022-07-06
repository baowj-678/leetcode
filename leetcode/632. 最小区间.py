# 你有 k 个升序排列的整数数组。找到一个最小区间，使得 k 个列表中的每个列表至少有一个数包含在其中。

# 我们定义如果 b-a < d-c 或者在 b-a == d-c 时 a < c，则区间 [a,b] 比 [c,d] 小。

# 示例 1:

# 输入:[[4,10,15,24,26], [0,9,12,20], [5,18,22,30]]
# 输出: [20,24]
# 解释:
# 列表 1：[4, 10, 15, 24, 26]，24 在区间 [20,24] 中。
# 列表 2：[0, 9, 12, 20]，20 在区间 [20,24] 中。
# 列表 3：[5, 18, 22, 30]，22 在区间 [20,24] 中。
# 注意:

# 给定的列表可能包含重复元素，所以在这里升序表示 >= 。
# 1 <= k <= 3500
# -105 <= 元素的值 <= 105
# 对于使用Java的用户，请注意传入类型已修改为List<List<Integer>>。重置代码模板后可以看到这项改动。

# 来源：力扣（LeetCode）
# 链接：https://leetcode-cn.com/problems/smallest-range-covering-elements-from-k-lists
# 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
import heapq


class Solution:
    def smallestRange(self, nums):
        left, right = int(-9e5), int(9e5)
        maxValue = max([vec[0] for vec in nums])
        priorityQueue = [(vec[0], i, 0) for i, vec in enumerate(nums)]
        heapq.heapify(priorityQueue)
        while True:
            minValue, r, i = heapq.heappop(priorityQueue)
            if(right - left > maxValue - minValue):
                left, right = minValue, maxValue
                print(minValue, maxValue)
            if(i >= len(nums[r]) - 1):
                break
            else:
                maxValue = max(maxValue, nums[r][i + 1])
                heapq.heappush(priorityQueue, (nums[r][i + 1], r, i + 1))
        return [left, right]


if __name__ == '__main__':
    solution = Solution()
    nums = [[4, 10, 15, 24, 26], [0, 9, 12, 20], [5, 18, 22, 30]]
    print(solution.smallestRange(nums))
