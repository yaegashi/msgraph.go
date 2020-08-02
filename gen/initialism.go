package main

import (
	"sort"
	"strings"
	"unicode"
)

// commonInitialisms is a set of common initialisms.
// Only add entries that are highly unlikely to be non-initialisms.
// For instance, "ID" is fine (Freudian code is rare), but "AND" is not.
// Taken from https://github.com/golang/lint/blob/master/lint.go
var commonInitialisms = map[string]string{
	"ACL":    "ACL",
	"API":    "API",
	"ASCII":  "ASCII",
	"CIDR":   "CIDR",
	"CPU":    "CPU",
	"CSS":    "CSS",
	"DMA":    "DMA",
	"DNS":    "DNS",
	"EOF":    "EOF",
	"GUID":   "GUID",
	"HTML":   "HTML",
	"HTTP":   "HTTP",
	"HTTPS":  "HTTPS",
	"ICMP":   "ICMP",
	"ID":     "ID",
	"IDS":    "IDs",
	"IGMP":   "IGMP",
	"IOS":    "IOS",
	"ITUNES": "ITunes",
	"IP":     "IP",
	"IPV4":   "IPv4",
	"IPV6":   "IPv6",
	"JSON":   "JSON",
	"LHS":    "LHS",
	"MDM":    "MDM",
	"MFA":    "MFA",
	"NDES":   "NDES",
	"OAUTH":  "OAuth",
	"OMA":    "OMA",
	"QPS":    "QPS",
	"RAM":    "RAM",
	"RHS":    "RHS",
	"RPC":    "RPC",
	"SKU":    "SKU",
	"SKUS":   "SKUs",
	"SLA":    "SLA",
	"SMTP":   "SMTP",
	"SQL":    "SQL",
	"SSH":    "SSH",
	"TCP":    "TCP",
	"TLS":    "TLS",
	"TTL":    "TTL",
	"UDP":    "UDP",
	"UI":     "UI",
	"UID":    "UID",
	"URI":    "URI",
	"URL":    "URL",
	"UTF8":   "UTF8",
	"UUID":   "UUID",
	"VM":     "VM",
	"VPN":    "VPN",
	"VPP":    "VPP",
	"WIFI":   "WiFi",
	"XML":    "XML",
	"XMPP":   "XMPP",
	"XSRF":   "XSRF",
	"XSS":    "XSS",
}

// commonInitialismValues is a list of the commonInitialisms values
// sorted by the length for prefix matches in splitName().
var commonInitialismValues []string

func init() {
	for _, value := range commonInitialisms {
		commonInitialismValues = append(commonInitialismValues, value)
	}
	sort.Slice(commonInitialismValues, func(i, j int) bool { return len(commonInitialismValues[i]) > len(commonInitialismValues[j]) })
}

// lintName returns a different name if it should be different.
// Taken from https://github.com/golang/lint/blob/master/lint.go
func lintName(name string) (should string) {
	// Split camelCase at any lower->upper transition, and split on underscores.
	// Check each word for common initialisms.
	runes := []rune(name)
	w, i := 0, 0 // index of start of word, scan
	for i+1 <= len(runes) {
		eow := false // whether we hit the end of a word
		if i+1 == len(runes) {
			eow = true
		} else if unicode.IsLower(runes[i]) && !unicode.IsLower(runes[i+1]) {
			// lower->non-lower
			eow = true
		}
		i++
		if !eow {
			continue
		}

		// [w,i) is a word.
		word := string(runes[w:i])
		if u, ok := commonInitialisms[strings.ToUpper(word)]; ok {
			// Keep consistent case, which is lowercase only at the start.
			if w == 0 && unicode.IsLower(runes[w]) {
				u = strings.ToLower(u)
			}
			// All the common initialisms are ASCII,
			// so we can replace the bytes exactly.
			copy(runes[w:], []rune(u))
		} else if w > 0 && strings.ToLower(word) == word {
			// already all lowercase, and not the first word, so uppercase the first character.
			runes[w] = unicode.ToUpper(runes[w])
		}
		w = i
	}
	return string(runes)
}

// splitName returns a slice of CamelCase words from name.
func splitName(name string) (words []string) {
loop:
	for i := 0; len(name) > 0; i++ {
		for _, v := range commonInitialismValues {
			if strings.HasPrefix(name, v) {
				words = append(words, v)
				name = name[len(v):]
				continue loop
			}
		}
		for i, c := range name {
			if i != 0 && unicode.IsUpper(c) {
				words = append(words, name[:i])
				name = name[i:]
				continue loop
			}
		}
		words = append(words, name)
		name = ""
	}
	return
}
