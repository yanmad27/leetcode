package main
/*
 * @lc app=leetcode id=874 lang=golang
 *
 * [874] Walking Robot Simulation
 */

// @lc code=start
func robotSim(commands []int, obstacles [][]int) int {
	result := 0
	direction := 0
	x := 0
	y := 0
	for _, command := range commands {
		switch command {
		case -2:
			direction--
			if direction < 0 {
				direction = 3
			}
		case -1:
			direction++
			if direction > 3 {
				direction = 0
			}
		default:
			switch direction {
			case 0:
				tmp := command
				for _, b := range obstacles {
					if x == b[0] && b[1]-y <= tmp {
						tmp = b[1] - y - 1
					}
				}
				y += tmp
			case 1:
				tmp := command
				for _, b := range obstacles {
					if y == b[1] && b[0]-x <= tmp {
						tmp = b[0] - x - 1
					}
				}
				x += tmp
			case 2:
				tmp := command
				for _, b := range obstacles {
					if x == b[0] && y-b[1] >= tmp {
						tmp = y - b[1] - 1
					}
				}
				y -= tmp
			case 3:
				tmp := command
				for _, b := range obstacles {
					if y == b[1] && x-b[0] >= tmp {
						tmp = x - b[0] - 1
					}
				}
				x -= tmp
			}

			if tmp := x*x + y*y; tmp > result {
				result = tmp
			}
		}
	}
	return result
}

// @lc code=end

