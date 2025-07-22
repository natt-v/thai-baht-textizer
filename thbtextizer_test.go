package thbtextizer

import (
	"testing"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "147521.19",
			expected: "หนึ่งแสนสี่หมื่นเจ็ดพันห้าร้อยยี่สิบเอ็ดบาทสิบเก้าสตางค์",
		},
		{
			input:    "147521",
			expected: "หนึ่งแสนสี่หมื่นเจ็ดพันห้าร้อยยี่สิบเอ็ดบาทถ้วน",
		},
		{
			input:    "147521.00",
			expected: "หนึ่งแสนสี่หมื่นเจ็ดพันห้าร้อยยี่สิบเอ็ดบาทถ้วน",
		},
		{
			input:    "0",
			expected: "ศูนย์บาทถ้วน",
		},
		{
			input:    "0.50",
			expected: "ศูนย์บาทห้าสิบสตางค์",
		},
		{
			input:    "1000000",
			expected: "หนึ่งล้านบาทถ้วน",
		},
		{
			input:    "1000000.25",
			expected: "หนึ่งล้านบาทยี่สิบห้าสตางค์",
		},
		{
			input:    "100.01",
			expected: "หนึ่งร้อยบาทหนึ่งสตางค์",
		},
		{
			input:    "50.05",
			expected: "ห้าสิบบาทห้าสตางค์",
		},
		{
			input:    "11",
			expected: "สิบเอ็ดบาทถ้วน",
		},
		{
			input:    "21",
			expected: "ยี่สิบเอ็ดบาทถ้วน",
		},
		{
			input:    "31",
			expected: "สามสิบเอ็ดบาทถ้วน",
		},
		{
			input:    "91",
			expected: "เก้าสิบเอ็ดบาทถ้วน",
		},
		{
			input:    "1",
			expected: "หนึ่งบาทถ้วน",
		},
		{
			input:    "101",
			expected: "หนึ่งร้อยเอ็ดบาทถ้วน",
		},
		{
			input:    "100.11",
			expected: "หนึ่งร้อยบาทสิบเอ็ดสตางค์",
		},
		{
			input:    "111",
			expected: "หนึ่งร้อยสิบเอ็ดบาทถ้วน",
		},
		{
			input:    "1001",
			expected: "หนึ่งพันเอ็ดบาทถ้วน",
		},
		{
			input:    "2501",
			expected: "สองพันห้าร้อยเอ็ดบาทถ้วน",
		},
		{
			input:    "100000001.01",
			expected: "หนึ่งร้อยล้านเอ็ดบาทหนึ่งสตางค์",
		},
		{
			input:    "100.21",
			expected: "หนึ่งร้อยบาทยี่สิบเอ็ดสตางค์",
		},
		{
			input:    "100.31",
			expected: "หนึ่งร้อยบาทสามสิบเอ็ดสตางค์",
		},
		{
			input:    "0",
			expected: "ศูนย์บาทถ้วน",
		},
		{
			input:    "21.25",
			expected: "ยี่สิบเอ็ดบาทยี่สิบห้าสตางค์",
		},
		{
			input:    "1234567.89",
			expected: "หนึ่งล้านสองแสนสามหมื่นสี่พันห้าร้อยหกสิบเจ็ดบาทแปดสิบเก้าสตางค์",
		},
		{
			input:    "500200300.00",
			expected: "ห้าร้อยล้านสองแสนสามร้อยบาทถ้วน",
		},
		{
			input:    "999999999.99",
			expected: "เก้าร้อยเก้าสิบเก้าล้านเก้าแสนเก้าหมื่นเก้าพันเก้าร้อยเก้าสิบเก้าบาทเก้าสิบเก้าสตางค์",
		},
		{
			input:    "1,234,567,889,999,999,999",
			expected: "หนึ่งล้านสองแสนสามหมื่นสี่พันห้าร้อยหกสิบเจ็ดล้านแปดแสนแปดหมื่นเก้าพันเก้าร้อยเก้าสิบเก้าล้านเก้าแสนเก้าหมื่นเก้าพันเก้าร้อยเก้าสิบเก้าบาทถ้วน",
		},
		{
			input:    "9,223,372,036,854,775,807",
			expected: "เก้าล้านสองแสนสองหมื่นสามพันสามร้อยเจ็ดสิบสองล้านสามหมื่นหกพันแปดร้อยห้าสิบสี่ล้านเจ็ดแสนเจ็ดหมื่นห้าพันแปดร้อยเจ็ดบาทถ้วน",
		},
		{
			input:    "1,000,000,000,000,000,000",
			expected: "หนึ่งล้านล้านล้านบาทถ้วน",
		},
		{
			input:    "100,000,000,000,000,000",
			expected: "หนึ่งแสนล้านล้านบาทถ้วน",
		},
		{
			input:    "10,000,000,000,000,000",
			expected: "หนึ่งหมื่นล้านล้านบาทถ้วน",
		},
		{
			input:    "1,000,000,000,000,000",
			expected: "หนึ่งพันล้านล้านบาทถ้วน",
		},
		{
			input:    "100,000,000,000,000",
			expected: "หนึ่งร้อยล้านล้านบาทถ้วน",
		},
		{
			input:    "10,000,000,000,000",
			expected: "สิบล้านล้านบาทถ้วน",
		},
		{
			input:    "1,000,000,000,000",
			expected: "หนึ่งล้านล้านบาทถ้วน",
		},
	}

	for _, test := range tests {
		result, err := Convert(test.input)
		if err != nil {
			t.Errorf("Convert(%s) returned error: %v", test.input, err)
			continue
		}
		if result != test.expected {
			t.Errorf("Convert(%s) = %s, expected %s", test.input, result, test.expected)
		}
	}
}

func TestConvertWithRounding(t *testing.T) {
	tests := []struct {
		input        string
		roundingMode DecimalRoundingMode
		expected     string
	}{
		{
			input:        "123.456",
			roundingMode: RoundHalf,
			expected:     "หนึ่งร้อยยี่สิบสามบาทสี่สิบหกสตางค์", // 0.456 -> 0.46
		},
		{
			input:        "123.454",
			roundingMode: RoundHalf,
			expected:     "หนึ่งร้อยยี่สิบสามบาทสี่สิบห้าสตางค์", // 0.454 -> 0.45
		},
		{
			input:        "123.455",
			roundingMode: RoundHalf,
			expected:     "หนึ่งร้อยยี่สิบสามบาทสี่สิบหกสตางค์", // 0.455 -> 0.46 (round half up)
		},
		// Test RoundDown (truncate)
		{
			input:        "123.456",
			roundingMode: RoundDown,
			expected:     "หนึ่งร้อยยี่สิบสามบาทสี่สิบห้าสตางค์", // 0.456 -> 0.45
		},
		{
			input:        "123.459",
			roundingMode: RoundDown,
			expected:     "หนึ่งร้อยยี่สิบสามบาทสี่สิบห้าสตางค์", // 0.459 -> 0.45
		},
		// Test RoundUp
		{
			input:        "123.451",
			roundingMode: RoundUp,
			expected:     "หนึ่งร้อยยี่สิบสามบาทสี่สิบหกสตางค์", // 0.451 -> 0.46
		},
		{
			input:        "123.459",
			roundingMode: RoundUp,
			expected:     "หนึ่งร้อยยี่สิบสามบาทสี่สิบหกสตางค์", // 0.459 -> 0.46
		},
		// Edge cases
		{
			input:        "100.995",
			roundingMode: RoundHalf,
			expected:     "หนึ่งร้อยบาทเก้าสิบเก้าสตางค์", // 0.995 -> 0.99 (capped)
		},
		{
			input:        "100.991",
			roundingMode: RoundUp,
			expected:     "หนึ่งร้อยบาทเก้าสิบเก้าสตางค์", // 0.991 -> 0.99 (capped)
		},
		// Test case for 100.990 with different rounding modes
		{
			input:        "100.990",
			roundingMode: RoundUp,
			expected:     "หนึ่งร้อยบาทเก้าสิบเก้าสตางค์", // 0.990 -> 0.99
		},
		{
			input:        "100.990",
			roundingMode: RoundHalf,
			expected:     "หนึ่งร้อยบาทเก้าสิบเก้าสตางค์", // 0.990 -> 0.99
		},
		{
			input:        "123.456",
			roundingMode: -1,                                    // Special marker for testing default behavior
			expected:     "หนึ่งร้อยยี่สิบสามบาทสี่สิบหกสตางค์", // Default should use RoundHalf
		},
	}

	for _, test := range tests {
		var result string
		var err error

		if test.roundingMode == -1 {
			// Test default behavior (no rounding mode specified)
			result, err = Convert(test.input)
		} else {
			result, err = Convert(test.input, test.roundingMode)
		}

		if err != nil {
			t.Errorf("Convert(%s, %v) returned error: %v", test.input, test.roundingMode, err)
			continue
		}
		if result != test.expected {
			t.Errorf("Convert(%s, %v) = %s, expected %s",
				test.input, test.roundingMode, result, test.expected)
		}
	}
}

func TestConvertWithNumericTypes(t *testing.T) {
	tests := []struct {
		input    any
		expected string
		name     string
	}{
		// String inputs
		{input: "123.45", expected: "หนึ่งร้อยยี่สิบสามบาทสี่สิบห้าสตางค์", name: "string"},

		// Integer inputs
		{input: 123, expected: "หนึ่งร้อยยี่สิบสามบาทถ้วน", name: "int"},
		{input: int8(50), expected: "ห้าสิบบาทถ้วน", name: "int8"},
		{input: int16(1000), expected: "หนึ่งพันบาทถ้วน", name: "int16"},
		{input: int32(25000), expected: "สองหมื่นห้าพันบาทถ้วน", name: "int32"},
		{input: int64(100000), expected: "หนึ่งแสนบาทถ้วน", name: "int64"},

		// Unsigned integer inputs
		{input: uint(456), expected: "สี่ร้อยห้าสิบหกบาทถ้วน", name: "uint"},
		{input: uint8(99), expected: "เก้าสิบเก้าบาทถ้วน", name: "uint8"},
		{input: uint16(2500), expected: "สองพันห้าร้อยบาทถ้วน", name: "uint16"},
		{input: uint32(50000), expected: "ห้าหมื่นบาทถ้วน", name: "uint32"},
		{input: uint64(1000000), expected: "หนึ่งล้านบาทถ้วน", name: "uint64"},

		// Float inputs (formatted to 2 decimal places)
		{input: float32(123.45), expected: "หนึ่งร้อยยี่สิบสามบาทสี่สิบห้าสตางค์", name: "float32"},
		{input: float64(999.99), expected: "เก้าร้อยเก้าสิบเก้าบาทเก้าสิบเก้าสตางค์", name: "float64"},
		{input: float64(100.5), expected: "หนึ่งร้อยบาทห้าสิบสตางค์", name: "float64 with .5"},
		{input: float64(50), expected: "ห้าสิบบาทถ้วน", name: "float64 whole number"},

		// Edge cases
		{input: 0, expected: "ศูนย์บาทถ้วน", name: "zero int"},
		{input: float64(0.0), expected: "ศูนย์บาทถ้วน", name: "zero float"},
	}

	for _, test := range tests {
		result, err := Convert(test.input)
		if err != nil {
			t.Errorf("Convert(%v) [%s] returned error: %v", test.input, test.name, err)
			continue
		}
		if result != test.expected {
			t.Errorf("Convert(%v) [%s] = %s, expected %s", test.input, test.name, result, test.expected)
		}
	}
}

func TestConvertWithInvalidTypes(t *testing.T) {
	// Test unsupported types
	result, err := Convert([]int{1, 2, 3})
	if err == nil {
		t.Errorf("Convert with unsupported type should return error, got result: %s", result)
	}
	if result != "" {
		t.Errorf("Convert with unsupported type should return empty string, got %s", result)
	}

	result, err = Convert(map[string]int{"test": 1})
	if err == nil {
		t.Errorf("Convert with unsupported type should return error, got result: %s", result)
	}
	if result != "" {
		t.Errorf("Convert with unsupported type should return empty string, got %s", result)
	}
}

func TestConvertWithExceedingMaxValue(t *testing.T) {
	tests := []struct {
		input       string
		expectError bool
		description string
	}{
		// Valid values (should not error)
		{input: MaxSupportedValue, expectError: false, description: "exact max value"},
		{input: "9223372036854775807", expectError: false, description: "int64 max value"},
		{input: "1234567889999999999", expectError: false, description: "19 digits under int64 max"},

		// Invalid values (should error)
		{input: "9223372036854775808", expectError: true, description: "19 digits exceeds int64 max by 1"},
		{input: "12345678901234567890", expectError: true, description: "20 digits exceeds max"},
		{input: "18446744073709551615", expectError: true, description: "uint64 max exceeds int64 max"},
		{input: "100000000000000000000", expectError: true, description: "21 digits - exceeds max"},
		{input: "999999999999999999999", expectError: true, description: "21 digits - much larger"},
		{input: "123456789012345678901", expectError: true, description: "21 digits - way over max"},
		{input: "999999999999999999999999999", expectError: true, description: "27 digits - extremely large"},

		// Edge cases
		{input: "000100000000000000000000", expectError: true, description: "leading zeros but exceeds when trimmed"},
		{input: "0009223372036854775807", expectError: false, description: "leading zeros, valid when trimmed"},
	}

	for _, test := range tests {
		result, err := Convert(test.input)

		if test.expectError {
			if err == nil {
				t.Errorf("%s: Expected error for input %s, but got result: %s", test.description, test.input, result)
			}
			if result != "" {
				t.Errorf("%s: Expected empty result for invalid input, got: %s", test.description, result)
			}
		} else {
			if err != nil {
				t.Errorf("%s: Unexpected error for valid input %s: %v", test.description, test.input, err)
			}
		}
	}
}

func TestConvertWithOverflowHandling(t *testing.T) {
	// Disable warning logs for cleaner test output
	originalLogSetting := EnableWarningLogs
	originalOverflowSetting := AllowOverflow
	EnableWarningLogs = false
	defer func() {
		EnableWarningLogs = originalLogSetting
		AllowOverflow = originalOverflowSetting
	}()

	tests := []struct {
		input         string
		roundingMode  DecimalRoundingMode
		allowOverflow bool
		expected      string
		name          string
	}{
		// Test 0.995 case with different rounding modes
		{input: "100.995", roundingMode: RoundHalf, allowOverflow: false, expected: "หนึ่งร้อยบาทเก้าสิบเก้าสตางค์", name: "0.995 with RoundHalf (forced down)"},
		{input: "100.995", roundingMode: RoundUp, allowOverflow: false, expected: "หนึ่งร้อยบาทเก้าสิบเก้าสตางค์", name: "0.995 with RoundUp (capped)"},
		{input: "100.995", roundingMode: RoundUp, allowOverflow: true, expected: "หนึ่งร้อยเอ็ดบาทถ้วน", name: "0.995 with RoundUp and overflow (overflow to 101)"},
		{input: "100.995", roundingMode: RoundHalf, allowOverflow: true, expected: "หนึ่งร้อยเอ็ดบาทถ้วน", name: "0.995 with RoundHalf and overflow (overflow to 101)"},
		{input: "100.999", roundingMode: RoundUp, allowOverflow: true, expected: "หนึ่งร้อยเอ็ดบาทถ้วน", name: "0.999 with RoundUp and overflow (overflow to 101)"},

		// Test other edge cases
		{input: "50.996", roundingMode: RoundHalf, allowOverflow: false, expected: "ห้าสิบบาทเก้าสิบเก้าสตางค์", name: "0.996 with RoundHalf (forced down)"},
		{input: "50.996", roundingMode: RoundUp, allowOverflow: true, expected: "ห้าสิบเอ็ดบาทถ้วน", name: "0.996 with RoundUp and overflow (overflow to 51)"},
		{input: "50.996", roundingMode: RoundHalf, allowOverflow: true, expected: "ห้าสิบเอ็ดบาทถ้วน", name: "0.996 with RoundHalf and overflow (overflow to 51)"},

		// Test RoundHalf vs overflow behavior difference
		{input: "100.994", roundingMode: RoundHalf, allowOverflow: false, expected: "หนึ่งร้อยบาทเก้าสิบเก้าสตางค์", name: "0.994 with RoundHalf (no overflow needed)"},
		{input: "100.994", roundingMode: RoundHalf, allowOverflow: true, expected: "หนึ่งร้อยบาทเก้าสิบเก้าสตางค์", name: "0.994 with RoundHalf and overflow (no overflow needed)"},
		{input: "100.991", roundingMode: RoundUp, allowOverflow: false, expected: "หนึ่งร้อยบาทเก้าสิบเก้าสตางค์", name: "0.991 with RoundUp (normal)"},
	}

	for _, test := range tests {
		AllowOverflow = test.allowOverflow
		result, err := Convert(test.input, test.roundingMode)
		if err != nil {
			t.Errorf("%s: Convert(%s, %v) returned error: %v", test.name, test.input, test.roundingMode, err)
			continue
		}
		if result != test.expected {
			t.Errorf("%s: Convert(%s, %v) = %s, expected %s", test.name, test.input, test.roundingMode, result, test.expected)
		}
	}
}

func TestWarningLogControl(t *testing.T) {
	// Test that warning logs can be enabled/disabled
	originalLogSetting := EnableWarningLogs
	originalOverflowSetting := AllowOverflow
	defer func() {
		EnableWarningLogs = originalLogSetting
		AllowOverflow = originalOverflowSetting
	}()

	// Test SetWarningLogs function
	SetWarningLogs(false)
	if EnableWarningLogs != false {
		t.Errorf("SetWarningLogs(false) failed, EnableWarningLogs = %v", EnableWarningLogs)
	}

	SetWarningLogs(true)
	if EnableWarningLogs != true {
		t.Errorf("SetWarningLogs(true) failed, EnableWarningLogs = %v", EnableWarningLogs)
	}

	// Test SetAllowOverflow function
	SetAllowOverflow(false)
	if AllowOverflow != false {
		t.Errorf("SetAllowOverflow(false) failed, AllowOverflow = %v", AllowOverflow)
	}

	SetAllowOverflow(true)
	if AllowOverflow != true {
		t.Errorf("SetAllowOverflow(true) failed, AllowOverflow = %v", AllowOverflow)
	}

	// Test that conversion still works with logging disabled
	SetWarningLogs(false)
	SetAllowOverflow(false)
	result, err := Convert("100.995", RoundHalf)
	if err != nil {
		t.Errorf("Convert with logging disabled returned error: %v", err)
	}
	expected := "หนึ่งร้อยบาทเก้าสิบเก้าสตางค์"
	if result != expected {
		t.Errorf("Convert with logging disabled = %s, expected %s", result, expected)
	}
}

func TestProblematicLargeNumbers(t *testing.T) {
	// Test the specific large numbers mentioned in the issue
	testCases := []struct {
		input       any
		expectError bool
		name        string
	}{
		{int64(9_223_372_036_854_775_807), false, "int64 max value"},
		{MaxSupportedValue, false, "string max supported value"},
		{"1234567889999999999", false, "string 1234567889999999999"},
		{"1000000000000000000", false, "string 1000000000000000000"},
		{"9000000000000000000", false, "string 9000000000000000000"},
		{"100000000000000000000", true, "21 digits - should fail"},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result, err := Convert(test.input)

			if test.expectError {
				if err == nil {
					t.Errorf("Expected error for %s, but got result: %s", test.name, result)
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error for %s: %v", test.name, err)
				return
			}

			t.Logf("Input: %v", test.input)
			t.Logf("Result: %s", result)

			// Basic validation - should contain "บาท" and not be empty
			if result == "" {
				t.Errorf("Result should not be empty for %s", test.name)
			}
			if !contains(result, "บาท") {
				t.Errorf("Result should contain 'บาท' for %s, got: %s", test.name, result)
			}
		})
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || containsMiddle(s, substr)))
}

func containsMiddle(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func TestDebugLargeNumbers(t *testing.T) {
	// Test digit level handling for different positions
	testCases := []struct {
		input       any
		description string
	}{
		// Test basic digit levels within 6-digit groups
		{"1000", "thousands"},
		{"10000", "ten thousands"},
		{"100000", "hundred thousands"},
		{"1000000", "1 million"},
		{"10000000", "10 million (7 digits)"},
		{"100000000", "100 million (8 digits)"},
		{"1000000000", "1 billion (10 digits)"},
		{"10000000000", "10 billion (11 digits)"},
		{"100000000000", "100 billion (12 digits)"},
		{"1000000000000", "1 trillion (13 digits) - debug grouping"},
		{"1000000000000000", "1 quadrillion (16 digits) - debug grouping"},

		// Test the problematic large numbers
		{int64(9_223_372_036_854_775_807), "max int64"},
		{MaxSupportedValue, "max supported value"},
		{"1000000000000000000", "1 * 10^18"},
		{"9000000000000000000", "9 * 10^18"},
	}

	for _, tc := range testCases {
		result, err := Convert(tc.input)
		if err != nil {
			t.Errorf("Convert(%v) [%s] returned error: %v", tc.input, tc.description, err)
			continue
		}
		t.Logf("%s (%v) → %s", tc.description, tc.input, result)
	}
}
