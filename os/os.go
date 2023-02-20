package os

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

func GetCpuPercent() float64 {
	percent, _ := cpu.Percent(time.Duration(time.Second), false)
	return percent[0]
}
func GetMemPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}
func GetDiskPercent() float64 {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return diskInfo.UsedPercent
}
func GetCpuInfo() []cpu.InfoStat {
	info, _ := cpu.Info()
	return info
}
func GetIOInfo() map[string]disk.IOCountersStat {
	info, _ := disk.IOCounters()
	return info
}
func GetNetInfo(str string) []net.ConnectionStat {
	info, _ := net.Connections(str)
	return info
}
