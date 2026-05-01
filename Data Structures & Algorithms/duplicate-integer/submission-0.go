func hasDuplicate(nums []int) bool {
    numsMap := make(map[int]bool)
    
    for _, v := range nums {
        if _, ok := numsMap[v]; ok {
            return true
        }
        numsMap[v] = true
    }

    return false
}
