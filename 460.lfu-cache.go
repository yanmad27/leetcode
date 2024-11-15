package main

/*
 * @lc app=leetcode id=460 lang=golang
 *
 * [460] LFU Cache
 */

// @lc code=start
type LFUCache struct {
}

func Constructor(capacity int) LFUCache {
	return LFUCache{}
}

func (this *LFUCache) Get(key int) int {
	return -1
}

func (this *LFUCache) Put(key int, value int) {

}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
