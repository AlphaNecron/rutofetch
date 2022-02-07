package rutofetch

type Color int8

const (
	Black   Color = 0
	Red     Color = 1
	Green   Color = 2
	Yellow  Color = 3
	Blue    Color = 4
	Magenta Color = 5
	Cyan    Color = 6
	White   Color = 7
)

func (c Color) Fg() string {
	return toEscape(int8(c) + 30)
}

func (c Color) Bg() string {
	return toEscape(int8(c) + 40)
}
