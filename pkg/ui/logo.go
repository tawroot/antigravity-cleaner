package ui

import (
	"fmt"
	"strings"
)

// RGB holds a 24-bit TrueColor channel tuple
type RGB struct {
	R, G, B uint8
}

var logoA = [10][12]*RGB{
	{nil, nil, nil, nil, nil, &RGB{242, 146, 46}, &RGB{240, 114, 54}, nil, nil, nil, nil, nil},
	{nil, nil, nil, nil, &RGB{219, 177, 49}, &RGB{246, 145, 46}, &RGB{243, 115, 55}, &RGB{240, 88, 59}, nil, nil, nil, nil},
	{nil, nil, nil, &RGB{158, 195, 69}, &RGB{181, 180, 62}, &RGB{226, 153, 61}, &RGB{246, 122, 52}, &RGB{248, 106, 53}, &RGB{239, 84, 66}, nil, nil, nil},
	{nil, nil, nil, &RGB{134, 198, 78}, &RGB{117, 180, 94}, &RGB{204, 149, 77}, &RGB{239, 121, 71}, &RGB{225, 102, 82}, &RGB{225, 79, 89}, nil, nil, nil},
	{nil, nil, &RGB{124, 194, 81}, &RGB{113, 194, 92}, &RGB{92, 169, 143}, &RGB{92, 145, 179}, &RGB{131, 115, 176}, &RGB{116, 111, 195}, &RGB{153, 93, 168}, &RGB{156, 91, 151}, nil, nil},
	{nil, nil, &RGB{128, 198, 84}, &RGB{84, 184, 129}, &RGB{64, 151, 222}, nil, nil, &RGB{74, 126, 228}, &RGB{112, 110, 206}, &RGB{143, 100, 180}, nil, nil},
	{nil, nil, &RGB{97, 195, 125}, &RGB{67, 174, 171}, nil, nil, nil, nil, &RGB{74, 128, 234}, &RGB{108, 115, 216}, nil, nil},
	{nil, &RGB{109, 198, 148}, &RGB{98, 186, 213}, &RGB{71, 168, 220}, nil, nil, nil, nil, &RGB{61, 137, 251}, &RGB{74, 129, 240}, &RGB{101, 121, 225}, nil},
	{nil, &RGB{107, 199, 163}, &RGB{100, 182, 246}, nil, nil, nil, nil, nil, nil, &RGB{56, 134, 251}, &RGB{72, 129, 244}, nil},
	{&RGB{103, 185, 244}, &RGB{100, 182, 246}, nil, nil, nil, nil, nil, nil, nil, nil, &RGB{56, 131, 249}, &RGB{61, 133, 252}},
}

var logoC = [10][11]*RGB{
	{nil, nil, nil, &RGB{246, 145, 46}, &RGB{243, 120, 50}, &RGB{240, 95, 55}, &RGB{238, 80, 65}, &RGB{235, 70, 75}, nil, nil, nil},
	{nil, nil, &RGB{220, 175, 45}, &RGB{240, 150, 45}, &RGB{245, 130, 50}, &RGB{242, 105, 55}, &RGB{238, 85, 65}, &RGB{235, 75, 75}, &RGB{230, 70, 85}, nil, nil},
	{nil, &RGB{175, 190, 60}, &RGB{195, 185, 50}, &RGB{220, 165, 45}, nil, nil, nil, &RGB{225, 75, 95}, &RGB{210, 75, 115}, &RGB{195, 80, 135}, nil},
	{nil, &RGB{135, 195, 75}, &RGB{155, 190, 65}, nil, nil, nil, nil, nil, nil, nil, nil},
	{nil, &RGB{100, 195, 110}, &RGB{115, 195, 95}, nil, nil, nil, nil, nil, nil, nil, nil},
	{nil, &RGB{70, 180, 160}, &RGB{85, 190, 135}, nil, nil, nil, nil, nil, nil, nil, nil},
	{nil, &RGB{65, 165, 220}, &RGB{68, 175, 190}, nil, nil, nil, nil, nil, nil, nil, nil},
	{nil, &RGB{60, 140, 245}, &RGB{62, 155, 235}, &RGB{65, 165, 220}, nil, nil, nil, &RGB{120, 105, 215}, &RGB{140, 95, 200}, &RGB{160, 90, 185}, nil},
	{nil, nil, &RGB{65, 130, 250}, &RGB{62, 140, 245}, &RGB{68, 145, 240}, &RGB{80, 135, 235}, &RGB{100, 120, 225}, &RGB{120, 110, 215}, &RGB{140, 100, 205}, nil, nil},
	{nil, nil, nil, &RGB{68, 130, 250}, &RGB{75, 125, 245}, &RGB{90, 120, 235}, &RGB{110, 110, 225}, &RGB{130, 100, 215}, nil, nil, nil},
}

// RenderACLogo returns the ANSI 24-bit truecolor rendered AC pixel emblem
func RenderACLogo(indent string) string {
	var sb strings.Builder
	for r := 0; r < 10; r++ {
		sb.WriteString(indent)
		// Render 'A'
		for _, pixel := range logoA[r] {
			if pixel == nil {
				sb.WriteString("  ")
			} else {
				sb.WriteString(fmt.Sprintf("\033[48;2;%d;%d;%dm  \033[0m", pixel.R, pixel.G, pixel.B))
			}
		}
		// Gap between A and C (2 blocks = 4 spaces)
		sb.WriteString("    ")
		// Render 'C'
		for _, pixel := range logoC[r] {
			if pixel == nil {
				sb.WriteString("  ")
			} else {
				sb.WriteString(fmt.Sprintf("\033[48;2;%d;%d;%dm  \033[0m", pixel.R, pixel.G, pixel.B))
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
