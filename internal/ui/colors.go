package ui

const (
	ColorBackground uint16 = 0xEF7D
	ColorBoard      uint16 = 0xB596
	ColorEmpty      uint16 = 0xCE17
	ColorDarkText   uint16 = 0x7B4D
	ColorLightText  uint16 = 0xFFFF
)

func TileColor(value int) uint16 {
	switch value {
	case 2:
		return 0xEF5B // Cream
	case 4:
		return 0xEF19 // Beige
	case 8:
		return 0xF58F // Light orange
	case 16:
		return 0xF4AC // Orange
	case 32:
		return 0xF36B // Coral
	case 64:
		return 0xF2E7 // Red-orange
	case 128:
		return 0xEE6E // Light gold
	case 256:
		return 0xEE2C // Gold
	case 512:
		return 0xEDCA // Deep gold
	case 1024:
		return 0xED88 // Golden yellow
	case 2048:
		return 0xED46 // Classic 2048 gold
	case 4096:
		return 0xB4D5 // Purple
	case 8192:
		return 0x9375 // Violet
	case 16384:
		return 0x7214 // Deep violet
	case 32768:
		return 0x50F3 // Dark purple
	case 65536:
		return 0x31B2 // Indigo
	case 131072:
		return 0x1971 // Deep indigo
	default:
		if value > 131072 {
			return 0x10A2 // Near-black blue
		}
		return ColorEmpty
	}
}

func TileTextColor(value int) uint16 {
	if value <= 4 {
		return ColorDarkText
	}
	return ColorLightText
}
