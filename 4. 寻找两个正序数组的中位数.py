# 给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。

# 请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

# 你可以假设 nums1 和 nums2 不会同时为空。

# 示例 1:

# nums1 = [1, 3]
# nums2 = [2]

# 则中位数是 2.0
# 示例 2:

# nums1 = [1, 2]
# nums2 = [3, 4]

# 则中位数是 (2 + 3)/2 = 2.5


class Solution:
    def findMedianSortedArrays(self, nums1, nums2):
        # m < n, len(nums1) < len(nums2)
        m, n = len(nums1), len(nums2)
        if m > n:
            nums1, nums2, m, n = nums2, nums1, n, m
        if n == 0:
            raise ValueError

        imin, imax, half_len = 0, m, (m + n + 1) // 2
        while imin <= imax:
            i = (imin + imax) // 2
            j = half_len - i
            if i < m and nums2[j-1] > nums1[i]:
                # i is too small, must increase it
                imin = i + 1
            elif i > 0 and nums1[i-1] > nums2[j]:
                # i is too big, must decrease it
                imax = i - 1
            else:
                # i is perfect

                if i == 0:
                    max_of_left = nums2[j-1]
                elif j == 0:
                    max_of_left = nums1[i-1]
                else:
                    max_of_left = max(nums1[i-1], nums2[j-1])

                if (m + n) % 2 == 1:
                    return max_of_left

                if i == m:
                    min_of_right = nums2[j]
                elif j == n:
                    min_of_right = nums1[i]
                else:
                    min_of_right = min(nums1[i], nums2[j])

                return (max_of_left + min_of_right) / 2.0


if __name__ == "__main__":
    s = Solution()
    nums1 = [1, 2]
    nums2 = [2, 4]
    print(s.findMedianSortedArrays(nums1, nums2))


# pointOneLeft = 0
# pointOneRight = len(nums1)
# lenOne = pointOneRight
# pointTwoLeft = 0
# pointTwoRight = len(nums2)
# lenTwo = pointTwoRight
# while(lenOne > 2 and lenTwo > 2):
#     half = min(lenOne, lenTwo) // 2
#     numOneLeft = nums1[pointOneLeft + half]
#     numOneRight = nums1[pointOneRight - half]
#     numTwoLeft = nums2[pointTwoLeft + half]
#     numTwoRight = nums2[pointTwoRight - half]
#     if(numOneLeft < numTwoLeft):
#         pointOneLeft += half
#     else:
#         pointTwoLeft += half
#     if(numOneRight > numTwoRight):
#         pointOneRight -= half
#     else:
#         pointTwoRight -= half
#     lenOne = pointOneRight - pointOneLeft
#     lenTwo = pointTwoRight - pointTwoLeft
# if(lenOne == 1):
#     half = lenTwo // 2
#     if(half*2 == lenTwo):
#         numTwoLeft = nums2[pointTwoLeft + half - 1]
#         numTwoRight = nums2[pointTwoLeft + half]
#         numOneLeft = nums1[pointOneLeft]
#         if(numOneLeft < numTwoLeft):
#             return numTwoRight
#         elif(numOneLeft > numTwoRight):
#             return numTwoLeft
#         else:
#             return numOneLeft
#     else:
