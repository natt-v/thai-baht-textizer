# Thai Baht Textizer

A Go library for converting numeric amounts to Thai text representation following proper Thai language rules for currency formatting.

## Features

✅ **Multiple Input Types**: Support for `string`, `int`, `uint`, `float32`, `float64` and their variants  
✅ **Thai Language Rules**: Proper use of "เอ็ด" vs "หนึ่ง" based on position  
✅ **Configurable Rounding**: Three decimal rounding modes (RoundHalf, RoundDown, RoundUp)  
✅ **Overflow Control**: Optional overflow behavior for precise financial calculations  
✅ **Large Numbers**: Support for numbers up to 9,223,372,036,854,775,807 (19 digits) with proper million grouping  
✅ **Input Validation**: Validates maximum supported values and input types  
✅ **Error Handling**: Returns descriptive errors for unsupported types and exceeded limits  
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

The library validates input and returns descriptive errors for various scenarios:

### Unsupported Types
```go
result, err := thbtextizer.Convert([]int{1, 2, 3})
if err != nil {
    fmt.Printf("Error: %v\n", err)
    // Error: unsupported type: only string, int, uint, float32, float64 and their variants are supported
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
    
    "github.com/natt-v/thai-baht-textizer"
)

func main() {
    amount := "123.456"
    
    // Test different rounding modes
    modes := []thbtextizer.DecimalRoundingMode{
        thbtextizer.RoundHalf,
        thbtextizer.RoundDown,
        thbtextizer.RoundUp,
    }
    
    modeNames := []string{"RoundHalf", "RoundDown", "RoundUp"}
    
    for i, mode := range modes {
        result, _ := thbtextizer.Convert(amount, mode)
        fmt.Printf("%s: %s\n", modeNames[i], result)
    }
    
    // Test overflow behavior
    thbtextizer.SetWarningLogs(false) // Disable warnings for cleaner output
    
    // Test with overflow disabled (default)
    thbtextizer.SetAllowOverflow(false)
    result1, _ := thbtextizer.Convert("100.995", thbtextizer.RoundHalf)
    fmt.Printf("No overflow: %s\n", result1)
    
    // Test with overflow enabled
    thbtextizer.SetAllowOverflow(true)
    result2, _ := thbtextizer.Convert("100.995", thbtextizer.RoundHalf)
    fmt.Printf("With overflow: %s\n", result2)
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

See [CHANGELOG.md](CHANGELOG.md) for detailed release history.

### Recent Releases

**v1.1.1** - Critical bug fixes for large number conversion logic  
**v1.1.0** - Simplified API with improved overflow control  
**v1.0.0** - Initial release with full Thai baht text conversion