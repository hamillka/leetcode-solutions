func sortColors(nums []int)  {
    cntRed, cntWhite := 0, 0
    for _, v := range nums {
        if v == 0 {
            cntRed++
        } else if v == 1 {
            cntWhite++
        }
    }

    i := 0
    for ; i < cntRed; i++ {
        nums[i] = 0
    }
    for ; i < cntRed+cntWhite; i++ {
        nums[i] = 1
    }
    for ; i < len(nums); i++ {
        nums[i] = 2
    }
}
