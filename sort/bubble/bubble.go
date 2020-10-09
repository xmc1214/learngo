package bubble

func Bubble(arr []int) {
	flag := false // 用于判断数组有序后提前结束排序
	len := len(arr)
	for i := 0; i < len; i++ {
		flag = false
		for j := 0; j < len-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
