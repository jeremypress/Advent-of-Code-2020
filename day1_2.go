package main

import (
	"fmt"
	"strconv"
)

func Run1_2() {

	var lines, _ = ReadFile("./input/day1.txt")
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
            for k := 0; k < len(lines); k++ {
			if j == i || k == j || k == i{
				continue;
			}

			iInt, _ := strconv.Atoi(lines[i]);
			jInt, _ := strconv.Atoi(lines[j]);
            kInt, _ := strconv.Atoi(lines[k]);
			sum := iInt + jInt + kInt
			if sum == 2020 {
				fmt.Println(iInt * jInt * kInt)
				return
			}
            }


		}

	}

}
