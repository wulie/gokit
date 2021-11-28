package main

import "fmt"

func main() {
	KeyType = makeKeyType()
	fmt.Println(KeyType)
}

var KeyType = make(map[string]int, 0)

func makeKeyType() map[string]int {
	m := map[string]int{
		"AA": 9, "AC": 4, "AD": 9, "AL": 4, "AN": 9, "AP": 4, "AR": 4, "AU": 4, "AV": 0, "AVGV": 7, "BP": 3, "BV": 6, "C1": 4, "C2": 4,
		"C3": 4, "C4": 4, "C5": 4, "C6": 4, "C7": 4, "C8": 4, "CD": 4, "CP": 3, "CT": 0,
		"DB": 6, "DL": 4, "DS": 3, "DT": 4, "ED": 9, "ET": 0, "EU": 9, "EX": 9, "FB": 6, "FK": 6, "FL": 4, "FLOW": 7, "FM": 3, "FQ": 3,
		"FS": 4, "GN": 9, "GP": 9, "H3": 6, "H4": 6, "HI": 9, "HL": 6, "HO": 9, "HW": 4, "ID": 4,
		"IL": 4, "IO": 4, "IP": 9, "IT": 4, "IV": 6, "KR": 9, "KS": 4, "KT": 4, "KZ": 4, "L3": 6, "L4": 6, "LC": 3, "LG": 4, "LI": 9,
		"LL": 6, "LO": 4, "LZ": 4, "MAXTIME": 0, "MAXV": 7, "MINTIME": 0, "MINV": 7, "MM": 4,
		"MT": 4, "ND": 4, "OF": 4, "OT": 3, "PD": 4, "PJ": 9, "PL": 9, "PN": 9, "PO": 4, "PS": 4, "PT": 4, "PW": 9, "RS": 9, "RT": 4,
		"SD": 9, "SG": 0, "SL": 4, "SR": 9, "ST": 9, "SY": 4, "SZ": 4, "TA": 0, "TD": 4, "TF": 0, "TI": 4,
		"TM": 0, "TP": 3, "TS": 4, "TT": 4, "TV": 6, "UD": 5, "US": 9, "VN": 9, "WT": 4, "ZH": 6, "ZL": 6}
	return m

}
