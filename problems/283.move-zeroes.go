package problems

/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int)  {
    NoZeroValue := 0

    for i := 0; i<len(nums);i++{
        if nums[i]!=0{
            nums[NoZeroValue]=nums[i]
            NoZeroValue++
        }
    }
    for i:= NoZeroValue;i<len(nums);i++{
        nums[i]=0
    }
    
}

// @lc code=end
