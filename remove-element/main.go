package main

func removeElement(nums []int, val int) int {
    k := 0
    for _,v := range nums {
        if v == val {
            k += 1
        }
    }
  
    for z := 0; z < k; z++ {
        for i := 0; i < len(nums) - 1; i++ {
        if nums[i] == val {
            nums[i] = nums[i + 1]
            for l := i + 1; l < len(nums) - 1; l++ {
                nums[l] = nums[l + 1]
            }
        }
        }
    }
    return len(nums) - k
}

func main() {
  testCase1 := []int{3, 2, 2, 3}
  testCase2 := []int{0, 1, 2, 2, 3, 0, 4, 2}

  removeElement(testCase1, 3)
  removeElement(testCase2, 2)
}

