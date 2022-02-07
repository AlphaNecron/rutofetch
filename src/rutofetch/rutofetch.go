package rutofetch

import (
	"fmt"
	"github.com/BourgeoisBear/rasterm"
	"github.com/ahmetalpbalkan/go-cursor"
)

func Rutofetch() {
	osInfo := getSysInfo()

	info := [6]Info{
		{"os", getOs()},
		{"arch", osInfo.KernelArch},
		{"kernel", osInfo.KernelVersion},
		{"uptime", parseUptime(osInfo.Uptime)},
		{"cpu", getCpuBrand()},
		{"mem", getMemInfo()},
	}
	fmt.Print(rasterm.ESC_ERASE_DISPLAY)
	canPrintImg := printArt()
	if canPrintImg {
		fmt.Print(cursor.MoveTo(1, 16))
	}
	fmt.Printf("%s%s%s%s@%s%s\n%s", toEscape(1), Yellow.Fg(), getUserName(), White.Fg(), Yellow.Fg(), osInfo.Hostname, toEscape(0))
	if canPrintImg {
		fmt.Print(cursor.MoveRight(15))
	}
	for _, inf := range info {
		inf.print()
		if canPrintImg {
			fmt.Print(cursor.MoveRight(15))
		}
	}
	fmt.Print(makeColorBlock())
	fmt.Print(cursor.MoveDown(1), cursor.MoveNextLine())
}
