package main

/*
 * @lc app=leetcode id=1352 lang=golang
 *
 * [1352] Product of the Last K Numbers
 */

// @lc code=start
type ProductOfNumbers struct {
	prefixProduct  []int
	currentProduct int
	lastZero       int
}

func Constructor1352() ProductOfNumbers {
	return ProductOfNumbers{
		prefixProduct:  []int{1},
		currentProduct: 1,
		lastZero:       0,
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num != 0 {
		this.currentProduct *= num
	} else {
		this.currentProduct = 1
		this.lastZero = len(this.prefixProduct)
	}
	this.prefixProduct = append(this.prefixProduct, this.currentProduct)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if len(this.prefixProduct)-k <= this.lastZero {
		return 0
	} else {
		return this.prefixProduct[len(this.prefixProduct)-1] / this.prefixProduct[len(this.prefixProduct)-k-1]
	}
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
// @lc code=end
