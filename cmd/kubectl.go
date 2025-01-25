// 实现类似 kubectl 命令的功能
package main

/*
  url: https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go#using-a-flag-to-change-a-program-s-behavior
  Explanations of flag:
  1. define variables to capture flag values
  2. define the falgs your Go application will use
  3. parse the flags provided to the application upon execution
*/

import (
	"flag"
	"fmt"
	"os"
)

// -f 参数
var file string

func main() {
	flag.StringVar(&file, "f", "", "输入清单")
	flag.Parse()
	fmt.Println(os.Args[1:])
	fmt.Printf("-f %v\n", file)
	fmt.Printf("len = %v\n", len(file))
	fmt.Printf("flag arg: %v\n", flag.Arg(1))
	fmt.Printf("flag args: %v\n", flag.Args())
}