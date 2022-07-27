/**
 * @author: baowj
 * @time: 2022/4/5 17:22:48
 */
package main

func countPrimeSetBits(left int, right int) int {
	sum := 0
	M := []bool{false, false, true, true, false, true, false, true, false, false, false, true, false, true, false, false, false, true, false, true, false, false, false, true, false, false, false, false, false, true, false, true}
	for i := left; i <= right; i++ {
		cnt := 0
		for tmp := i; tmp > 0; tmp >>= 1 {
			if tmp&1 == 1 {
				cnt++
			}
		}
		if M[cnt] {
			sum++
		}
	}
	return sum
}

func main() {
	M := [32]bool{}
	for i := 0; i < 32; i++ {
		M[i] = false
	}
	M[2] = true
	M[3] = true
	M[5] = true
	M[7] = true
	M[11] = true
	M[13] = true
	M[17] = true
	M[19] = true
	M[23] = true
	M[29] = true
	M[31] = true
	for i := 0; i < 32; i++ {
		print(M[i])
		print(",")
	}
}
