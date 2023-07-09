/**
 * @author: baowj
 * @time: 2023/7/9 14:44
 */
package main

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	if numOnes >= k {
		return k
	} else if numOnes+numZeros >= k {
		return numOnes
	} else {
		return numOnes - (k - numOnes - numZeros)
	}
}
