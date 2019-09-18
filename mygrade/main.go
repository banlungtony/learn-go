package main

func main() {
	showg(4)
}

func showg(grade int) string {
	var a,b,c,d,f = "A", "B", "C", "D", "F"
	if (grade == 4) {
		return a
	}
	if (grade == 3) {
		return b
	}
	if (grade == 2) {
		return c
	}
	if (grade == 1) {
		return d
	}
	return f
}