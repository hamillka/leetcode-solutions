func peakIndexInMountainArray(arr []int) int {
    l, r := 0, len(arr) - 1
    for l <= r {
        m := (l + r) / 2
        if arr[m] < arr[m+1] {
            l = m + 1
        } else {
            r = m - 1
        }
    }

    return l
}
