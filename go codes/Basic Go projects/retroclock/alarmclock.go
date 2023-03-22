package main

import (
	"fmt"
	"time"
)

func main() {
	type placeholder [5]string

	var zero = placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	var one = placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	var two = placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	var three = placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	var four = placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	var five = placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	var six = placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	var seven = placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	var eight = placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	var nine = placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	var colon = placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	var alarm = [...]placeholder{
		{
			"   ",
			"   ",
			"   ",
			"   ",
			"   ",
		},
		{
			"███",
			"█ █",
			"███",
			"█ █",
			"█ █",
		},
		{
			"█  ",
			"█  ",
			"█  ",
			"█  ",
			"███",
		},
		{
			"███",
			"█ █",
			"███",
			"█ █",
			"█ █",
		},
		{
			"██ ",
			"█ █",
			"██ ",
			"█ █",
			"█ █",
		},
		{
			"█ █",
			"███",
			"█ █",
			"█ █",
			"█ █",
		},
		{
			" █ ",
			" █ ",
			" █ ",
			"   ",
			" █ ",
		},
		{
			"   ",
			"   ",
			"   ",
			"   ",
			"   ",
		},
	}

	var digits = [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}
	for {

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		alarmed := sec%10 == 0

		for line := range clock[0] {
			if alarmed {
				clock = alarm
			}

			for index, digit := range clock {
				next := clock[index][line]
				if digit == colon && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}

		time.Sleep(time.Second)
	}
}
