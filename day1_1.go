package main

import (
    "fmt"
    "strconv"
    )

func Run1_1() {
	var lines, _ = ReadFile("./input/day1.txt");
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if j == i {
				continue
			}

            iInt, _ := strconv.Atoi(lines[i])
            jInt, _ := strconv.Atoi(lines[j]);
            sum := iInt + jInt
            if sum == 2020 {
                fmt.Println(iInt * jInt);
                return;
            }
		}

	}
}
