package main

func makeEqual(words []string) bool {
	m := make(map[byte]int)
	for _,v := range words{
			for i:=0 ;i<len(v);i++{
					m[v[i]]=m[v[i]]+1
			}
	}
	for _,v := range m{
			if v%len(words)!=0{
					return false
			}
	}
	return true
}