/**
 * @author: baowj
 * @time: 2023/2/8 21:33:36
 * @email: bwj_678@qq.com
 */
package main

func removeSubfolders(folder []string) []string {
	fs := make(map[string]struct{})
	for _, dir := range folder {
		fs[dir] = struct{}{}
	}
	res := []string{}
	for _, dir := range folder {
		i := len(dir) - 1
		isSubFolder := false
		for ; i >= 0; i-- {
			if dir[i] == '/' {
				if i != 0 {
					if _, ok := fs[dir[:i]]; ok {
						isSubFolder = true
						break
					}
				} else {
					break
				}
			}
		}
		if !isSubFolder {
			res = append(res, dir)
		}
	}
	return res
}
