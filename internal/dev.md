# Dev

## 生成README

```bash
make readme
```

或者：

```bash
go run ./internal/gendoc -o README.md
go run ./internal/gendoc -o README.zh-CN.md -l zh-CN
```

## AI 生成单元测试

- 先选中要测试的方法体
- 输入下面的提示语

```text
/unittest 生成测试
使用 "github.com/gookit/goutil/testutil/assert" 断言
一个方法如果有多个case，使用 t.Run 方式测试
```

**提示2**：

选中要测试的文件，输入下面的提示语

```text
为 SliceToSMap, SliceToMap, SliceToTypeMap  生成单元测试
- 测试框架使用 github.com/gookit/goutil/testutil/assert
```