/**
 * @author: baowj
 * @time: 2022/12/10 22:11:40
 * @email: bwj_678@qq.com
 */
package main

func squareIsWhite(coordinates string) bool {
	return (coordinates[0]-'a'+coordinates[1]-'1')%2 == 1
}
