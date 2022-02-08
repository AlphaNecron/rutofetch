package rutofetch

import (
	"bytes"
	"fmt"
	"github.com/BourgeoisBear/rasterm"
	"image"
	_ "image/png"
	"os"
	"regexp"
	"strings"
	"time"
)

func makeColorBlock() string {
	var block string
	for i := 0; i < 8; i++ {
		block += fmt.Sprintf("%s%s", Color(i).Fg(), strings.Repeat("\xe2\x96\x81", 2))
	}
	return block
}

func loadImage() (iImg image.Image, imgFmt string, E error) {
	return image.Decode(bytes.NewReader(naruto))
}

func printArt() bool {

	s := rasterm.Settings{}
	if rasterm.IsTermKitty() {
		r := bytes.NewReader(naruto)
		err := s.KittyCopyPNGInline(os.Stdout, r, r.Size())
		if err != nil {
			return false
		}
		return true
	} else if sixelCapable, _ := rasterm.IsSixelCapable(); sixelCapable {
		img, _, err := loadImage()
		if err != nil {
			return false
		}
		if ip, ok := img.(*image.Paletted); ok {
			err = s.SixelWriteImage(os.Stdout, ip)
			if err != nil {
				return false
			}
		}
	}
	return false
}

func shortenCpu(cpu string) string {
	r := regexp.MustCompile(`(?i)(\((R|TM))\)|@\s[0-9]\.[0-9]{2}[MG]Hz$|Core|Core? Duo|CPU|Intel|AMD|Qualcomm|Processor|(Dual|Quad|Six|Eight|[0-9]+)-Core`)
	return strings.Replace(strings.TrimSpace(r.ReplaceAllString(cpu, "")), "  ", " ", 1)
}

func formatUnit(b uint64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func parseUptime(secs uint64) string {
	return time.Duration(secs * 1e9).String()
}

func toEscape(c int8) string {
	return fmt.Sprintf("\033[%dm", c)
}
