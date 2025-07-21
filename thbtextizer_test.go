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
			input:    "1234567889999999999",
			expected: "หนึ่งล้านสองแสนสามหมื่นสี่พันห้าร้อยหกสิบเจ็ดล้านแปดแสนแปดหมื่นเก้าพันเก้าร้อยเก้าสิบเก้าล้านเก้าแสนเก้าหมื่นเก้าพันเก้าร้อยเก้าสิบเก้าบาทถ้วน",
		},
		{
			input:    "18446744073709551615",
			expected: "สิบแปดล้านสี่แสนสี่หมื่นหกพันเจ็ดร้อยสี่สิบสี่ล้านเจ็ดหมื่นสามพันเจ็ดร้อยเก้าล้านห้าแสนห้าหมื่นหนึ่งพันหกร้อยสิบห้าบาทถ้วน",
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
		{input: "18446744073709551615", expectError: false, description: "uint64 max value"},
		{input: "1234567889999999999", expectError: false, description: "19 digits under max"},
		{input: "12345678901234567890", expectError: false, description: "20 digits at max"},

		// Invalid values (should error)
		{input: "100000000000000000000", expectError: true, description: "21 digits - exceeds max"},
		{input: "999999999999999999999", expectError: true, description: "21 digits - much larger"},
		{input: "123456789012345678901", expectError: true, description: "21 digits - way over max"},
		{input: "999999999999999999999999999", expectError: true, description: "27 digits - extremely large"},

		// Edge cases
		{input: "000100000000000000000000", expectError: true, description: "leading zeros but exceeds when trimmed"},
		{input: "000099999999999999999999", expectError: false, description: "leading zeros, valid when trimmed"},
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
