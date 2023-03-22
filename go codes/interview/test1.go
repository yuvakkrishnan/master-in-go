
Yuvakkrishnan
Karunakar
Karunakar
Meeting Chat
Karunakar
to
Everyone
01:05 PM
K
What might be wrong with the following small program?

func main() {
	scanner := bufio.NewScanner(strings.NewReader(`one
two
three
four
`))
	var (
		text string
		n    int
	)
	for scanner.Scan() {
		n++
		text += fmt.Sprintf("%d. %s\n", n, scanner.Text())
	}
	fmt.Print(text)

	// Output:
	// 1. One
	// 2. Two
	// 3. Three
	// 4. Four
}
What is wrong with the following code snippet?

type Orange struct {
	Quantity int
}

func (o *Orange) Increase(n int) {
	o.Quantity += n
}

func (o *Orange) Decrease(n int) {
	o.Quantity -= n
}

func (o *Orange) String() string {
	return fmt.Sprintf("%v", o.Quantity)
}

func main() {
	var orange Orange
	orange.Increase(10)
	orange.Decrease(5)
	fmt.Println(orange)
}
How would you implement a stack and a queue in Go? Explain and provide code examples.