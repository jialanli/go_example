package test1_pkg

import (
	"fmt"
	"go_example/test1_pkg/test1_inner_pkg"
	"go_example/test2_pkg"
)

func init() {
	fmt.Println("test1_pkg:a1")
}

func Test1_pkg_a1() {
	fmt.Println("func test1_pkg:a1")
	test2_pkg.Test2_pkg_a1()
	fmt.Println("---------------")
	test1_inner_pkg.Test1_inner_pkg_a1()
}
