package main

import "sort"

/*
 * @lc app=leetcode id=1233 lang=golang
 *
 * [1233] Remove Sub-Folders from the Filesystem
 */

// @lc code=start

func checkExist(rootMap map[string]bool, folder string) bool {
	var tmp string
	folder += "/"
	for i := 0; i < len(folder); i++ {
		if folder[i] == '/' {
			if rootMap[tmp] {
				return true
			}
		}
		tmp = folder[:i+1]
	}
	return false
}

func removeSubfolders(folders []string) []string {

	rootMap := make(map[string]bool)
	rootArr := []string{}

	sort.Slice(folders, func(i, j int) bool {
		return len(folders[i]) < len(folders[j])
	})

	for _, folder := range folders {
		if !checkExist(rootMap, folder) {
			rootMap[folder] = true
			rootArr = append(rootArr, folder)
		}
	}

	return rootArr

}

// @lc code=end
