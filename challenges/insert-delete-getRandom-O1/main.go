package main

import (
    "fmt"
	"math/rand"
)

type RandomizedSet struct {
    indexMap map[int]int
    values []int
}

func Constructor() RandomizedSet {
    return RandomizedSet{
        indexMap: make(map[int]int),
        values: []int{},
    }
}

func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.indexMap[val]; ok {
        return false
    }

    this.values = append(this.values, val)
    this.indexMap[val] = len(this.values) - 1
    return true
}

func (this *RandomizedSet) Remove(val int) bool {
    i, ok := this.indexMap[val]
    if !ok {
        return false
    }

    count := len(this.values)
    
    swap(this.values, i, count - 1)
    lastValue := this.values[i]
    this.values = this.values[:count - 1]

    this.indexMap[lastValue] = i
    delete(this.indexMap, val)

    return true
}

func (this *RandomizedSet) GetRandom() int {
    randomIndex := rand.Intn(len(this.values))
    return this.values[randomIndex]
}

func swap(arr []int, i, j int) {
    arr[i], arr[j] = arr[j], arr[i]
} 

func main() {
	randomizedSet := Constructor()

	fmt.Println(randomizedSet.Insert(1))   // true
	fmt.Println(randomizedSet.Remove(2))   // false
	fmt.Println(randomizedSet.Insert(2))   // true
	fmt.Println(randomizedSet.GetRandom()) // 1 or 2
	fmt.Println(randomizedSet.Remove(1))   // true
	fmt.Println(randomizedSet.Insert(2))   // false
	fmt.Println(randomizedSet.GetRandom()) // 2
}