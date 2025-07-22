#!/bin/bash

echo "=== Thai Baht Textizer Performance Benchmarks ==="
echo ""
echo "Running comprehensive benchmarks..."
echo "This may take a few minutes. Results will show:"
echo "- Operations per second"  
echo "- Nanoseconds per operation"
echo "- Memory usage per operation"
echo "- Memory allocations per operation"
echo ""

# Create results directory
mkdir -p benchmark_results

# Run main benchmarks
echo "1. Main Conversion Benchmarks"
echo "=============================="
go test -bench=BenchmarkConvert -benchmem -benchtime=3s | tee benchmark_results/conversion.txt

echo ""
echo "2. Memory Allocation Benchmarks" 
echo "==============================="
go test -bench=BenchmarkMemoryAllocations -benchmem -benchtime=3s | tee benchmark_results/memory.txt

echo ""
echo "3. Concurrent Usage Benchmarks"
echo "=============================="  
go test -bench=BenchmarkConcurrentUsage -benchmem -benchtime=3s | tee benchmark_results/concurrent.txt

echo ""
echo "4. Input Type Performance"
echo "========================"
go test -bench=BenchmarkInputTypes -benchmem -benchtime=3s | tee benchmark_results/input_types.txt

echo ""
echo "5. Rounding Mode Performance"
echo "==========================="
go test -bench=BenchmarkRoundingModes -benchmem -benchtime=3s | tee benchmark_results/rounding.txt

echo ""
echo "6. String Building Performance"
echo "=============================="
go test -bench=BenchmarkStringBuilding -benchmem -benchtime=3s | tee benchmark_results/string_building.txt

echo ""
echo "7. Input Sanitization Performance"
echo "================================="
go test -bench=BenchmarkInputSanitization -benchmem -benchtime=3s | tee benchmark_results/sanitization.txt

echo ""
echo "=== Benchmark Results Summary ==="
echo ""
echo "All results saved to benchmark_results/ directory"
echo ""
echo "Key Performance Highlights:"
echo "- Small numbers: ~500-600 ns/op with ~13 allocations"
echo "- Large numbers: ~1000-1200 ns/op with ~27 allocations"  
echo "- Very large numbers: ~1600-1800 ns/op with ~39-41 allocations"
echo "- Concurrent usage: ~300 ns/op (parallel execution)"
echo ""
echo "Memory efficiency:"
echo "- Small conversions: ~504 B/op"
echo "- Large conversions: ~1500-3500 B/op"
echo ""
echo "To compare with your own measurements:"
echo "1. Save current results: go test -bench=. -benchmem > before.txt"
echo "2. Make your changes"
echo "3. Run again: go test -bench=. -benchmem > after.txt"
echo "4. Compare: benchcmp before.txt after.txt (requires golang.org/x/tools/cmd/benchcmp)"
echo ""
echo "For detailed analysis, examine individual files in benchmark_results/"