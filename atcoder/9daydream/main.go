package main

import (
	"bytes"
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	sb := []byte(s)
	rev(sb)

	w := [][]byte{[]byte("dream"), []byte("dreamer"), []byte("erase"), []byte("eraser")}

	for i := range w {
		rev(w[i])
	}

	i := 0
	for i < len(sb) {
		var ok bool
		for _, t := range w {
			if bytes.HasPrefix(sb[i:], t) {
				ok = true
				i += len(t)
				break
			}
		}
		if !ok {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}

func rev(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
