# Thai Baht Textizer

A Go library for converting numeric amounts to Thai text representation following proper Thai language rules for currency formatting.

## Features

✅ **Multiple Input Types**: Support for `string`, `int`, `uint`, `float32`, `float64` and their variants  
✅ **Thai Language Rules**: Proper use of "เอ็ด" vs "หนึ่ง" based on position  
✅ **Configurable Rounding**: Multiple decimal rounding modes with overflow handling  
✅ **Large Numbers**: Support for millions with proper grouping  
✅ **Error Handling**: Returns descriptive errors for unsupported types  
✅ **Warning Logs**: Optional logging for satang capping edge cases  
✅ **Comprehensive Testing**: Full test coverage with edge cases  

## Installation

```bash
go get github.com/natt-v/thai-baht-textizer
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/natt-v/thai-baht-textizer"
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

## API Reference

### Main Function

```go
func Convert(amount any, roundingMode ...DecimalRoundingMode) (string, error)
```

**Parameters:**
- `amount`: Numeric value (string, int, uint, float32, float64, and their variants)
- `roundingMode`: Optional rounding mode (defaults to `RoundHalf`)

**Returns:**
- `string`: Thai text representation
- `error`: Error for unsupported types or invalid input

## Rounding Modes

```go
const (
    RoundHalf         // Round half (0.455 → 0.46, 0.454 → 0.45) - DEFAULT
    RoundDown         // Always truncate (0.456 → 0.45)
    RoundUp           // Always round up (0.451 → 0.46), caps at 99 satang
    RoundUpOverflow   // Round up with overflow (can increase baht amount)
    RoundHalfOverflow // Round half with overflow (can increase baht amount)
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

### Special Case: 99 Satang Cap Handling

When rounding would result in 100+ satang:

```go
// Standard modes cap at 99 satang with warning log
result, _ := thbtextizer.Convert("100.995", thbtextizer.RoundHalf)
// Output: "หนึ่งร้อยบาทเก้าสิบเก้าสตางค์" (100.99)
// Log: "Warning: 995 rounds to 100 satang, forced to round down..."

// Overflow modes properly overflow to next baht
result, _ := thbtextizer.Convert("100.995", thbtextizer.RoundUpOverflow)
// Output: "หนึ่งร้อยเอ็ดบาทถ้วน" (101.00)

result, _ = thbtextizer.Convert("100.995", thbtextizer.RoundHalfOverflow)
// Output: "หนึ่งร้อยเอ็ดบาทถ้วน" (101.00)
```

## Input Type Support

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

```go
// Millions with proper grouping
result, _ := thbtextizer.Convert("999999999.99")
// Output: "เก้าร้อยเก้าสิบเก้าล้านเก้าแสนเก้าหมื่นเก้าพันเก้าร้อยเก้าสิบเก้าบาทเก้าสิบเก้าสตางค์"

result, _ = thbtextizer.Convert("100000001.01") 
// Output: "หนึ่งร้อยล้านเอ็ดบาทหนึ่งสตางค์"
```

## Warning Log Control

By default, the library logs warnings when satang is capped at 99. You can control this behavior:

### Disable Warning Logs

```go
// Method 1: Using SetWarningLogs function
thbtextizer.SetWarningLogs(false)

// Method 2: Direct variable access
thbtextizer.EnableWarningLogs = false

// Now conversions won't print warnings
result, _ := thbtextizer.Convert("100.995") // No warning log
```

### Re-enable Warning Logs

```go
thbtextizer.SetWarningLogs(true)
// or
thbtextizer.EnableWarningLogs = true
```

### Example Warning Log

```
2025/07/21 18:39:12 Warning: 995 rounds to 100 satang, forced to round down to 99 satang to maintain currency format. Consider using RoundHalfOverflow or RoundUpOverflow mode.
```

## Error Handling

```go
result, err := thbtextizer.Convert([]int{1, 2, 3})
if err != nil {
    fmt.Printf("Error: %v\n", err)
    // Error: unsupported type: only string, int, uint, float32, float64 and their variants are supported
}
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

## Examples

### Basic Usage

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/natt-v/thai-baht-textizer"
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

### Advanced Usage with Rounding Modes

```go
package main

import (
    "fmt"
    
    "github.com/natt-v/thai-baht-textizer"
)

func main() {
    amount := "123.456"
    
    modes := []thbtextizer.DecimalRoundingMode{
        thbtextizer.RoundHalf,
        thbtextizer.RoundDown,
        thbtextizer.RoundUp,
        thbtextizer.RoundUpOverflow,
        thbtextizer.RoundHalfOverflow,
    }
    
    modeNames := []string{"RoundHalf", "RoundDown", "RoundUp", "RoundUpOverflow", "RoundHalfOverflow"}
    
    for i, mode := range modes {
        result, _ := thbtextizer.Convert(amount, mode)
        fmt.Printf("%s: %s\n", modeNames[i], result)
    }
    
    // Disable warnings for cleaner output
    thbtextizer.SetWarningLogs(false)
    
    // Test edge case
    result, _ := thbtextizer.Convert("100.995", thbtextizer.RoundHalfOverflow)
    fmt.Printf("Edge case (100.995 with RoundHalfOverflow): %s\n", result)
}
```

## Performance

The library is optimized for typical currency amounts (up to 999,999,999 baht) with minimal memory allocations and no global state caching.

## License

MIT License - see LICENSE file for details.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality  
4. Ensure all tests pass: `go test -v`
5. Submit a pull request

## Changelog

### v1.0.0
- Initial release with full Thai baht text conversion
- Support for multiple input types
- Configurable rounding modes with overflow handling
- Warning log control
- Comprehensive test coverage