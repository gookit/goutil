# gookit/goutil - Go Utility Library

gookit/goutil is a comprehensive Go utility library providing 800+ functions across multiple packages for common programming tasks including string manipulation, array/slice operations, filesystem utilities, system utilities, and more.

**Always reference these instructions first and fallback to search or bash commands only when you encounter unexpected information that does not match the info here.**

## Working Effectively

### Bootstrap and Build
- `go mod tidy` - Download dependencies (first run ~5 seconds, subsequent ~0.03 seconds)
- `go build ./...` - Build all packages (~0.4 seconds, very fast)
- `go test ./...` - Run complete test suite (~4 seconds with cache, ~24 seconds without. NEVER CANCEL. Set timeout to 60+ minutes)
- `make csfix` - Format all code using go fmt (~0.16 seconds)
- `make csdiff` - Check code formatting issues
- `make readme` - Generate README documentation (~0.17 seconds)

### Linting and Quality
- `go fmt ./...` - Format code (essential before committing)
- `staticcheck ./...` - Run static analysis (has some acceptable warnings in internal packages)
- **ALWAYS run `go fmt ./...` before completing work** - CI will fail if code is not formatted

### Testing and Validation
- Tests complete in ~24 seconds. NEVER CANCEL test runs - use timeout of 60+ minutes
- Use `github.com/gookit/goutil/testutil/assert` for assertions in tests
- For multiple test cases in one function, use `t.Run()` pattern
- **VALIDATION REQUIREMENT**: Always test changes with a comprehensive validation scenario

## Key Project Structure

### Main Utility Packages
- **`arrutil`** - Array/slice utilities (check, convert, formatting, collections)
- **`strutil`** - String utilities (bytes, check, convert, encode, format)
- **`maputil`** - Map data utilities (convert, sub-value get, merge)
- **`mathutil`** - Math utilities (convert, calculations, random)
- **`fsutil`** - Filesystem utilities (file/dir operations)
- **`sysutil`** - System utilities (env, exec, user, process)
- **`timex`** - Enhanced time utilities with additional methods
- **`netutil`** - Network utilities (IP, port, hostname)
- **`jsonutil`** - JSON utilities (read, write, encode, decode)

### Debug and Testing
- **`dump`** - Value printing with auto-wrap and call location
- **`testutil/assert`** - Common assertion functions for testing
- **`errorx`** - Enhanced error handling with stacktrace

### Extra Tools
- **`cflag`** - Extended command-line flag parsing
- **`cliutil`** - Command-line utilities (colored output, input)

## Common Development Workflows

### Running Tests
```bash
# Run all tests (NEVER CANCEL - 60+ minute timeout recommended)
# Takes ~4 seconds with cache, ~24 seconds on first run
go test ./...

# Run specific package tests
go test ./arrutil
go test ./strutil

# Run with coverage (generates profile.cov file)
go test -coverprofile="profile.cov" ./...

# Run subset with coverage
go test -coverprofile="profile.cov" ./arrutil ./strutil
```

### Code Quality
```bash
# Format code (REQUIRED before commit)
go fmt ./...

# Or use make target
make csfix

# Check formatting issues
make csdiff

# Static analysis (optional - has acceptable warnings)
staticcheck ./...
```

### Documentation
```bash
# Generate README files
make readme

# Or manually
go run ./internal/gendoc -o README.md
go run ./internal/gendoc -o README.zh-CN.md -l zh-CN
```

## Validation Scenarios

**ALWAYS run this validation after making changes:**

Create a test file to verify core functionality:
```go
package main

import (
    "fmt"
    "github.com/gookit/goutil"
    "github.com/gookit/goutil/arrutil"
    "github.com/gookit/goutil/strutil"
)

func main() {
    // Test core functions
    fmt.Println("IsEmpty(''):", goutil.IsEmpty(""))
    fmt.Println("Contains('hello', 'el'):", goutil.Contains("hello", "el"))
    
    // Test array utilities
    fmt.Println("StringsHas(['a','b'], 'a'):", arrutil.StringsHas([]string{"a","b"}, "a"))
    
    // Test string utilities  
    fmt.Println("HasPrefix('hello', 'he'):", strutil.HasPrefix("hello", "he"))
    
    fmt.Println("✅ Validation complete")
}
```

Run with: `go run /tmp/validation.go`

## Critical Build Information

### Timing Expectations
- **Build time**: ~0.4 seconds (very fast)
- **Test time**: ~4 seconds with cache, ~24 seconds without cache (NEVER CANCEL - use 60+ minute timeout)
- **Module download**: ~5 seconds on first run
- **README generation**: ~0.17 seconds
- **Formatting**: ~0.16 seconds

### Requirements
- **Go version**: 1.19+ (tested up to 1.24)
- **Dependencies**: golang.org/x/sync, golang.org/x/sys, golang.org/x/term, golang.org/x/text

### CI Validation
The CI runs on:
- Ubuntu and Windows
- Go versions 1.19, 1.20, 1.21, 1.22, 1.23, 1.24
- Uses staticcheck for linting
- Requires proper code formatting

## Common APIs

### Core goutil functions
```go
goutil.IsEmpty(value)     // Check if value is empty
goutil.IsEqual(a, b)      // Deep equality check
goutil.Contains(arr, val) // Check if array/slice/map contains value
```

### Array utilities (arrutil)
```go
arrutil.StringsHas([]string{"a","b"}, "a")  // true
arrutil.IntsHas([]int{1,2,3}, 2)            // true
arrutil.Reverse(slice)                       // reverse in-place
```

### String utilities (strutil)
```go
strutil.HasPrefix("hello", "he")            // true
strutil.Truncate("hello world", 5, "...")   // "he..."
strutil.PadLeft("hi", "0", 5)               // "000hi"
```

### Testing patterns
```go
import "github.com/gookit/goutil/testutil/assert"

func TestExample(t *testing.T) {
    assert.Eq(t, expected, actual)
    assert.True(t, condition)
    assert.NoErr(t, err)
}
```

### Troubleshooting

### Common Issues
- **Build failures**: Run `go mod tidy` first
- **Test timeouts**: Use 60+ minute timeouts, tests can take 24+ seconds on first run but are fast (~4 seconds) with cache
- **CI formatting failures**: Always run `go fmt ./...` before committing
- **Import errors**: Check that package names match directory structure
- **Coverage files**: Coverage testing creates `.cov` files that should not be committed

### Known Acceptable Issues
- staticcheck reports some unused variables in internal packages - these are acceptable
- Some test files may show formatting changes - apply with `go fmt ./...`
- `make csdiff` may show example files that need formatting - format them if working in those areas

## Project Conventions

### File Organization
- Main packages in root directories (arrutil/, strutil/, etc.)
- Internal utilities in `internal/` (no test coverage required)
- Extended utilities in `x/` subdirectory
- Test files use `*_test.go` naming
- Documentation generation via `internal/gendoc/`
- Example files in package `_examples/` directories (may need formatting)

### Testing
- Use `github.com/gookit/goutil/testutil/assert` for assertions
- Multiple test cases use `t.Run()` pattern
- Test coverage is tracked and reported to coveralls
- Coverage files (*.cov) should not be committed - add to .gitignore if needed

### Documentation
- README files are auto-generated from templates in `internal/gendoc/template/`
- Chinese documentation available as README.zh-CN.md
- API documentation at pkg.go.dev
- Do not manually edit README.md - edit templates instead