package main

import (
	"fmt"
)

func quickSort(values []int, left int, right int) {
	if left < right {
		
		temp := values[left]
		i, j := left, right
		for {
			// 从右向左找，找到第一个比分水岭小的数
			for values[j] >= temp && i < j {
				j--
			}
			// 从左向右找，找到第s一个比分水岭大的数
			for values[i] <= temp && i < j {
				i++
			}
			// 如果哨兵相遇，则退出循环
			if i >= j {
				break
			}
			// 交换左右两侧的值
			values[i], values[j] = values[j], values[i]
		}
        // 将分水岭移到哨兵相遇点
        // 上述先从右向左，所以values[left] >= values[i]
		values[left], values[i] = values[i], values[left]
		// 递归，左右两侧分别排序
		quickSort(values, left, i-1)
		quickSort(values, i+1, right)
	}
}

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}

func main() {
	list := []int{2,2,1,1,3,3,3,4,4,1,1,1}
	QuickSort(list)
	fmt.Println(list)
}





