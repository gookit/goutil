# Obox - 轻量级依赖管理包设计

## 概述

`obox` 是一个轻量级的依赖管理包，使用包级变量和泛型实现，支持命名注册、多种生命周期模式。

## 目标

- 简洁的泛型 API
- 支持命名注册（同一类型可注册多个实例）
- 支持三种生命周期：单例、瞬态、惰性单例
- 类型安全，类型不匹配时返回错误
- 线程安全

## API 设计

### 包级泛型函数

```go
// 注册依赖
func Set[T any](name string, val T, opts ...Option)

// 安全获取，返回错误
func Get[T any](name string) (T, error)

// Must 版本，不存在或类型错误时 panic
func MustGet[T any](name string) T

// 检查是否存在
func Has(name string) bool

// 删除注册
func Delete(name string)

// 清空所有
func Reset()
```

### 生命周期配置

```go
type Option func(*entry)

// 瞬态模式：每次 Get 创建新实例
func Transient() Option

// 惰性单例：首次 Get 时通过工厂函数创建
func Lazy(factory func() any) Option
```

### 错误定义

```go
var (
    ErrNotFound     = errors.New("obox: dependency not found")
    ErrTypeMismatch = errors.New("obox: type mismatch")
)
```

## 内部实现

### 核心结构

```go
type Lifetime int

const (
    Singleton Lifetime = iota  // 默认，单例
    Transient                  // 每次获取创建新实例
    Lazy                       // 惰性单例，首次 Get 时创建
)

type entry struct {
    value    any
    lifetime Lifetime
    factory  func() any
}

type container struct {
    mu      sync.RWMutex
    entries map[string]*entry
}

var defaultContainer = &container{
    entries: make(map[string]*entry),
}
```

### Get 逻辑

1. 加读锁查找 entry
2. 不存在返回 `ErrNotFound`
3. 存在则根据生命周期处理：
   - Singleton: 直接返回存储的值
   - Transient: 调用 factory 创建新实例
   - Lazy: 首次调用 factory 创建并缓存，后续返回缓存值
4. 类型断言，失败返回 `ErrTypeMismatch`

## 使用示例

```go
package main

import "github.com/gookit/goutil/x/obox"

func main() {
    // 单例注册（默认）
    obox.Set("db", &Database{})

    // 惰性单例
    obox.Set("cache", nil, obox.Lazy(func() any {
        return NewCache()
    }))

    // 瞬态/工厂模式
    obox.Set("request", &Request{}, obox.Transient())

    // 安全获取
    db, err := obox.Get[*Database]("db")
    if err != nil {
        panic(err)
    }

    // Must 获取
    cache := obox.MustGet[Cache]("cache")

    // 检查存在
    if obox.Has("db") {
        // ...
    }

    // 删除
    obox.Delete("request")

    // 清空所有
    obox.Reset()
}
```

## 文件结构

```
goutil/x/obox/
├── obox.go      # 核心实现
├── obox_test.go # 单元测试
└── README.md    # 文档（可选）
```

## 测试用例

1. 基本注册和获取
2. 类型不匹配返回错误
3. MustGet 在错误时 panic
4. 单例模式：多次获取返回同一实例
5. 瞬态模式：每次获取创建新实例
6. 惰性单例：首次获取时创建
7. 并发安全测试
8. Has/Delete/Reset 功能测试