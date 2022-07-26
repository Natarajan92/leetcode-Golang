package twosum

//Step 1: Creating Hashmap
//By creating hashmap the complexity of running for loop multiple time is avoided and at the same time the less memory utilizationa and storage
//Add current value as key and its index as value in hashmap
//Step 2: Calculate the required number to acheive the target
//Step 3: Check whether requiredNumber is available in map
//Step 4: Return index of the numbers

func twoSum(nums []int, target int) []int {

	mapValueIndex := make(map[int]int)
	for currentIndex, currentValue := range nums {
		requiredNumber := target - currentValue
		if requiredIndex, ok := mapValueIndex[requiredNumber]; ok {
			return []int{requiredIndex, currentIndex}
		}
		mapValueIndex[currentValue] = currentIndex
	}
	return []int{}
}
