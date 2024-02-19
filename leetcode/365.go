package main

func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	if jug1Capacity+jug2Capacity < targetCapacity {
		return false
	}
	if jug1Capacity < jug2Capacity {
		jug1Capacity, jug2Capacity = jug2Capacity, jug1Capacity
	}
	for jug2Capacity != 0 {
		jug1Capacity, jug2Capacity = jug2Capacity, jug1Capacity%jug2Capacity
	}

	return targetCapacity%jug1Capacity == 0
}
