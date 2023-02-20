package minitor

import (
	"fmt"
	"minitor/os"
	_ "minitor/os"
)

func main() {
	fmt.Println(os.GetCpuPercent())
	fmt.Println(os.GetMemPercent())
	fmt.Println(os.GetDiskPercent())
	fmt.Println(os.GetCpuInfo())
	fmt.Println(os.GetNetInfo())
}
