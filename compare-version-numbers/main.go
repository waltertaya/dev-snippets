package main

import (
	"fmt"
	"strconv"
	"strings"
)

// split into arrays method
func compareVersion(version1 string, version2 string) int {
	var result int
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	var v1Extra, v2Extra string

	if len(v1) > len(v2) {
		v1Extra = strings.Join(v1[len(v2):], ".")
		v1 = v1[:len(v2)]
	} else if len(v2) > len(v1) {
		v2Extra = strings.Join(v2[len(v1):], ".")
		v2 = v2[:len(v1)]
	}

	for i := 0; i < len(v1); i++ {
		v1_, _ := strconv.Atoi(v1[i])
		v2_, _ := strconv.Atoi(v2[i])
		if v1_ > v2_ {
			return 1
		} else if v1_ < v2_ {
			return -1
		}
	}

	if strings.ContainsAny(v1Extra, "123456789") {
		return 1
	}
	if strings.ContainsAny(v2Extra, "123456789") {
		return -1
	}

	return result
}

// two pointers solution
func compareVersion2Pointer(version1 string, version2 string) int {
    i, j := 0, 0
    n, m := len(version1), len(version2)

    for i < n || j < m {
        num1, num2 := 0, 0

        for i < n && version1[i] != '.' {
            num1 = num1*10 + int(version1[i]-'0')
            i++
        }
        if i < n && version1[i] == '.' {
            i++
        }

        for j < m && version2[j] != '.' {
            num2 = num2*10 + int(version2[j]-'0')
            j++
        }
        if j < m && version2[j] == '.' {
            j++
        }

        if num1 > num2 {
            return 1
        }
        if num1 < num2 {
            return -1
        }
    }

    return 0
}

func main() {

	// // Test case for method 1
	// version1, version2 := "1.2", "1.10"
	// fmt.Println(compareVersion(version1, version2)) // -1

	// version1, version2 = "1.01", "1.001"
	// fmt.Println(compareVersion(version1, version2)) // 0

	// version1, version2 = "1.0", "1.0.0.0"
	// fmt.Println(compareVersion(version1, version2)) // 0

	// // Test case for method 2 (two pointer solution)
	// version1, version2 = "1.2", "1.10"
	// fmt.Println(compareVersion2Pointer(version1, version2)) // -1

	// version1, version2 = "1.01", "1.001"
	// fmt.Println(compareVersion2Pointer(version1, version2)) // 0

	// version1, version2 = "1.0", "1.0.0.0"
	// fmt.Println(compareVersion2Pointer(version1, version2)) // 0

	version1, version2 := "1.2", "1.10"
	fmt.Println(compareVersion(version1, version2) == compareVersion2Pointer(version1, version2)) // -1

	version1, version2 = "1.01", "1.001"
	fmt.Println(compareVersion(version1, version2) == compareVersion2Pointer(version1, version2)) // 0

	version1, version2 = "1.0", "1.0.0.0"
	fmt.Println(compareVersion(version1, version2) == compareVersion2Pointer(version1, version2)) // 0
}
