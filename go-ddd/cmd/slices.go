package cmd

import "fmt"

// Slices cmd
func Slices() {
	s := make([]string, 1)
	s[0] = "hey"
	s = append(s, "foo", "bar")
	c := make([]string, len(s))
	copy(c, s)
	sub := s[:3]
	sub[0] = "boom"
	fmt.Println(s, sub, c)

	m := make(map[string]string)
	m["first"] = "one"
	m["second"] = "two"
	delete(m, "second")
	fmt.Println(m, m["first"])
}
