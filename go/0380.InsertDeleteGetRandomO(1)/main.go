package main

import (
	"math/rand"
)

type RandomizedSet struct {
	// 数组
	nums       []int
	// HashMap key是数组的元素,value是数组元素在数组里的索引
	valToIndex map[int]int
}


func Constructor() RandomizedSet {
	// 初始化
	return RandomizedSet{
		nums:       []int{},
		valToIndex: map[int]int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	// 判断是否有这个值，有就返回false。
	if _, ok := this.valToIndex[val]; ok {
		return false
	}
	// 将新的值添加到数组最后
	this.nums = append(this.nums, val)
	// 更新HashMap
	this.valToIndex[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	// 判断HashMap里是否有这个值
	if index, ok := this.valToIndex[val]; ok {
		lastIndex := len(this.nums) - 1
		// 更新最后一个元素的索引
		this.valToIndex[this.nums[lastIndex]] = index
		// 将要删除的元素与最后一个交换
		this.nums[index], this.nums[lastIndex] = this.nums[lastIndex], this.nums[index]
		// 删除最后一个元素
		this.nums = this.nums[:lastIndex]
		// 删除HashMap里的元素
		delete(this.valToIndex, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	// 随机返回数组里的元素
	return this.nums[rand.Intn(len(this.nums))]
}