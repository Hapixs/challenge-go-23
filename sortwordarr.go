package piscine

func SortWordArr(s []string) {
	for i := 0; i < len(s); i++ {
		for j := i; j > 0 && rune(s[j-1][0]) > rune(s[j][0]); j-- {
			s[j-1], s[j] = s[j], s[j-1]
		}
	}
}

/*
[3 2j9xRyVdyUe j9 ctpX bwwgs QpCG nlFDZwK5 iUuZ532NG 3Hm 8owmxMG 47V   V7O89 n5LsL V Iy NWc2ecd je V pFH K UKkLl]
[3 2j9xRyVdyUe 8owmxMG 47V   Iy NWc2ecd je QpCG nlFDZwK5 V7O89 n5LsL V V pFH K UKkLl iUuZ532NG 3Hm j9 ctpX bwwgs]
[3 2j9xRyVdyUe 8owmxMG 47V   Iy NWc2ecd je QpCG nlFDZwK5 V pFH K UKkLl V7O89 n5LsL V iUuZ532NG 3Hm j9 ctpX bwwgs]

[ HlNj8yy00Cdp 7Ibkunql  Kkm 8 6sd PcZpfx3 HkjbLGnw5n NG MBqFmM9 V2CBP glw3lNkQ fzss g5URlYJ8oCZ5G lbLEZV6npYp5L]
[ HlNj8yy00Cdp 7Ibkunql  Kkm 8 6sd PcZpfx3 HkjbLGnw5n NG MBqFmM9 V2CBP g5URlYJ8oCZ5G glw3lNkQ fzss lbLEZV6npYp5L]
*/
