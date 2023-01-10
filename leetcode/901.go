/**
 * @author: baowj
 * @time: 2022/10/21 21:28:41
 */
package main

import "fmt"

type StockSpanner struct {
	array []int
	cnt   []int
}

func Constructor901() StockSpanner {
	return StockSpanner{
		array: []int{},
		cnt:   []int{},
	}

}

func (this *StockSpanner) Next(price int) int {
	if len(this.array) == 0 {
		this.array = append(this.array, price)
		this.cnt = append(this.cnt, 1)
		return 1
	}
	if price < this.array[len(this.array)-1] {
		this.array = append(this.array, price)
		this.cnt = append(this.cnt, 1)
	} else if price == this.array[len(this.array)-1] {
		this.array = append(this.array, price)
		this.cnt = append(this.cnt, this.cnt[len(this.cnt)-1]+1)
	} else {
		i := len(this.cnt) - 1
		fmt.Println(this.cnt)
		fmt.Println(this.array)
		for i >= 0 && this.array[i] < price {
			i -= this.cnt[i]
		}
		if i == -1 {
			i = 0
		} else {
			if this.array[i] <= price {
				if i != 0 {
					i -= this.array[i]
				}
			} else {
				i++
			}
		}
		this.array = append(this.array, price)
		this.cnt = append(this.cnt, len(this.cnt)-i+1)
	}

	return this.cnt[len(this.cnt)-1]
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
