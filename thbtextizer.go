package thbtextizer

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type DecimalRoundingMode int

const (
	RoundHalf         DecimalRoundingMode = iota // Always round half - DEFAULT
	RoundDown                                    // Always round down (truncate)
	RoundUp                                      // Always round up (capping satang at 99)
	RoundUpOverflow                              // Round up with overflow (can increase baht amount)
	RoundHalfOverflow                            // Round half with overflow (can increase baht amount)
)

var digitNames = map[int]string{
	1: "หนึ่ง", 2: "สอง", 3: "สาม", 4: "สี่", 5: "ห้า",
	6: "หก", 7: "เจ็ด", 8: "แปด", 9: "เก้า",
}

var unitNames = map[int]string{
	0: "", 1: "สิบ", 2: "ร้อย", 3: "พัน", 4: "หมื่น", 5: "แสน", 6: "ล้าน",
}

// EnableWarningLogs controls whether warning logs are printed when satang is capped at 99
var EnableWarningLogs = true

// SetWarningLogs enables or disables warning logs for satang capping
func SetWarningLogs(enabled bool) {
	EnableWarningLogs = enabled
}

func Convert(amount any, roundingMode ...DecimalRoundingMode) (string, error) {
	// Default to RoundHalf if no mode specified
	mode := RoundHalf
	if len(roundingMode) > 0 {
		mode = roundingMode[0]
	}

	// Convert any numeric type to string
	amountStr, err := convertToString(amount)
	if err != nil {
		return "", err
	}

	parts := strings.Split(amountStr, ".")
	wholePart := parts[0]

	var decimalPart string
	var overflow bool
	if len(parts) > 1 {
		decimalPart, overflow = formatDecimalPartWithRounding(parts[1], mode)

		// Handle overflow case where satang rounds up to 100
		if overflow {
			// Increment whole part by 1
			wholeNum, err := strconv.Atoi(wholePart)
			if err == nil {
				wholePart = strconv.Itoa(wholeNum + 1)
				decimalPart = "00" // Reset to 00 satang
			}
		}
	}

	bahtText := convertWholeNumber(wholePart)
	if bahtText == "" {
		bahtText = "ศูนย์"
	}
	bahtText += "บาท"

	if decimalPart == "" || decimalPart == "00" {
		bahtText += "ถ้วน"
	} else {
		satangText := convertDecimalPart(decimalPart)
		if satangText == "" {
			satangText = "ศูนย์"
		}
		bahtText += satangText + "สตางค์"
	}

	return bahtText, nil
}

// convertToString converts various numeric types to string representation
func convertToString(amount any) (string, error) {
	switch v := amount.(type) {
	case string:
		return v, nil
	case int:
		return fmt.Sprintf("%d", v), nil
	case int8:
		return fmt.Sprintf("%d", v), nil
	case int16:
		return fmt.Sprintf("%d", v), nil
	case int32:
		return fmt.Sprintf("%d", v), nil
	case int64:
		return fmt.Sprintf("%d", v), nil
	case uint:
		return fmt.Sprintf("%d", v), nil
	case uint8:
		return fmt.Sprintf("%d", v), nil
	case uint16:
		return fmt.Sprintf("%d", v), nil
	case uint32:
		return fmt.Sprintf("%d", v), nil
	case uint64:
		return fmt.Sprintf("%d", v), nil
	case float32:
		return fmt.Sprintf("%.2f", v), nil
	case float64:
		return fmt.Sprintf("%.2f", v), nil
	default:
		return "", errors.New("unsupported type: only string, int, uint, float32, float64 and their variants are supported")
	}
}

func formatDecimalPartWithRounding(decimal string, roundingMode DecimalRoundingMode) (string, bool) {
	if len(decimal) == 0 {
		return "00", false
	}
	if len(decimal) == 1 {
		return decimal + "0", false
	}
	if len(decimal) == 2 {
		return decimal, false
	}

	// Handle more than 2 decimal places with rounding
	if len(decimal) > 2 {
		// Get first 2 digits and the third digit for rounding decision
		first2 := decimal[:2]
		thirdDigit, _ := strconv.Atoi(string(decimal[2]))

		// Convert first 2 digits to integer for rounding calculation
		value, _ := strconv.Atoi(first2)
		originalValue := value

		warningMsg := "Warning: %.3s rounds to 100 satang, forced to round down to 99 satang to maintain currency format. Consider using %s mode."

		switch roundingMode {
		case RoundDown:
			// Always truncate (keep as is)
			return first2, false
		case RoundUp:
			// Always round up if there are any additional digits
			if len(decimal) > 2 && thirdDigit > 0 {
				value++
				if value >= 100 {
					if roundingMode == RoundUpOverflow {
						return "00", true // Overflow to next baht
					} else {
						// Capping satang at 99 and log warning for special cases
						if originalValue == 99 && EnableWarningLogs {
							log.Printf(warningMsg, decimal, "RoundUpOverflow")
						}
						value = 99
					}
				}
			}
		case RoundUpOverflow:
			// Round up with overflow capability
			if len(decimal) > 2 {
				value++
				if value >= 100 {
					return "00", true // Overflow to next baht
				}
			}
		case RoundHalf:
			// Round half: round up if third digit >= 5
			if thirdDigit >= 5 {
				value++
				if value >= 100 {
					// Special case for 0.995+ - force round down and log
					if originalValue == 99 && EnableWarningLogs {
						log.Printf(warningMsg, decimal, "RoundHalfOverflow")
					}
					value = 99 // Cap at 99 satang
				}
			}
		case RoundHalfOverflow:
			// Round half with overflow capability
			if thirdDigit >= 5 {
				value++
				if value >= 100 {
					return "00", true // Overflow to next baht
				}
			}
		}

		// Format back to 2-digit string with leading zero if needed
		return fmt.Sprintf("%02d", value), false
	}

	return decimal, false
}

func convertWholeNumber(numberStr string) string {
	if !isValidNumber(numberStr) {
		return ""
	}

	digits := parseDigits(numberStr)
	if len(digits) == 0 {
		return ""
	}

	return buildThaiText(digits)
}

func isValidNumber(str string) bool {
	for _, char := range str {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func parseDigits(numberStr string) []int {
	digits := make([]int, 0, len(numberStr))
	for _, char := range numberStr {
		digit, _ := strconv.Atoi(string(char))
		digits = append(digits, digit)
	}
	return digits
}

func buildThaiText(digits []int) string {
	digitCount := len(digits)
	if digitCount <= 6 {
		// Handle numbers up to 6 digits directly
		return convertSixDigitGroup(digits)
	}

	var result []string

	// Process in groups of 6 digits from right to left
	for startPos := digitCount; startPos > 0; startPos -= 6 {
		endPos := max(startPos-6, 0)
		group := digits[endPos:startPos]
		groupText := convertSixDigitGroup(group)

		if groupText != "" {
			millionLevel := (digitCount - startPos) / 6

			if millionLevel > 0 {
				groupText += "ล้าน"
			}

			result = append([]string{groupText}, result...)
		}
	}

	return strings.Join(result, "")
}

func convertSixDigitGroup(digits []int) string {
	var result []string
	digitCount := len(digits)

	for position, digit := range digits {
		if digit == 0 {
			continue
		}

		positionFromRight := digitCount - position - 1
		unitIndex := positionFromRight % 6

		text := convertDigitAtPosition(digit, unitIndex, positionFromRight, len(digits))
		if text != "" {
			result = append(result, text)
		}
	}

	return strings.Join(result, "")
}

func convertDigitAtPosition(digit, unitIndex, positionFromRight, totalDigits int) string {
	digitName := digitNames[digit]
	unitName := unitNames[unitIndex]

	switch unitIndex {
	case 0: // ones place
		if digit == 1 && totalDigits > 1 && positionFromRight == 0 {
			return "เอ็ด" + unitName
		}
		return digitName + unitName

	case 1: // tens place
		switch digit {
		case 1:
			return unitName
		case 2:
			return "ยี่" + unitName
		default:
			return digitName + unitName
		}

	default: // hundreds, thousands, etc.
		return digitName + unitName
	}
}

func convertDecimalPart(decimalStr string) string {
	if !isValidNumber(decimalStr) {
		return ""
	}

	// Convert to integer for easy comparison
	value, _ := strconv.Atoi(decimalStr)

	// Special cases for decimal satang conversion
	switch {
	case value == 1:
		return "หนึ่ง" // 01 -> หนึ่งสตางค์
	case value == 11:
		return "สิบเอ็ด" // 11 -> สิบเอ็ดสตางค์
	case value >= 12 && value <= 19:
		// 12-19: regular conversion (สิบสอง, สิบสาม, etc.)
		ones := value - 10
		return "สิบ" + digitNames[ones]
	case value >= 21 && value <= 99 && value%10 == 1:
		// 21, 31, 41, etc.: use เอ็ด for ones place
		tens := value / 10
		if tens == 2 {
			return "ยี่สิบเอ็ด"
		}
		return digitNames[tens] + "สิบเอ็ด"
	default:
		// For all other cases, use regular conversion
		return convertWholeNumber(decimalStr)
	}
}
