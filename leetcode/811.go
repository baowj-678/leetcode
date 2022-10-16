/**
 * @author: baowj
 * @time: 2022/10/8 21:29:48
 */
package main

import (
	"strconv"
	"strings"
)

type struct811 struct {
	s     string
	times int
	next  map[string]*struct811
}

func subdomainVisits(cpdomains []string) []string {
	root := &struct811{
		s:     "",
		times: 0,
		next:  map[string]*struct811{},
	}
	for _, s := range cpdomains {
		domains, cnt := parser811(s)
		addDomain(domains, cnt, root)
	}
	res := []string{}
	res = dfs811(root, res, []string{})
	return res
}

func addDomain(domains []string, cnt int, root *struct811) {
	root.times += cnt
	for i := len(domains) - 1; i >= 0; i-- {
		domain := domains[i]
		if v, ok := root.next[domain]; ok {
			root = v
			root.times += cnt
		} else {
			root.next[domain] = &struct811{
				s:     domain,
				times: cnt,
				next:  map[string]*struct811{},
			}
			root = root.next[domain]
		}
	}
}

func parser811(s string) (domains []string, cnt int) {
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			cnt, _ = strconv.Atoi(s[:i])
			domains = strings.Split(s[i+1:], ".")
		}
	}
	return
}

func dfs811(root *struct811, res []string, domains []string) []string {
	if root.s != "" {
		domains = append(domains, root.s)
		res = append(res, generateRes(domains, root.times))
	}
	for _, v := range root.next {
		res = dfs811(v, res, domains)
	}
	if root.s != "" {
		domains = domains[:len(domains)-1]
	}
	return res
}

func generateRes(domains []string, cnt int) string {
	res := ""
	res += strconv.Itoa(cnt)
	res += " "
	for i := len(domains) - 1; i > 0; i-- {
		res += domains[i]
		res += "."
	}
	res += domains[0]
	return res
}
