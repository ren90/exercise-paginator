package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ren90/paginator/pagination"
)

func main() {
	argsForCall := os.Args[1:]

	if len(argsForCall) < 4 {
		panic("Not enough arguments: please call with <current_page>, <total_pages>, <boundaries>, <around>")
	}

	cp, err := strconv.Atoi(argsForCall[0])
	if err != nil {
		panic("<current_page> must be an integer")
	}

	tp, err := strconv.Atoi(argsForCall[1])
	if err != nil {
		panic("<total_pages> must be an integer")
	}

	b, err := strconv.Atoi(argsForCall[2])
	if err != nil {
		panic("<boundaries> must be an integer")
	}

	a, err := strconv.Atoi(argsForCall[3])
	if err != nil {
		panic("<around> must be an integer")
	}

	fp := &pagination.Footer{CurrentPage: cp, TotalPages: tp, Boundaries: b, Around: a}
	fmt.Println(strings.Join(fp.ComputePagination(), " "))
}
