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
	RoundHalf DecimalRoundingMode = iota
	RoundDown
	RoundUp
)

// MaxSupportedValue is the maximum number we can reliably convert to Thai text
// This is set to 9,223,372,036,854,775,807 (19 digits) which is int64 maximum
// and a practical limit for Thai currency representation
const MaxSupportedValue = "9223372036854775807"

var digitNames = map[int]string{
	1: "หนึ่ง", 2: "สอง", 3: "สาม", 4: "สี่", 5: "ห้า",
	6: "หก", 7: "เจ็ด", 8: "แปด", 9: "เก้า",
}

var unitNames = map[int]string{
	0: "", 1: "สิบ", 2: "ร้อย", 3: "พัน", 4: "หมื่น", 5: "แสน", 6: "ล้าน",
}

// EnableWarningLogs controls whether warning logs are printed when satang is capped at 99
var EnableWarningLogs = true

// AllowOverflow controls whether rounding can overflow to the next baht amount
var AllowOverflow = false

// SetWarningLogs enables or disables warning logs for satang capping
func SetWarningLogs(enabled bool) {
	EnableWarningLogs = enabled
}

// SetAllowOverflow enables or disables overflow behavior for rounding
func SetAllowOverflow(enabled bool) {
	AllowOverflow = enabled
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

	// Remove commas from input (e.g., "1,234,567" -> "1234567")
	amountStr = strings.ReplaceAll(amountStr, ",", "")

	// Validate that the number doesn't exceed our maximum supported value
	if err := validateMaxValue(amountStr); err != nil {
		return "", err
	}

	parts := strings.Split(amountStr, ".")
	integerPart := parts[0]

	var decimalPart string
	var overflow bool
	if len(parts) > 1 {
		decimalPart, overflow = formatDecimalPartWithRounding(parts[1], mode)

		// Handle overflow case where satang rounds up to 100
		if overflow {
			integerNum, err := strconv.Atoi(integerPart)
			if err == nil {
				decimalPart = "00" // Reset to 00 satang
				integerPart = strconv.Itoa(integerNum + 1)
			}
		}
	}

	bahtText := convertIntegerNumber(integerPart)
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

// validateMaxValue checks if the input number exceeds our maximum supported value
func validateMaxValue(amountStr string) error {
	// Extract just the integer part (before decimal point)
	parts := strings.Split(amountStr, ".")
	integerPart := parts[0]

	// Remove any leading zeros for comparison
	integerPart = strings.TrimLeft(integerPart, "0")
	if integerPart == "" {
		integerPart = "0"
	}

	// Check if the number of digits exceeds our maximum
	if len(integerPart) > len(MaxSupportedValue) {
		return fmt.Errorf("input number exceeds maximum supported value of %s (got %d digits, max %d digits)", MaxSupportedValue, len(integerPart), len(MaxSupportedValue))
	}

	// If same number of digits, do numeric comparison
	if len(integerPart) == len(MaxSupportedValue) {
		// Parse both as big integers for proper comparison
		inputNum, err1 := strconv.ParseUint(integerPart, 10, 64)
		maxNum, err2 := strconv.ParseUint(MaxSupportedValue, 10, 64)

		// If either parsing fails, fall back to string comparison
		if err1 != nil || err2 != nil {
			if integerPart > MaxSupportedValue {
				return fmt.Errorf("input number exceeds maximum supported value of %s", MaxSupportedValue)
			}
		} else if inputNum > maxNum {
			return fmt.Errorf("input number exceeds maximum supported value of %s", MaxSupportedValue)
		}
	}

	return nil
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
		first2Digits := decimal[:2]
		thirdDigit, _ := strconv.Atoi(string(decimal[2]))

		// Convert first 2 digits to integer for rounding calculation
		value, _ := strconv.Atoi(first2Digits)
		originalValue := value
		warningMsg := "Warning: %s rounds to 100 satang, forced to round down to 99 satang to maintain currency format. Consider enabling AllowOverflow."

		switch roundingMode {
		case RoundDown:
			return first2Digits, false
		case RoundUp:
			if len(decimal) > 2 && thirdDigit > 0 {
				value++
				if value >= 100 {
					if AllowOverflow {
						return "00", true
					} else {
						if originalValue == 99 && EnableWarningLogs {
							log.Printf(warningMsg, decimal)
						}
						value = 99
					}
				}
			}
		case RoundHalf:
			if thirdDigit >= 5 {
				value++
				if value >= 100 {
					if AllowOverflow {
						return "00", true
					} else {
						if originalValue == 99 && EnableWarningLogs {
							log.Printf(warningMsg, decimal)
						}
						value = 99
					}
				}
			}
		}

		return fmt.Sprintf("%02d", value), false
	}

	return decimal, false
}

func convertIntegerNumber(numberStr string) string {
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

// countNonZeroGroups counts how many 6-digit groups contain non-zero digits
func countNonZeroGroups(digits []int) int {
	digitCount := len(digits)
	count := 0

	for startPos := digitCount; startPos > 0; startPos -= 6 {
		endPos := max(startPos-6, 0)
		group := digits[endPos:startPos]

		// Check if group has any non-zero digits
		hasNonZero := false
		for _, digit := range group {
			if digit != 0 {
				hasNonZero = true
				break
			}
		}

		if hasNonZero {
			count++
		}
	}

	return count
}

func buildThaiText(digits []int) string {
	digitCount := len(digits)
	if digitCount <= 6 {
		return convertSixDigitGroup(digits)
	}

	var result []string

	// Process in groups of 6 digits from right to left
	groupsFromRight := 0
	for startPos := digitCount; startPos > 0; startPos -= 6 {
		endPos := max(startPos-6, 0)
		group := digits[endPos:startPos]
		groupText := convertSixDigitGroup(group)

		if groupText != "" {
			// Add "ล้าน" suffix based on pattern:
			// - For numbers where most groups are zeros (like 1,000,000,000,000):
			//   the non-zero group gets multiple ล้าน based on total groups
			// - For numbers with digits in multiple groups:
			//   each group gets single ล้าน except rightmost

			// Check if this is a "telescoping zeros" pattern by counting non-zero groups
			hasMultipleNonZeroGroups := countNonZeroGroups(digits)

			if hasMultipleNonZeroGroups > 1 {
				// Multiple groups have non-zero digits: use single ล้าน rule
				if groupsFromRight > 0 {
					groupText += "ล้าน"
				}
			} else {
				// Only one group has non-zero digits: use multiple ล้าน rule
				for i := 0; i < groupsFromRight; i++ {
					groupText += "ล้าน"
				}
			}

			result = append([]string{groupText}, result...)
		}
		groupsFromRight++
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
		return convertIntegerNumber(decimalStr)
	}
}
