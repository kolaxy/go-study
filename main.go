package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	cows := [100]string{"1 korova", "2 korovy", "3 korovy", "4 korovy", "5 korov", "6 korov", "7 korov", "8 korov", "9 korov", "10 korov", "11 korov", "12 korov", "13 korov", "14 korov", "15 korov", "16 korov", "17 korov", "18 korov", "19 korov", "20 korov", "21 korova", "22 korovy", "23 korovy", "24 korovy", "25 korov", "26 korov", "27 korov", "28 korov", "29 korov", "30 korov", "31 korova", "32 korovy", "33 korovy", "34 korovy", "35 korov", "36 korov", "37 korov", "38 korov", "39 korov", "40 korov", "41 korova", "42 korovy", "43 korovy", "44 korovy", "45 korov", "46 korov", "47 korov", "48 korov", "49 korov", "50 korov", "51 korova", "52 korovy", "53 korovy", "54 korovy", "55 korov", "56 korov", "57 korov", "58 korov", "59 korov", "60 korov", "61 korova", "62 korovy", "63 korovy", "64 korovy", "65 korov", "66 korov", "67 korov", "68 korov", "69 korov", "70 korov", "71 korova", "72 korovy", "73 korovy", "74 korovy", "75 korov", "76 korov", "77 korov", "78 korov", "79 korov", "80 korov", "81 korova", "82 korovy", "83 korovy", "84 korovy", "85 korov", "86 korov", "87 korov", "88 korov", "89 korov", "90 korov", "91 korova", "92 korovy", "93 korovy", "94 korovy", "95 korov", "96 korov", "97 korov", "98 korov", "99 korov", "100 korov"}
	fmt.Println(cows[num-1])
}