package manage_str

import "strings"

func commonSubStr(str1, str2 string, minLength int) string {
	len1, len2 := len(str1), len(str2)
	longest := 0
	lastSubArray := make([]int, len2)
	startSubStr := 0

	for i := 0; i < len1; i++ {
		curSubArray := make([]int, len2)
		for j := 0; j < len2; j++ {
			if str1[i] != str2[j] {
				curSubArray[j] = 0
			} else {
				if i == 0 || j == 0 {
					curSubArray[j] = 1
				} else {
					curSubArray[j] = lastSubArray[j-1] + 1
				}

				if curSubArray[j] > longest {
					longest = curSubArray[j]
					if longest >= minLength {
						startSubStr = i - longest + 1
					}
				}
			}
		}
		lastSubArray = curSubArray
	}

	if longest < minLength {
		return ""
	}

	return str1[startSubStr : startSubStr+longest]
}

func mergeStringsBasedOnCommonParts(stringsArr []string, minLength int) []string {
	mergedStrings := []string{}

	for i := 0; i < len(stringsArr); i++ {
		for j := i + 1; j < len(stringsArr); j++ {
			common := commonSubStr(stringsArr[i], stringsArr[j], minLength)
			if common != "" {
				index := strings.Index(stringsArr[i], common)
				merged := stringsArr[i][:index] + stringsArr[j]
				mergedStrings = append(mergedStrings, merged)
			}
		}
	}

	return mergedStrings
}
