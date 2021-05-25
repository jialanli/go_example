package main
/*
by jialanli 2021
yuluoshanlan
*/
import (
	"fmt"
	mypkg "github.com/jialanli/lacia/utils"
	"go_example/test1_pkg"
	_ "io"
	"time"
)

func init() { // IDE 中输入init根据提示直接回车可打出
	fmt.Println("main")
}

const const_var1 = "A"
var s string

// 注释  按ctrl + l 快捷键可打出
func main() {
	fmt.Println("lan: start the world")
	test1_pkg.Test1_pkg_a1()
	/* 注释 按ctrl + shift + l 快捷键可打出*/
	ts := time.Now().Unix()
	str := mypkg.GetTimeStrOfDayTimeByTs(ts)
	fmt.Println(str)
}
