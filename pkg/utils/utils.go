// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package utils

import (
	"github.com/urfave/cli/v2"
	"os"
	"regexp"
	"strings"
)

//https://stackoverflow.com/a/37293398/5685796
var rgxLeadClose = regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
var rgxInside = regexp.MustCompile(`[\s\p{Zs}]{2,}`)

// IsContains returns True if given slice contains given string item.
// Otherwise, returns False.
func IsContains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func GetArgString(args cli.Args) string {
	return TrimSliceString(args.Slice())
}

func TrimSliceString(str []string) string {
	final := rgxLeadClose.ReplaceAllString(strings.Join(str, " "), "")
	final = rgxInside.ReplaceAllString(final, " ")
	final = strings.Replace(final, " ", "+", -1)
	return final
}

func IsInstalledCMD(cmd string) (string, bool) {
	paths := []string{"/usr/bin/", "/usr/local/bin/"}
	for i := 0; i < len(paths); i++ {
		if fileExists(paths[i] + cmd) {
			return paths[i], true
		}
	}
	return "", false
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
