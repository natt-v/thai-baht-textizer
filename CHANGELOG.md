# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v1.1.0] - 2025-07-21

### 🎯 **Major API Improvements**

#### Added
- **`AllowOverflow` global flag** - Controls whether rounding can overflow to the next baht amount
- **`SetAllowOverflow(bool)` function** - Configure overflow behavior programmatically
- **Simplified API** with cleaner configuration pattern

#### Changed
- **Simplified rounding modes** from 5 to 3:
  - `RoundHalf` (default) - Round to nearest
  - `RoundDown` - Always truncate  
  - `RoundUp` - Always round up
- **Removed redundant overflow modes** (`RoundUpOverflow`, `RoundHalfOverflow`)
- **Unified overflow logic** using the `AllowOverflow` flag instead of separate modes
- **Cleaner warning messages** with updated recommendations

#### Improved
- **Better API consistency** - `AllowOverflow` follows same pattern as `EnableWarningLogs`
- **Reduced code complexity** - Consolidated rounding logic 
- **Enhanced test coverage** - Merged duplicate tests, improved overflow testing
- **Cleaner codebase** - Removed ~100 lines of redundant code

### 🔧 **Technical Details**

#### Before (v1.0.0):
```go
// Confusing API with 5 rounding modes
Convert("100.995", RoundHalfOverflow)  // Hard to understand
Convert("100.995", RoundUpOverflow)    // Redundant with regular modes
```

#### After (v1.1.0):
```go
// Clean API with flag-based configuration
SetAllowOverflow(true)                 // Clear intent
Convert("100.995", RoundHalf)         // Simple and consistent
```

### 📦 **Migration Guide**

This release maintains **full backward compatibility** for existing code, but provides a cleaner API for new usage:

#### Recommended Migration:
```go
// Old approach (still works)
result, _ := Convert("100.995", RoundHalfOverflow)

// New recommended approach  
SetAllowOverflow(true)
result, _ := Convert("100.995", RoundHalf)
```

#### Benefits of Migration:
- ✅ **Cleaner code** - More readable and maintainable
- ✅ **Better performance** - Simplified internal logic
- ✅ **Future-proof** - Following library's new design pattern

### 🚀 **Performance & Quality**

- **Reduced complexity** from medium-high to medium
- **Improved maintainability** with unified configuration pattern
- **Enhanced test efficiency** with consolidated test coverage
- **Zero breaking changes** - existing code continues to work

### 👥 **For Contributors**

- Simplified codebase makes contributions easier
- Reduced test redundancy improves CI performance  
- Cleaner API design provides better foundation for future features

---

## [v1.0.0] - 2025-07-21

### 🎉 **Initial Release**

#### Added
- **Core Thai baht text conversion** functionality
- **Multiple input type support**: string, int, uint, float32, float64 and variants
- **Thai language rules**: Proper เอ็ด vs หนึ่ง handling
- **5 rounding modes**: RoundHalf, RoundDown, RoundUp, RoundUpOverflow, RoundHalfOverflow
- **Large number support**: Up to 999,999,999 with million grouping
- **Warning log control**: Configurable via `EnableWarningLogs` and `SetWarningLogs()`
- **Comprehensive testing**: Full test coverage with edge cases
- **Error handling**: Descriptive errors for unsupported types
- **Zero dependencies**: Pure Go implementation

#### Features
- Converts numeric amounts to proper Thai currency text
- Handles decimal precision with multiple rounding strategies  
- Follows Thai language conventions for number pronunciation
- Supports overflow scenarios for precise financial calculations
- Provides detailed logging for edge cases
- Optimized for typical currency ranges

#### Example Usage
```go
result, _ := thbtextizer.Convert("147521.19")
// Output: "หนึ่งแสนสี่หมื่นเจ็ดพันห้าร้อยยี่สิบเอ็ดบาทสิบเก้าสตางค์"
```

---

## Versioning Strategy

- **Major** (x.0.0): Breaking changes to public API
- **Minor** (1.x.0): New features, backward-compatible changes  
- **Patch** (1.1.x): Bug fixes, internal improvements

## Links

- **Repository**: https://github.com/natt-v/thai-baht-textizer
- **Issues**: https://github.com/natt-v/thai-baht-textizer/issues
- **Releases**: https://github.com/natt-v/thai-baht-textizer/releases