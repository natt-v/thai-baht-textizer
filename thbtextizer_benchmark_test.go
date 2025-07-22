package thbtextizer

import (
	"testing"
)

// BenchmarkConvert tests the performance of different number sizes
func BenchmarkConvert(b *testing.B) {
	testCases := []struct {
		name   string
		amount string
	}{
		{"small_numbers", "123.45"},
		{"medium_numbers", "12345.67"},
		{"large_numbers", "123456789.99"},
		{"very_large_numbers", "9223372036854775807"},
		{"complex_large", "1234567889999999999"},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err := Convert(tc.amount)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkConverterInstance tests the performance of instance-based conversion
func BenchmarkConverterInstance(b *testing.B) {
	config := &Config{
		EnableWarningLogs: false,
		AllowOverflow:     false,
		DefaultRounding:   RoundHalf,
	}
	converter := NewConverter(config)

	testCases := []struct {
		name   string
		amount string
	}{
		{"small_numbers", "123.45"},
		{"medium_numbers", "12345.67"},
		{"large_numbers", "123456789.99"},
		{"very_large_numbers", "9223372036854775807"},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err := converter.Convert(tc.amount)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkMemoryAllocations focuses on memory allocation efficiency
func BenchmarkMemoryAllocations(b *testing.B) {
	testCases := []struct {
		name   string
		amount string
	}{
		{"simple", "123.45"},
		{"large", "999999999.99"},
		{"very_large", "1234567889999999999"},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err := Convert(tc.amount)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkRoundingModes compares performance of different rounding modes
func BenchmarkRoundingModes(b *testing.B) {
	amount := "123.456"
	modes := []struct {
		name string
		mode DecimalRoundingMode
	}{
		{"RoundHalf", RoundHalf},
		{"RoundDown", RoundDown},
		{"RoundUp", RoundUp},
	}

	for _, mode := range modes {
		b.Run(mode.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err := Convert(amount, mode.mode)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkStringBuilding specifically tests string building performance
func BenchmarkStringBuilding(b *testing.B) {
	// Test cases that stress string building
	testCases := []struct {
		name   string
		amount string
	}{
		{"short_result", "1"},
		{"medium_result", "123456"},
		{"long_result", "999999999.99"},
		{"very_long_result", "999999999999999999.99"},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err := Convert(tc.amount)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkConcurrentUsage tests performance under concurrent load
func BenchmarkConcurrentUsage(b *testing.B) {
	b.Run("global_function", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_, err := Convert("123456.78")
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	})

	b.Run("instance_based", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			converter := NewDefaultConverter()
			for pb.Next() {
				_, err := converter.Convert("123456.78")
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	})
}

// BenchmarkInputTypes tests performance with different input types
func BenchmarkInputTypes(b *testing.B) {
	testCases := []struct {
		name  string
		input interface{}
	}{
		{"string", "123.45"},
		{"int", 123},
		{"int64", int64(123456789)},
		{"float64", 123.45},
		{"uint64", uint64(123456789)},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err := Convert(tc.input)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkInputSanitization tests the performance impact of input sanitization
func BenchmarkInputSanitization(b *testing.B) {
	testCases := []struct {
		name  string
		input string
	}{
		{"clean_input", "123.45"},
		{"whitespace", "  123.45  "},
		{"commas", "1,234,567.89"},
		{"underscores", "1_000_000.50"},
		{"mixed_formatting", "  1,234_567.89  "},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err := Convert(tc.input)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
