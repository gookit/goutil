package cliutil

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/gookit/goutil/x/ccolor"
	"github.com/gookit/goutil/x/stdio"
)

// SimpleTable 使用标准库 text/tabwriter 实现简单的表格
func SimpleTable(cols []string, rows [][]any) {
	var sb strings.Builder
	// 参数说明：输出流, 最小单元宽度, 制表符宽度, 填充空格数, 填充字节, 标志位(Debug加"|")
	w := tabwriter.NewWriter(&sb, 0, 2, 3, ' ', tabwriter.Debug)

	// 渲染表头
	header := strings.Join(cols, "\t  ")
	stdio.Fprintln(w, header)
	// stdio.Fprintln(w, strings.Repeat("-", len(header) + len(cols)*2))

	// 渲染行
	for _, row := range rows {
		var rowStr []string
		for _, cell := range row {
			rowStr = append(rowStr, fmt.Sprintf("%v", cell))
		}
		stdio.Fprintln(w, strings.Join(rowStr, "\t  "))
	}

	_ = w.Flush()
	fmt.Println(sb.String())
}

// TableStyle 定义表格的样式
type TableStyle struct {
	// 边框字符
	BorderTop     string // 顶部边框
	BorderBottom  string // 底部边框
	BorderLeft    string // 左边框
	BorderRight   string // 右边框
	BorderJoin    string // 内部交叉点
	BorderRowSep  string // 行分隔符
	HeaderSep     string // 表头分隔符
	ColumnSep     string // 列分隔符
	ColumnPadding int    // 列内边距

	// 默认的对齐方式 0: 左对齐, 1: 居中, 2: 右对齐
	//  - 默认左对齐
	DefaultAlign int
	// 定义每列的对齐方式 (0: 左对齐, 1: 居中, 2: 右对齐)
	Align []int

	// 颜色设置 (需要终端支持 ANSI 颜色)
	HeaderColor string
	RowColor    string
	BorderColor string
}

// 默认的表格样式实例
var defaultStyle = DefaultStyle()

// MinimalStyle 极简风格 (无边框)
var MinimalStyle = &TableStyle{
	BorderTop:     "",
	BorderBottom:  "",
	BorderLeft:    "",
	BorderRight:   "",
	BorderJoin:    " ",
	HeaderSep:     "-",
	ColumnSep:     " ",
	BorderRowSep:  "",
	ColumnPadding: 1,
	HeaderColor:   "ylw", // 黄色
}

// DefaultStyle 返回默认的表格样式
func DefaultStyle() *TableStyle {
	return &TableStyle{
		BorderTop:     "-", // 使用横线作为顶部边框
		BorderBottom:  "-", // 使用横线作为底部边框
		BorderLeft:    "|", // 使用竖线作为左边框
		BorderRight:   "|", // 使用竖线作为右边框
		BorderJoin:    "+", // 使用加号作为交叉点
		BorderRowSep:  "-", // 使用横线作为行分隔
		HeaderSep:     "=", // 使用等号作为表头分隔
		ColumnSep:     "|",
		ColumnPadding: 1,      // 默认1个空格的边距
		HeaderColor:   "bold", // 粗体
		RowColor:      "",     // 默认无颜色
		BorderColor:   "",     // 默认无颜色
	}
}

// ShowTable CLI渲染显示显示表格
func ShowTable(cols []string, rows [][]any, options ...*TableStyle) {
	ccolor.Println(FormatTable(cols, rows, options...))
}

// FormatTable CLI渲染显示显示表格
func FormatTable(cols []string, rows [][]any, options ...*TableStyle) string {
	return NewTableBuilder(options...).Format(cols, rows)
}

type TableBuilder struct {
	*TableStyle
	// sb strings.Builder
	// context data
	colWidths  []int
	totalWidth int
}

func NewTableBuilder(options ...*TableStyle) *TableBuilder {
	t := &TableBuilder{
		TableStyle: defaultStyle,
	}

	// 处理选项
	if len(options) > 0 && options[0] != nil {
		t.TableStyle = options[0]
	}
	return t
}

func (t *TableBuilder) prepare(cols []string, rows [][]any) {
	// 计算每列的最大宽度 (包含padding)
	colWidths := make([]int, len(cols))
	for i, col := range cols {
		colWidths[i] = len(col)
	}
	for _, row := range rows {
		for i, cell := range row {
			str := fmt.Sprintf("%v", cell)
			if len(str) > colWidths[i] {
				colWidths[i] = len(str)
			}
		}
	}
	t.colWidths = colWidths

	// 计算总宽度
	padding := t.ColumnPadding
	for _, w := range colWidths {
		t.totalWidth += w + padding*2 + 1 // 内容 + 左右padding + 分隔符
	}
	t.totalWidth += 1 // 最后的右边框

}

// Format 格式化构建CLI表格
func (t *TableBuilder) Format(cols []string, rows [][]any) string {
	if len(cols) == 0 {
		return ""
	}

	t.prepare(cols, rows)

	var buf bytes.Buffer

	// 1. 顶部边框
	if t.BorderTop != "" {
		buf.WriteString(t.buildSep(t.BorderTop))
		buf.WriteRune('\n')
	}

	// 2. 表头
	buf.WriteString(t.buildRow(cols, true))

	// 3. 表头分隔符
	if t.HeaderSep != "" {
		buf.WriteString(t.buildSep(t.HeaderSep))
		buf.WriteRune('\n')
	}

	// 4. 数据行
	for _, row := range rows {
		rowStrs := make([]string, len(cols))
		for i, cell := range row {
			rowStrs[i] = fmt.Sprintf("%v", cell)
		}
		buf.WriteString(t.buildRow(rowStrs, false))
	}

	// 底部边框
	if t.BorderBottom != "" {
		buf.WriteString(t.buildSep(t.BorderBottom))
	}

	return buf.String()
}

// 构建辅助函数：构建分隔线
func (t *TableBuilder) buildSep(char string) string {
	if char == "" {
		return ""
	}

	var sb strings.Builder
	sb.WriteString(t.BorderLeft)
	for i, w := range t.colWidths {
		if i > 0 {
			sb.WriteString(t.BorderJoin)
		}
		sb.WriteString(strings.Repeat(char, w+t.ColumnPadding*2))
	}

	sb.WriteString(t.BorderRight)
	return sb.String()
}

// 构建辅助函数：构建一行
func (t *TableBuilder) buildRow(cells []string, isHeader bool) string {
	var sb strings.Builder
	sb.WriteString(t.BorderLeft)
	padding := t.ColumnPadding

	for i, cell := range cells {
		if i > 0 {
			// sb.WriteString(t.BorderColor)
			sb.WriteString(t.ColumnSep)
		}

		// 应用对齐
		width := t.colWidths[i]
		alignment := 0 // 默认左对齐
		if t.Align != nil && i < len(t.Align) {
			alignment = t.Align[i]
		}

		padLeft := strings.Repeat(" ", padding)
		padRight := strings.Repeat(" ", padding)

		switch alignment {
		case 1: // 居中
			totalPad := width - len(cell)
			if totalPad > 0 {
				leftPad := totalPad / 2
				rightPad := totalPad - leftPad
				padLeft = strings.Repeat(" ", padding+leftPad)
				padRight = strings.Repeat(" ", padding+rightPad)
			}
		case 2: // 右对齐
			padLeft = strings.Repeat(" ", width+padding-len(cell))
			padRight = strings.Repeat(" ", padding)
		default: // 左对齐
			padRight = strings.Repeat(" ", width+padding-len(cell))
		}

		sb.WriteString(padLeft)
		// sb.WriteString(cell)
		if isHeader {
			sb.WriteString(ccolor.WrapTag(cell, t.HeaderColor))
		} else {
			sb.WriteString(ccolor.WrapTag(cell, t.RowColor))
		}

		sb.WriteString(padRight)
	}

	sb.WriteString(t.BorderRight)
	sb.WriteString("\n")

	return sb.String()
}
