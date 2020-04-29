/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
    len := mountainArr.length()
    left := 0
    right := len - 1 
    for left != right {
        mid := (left + right) >> 1
        if mountainArr.get(mid)> mountainArr.get(mid+1){
            right = mid 
        }else{
            left = mid + 1
        }
    }
    peak := left
    left, right = 0, peak
    for left != right{
        mid := (left+right) / 2
        val := mountainArr.get(mid)
        if val == target {return mid}
        if target > val {
            left = mid + 1
        }else{
            right = mid
        }
    }
    if target == mountainArr.get(left) {return left}
    //再在右边找
    left,right=peak,mountainArr.length()-1
    for left!=right{
        mid:=(left+right)/2
        val:=mountainArr.get(mid)
        if val==target {return mid}
        if target<val {
            left=mid+1
        }else{
            right=mid
        }
    }
    if target==mountainArr.get(left) {return left}
    return -1
}



