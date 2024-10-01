package main

// /*
//  * @lc app=leetcode id=729 lang=golang
//  *
//  * [729] My Calendar I
//  */

// // @lc code=start
// type MyCalendar struct {
// 	booking [][2]int
// }

// func Constructor() MyCalendar {
// 	return MyCalendar{}
// }

// func (this *MyCalendar) Book(start int, end int) bool {
// 	count := 0
// 	for _, period := range this.booking {
// 		bookedStart := period[0]
// 		bookedEnd := period[1]
// 		if start < bookedEnd && end >= bookedEnd {
// 			count++
// 		} else if start <= bookedStart && end > bookedStart {
// 			count++
// 		} else if start >= bookedStart && end < bookedEnd {
// 			count++
// 		} else if start < bookedStart && end > bookedEnd {
// 			count++
// 		}
// 		if count == 1 {
// 			return false
// 		}
// 	}
// 	this.booking = append(this.booking, [2]int{start, end})
// 	return true
// }

// /**
//  * Your MyCalendar object will be instantiated and called as such:
//  * obj := Constructor();
//  * param_1 := obj.Book(start,end);
//  */
// // @lc code=end
