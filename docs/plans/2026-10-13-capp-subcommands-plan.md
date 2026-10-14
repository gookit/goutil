<!-- template_id: plan; template_version: 1.2.0 -->
# capp 两级子命令支持实施计划

> 状态：Draft 0.2 / 待人工计划批准
> 实施进度：T1–T8 已全部完成（2026-10-13），对应提交 `a2811a9`…`c926560`；另有一项计划外的独立修复 `17e1775`（命令描述空值安全 + 命令帮助跟随 app 输出）。

## 修订记录

| 版本 | 日期 | 作者 | 摘要 |
|---|---|---|---|
| 0.1 | 2026-10-13 | Claude（inhere 监督） | 初稿：按设计 0.3 拆 8 个任务（基线 → 重构 → 注册 → 分发 → 帮助 → 回归 → 文档 → 收口），含能力发现、基线、回滚与完成 Gate |
| 0.2 | 2026-10-13 | Claude（inhere 监督） | 按一轮独立评审修订：组命令帮助实现与 `HelpFunc` 安装并入 T3（去掉无法判定的「占位实现」，签名统一为 `func(*cflag.CFlags)`）；T4 明确 `App.BeforeRun/AfterRun` 挂到「实际执行的命令」分支；验证命令统一为整包运行并规定新用例命名（避免 `-run` 过滤漏测）；补 `App.NameWidth` 的宽度写回与注册抖动（`app.propagate`）；示例文件确定为 `_example/group.go`；完成 Gate 顺序与任务计数口径修正 |

> 仅语义变化递增版本；纯 identity/provenance/元数据纠正沿用原版本，并在 Git/进度记录中留痕。

## 规划声明

- 思考模式：`RIGOROUS`（SR1408；计划的编写只允许严谨模式）。
- 核心目标：把设计 0.3 的 `cmdSet` 抽象、`Cmd` 子命令注册、逐级下钻与组帮助落到 `gookit2/goutil/cflag/capp`，每一步都有可执行验证。
- scope freeze：唯一 owner 是 `gookit2/goutil/cflag/capp`（含 `_example/`、`README.md`、`README.zh-CN.md`）与本次新增文档。不改 `cflag` 包行为，不改任何消费方仓库。
- 扩围策略：`DEFER_OR_REQUEST`；实施中发现计划外路径/符号按 plan-contract 的四类（`Operational Discovery`/`Corrective`/`Semantic Amendment`/`Ownership Conflict`）分类处理。
- 交付档位：Full（新增导出 API）；治理暴露度低（SR1409）→ 评审封顶一轮合并轴评审。
- 评审预算：本计划一轮合并轴独立评审（2026-10-13 已完成，结论“需修订”，5 条 `CORE_BLOCKING` 已在 0.2 处理）。
- 停止条件：计划获人工批准前不实施；出现语义/验收/授权变化时停止并回到 Gate。

## 目标与完成定义

目标：交付设计 0.3 的对外结果——`app.Add(group)` + `group.Add(sub)` 可用；`app group sub -x val arg` 正确分发；`app group`（无位置参数）按 `Func` 是否为空执行或显示组帮助；`app group -h`/`help` 显示组帮助；现有单级应用零改动继续工作。

完成定义（全部满足才算完成）：

1. `cflag/capp` 三个文件按设计 0.3 落地：`cmdset.go`（新增）、`cmd.go`、`app.go`。
2. 对外 API 只增不改：`App.Add/AddOrErr/NameWidth`、`Cmd.Parse/MustParse/QuickRun`、既有错误文案与帮助外形不变（含 `App.NameWidth` 仍可外部设置且仍被顶层帮助使用）。
3. 新增导出符号：`Cmd.Add`、`Cmd.AddOrErr`、`Cmd.Sub`、`Cmd.HasSubs`、`Cmd.Parent`、`Cmd.Path`。
4. 设计 0.3「安全」节列的四个评审触发点各有回归用例：repeatable 选项不重复写入、`AfterFlagParse` 只触发一次、组命令位置参数冲突在注册期报错、两种注册顺序下组帮助与 `Path()` 均正确。
5. `go test ./...`、`go vet ./cflag/...`、`gofmt -l ./cflag` 全部无问题。
6. 中英文 README 都有内容一致的「两级子命令」小节；示例 `cflag/_example/group.go` 的输出满足 T7 列出的两条可核对断言（`-h` 输出含 `Commands:` 与两个子命令名；`sub1 -x val arg` 输出含 `x=val`）。

## 范围、排除项与授权

范围（本计划的任务只允许改这些路径）：

- `gookit2/goutil/cflag/capp/cmdset.go`（新增）
- `gookit2/goutil/cflag/capp/cmd.go`
- `gookit2/goutil/cflag/capp/app.go`
- `gookit2/goutil/cflag/capp/*_test.go`
- `gookit2/goutil/cflag/README.md`、`gookit2/goutil/cflag/README.zh-CN.md`
- `gookit2/goutil/cflag/_example/group.go`（新增示例）

排除项：

- 不改 `cflag` 包任何文件（设计 D7）。
- 不改 `suag`、`blogshare`、`homepagex`、`miglite` 等消费方，也不替它们拆分命令。
- 不做跨级命令名/别名冲突校验、不做 `app help <cmd> <sub>`、不做 `MustSub`、不做 category/命令 ID、不做 shell 补全（设计 deferred）。
- 不 push、不发版、不改 `go.mod`/依赖。
- 不改动设计文档 0.3（设计是已批准的输入；若实施中必须改设计，按 `Semantic Amendment` 回到 Gate）。

授权边界：本地按功能点提交是默认授权的；本次任务的执行仍需单独的当前执行请求。计划获批不授权 push、release、部署或改动其他项目。

- `host_or_non_offline_action=NOT_APPLICABLE`

（本计划范围内没有任何出机器动作：不 push、不发版、不部署、不改主机配置、不产生外部消息；`go run`/`go test`/`go vet` 都在本机与本地仓库内。）

## 输入与批准证据

- 设计：`gookit2/goutil/docs/design/2026-10-13-capp-subcommands-design.md`（Draft 0.3），设计提交 `b29ee8c`。
- 设计阶段的人工决定（用户原话 2026-10-13）：「组命令无子名时 - Func 不为空则执行，否则显示帮助；其他按建议；独立评审一次」。
- 设计评审（一轮合并轴，2026-10-13）：结论“需修订”，4 条 `CORE_BLOCKING`（A1–A4）已在设计 0.3 处理，5 条 `SUGGESTION`（S5–S9）已合入。
- 编写计划的授权（用户原话 2026-10-13）：「批准编写计划」。
- 本计划评审（一轮合并轴，2026-10-13）：结论“需修订”，P1–P5 五条 `CORE_BLOCKING` 与 P6–P13 建议已在 0.2 处理。
- 规范输入：`SR1403`、`SR1409`、`GR102`、`GR104`、`GR106`。

## Capability Discovery

### Capability decisions

| capability_id | required_capability | searched_candidates | direct_reuse | thin_adapter_or_owner_extension | decision | proven_gap | duplication_and_lifecycle_risk |
|---|---|---|---|---|---|---|---|
| CAP-01 | 组命令分发：用组自身选项集解析出子命令名，且只解析一次 | `cflag.CFlags.Prepare`、`DoParseStopOnArg`、`Args`、`ShowHelp`、`ErrStopRun`、`AfterFlagParse`、`HelpFunc`；标准库 `flag.FlagSet.Parse`；`gcli` 命令树（`Command.Add`/`innerDispatch`） | 是：`Prepare`+`DoParseStopOnArg`+`Args`+`ShowHelp` 已导出，评审逐项核对可用（设计 D7） | 无需 Adapter：在 `capp` 内新增 `App.runCmd` 原位扩展 | DIRECT_REUSE | none | 无：不新增 Module、不引依赖，`capp` 仍是唯一 owner |
| CAP-02 | 命令集注册表（名字/命令表/别名/列宽）与子命令列表渲染 | 现有 `App.names`/`cmds`/`cmdAs`/`NameWidth`；`maputil.Aliases`；`CFlags.RenderOptionsHelp`；`strutil.Buffer`/`PadRight`/`UpperFirst` | 是：别名表、选项帮助渲染、字符串补齐全部复用 | 在同一 owner `capp` 内把命令集抽成未导出 `cmdSet`（携带 `add`/`find`/`render`/宽度公式） | OWNER_EXTENSION | none | 无：同包内重构，无第二 registry；`App.NameWidth` 语义不变 |
| CAP-03 | 注册期 fail-fast：重名、别名重复、组命令与位置参数声明冲突 | `App.addCmd`/`Add` 现有 error+panic 分工（`app.go:118`、`app.go:137-139`）；`maputil.Aliases.AddAlias` 冲突 panic（`alias.go:13-18`） | 是：沿用既有 error/panic 分工与别名表语义 | 在 `Cmd.AddOrErr`/`Cmd.AddArg`/`Cmd.BindArg` 原位加校验 | OWNER_EXTENSION | none | 无：规则只有一份，`App` 与 `Cmd` 共用 `cmdSet.add` |
| CAP-04 | 单元测试与断言 | `github.com/gookit/goutil/x/assert`（GR106 指定）；现有 `cflag/capp/*_test.go` | 是 | 无 | DIRECT_REUSE | none | 无 |

### New module candidates

| candidate_id | capability_id | proposed_module | searched_candidates | direct_reuse_gap | thin_adapter_or_owner_extension_gap | proven_gap | unique_owner_and_lifecycle | deletion_or_merge_handling |
|---|---|---|---|---|---|---|---|---|
| NONE | none | none | none | none | none | none | none | none |

### Rejected new tools

| rejected_candidate | capability_id | deletion_test_and_reason |
|---|---|---|
| 引入 `gookit2/gcli` 的命令树作为实现 | CAP-01 | 删除测试：若 `capp` 需要 category/命令 ID/独立解析器等 gcli 独有能力，则 gcli 才是需要的层；实测 `cmdSet`（约 40 行）即可满足验收，引入跨 Module 依赖只增加版本与生命周期负担，故不引入 |
| 新建 `capp/subcmd` 子包或独立 `SubCmd` 类型 | CAP-01 | 删除测试：把子命令能力从 `Cmd` 上删掉后，帮助、别名、选项归属都无法表达（关键功能丢失）→ 说明不需要第二 owner；能力并入 `capp/cmd.go` |
| 给 `cflag` 新增导出的组解析 API | CAP-01 | 删除测试：删掉该 API 后，`Prepare`+`DoParseStopOnArg`+`Args` 仍能完成组分发（评审 A1–A4 修订后成立）→ 不必要，且会扩大 `cflag` 回归面（设计 D7） |

## 前置检查与 fail-closed 条件

workspace baseline（2026-10-13 实测）：

- Git root：本仓库根（工作区内的 `gookit2/goutil`，独立 Git 仓库，本计划的唯一 Git 边界）；branch `master`；HEAD `b29ee8c`。
- HEAD 说明：设计文档盘点时记录的是代码 HEAD `068984d`；`b29ee8c` 只多出设计文档提交，`cflag` 代码未变（T1 复核）。
- `git status --short` 为空：无 in-scope dirty/untracked，无需要保留的无关 dirty 路径。
- 工具链：`go1.25.10 windows/amd64`；`goutil` 不在任何 `go.work` 中（工作区内的 `go.work` 只在 `gookit/clitools`、`gookit2/gcli`）。
- 基线测试：`go test ./cflag/...` 全绿（运行目录 `gookit2/goutil`；`cflag`、`cflag/capp` 两包 ok）。
- 收口命令基线（同为 2026-10-13 实测，运行目录 `gookit2/goutil`）：`gofmt -l ./cflag` 无输出；`go vet ./cflag/...` 无输出；`go test ./...` 全部 `ok`（无预存在失败）。
- expected paths/symbols（planning evidence，不是封闭白名单）：`cflag/capp/{cmdset.go,cmd.go,app.go}`；新增符号 `Cmd.Add/AddOrErr/Sub/HasSubs/Parent/Path`；既有符号 `App.Add/AddOrErr/NameWidth`、`Cmd.Parse`、`CFlags.Prepare/DoParseStopOnArg/Args/ShowHelp/HelpFunc/RenderOptionsHelp`。
- 能力与授权边界：只用 stdlib 与 `goutil` 自身包；本地提交已授权，push/发版未授权。

约定（减少验证歧义）：

- 新增用例统一放 `cflag/capp/subcmd_test.go`，命名 `TestSubCmd_*`（注册/分发/校验）与 `TestGroupHelp_*`（帮助/宽度）。
- 每个任务的验证命令都用**整包**运行 `go test ./cflag/capp/ -v`，不使用 `-run` 过滤，避免过滤漏测。

fail-closed 条件（任一成立即停止并回报，不自行扩围）：

1. 开工时 `git status --short` 非空，或出现计划外 dirty/untracked 归属不明。
2. 基线 `go test ./cflag/...` 变红。
3. 设计 0.3 引用的任一 `cflag` 导出 API（`Prepare`/`DoParseStopOnArg`/`Args`/`ShowHelp`/`HelpFunc`）在当前 HEAD 上不存在或语义不符。
4. 需要改 `cflag` 包、需要引新依赖、或需要改消费方才能完成验收（属 `Semantic Amendment`）。

## 波次与依赖

```text
W0 T1 基线确认（只读，无提交）
  └─► W1 T2 抽 cmdSet（纯重构）─► W2 T3 Cmd 子命令注册与组帮助 ─► W3 T4 runCmd 逐级下钻
        ─► W4 T5 帮助渲染验收 ─► W5 T6 评审触发点回归用例 ─► W6 T7 文档与示例 ─► W7 T8 收口验证与提交
```

依赖链为线性：T1 → T2 → T3 → T4 → T5 → T6 → T7 → T8。每个任务以「测试绿 + 一次原子提交」结束，形成恢复点。

## 任务

### T1 基线确认

- 文件: 无（只读）
- 动作: 在运行目录 `gookit2/goutil` 执行 `git rev-parse --show-toplevel`、`git rev-parse --abbrev-ref HEAD`、`git rev-parse --short HEAD`、`git status --short`、`go version`；核对 `cflag/capp/{app.go,cmd.go}` 与设计 0.3 引用的行号/符号仍在（`App.addCmd` 宽度更新、`App.findCmd`、`CFlags.HelpFunc`）。不改任何文件。
- 验证: 运行目录 `gookit2/goutil`：`go test ./cflag/...`；期望两个包 `ok`。对照命令：`git status --short` 期望无输出。
- 完成标准: 上述命令输出与本计划 baseline 一致；不一致时按「前置检查与 fail-closed 条件」停止。
- 依赖: 无

### T2 抽出 cmdSet（纯重构，行为不变）

- 文件: `cflag/capp/cmdset.go`（新增）、`cflag/capp/app.go`
- 动作: 新增未导出 `type cmdSet struct { names []string; cmds map[string]*Cmd; cmdAs maputil.Aliases; nameWidth int }` 与 `newCmdSet()`（`nameWidth` 初始 12）、`add(c *Cmd) error`、`find(name string) (*Cmd, bool)`、`render(buf *strutil.Buffer, width int)`、`nameLen(c *Cmd) int`；把 `App.addCmd`（`app.go:131-158`）与 `App.findCmd`（`app.go:255-264`）的逻辑搬入 `add`/`find`。**宽度处理**（评审 P5）：宽度公式抽到 `cmdSet.nameLen(c *Cmd) int`（即现有 `ln` 计算：命令名长度 + 别名串长度 + 4）；`cmdSet.add` 用它更新 `s.nameWidth`，`App.addCmd` 用它更新导出的 `a.NameWidth`（保持 `app.go:149` 的现有行为，`App.NameWidth` 仍可外部设置）；`App` 用 `commands *cmdSet` 替换 `names/cmds/cmdAs` 三个字段；`App.showHelp` 命令列表段改调 `commands.render(buf, a.NameWidth)`，渲染沿用现有规则（`sort.Strings` 就地排序、别名展示、`UpperFirst`、右侧补齐），输出与现状逐字节一致。本任务只重构，不新增行为。
- 验证: 运行目录 `gookit2/goutil`：`go test ./cflag/...`、`go vet ./cflag/...`；`git diff --stat` 只含上述两文件。
- 完成标准: 既有 `cflag/capp` 用例未做任何修改即全绿；`App` 导出 API 与错误文案不变；`App.NameWidth` 仍生效（`app_test.go` 中长命令名参与帮助断言的部分不变）；在 `subcmd_test.go` 增加一条 `TestGroupHelp_NameWidth` 断言：设置 `app.NameWidth = 24` 后帮助中命令名列宽随之变化，且连续两次 `showHelp` 输出一致（覆盖就地排序不可见性）。
- 依赖: T1

### T3 Cmd 子命令注册、父子关系、注册期校验与组帮助实现

- 文件: `cflag/capp/cmd.go`、`cflag/capp/cmdset.go`、`cflag/capp/subcmd_test.go`（新增）
- 动作:
  - `Cmd` 增加未导出字段 `app *App`、`parent *Cmd`、`subs *cmdSet`、`hasPosArg bool`。
  - 新增 `Add(cmds ...*Cmd)`、`AddOrErr(cmds ...*Cmd) error`、`Sub(name string) (*Cmd, bool)`、`HasSubs() bool`、`Parent() *Cmd`、`Path() string`。`Path()` 沿 `parent` 链拼接自身路径，有 `app` 时以 `app.Name` 作前缀。
  - `Cmd.AddOrErr` 复用 `cmdSet.add`，设置 `sub.parent`/`sub.app`；**注册抖动（评审 P7）**：`App.addCmd` 在设置 `c.app = a` 后递归把 `app` 传播给 `c.subs` 下所有子命令（`Cmd.setApp(a)` 递归），保证「先 `group.Add(sub)` 再 `app.Add(group)`」时 `Path()` 也带 app 前缀。
  - 注册期校验：`Cmd.AddOrErr` 在「该命令已声明位置参数」时 panic（文案 `group command %q cannot declare positional args`）；覆盖 `AddArg`/`BindArg` 记录 `hasPosArg`，在 `HasSubs()` 为真时 panic。重名沿用 `addCmd` 的 error（`App.Add` 再 panic，`app.go:118`），别名重复沿用 `maputil.Aliases` 的 panic（`alias.go:13-18`）。
  - **组帮助**：实现 `func (c *Cmd) renderGroupHelp(cf *cflag.CFlags)`（签名与 `CFlags.HelpFunc` 一致，`cflag.go:68`，因此可直接赋值），输出：描述取 `getDesc()`；Usage 用 `Path()` 拼 `<root> group SUBCOMMAND [--Options...] [...Arguments]`；选项段复用 `cf.RenderOptionsHelp(buf)` 并追加 `--help, -h` 行；子命令段用 `c.subs.render(buf, c.subs.nameWidth)`；末行 `Use "<root> group SUBCOMMAND --help" for about a command`；`LongHelp`/`Example` 按叶子同样渲染；输出用 `ccolor.Fprintln(cf.Output(), ...)`（与 `CFlags.showHelp` 一致）。`Cmd.AddOrErr` 在首次注册子命令且用户未设 `HelpFunc` 时安装 `c.CFlags.HelpFunc = c.renderGroupHelp`（`CFlags` 为 nil 时先按 `initCmd` 的方式补齐）。
- 验证: 运行目录 `gookit2/goutil`：`go test ./cflag/capp/ -v`（新增 `TestSubCmd_*`、`TestGroupHelp_*`）；`go build ./cflag/...`。
- 完成标准: 新增用例判定以下事实——注册后 `HasSubs()`/`Sub`/`Parent`/`Path` 正确；两种注册顺序（先 `app.Add(group)` 后 `group.Add(sub)`，以及反向）下 `HelpFunc` 均指向组帮助且 `Path()` 都带 app 前缀；重名返回 error（`App.Add` 时 panic）；别名重复 panic；组命令与 `AddArg` 双向 panic（先声明后成组、先成组后声明）；独立使用 `Cmd`（`app == nil`）时 `Add`/`Sub`/`Path` 正常且 `Path()` 不带前缀；`app group -h` 的输出含组描述、组选项、子命令名（含别名）与末行提示。既有用例不变仍全绿。
- 依赖: T2

### T4 runCmd 逐级下钻与 App 钩子挂载点

- 文件: `cflag/capp/app.go`、`cflag/capp/subcmd_test.go`
- 动作:
  - 新增 `func (a *App) runCmd(c *Cmd, args []string) error`，形状按设计 0.3「总体方案 3」：叶子 → 执行分支；组命令 → `c.Prepare()` + `c.DoParseStopOnArg(args)`（只解析一次）、`-h`/`--help` 的 `flag.ErrHelp` 与 `cflag.ErrStopRun` 按既定分支处理；`help` 保留名 → `c.ShowHelp()`；子命令名取 `c.Args()[0]`（不用 `RemainArgs()`），命中则 `a.runCmd(sub, c.Args()[1:])`；未命中的第一个位置参数报 `command %q not exists in %q`；无位置参数时 `Func == nil` → `c.ShowHelp()`，否则执行组命令自身 `Func`。
  - **App 钩子挂载点（评审 P3）**：把现有位于 `RunWithArgs` 的 `App.BeforeRun`/`App.AfterRun` 包裹（`app.go:208-219`）改成包裹「实际执行的命令」：叶子命令执行分支，以及组命令执行自身 `Func` 的分支；组命令下钻分支不触发。参数 `cmdArgs` 用该分支实际收到的参数。组命令执行自身 `Func` 时**不**再调用 `c.Parse`（避免二次解析），只跑 `CFlags.BeforeRun`（若设）后调用 `c.Func(c)`。
  - `RunWithArgs` 把 `cmd.Parse(cmdArgs)`（`app.go:214`）替换为 `a.runCmd(cmd, cmdArgs)`，其余（`preRun`、`OnAppFlagParsed`、`RemainArgs` 取命令名、`showHelp` 分支）不变。
- 验证: 运行目录 `gookit2/goutil`：`go test ./cflag/capp/ -v`。
- 完成标准: 新增用例判定——两级分发与位置参数透传；子命令别名命中；父级选项写在子命令名之前可用、写在之后报 `flag provided but not defined`；未知子命令错误文案；`app group -- sub` 与叶子 `--` 行为；三级下钻；钩子用例：叶子命令上 `BeforeRun`/`AfterRun` 各 1 次且收到叶子实收参数，组命令下钻时**不**触发，组命令执行自身 `Func` 时各 1 次。单级既有用例（`TestApp_Run`、`TestApp_Run_cmdOptionsAfterArgs`）未修改仍全绿。
- 依赖: T3

### T5 帮助渲染验收（组 / 顶层 / 叶子）

- 文件: `cflag/capp/subcmd_test.go`
- 动作: 补齐帮助相关的验收断言（实现已在 T2/T3 完成，本任务只加断言与修正）：顶层帮助仍平铺一级命令、含别名与描述、`NameWidth` 生效；组帮助三个入口一致（`app group -h`、`app group help`、无 `Func` 的 `app group`）；组命令设置了 `LongHelp`/`Example` 时按叶子同样渲染；叶子命令帮助不受影响；组帮助中出现 `Use "<root> group SUBCOMMAND --help"` 行。
- 验证: 运行目录 `gookit2/goutil`：`go test ./cflag/capp/ -v`。
- 完成标准: 上述断言全部通过；`TestApp_Run` 内既有的顶层帮助断言（`app_test.go:116-134`）与 `TestApp_AfterFlagParse` 未修改仍通过。
- 依赖: T4

### T6 评审触发点的回归用例

- 文件: `cflag/capp/subcmd_test.go`
- 动作: 补齐设计 0.3「安全」节列出的用例：组命令带 repeatable 选项（如 `--names`，append 语义）断言只收到一次值（A2）；组命令 `AfterFlagParse` 计数为 1（A3）；组命令与 `AddArg` 冲突的 panic 路径（A1/D10）；两种注册顺序下组帮助可用（A4）；组命令短选项 `-v` 与子命令短选项 `-x` 共存（D2）；组命令有 `required` 选项未给时（输入形如 `app session`，设计边界表该行输入列措辞不精确，以本条为准）报 `flag option 'x' is required`，而 `app session -h` 先显示组帮助（S6）。
- 验证: 运行目录 `gookit2/goutil`：`go test ./cflag/capp/ -v`。
- 完成标准: 上述 6 类断言全部通过；无 `t.Skip`。
- 依赖: T5

### T7 文档与示例

- 文件: `cflag/README.md`、`cflag/README.zh-CN.md`、`cflag/_example/group.go`（新增）
- 动作: 中英文 README 各增加「两级子命令」小节，内容一致：注册方式（`app.Add(group)` + `group.Add(sub)`）；父级选项必须写在子命令名之前；组命令无位置参数时按 `Func` 是否为空执行或显示组帮助；组命令不支持位置参数绑定；别名必须在 `Add` **之前**赋值（设计 S7）；只承诺两级（实现递归支持任意级）。示例 `_example/group.go`（`package main`，与既有 `_example/app.go`/`cmd.go` 一致）注册一个组与两个子命令，其中一个带别名，打印收到的选项与参数。
- 验证: 运行目录 `gookit2/goutil`：`go run ./cflag/_example/group.go -h`（期望含 `Commands:` 与两个子命令名）；`go run ./cflag/_example/group.go sub1 -x val arg`（期望输出含 `x=val`）。说明：`_example` 目录名以下划线开头，`go build ./cflag/...` 不会构建它，因此示例以 `go run <file>` 为验证手段。
- 完成标准: 两条 `go run` 断言成立；中英文 README 小节描述一致且与示例行为相符。
- 依赖: T6

### T8 收口验证与提交

- 文件: 无新增（只跑命令与提交）
- 动作: 运行目录 `gookit2/goutil`：`gofmt -l ./cflag`、`go vet ./cflag/...`、`go test ./...`；核对 `git status --short` 只含本计划范围内的路径；按功能点做本地提交（不 push）。
- 验证: 运行目录 `gookit2/goutil`：`go test ./...` 全绿；`gofmt -l ./cflag` 无输出；`git status --short` 只含预期路径。
- 完成标准: goutil 全包测试通过、无格式问题、无计划外文件进入提交。
- 依赖: T7

## 回滚与恢复

- Git 根边界：`gookit2/goutil`。所有改动与提交只在该仓库内；不 `git add -A`，只提交本计划范围内的路径。
- dirty 保护：开工前 `git status --short` 必须为空；实施中若出现计划外 dirty/untracked，先判断归属（`Ownership Conflict` 时停止）。
- 恢复点：T2–T8 每个任务结束且测试绿后做一次原子提交，共 7 个恢复点；任一步失败可从上一个提交继续。
- 回滚策略：整批改动回退用 `git revert` 对应提交链；`capp` 的改动不涉及数据结构、迁移或外部状态，回退后旧行为立即恢复；若示例或 README 有问题，只回退该提交，不影响代码提交。
- 失败恢复：测试红时先定位到最近一个恢复点，按「前置检查与 fail-closed 条件」判断是否属于需要回到设计/计划 Gate 的问题（`Semantic Amendment`）。

## 人工 Gate

- Gate 1（开工前）：计划批准。需用户明确批准本计划后才解除实施 Gate。
- Gate 2（开工前）：当前执行请求。执行本计划需要单独的一次请求；批准计划不等于授权执行（plan-contract Gate）。
- 外部动作：本计划 `host_or_non_offline_action=NOT_APPLICABLE`，范围内无 push/发版/部署/主机配置/外部消息；若实施后需要 push 到远端或发布，需单独具名授权。
- 评审：本计划一轮合并轴独立评审（已完成，P1–P5 已处理）；实施中的发现按 plan-contract 四类分类，不自动开启新一轮 review。

## 可追溯性

| 设计来源 | 计划任务 | 验证 |
|---|---|---|
| 设计 D1/D2、总体方案 1（`cmdSet` 抽象、宽度权威） | T2 | `go test ./cflag/...` 未改既有用例全绿；`TestGroupHelp_NameWidth` |
| 设计 D4/D10/D11、总体方案 2/4（子命令注册、冲突校验、`HelpFunc` 安装、组帮助） | T3 | `go test ./cflag/capp/ -v`（注册/校验/两种顺序/`app group -h`） |
| 设计 D2/D3/D5/D6、总体方案 3/6（逐级下钻、单次解析、钩子语义） | T4 | `go test ./cflag/capp/ -v`（分发、钩子计数） |
| 设计总体方案 4/5、`cmdSet.render` 规则 | T5 | `go test ./cflag/capp/ -v`（帮助三处断言） |
| 设计「安全」节 A1–A4、S6 | T6 | `go test ./cflag/capp/ -v`（6 类回归） |
| 设计「范围」节文档/示例条目、S7（别名时序） | T7 | `go run ./cflag/_example/group.go -h`、`... sub1 -x val arg` |
| 设计「安全」节独立 `Cmd` 用法、`sort.Strings` 就地排序 | T2（排序一致性）、T3（独立 `Cmd`） | `TestGroupHelp_NameWidth`、`TestSubCmd_StandaloneCmd` |
| 设计「安全」节回滚、SR1407 最小验证 | T8 | `go test ./...`、`gofmt -l ./cflag` |
| 设计「待确认事项」deferred 项 | 无任务（保持排除项） | 计划「范围、排除项与授权」已声明不实施 |

## 完成 Gate 与剩余工作

完成 Gate：T1–T8 全部满足各自「完成标准」，且「目标与完成定义」1–6 全部成立；`go test ./...` 全绿、`gofmt` 与 `go vet` 无输出；`host_or_non_offline_action=NOT_APPLICABLE` 范围内的动作全部在本地完成。达成后向用户回报结果（实施/验证/提交证据），并按需请求后续授权——本计划内没有 push/发版，若需要则另提外部动作 Gate。

剩余工作（不在本计划内，实施完成后另行按需开启）：

- 消费方采纳：`suag` 把 `ps/send/gc/sessions` 拆成 `session` 组、`blogshare`/`homepagex`/`miglite` 是否需要分组——各自项目内另立请求。
- 设计 deferred 项：跨级命令名/别名冲突校验、`app help <cmd> <sub>`、`MustSub`、category/命令 ID、shell 补全。
- push 与发版：需单独的具名外部动作授权。
