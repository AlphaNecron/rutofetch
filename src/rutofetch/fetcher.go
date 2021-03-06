package rutofetch

import (
	"fmt"
	"github.com/dekobon/distro-detect/linux"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"log"
	"os/user"
	"runtime"
)

func getOs() string {
	switch os := runtime.GOOS; os {
	case "windows":
		return "Windows"
	case "linux":
		linux.FileSystemRoot = "/"
		distro := linux.DiscoverDistro()
		return distro.Name
	}
	return "Unknown"
}

func getSysInfo() host.InfoStat {
	info, err := host.Info()
	if err != nil {
		log.Fatalln(err.Error())
	}
	return *info
}

func getCpuBrand() string {
	cpuInfo, err := cpu.Info()
	if err != nil {
		log.Fatalln(err)
	}
	return shortenCpu(cpuInfo[0].ModelName)
}

func getMemInfo() string {
	mInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Fatalln(err.Error())
	}
	return fmt.Sprintf("%s / %s", formatUnit(mInfo.Used), formatUnit(mInfo.Total))
}

func getUserName() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	return usr.Username
}
