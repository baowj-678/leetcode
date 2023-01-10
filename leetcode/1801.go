/**
 * @author: baowj
 * @time: 2023/1/8 11:44:39
 * @email: bwj_678@qq.com
 */
package main

import (
	"container/heap"
)

const MOD1801 int = 1e9 + 7

type Order struct {
	price  int
	amount int
}

type MaxHeap []Order
type MinHeap []Order

func (hp MaxHeap) Len() int {
	return len(hp)
}

func (hp MaxHeap) Less(i, j int) bool {
	return hp[i].price > hp[j].price
}

func (hp MaxHeap) Swap(i, j int) {
	hp[i], hp[j] = hp[j], hp[i]
}

func (hp *MaxHeap) Push(x interface{}) {
	*hp = append(*hp, x.(Order))
}

func (hp *MaxHeap) Pop() interface{} {
	v := (*hp)[len(*hp)-1]
	*hp = (*hp)[:len(*hp)-1]
	return v
}

func (hp MinHeap) Len() int {
	return len(hp)
}

func (hp MinHeap) Less(i, j int) bool {
	return hp[i].price < hp[j].price
}

func (hp MinHeap) Swap(i, j int) {
	hp[i], hp[j] = hp[j], hp[i]
}

func (hp *MinHeap) Push(x interface{}) {
	*hp = append(*hp, x.(Order))
}

func (hp *MinHeap) Pop() interface{} {
	v := (*hp)[len(*hp)-1]
	*hp = (*hp)[:len(*hp)-1]
	return v
}

func getNumberOfBacklogOrders(orders [][]int) int {
	sell, buy := &MinHeap{}, &MaxHeap{}
	heap.Init(sell)
	heap.Init(buy)
	for _, order := range orders {
		// buy
		if order[2] == 0 {
			for order[1] > 0 && len(*sell) > 0 {
				minOrder := heap.Pop(sell).(Order)
				if minOrder.price > order[0] {
					heap.Push(sell, minOrder)
					break
				}
				if order[1] >= minOrder.amount {
					order[1] -= minOrder.amount
				} else {
					minOrder.amount -= order[1]
					order[1] = 0
					heap.Push(sell, minOrder)
				}
			}
			if order[1] != 0 {
				heap.Push(buy, Order{
					price:  order[0],
					amount: order[1],
				})
			}
		} else {
			// sell
			for order[1] > 0 && len(*buy) > 0 {
				maxOrder := heap.Pop(buy).(Order)
				if maxOrder.price < order[0] {
					heap.Push(buy, maxOrder)
					break
				}
				if order[1] >= maxOrder.amount {
					order[1] -= maxOrder.amount
				} else {
					maxOrder.amount -= order[1]
					order[1] = 0
					heap.Push(buy, maxOrder)
				}
			}
			if order[1] != 0 {
				heap.Push(sell, Order{
					price:  order[0],
					amount: order[1],
				})
			}
		}
	}
	res := 0
	for _, order := range *sell {
		res = (res + order.amount) % MOD1801
	}
	for _, order := range *buy {
		res = (res + order.amount) % MOD1801
	}
	return res
}

func main() {
	orders := [][]int{{10, 5, 0}, {15, 2, 1}, {25, 1, 1}, {30, 4, 0}}
	println(getNumberOfBacklogOrders(orders))
}
