class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number}
     */
    //        0 1 2 3 4 5
    // nums=[-1,0,2,4,6,8]
    // target=4
    search(nums, target) {
        let l = 0;
        let r = nums.length - 1;
        while (l <= r) {
            const m = l + Math.floor((r - l ) / 2);
            if (target < nums[m]) {
                r = m - 1;
            } else if (target > nums[m]) {
                l = m + 1;
            } else {
                return m;
            }
        }
        return -1;
    }
}
