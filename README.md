# Thai Baht Textizer

A high-performance Go library for converting numeric amounts to Thai text representation following proper Thai language rules for currency formatting.

## Features

### 🚀 **v1.2.0 New Features**
✅ **Thread-Safe Configuration**: Instance-based converters with isolated settings  
✅ **Enhanced Error Handling**: Specific error codes with helpful hints  
✅ **Advanced Input Sanitization**: Auto-correction and robust validation  
✅ **Performance Optimized**: 20-40% faster with `strings.Builder`  

### 🔧 **Core Features**
✅ **Multiple Input Types**: Support for `string`, `int`, `uint`, `float32`, `float64` and their variants  
✅ **Thai Language Rules**: Proper use of "เอ็ด" vs "หนึ่ง" based on position  
✅ **Configurable Rounding**: Three decimal rounding modes (RoundHalf, RoundDown, RoundUp)  
✅ **Overflow Control**: Optional overflow behavior for precise financial calculations  
✅ **Large Numbers**: Support for numbers up to 9,223,372,036,854,775,807 (19 digits) with proper million grouping  
✅ **Input Validation**: Validates maximum supported values and input types  
✅ **Warning Logs**: Optional logging for satang capping edge cases  
✅ **Comprehensive Testing**: 168+ tests with full edge case coverage  

## Installation

```bash
go get github.com/natt-v/thai-baht-textizer
```

## Quick Start

### Basic Usage
```go
package main

import (
    "fmt"
    "log"
    
    thbtextizer "github.com/natt-v/thai-baht-textizer"
)

func main() {
    // Basic usage
    result, err := thbtextizer.Convert(123.45)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result) 
    // Output: "หนึ่งร้อยยี่สิบสามบาทสี่สิบห้าสตางค์"
}
```

### Thread-Safe Usage (v1.2.0+)
```go
package main

import (
    "fmt"

    thbtextizer "github.com/natt-v/thai-baht-textizer"
)

func main() {
    // Create thread-safe converter with custom config
    config := &thbtextizer.Config{
        EnableWarningLogs: false,
        AllowOverflow:     true,
        DefaultRounding:   thbtextizer.RoundUp,
    }
    converter := thbtextizer.NewConverter(config)
    
    // Use instance-based conversion
    result, _ := converter.Convert("100.994")
    fmt.Println(result)
    // Output: "หนึ่งร้อยเอ็ดบาทถ้วน" (rounds up with overflow)
}
```

## API Reference

### Global Functions (Backward Compatible)

```go
func Convert(amount any, roundingMode ...DecimalRoundingMode) (string, error)
```

**Parameters:**
- `amount`: Numeric value (string, int, uint, float32, float64, and their variants)
- `roundingMode`: Optional rounding mode (defaults to `RoundHalf`)

**Returns:**
- `string`: Thai text representation
- `error`: Error for unsupported types or invalid input

### Thread-Safe API (v1.2.0+)

```go
// Configuration
type Config struct {
    EnableWarningLogs bool
    AllowOverflow     bool
    DefaultRounding   DecimalRoundingMode
}

func DefaultConfig() *Config
func NewConverter(config *Config) *Converter
func NewDefaultConverter() *Converter

// Instance-based conversion
func (c *Converter) Convert(amount any, roundingMode ...DecimalRoundingMode) (string, error)
```

### Enhanced Error Handling (v1.2.0+)

```go
type ErrorCode int
type ConversionError struct {
    Code    ErrorCode
    Message string
    Input   string
    Hint    string
}

const (
    ErrorCodeUnsupportedType ErrorCode = iota
    ErrorCodeExceedsMaxValue
    ErrorCodeInvalidInput
    ErrorCodeParseError
)
```

## Rounding Modes

```go
const (
    RoundHalf DecimalRoundingMode = iota // Round to nearest (default)
    RoundDown                            // Always truncate
    RoundUp                              // Always round up
)
```

### Rounding Mode Examples

```go
// Default RoundHalf behavior
result, _ := thbtextizer.Convert("123.456")
// Output: "หนึ่งร้อยยี่สิบสามบาทสี่สิบหกสตางค์" (123.46)

// Explicit rounding modes
result, _ = thbtextizer.Convert("123.456", thbtextizer.RoundDown)
// Output: "หนึ่งร้อยยี่สิบสามบาทสี่สิบห้าสตางค์" (123.45)

result, _ = thbtextizer.Convert("123.451", thbtextizer.RoundUp)
// Output: "หนึ่งร้อยยี่สิบสามบาทสี่สิบหกสตางค์" (123.46)
```

## Configuration

### Thread-Safe Configuration (v1.2.0+)

For applications requiring thread safety or instance-based configuration:

```go
// Create converter with custom configuration
config := &thbtextizer.Config{
    EnableWarningLogs: false,
    AllowOverflow:     true,
    DefaultRounding:   thbtextizer.RoundUp,
}
converter := thbtextizer.NewConverter(config)

// Use instance-based conversion (thread-safe)
result, _ := converter.Convert("100.994")
// Output: "หนึ่งร้อยเอ็ดบาทถ้วน" (rounds up with overflow)

// Create converter with default settings
defaultConverter := thbtextizer.NewDefaultConverter()
result, _ = defaultConverter.Convert("123.45")
```

### Global Configuration (Legacy)

### Overflow Control

Control whether rounding can overflow to the next baht amount:

```go
// Default: cap satang at 99
thbtextizer.SetAllowOverflow(false) // Default behavior
result, _ := thbtextizer.Convert("100.995", thbtextizer.RoundHalf)
// Output: "หนึ่งร้อยบาทเก้าสิบเก้าสตางค์" (100.99)
// Log: "Warning: 995 rounds to 100 satang, forced to round down..."

// Enable overflow to next baht
thbtextizer.SetAllowOverflow(true)
result, _ = thbtextizer.Convert("100.995", thbtextizer.RoundHalf)
// Output: "หนึ่งร้อยเอ็ดบาทถ้วน" (101.00)
```

### Warning Control

```go
// Disable warning logs for cleaner output
thbtextizer.SetWarningLogs(false)

// Re-enable warnings (default)
thbtextizer.SetWarningLogs(true)
```

## Input Handling

### Advanced Input Sanitization (v1.2.0+)

The library automatically cleans and validates input with smart correction:

```go
// Automatic whitespace and formatting cleanup
result, _ := thbtextizer.Convert("  1,234.56  ")  // Trims and handles commas
result, _ = thbtextizer.Convert("1_000_000.25")   // Removes underscores
result, _ = thbtextizer.Convert("\t987.65\t")     // Handles tabs

// Smart decimal correction
result, _ = thbtextizer.Convert(".45")            // "0.45" - adds leading zero
result, _ = thbtextizer.Convert("123.")           // "123.0" - adds trailing zero

// Sign handling (for future negative support)
result, _ = thbtextizer.Convert("+987.65")        // Removes positive sign
result, _ = thbtextizer.Convert("-123.45")        // Removes negative sign (abs value)

// Enhanced validation with specific errors
_, err := thbtextizer.Convert("12.34.56")         // Multiple decimal points
_, err = thbtextizer.Convert("abc123")            // Invalid characters
_, err = thbtextizer.Convert("")                  // Empty input
```

### Input Type Support

The library accepts various numeric types:

```go
// String inputs
result, _ := thbtextizer.Convert("123.45")

// Integer types
result, _ = thbtextizer.Convert(123)
result, _ = thbtextizer.Convert(int64(1000000))
result, _ = thbtextizer.Convert(uint32(50000))

// Float types  
result, _ = thbtextizer.Convert(float64(999.99))
result, _ = thbtextizer.Convert(float32(50.5))

// Unsupported types return error
result, err := thbtextizer.Convert([]int{1, 2, 3})
// err: "unsupported type: only string, int, uint, float32, float64 and their variants are supported"
```

## Thai Language Rules

### เอ็ด vs หนึ่ง Rule

The library correctly applies Thai language rules for "1":

```go
// หนึ่ง: Used for single digit 1 or non-ones positions
result, _ := thbtextizer.Convert("1")
// Output: "หนึ่งบาทถ้วน"

result, _ = thbtextizer.Convert("100.01") 
// Output: "หนึ่งร้อยบาทหนึ่งสตางค์"

// เอ็ด: Used when 1 is at the end of multi-digit numbers
result, _ = thbtextizer.Convert("11")
// Output: "สิบเอ็ดบาทถ้วน"

result, _ = thbtextizer.Convert("101")
// Output: "หนึ่งร้อยเอ็ดบาทถ้วน"

result, _ = thbtextizer.Convert("100.11")
// Output: "หนึ่งร้อยบาทสิบเอ็ดสตางค์"
```

### Special Cases

```go
// Zero
result, _ := thbtextizer.Convert("0")
// Output: "ศูนย์บาทถ้วน"

// Twenty (uses ยี่ instead of สองสิบ)
result, _ = thbtextizer.Convert("20")
// Output: "ยี่สิบบาทถ้วน"

// Whole numbers
result, _ = thbtextizer.Convert("100.00")
// Output: "หนึ่งร้อยบาทถ้วน"
```

## Large Number Support

The library correctly handles very large numbers with proper Thai million grouping:

```go
// Basic millions
result, _ := thbtextizer.Convert("1000000")
// Output: "หนึ่งล้านบาทถ้วน"

// Complex large numbers with mixed digits
result, _ = thbtextizer.Convert("1234567889999999999")
// Output: "หนึ่งล้านสองแสนสามหมื่นสี่พันห้าร้อยหกสิบเจ็ดล้านแปดแสนแปดหมื่นเก้าพันเก้าร้อยเก้าสิบเก้าล้านเก้าแสนเก้าหมื่นเก้าพันเก้าร้อยเก้าสิบเก้าบาทถ้วน"

// "Telescoping zeros" pattern - multiple millions
result, _ = thbtextizer.Convert("1000000000000000000")
// Output: "หนึ่งล้านล้านล้านบาทถ้วน"

// Maximum supported value (int64 max)
result, _ = thbtextizer.Convert("9223372036854775807")
// Output: "เก้าล้านสองแสนสองหมื่นสามพันสามร้อยเจ็ดสิบสองล้านสามหมื่นหกพันแปดร้อยห้าสิบสี่ล้านเจ็ดแสนเจ็ดหมื่นห้าพันแปดร้อยเจ็ดบาทถ้วน"
```

### Thai Million Grouping Rules

The library implements sophisticated logic for Thai number grouping:

- **Single non-zero group**: Numbers like `1,000,000,000,000` get multiple "ล้าน" (หนึ่งล้านล้าน)
- **Multiple non-zero groups**: Numbers like `1,234,567,889` get single "ล้าน" per group (หนึ่งล้านสองแสนสามหมื่น...)


## Error Handling

### Enhanced Error Handling (v1.2.0+)

The library provides detailed error information with specific error codes and helpful hints:

```go
// Type-specific error handling
_, err := thbtextizer.Convert([]int{1, 2, 3})
if convErr, ok := err.(*thbtextizer.ConversionError); ok {
    fmt.Printf("Error Code: %d\n", convErr.Code)        // 0 (ErrorCodeUnsupportedType)
    fmt.Printf("Message: %s\n", convErr.Message)        // Descriptive error
    fmt.Printf("Input: %s\n", convErr.Input)            // "[]int"
    fmt.Printf("Hint: %s\n", convErr.Hint)              // Actionable suggestion
}

// Input validation errors
_, err = thbtextizer.Convert("12.34.56")
if convErr, ok := err.(*thbtextizer.ConversionError); ok {
    switch convErr.Code {
    case thbtextizer.ErrorCodeInvalidInput:
        fmt.Println("Invalid input format")
    case thbtextizer.ErrorCodeExceedsMaxValue:
        fmt.Println("Number too large")
    }
}
```

### Legacy Error Handling
```go
result, err := thbtextizer.Convert([]int{1, 2, 3})
if err != nil {
    fmt.Printf("Error: %v\n", err)
    // Error: unsupported type: only string, int, uint, float32, float64 and their variants are supported. Hint: convert your input to one of the supported types
}
```

### Maximum Value Exceeded
```go
result, err := thbtextizer.Convert("100000000000000000000") // 21 digits
if err != nil {
    fmt.Printf("Error: %v\n", err)
    // Error: input number exceeds maximum supported value of 9223372036854775807 (got 21 digits, max 19 digits)
}

// Maximum supported value (19 digits - int64 max)
result, err = thbtextizer.Convert("9223372036854775807")
// This works fine
```

## Testing

### Run All Tests

```bash
go test -v
```

### Run Specific Tests

```bash
# Test basic conversion
go test -v -run TestConvert

# Test rounding modes
go test -v -run TestConvertWithRounding  

# Test overflow handling
go test -v -run TestConvertWithOverflowHandling

# Test different input types
go test -v -run TestConvertWithNumericTypes

# Test warning log control
go test -v -run TestWarningLogControl
```

### Test Coverage

```bash
go test -cover
```

### Benchmarks

```bash
# Run performance benchmarks
go test -bench=. -benchmem

# Specific benchmark categories
go test -bench=BenchmarkConvert -benchmem         # Conversion performance
go test -bench=BenchmarkMemoryAllocations -benchmem  # Memory efficiency  
go test -bench=BenchmarkConcurrentUsage -benchmem    # Concurrent performance
go test -bench=BenchmarkInputTypes -benchmem         # Input type performance
go test -bench=BenchmarkRoundingModes -benchmem      # Rounding mode performance
```

## Examples

### Basic Usage

```go
package main

import (
    "fmt"
    "log"

    thbtextizer "github.com/natt-v/thai-baht-textizer"
)

func main() {
    examples := []interface{}{
        147521.19,
        1000000.25,
        100.01,
        "50.05",
        int64(999999999),
    }
    
    for _, amount := range examples {
        result, err := thbtextizer.Convert(amount)
        if err != nil {
            log.Printf("Error converting %v: %v", amount, err)
            continue
        }
        fmt.Printf("%.2v → %s\n", amount, result)
    }
}
```

### Advanced Usage with Configuration

```go
package main

import (
    "fmt"
    "sync"
    
    thbtextizer "github.com/natt-v/thai-baht-textizer"
)

func main() {
    // Traditional approach (still works)
    legacyExample()
    
    // New thread-safe approach (v1.2.0+)
    threadSafeExample()
    
    // Concurrent usage example
    concurrentExample()
}

func legacyExample() {
    amount := "123.456"
    
    // Test different rounding modes
    modes := []thbtextizer.DecimalRoundingMode{
        thbtextizer.RoundHalf, thbtextizer.RoundDown, thbtextizer.RoundUp,
    }
    modeNames := []string{"RoundHalf", "RoundDown", "RoundUp"}
    
    for i, mode := range modes {
        result, _ := thbtextizer.Convert(amount, mode)
        fmt.Printf("%s: %s\n", modeNames[i], result)
    }
}

func threadSafeExample() {
    // Create converter with specific configuration
    config := &thbtextizer.Config{
        EnableWarningLogs: false,
        AllowOverflow:     true,
        DefaultRounding:   thbtextizer.RoundUp,
    }
    converter := thbtextizer.NewConverter(config)
    
    // Use instance-based conversion
    result1, _ := converter.Convert("100.995") // Uses RoundUp + overflow
    fmt.Printf("Converter result: %s\n", result1)
    
    // Global function remains unaffected
    result2, _ := thbtextizer.Convert("100.995", thbtextizer.RoundDown)
    fmt.Printf("Global result: %s\n", result2)
}

func concurrentExample() {
    var wg sync.WaitGroup
    
    // Create converters with different configurations
    configs := []*thbtextizer.Config{
        {DefaultRounding: thbtextizer.RoundDown, AllowOverflow: false},
        {DefaultRounding: thbtextizer.RoundUp, AllowOverflow: true},
        {DefaultRounding: thbtextizer.RoundHalf, AllowOverflow: false},
    }
    
    for i, config := range configs {
        wg.Add(1)
        go func(id int, cfg *thbtextizer.Config) {
            defer wg.Done()
            converter := thbtextizer.NewConverter(cfg)
            result, _ := converter.Convert("100.995")
            fmt.Printf("Goroutine %d: %s\n", id, result)
        }(i, config)
    }
    
    wg.Wait()
}
```

## Performance

### v1.2.0 Optimizations

The library has been significantly optimized for production use:

- **20-40% faster string building** using `strings.Builder` instead of concatenation
- **15-30% reduction in memory allocations** with pre-allocated slices
- **Zero overhead thread safety** for concurrent usage with instance-based converters
- **Reduced garbage collection pressure** for high-throughput scenarios

### Benchmarks

Real performance benchmarks on Apple M1 (your results may vary):

```bash
# Main conversion benchmarks
BenchmarkConvert/small_numbers-8       2435452    601.3 ns/op    504 B/op    13 allocs/op
BenchmarkConvert/medium_numbers-8      1536578    721.1 ns/op    888 B/op    16 allocs/op  
BenchmarkConvert/large_numbers-8       1000000   1249 ns/op    1520 B/op    27 allocs/op
BenchmarkConvert/very_large_numbers-8   745856   1603 ns/op    3152 B/op    39 allocs/op

# Memory allocation efficiency
BenchmarkMemoryAllocations/simple-8    2447869    482.1 ns/op    504 B/op    13 allocs/op
BenchmarkMemoryAllocations/large-8     1000000   1068 ns/op    1560 B/op    27 allocs/op

# Concurrent usage (parallel execution)
BenchmarkConcurrentUsage/global_function-8   3838452   293.4 ns/op   928 B/op   16 allocs/op
BenchmarkConcurrentUsage/instance_based-8    3837180   313.7 ns/op   928 B/op   16 allocs/op
```

### Running Benchmarks Yourself

```bash
# Quick start - run comprehensive benchmark suite
./run_benchmarks.sh

# Or run individual benchmark categories manually:

# Run all benchmarks
go test -bench=. -benchmem

# Run specific benchmark categories
go test -bench=BenchmarkConvert -benchmem           # Conversion performance
go test -bench=BenchmarkMemoryAllocations -benchmem # Memory efficiency
go test -bench=BenchmarkConcurrentUsage -benchmem   # Concurrent performance
go test -bench=BenchmarkInputTypes -benchmem        # Input type performance
go test -bench=BenchmarkRoundingModes -benchmem     # Rounding mode performance

# Run benchmarks multiple times for stability
go test -bench=BenchmarkConvert -benchmem -count=5

# Run longer benchmarks for more accurate results
go test -bench=BenchmarkConvert -benchmem -benchtime=5s

# Save benchmark results for comparison
go test -bench=. -benchmem > benchmark_results.txt
```

### Recommendations

- **High-throughput applications**: Use instance-based converters (`NewConverter()`)
- **Concurrent usage**: Create separate converter instances per goroutine
- **Memory-sensitive environments**: Benefits from reduced allocations in v1.2.0

## License

MIT License - see LICENSE file for details.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality  
4. Ensure all tests pass: `go test -v`
5. Submit a pull request

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for detailed release history.

### Recent Releases

**v1.2.0** - Major performance & API enhancements with thread-safe configuration  
**v1.1.1** - Critical bug fixes for large number conversion logic  
**v1.1.0** - Simplified API with improved overflow control  
**v1.0.0** - Initial release with full Thai baht text conversion