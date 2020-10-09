package main

import (
	"fmt"
	"goWorkspace/sort/bubble"
	"math/rand"
	"sort"
	"time"
)

const (
	lenArray = 10000   // 测试数组长度
	rangeNum = 1000000 //数组元素大小范围
)

func main() {
	arr := GenerateRand() // 生成随机数组

	//每次测试前复制一份原数组
	testArr := make([]int, lenArray)
	copy(testArr, arr)

	// 使用系统内置排序模块对原数组排序，用于验证手写排序是否正确
	sort.Ints(testArr)

	//测试冒泡排序
	bubble.Bubble(arr)
	//打印前15个数,并比较是否相同
	fmt.Println(arr[:15])
	fmt.Println(testArr[:15])
	fmt.Println(IsSame(arr, testArr))
}

//生成随机数组
func GenerateRand() []int {
	randSeed := rand.New(rand.NewSource(time.Now().Unix() + time.Now().UnixNano()))
	arr := make([]int, lenArray)
	for i := 0; i < lenArray; i++ {
		arr[i] = randSeed.Intn(rangeNum)
	}
	return arr
}

//比较两个切片
func IsSame(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	if (slice1 == nil) != (slice2 == nil) {
		return false
	}

	for i, num := range slice1 {
		if num != slice2[i] {
			return false
		}
	}
	return true
}
