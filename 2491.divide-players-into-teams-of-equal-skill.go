package main

/*
 * @lc app=leetcode id=2491 lang=golang
 *
 * [2491] Divide Players Into Teams of Equal Skill
 */

// @lc code=start
func dividePlayers(skill []int) int64 {
	totalSum := 0
	for _, num := range skill {
		totalSum += num
	}
	if len(skill) == 2 {
		return int64(skill[0] * skill[1])
	}
	singleSkill := totalSum / (len(skill) / 2)
	if len(skill) > 2 && totalSum%singleSkill != 0 {
		return -1
	}
	teamMap := make(map[int]int)
	for _, num := range skill {
		key := min(singleSkill-num, num)
		teamMap[key]++
	}
	chemistry := 0
	for key, team := range teamMap {
		if team%2 != 0 {
			return -1
		}
		chemistry += (singleSkill - key) * key * team / 2
	}
	return int64(chemistry)

}

// @lc code=end
