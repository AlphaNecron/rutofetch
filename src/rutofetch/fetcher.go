package rutofetch

import (
	"fmt"
	"github.com/dekobon/distro-detect/linux"
	cpu2 "github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"log"
	"os/user"
	"runtime"
	"strings"
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
	cpu, err := cpu2.Info()
	if err != nil {
		log.Fatalln(err)
	}
	return strings.Split(cpu[0].ModelName, " ")[2]
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
