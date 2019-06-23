// Given an array of integers, return indices of 
// the two numbers such that they add up to a specific target. 
// You may assume that each input would have 
// exactly one solution, and you may not use the same element twice.


func twoSum(nums []int, target int) []int {
    sum := 0
    result := make([]int, 2)
    count := len(nums)
    for i:=0; i<count; i++ {
        sum = 0 
        for j:=i+1; j<count; j++ {
            sum = nums[i] + nums[j]
            if sum == target {
                result[0] = i
                result[1] = j
                return result
            }
        }
    }
    return nil
}
