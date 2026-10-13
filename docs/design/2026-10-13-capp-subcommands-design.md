<!-- template_id: design; template_version: 1.1.1 -->
# capp 两级子命令支持设计

> 状态：Draft 0.3 / 待人工计划批准

## 修订记录

| 版本 | 日期 | 作者 | 摘要 |
|---|---|---|---|
| 0.1 | 2026-10-13 | Claude（inhere 监督） | 初稿：现状盘点、命令集抽象、逐级分发与父级帮助、兼容边界与待确认项 |
| 0.2 | 2026-10-13 | Claude（inhere 监督） | 按用户 2026-10-13 决定：组命令无子命令名时 `Func` 非空则执行、否则显示组帮助（D6、D5、关键流程随之调整）；确认不做 `app help <cmd> <sub>`、不加分组计数标注、只承诺两级 |
| 0.3 | 2026-10-13 | Claude（inhere 监督） | 按一轮独立评审修订：去掉组命令的二次解析（repeatable 选项翻倍、`AfterFlagParse` 二次触发）改为复用执行尾部；子命令名改从 `Args()` 取并禁止组命令绑定位置参数（`AddArg` 会吞子命令名）；组帮助的 `HelpFunc` 改为在子命令注册时安装（注册顺序依赖会导致组帮助失效）；另处理宽度同步、required 边界、别名时序、边界表措辞与测试清单 5 条建议 |

> 仅语义变化递增版本；纯 identity/provenance/元数据纠正沿用原版本，并在 Git/进度记录中留痕。

## 规划声明

- 思考模式：`RIGOROUS`（SR1408）。
- 核心目标：让 `cflag/capp` 在不破坏现有单级用法的前提下支持两级子命令（`app group sub --opts args`），并给出可用的分组帮助；对外 API 只增不改。
- 范围冻结：只改 `gookit2/goutil/cflag/capp`（含 `_example/`、README、测试）。不改 `cflag` 包的对外行为，不改任何消费方（suag / blogshare / homepagex / miglite）代码。
- 扩围策略：`DEFER_OR_REQUEST`。命令名与别名跨命名空间校验、`app help <cmd> <sub>`、分组计数标注等记入待确认或 deferred，不在本设计内实施。
- 交付档位：Full（新增导出 API，属外部 Interface 变化）。治理暴露度按 SR1409 判为**低**：该能力尚未发布、无生产用户与生产数据、错误用还原提交即可恢复、不涉 Secret/部署/出机器动作；因此评审封顶为一轮合并轴评审，不建 review ledger、不做 candidate commit 定位。
- 评审预算：设计一轮合并轴独立评审（2026-10-13 已完成，结论“需修订”，4 条 `CORE_BLOCKING` 已在 0.3 处理）；实施计划一轮。
- 停止条件：用户批准本设计前保持 `Draft`；批准只授权编写实施计划，不授权实施、提交或推送。

## 背景与目标

`cflag/capp` 目前只能构建**一层**命令应用：`App` 持有一批扁平的 `Cmd`，运行时分发只取 `args[0]` 当命令名。命令一多，同一业务域的动词只能平铺在顶层，帮助里也看不出分组。

用户 2026-10-13 提出：`capp` 现在支持一级命令，如何改造增强支持两级子命令。用户可见的结果：

- 能写 `app.Add(group)`，`group.Add(sub1, sub2)`，运行 `app group sub1 -x val arg`。
- `app group`（不给子命令）与 `app group -h` 显示**组帮助**：组的描述、组自己的选项、子命令列表、用法提示。
- 现有单级应用零改动继续工作：`App.Add`/`App.Run`/`Cmd.Func`/全局选项与命令选项的现有位置语义（含 `--` 停止解析）全部不变。
- 子命令支持自己的选项、位置参数、别名、帮助，与今天的一级命令完全一致。

消费方现状（本次不修改，仅说明动机）：suag 12 个命令平铺注册，其中 `ps/send/gc/sessions` 属于同一个会话域（`inhere-tools/suag/internal/cli/root.go:48-61`）；blogshare（`inhere.github.io/share/blogshare/internal/cli/app.go:39-51`）、homepagex（`inhere-tools/homepagex/cmd/homepagex/main.go:53-58`）、miglite（`gookit/miglite/pkg/command/cliapp.go:62-71`）同样是平铺列表；suag 还已在 `cmd/suag/main.go:24-27` 手工模拟“默认子命令”。这些都属于命令数增长后的组织压力，但不构成本次范围。

## 名词

- **一级命令 / 子命令**：`app group` 中的 `group` 是一级命令，`app group sub` 中的 `sub` 是二级命令（子命令）。本设计中“命令”默认指一级命令。
- **组命令（父命令）**：注册了子命令的命令。组命令可以有自己的选项，但不绑定位置参数（D10）；是否承担业务执行由它是否设置 `Func` 决定（D6）。
- **叶子命令**：没有子命令的命令，即今天的一级命令，执行 `Cmd.Func`。
- **命令集 `cmdSet`**：名字、命令表、别名表、名字宽度这四件套的封装。`App` 与组命令各持有一个，互不干扰。
- **命令路径**：从根到当前命令的名字链，如 `suag session list`。用于帮助的 Usage 行与错误信息。
- **别名空间**：一级命令别名与子命令别名各自在所属 `cmdSet` 内解析，互不干扰。

## 范围与非目标

范围：

- 新增 `capp/cmdset.go`（`cmdSet` 与命令表渲染），改造 `capp/app.go`、`capp/cmd.go`。
- `Cmd` 增加子命令注册与父子关系：`Add`、`AddOrErr`、`Sub`、`HasSubs`、`Path`。
- 分发改造：`App.RunWithArgs` 中“取 `args[0]` 当命令名”改为逐级下钻，组命令用自己的选项集解析出子命令名。
- 组命令帮助渲染（组描述 / Usage / 组选项 / 子命令列表 / 提示行）。
- 测试：两级分发、父子别名、父级选项位置、未知子命令、组帮助、组命令注册顺序（先加组/先加子两种）、`--` 边界、三级下钻（实现递归则顺带覆盖）；回归用例覆盖评审提出的四个触发点——组命令带 `AddArg`、组命令带 repeatable 选项、组命令带 `AfterFlagParse`、两种注册顺序下的组帮助。
- 文档：`cflag/README.md` 与 `cflag/README.zh-CN.md` 增加“两级子命令”小节；`cflag/_example/app.go` 增加分组示例。

非目标：

- 不改 `cflag` 包：`Parse`/`DoParse`/`DoParseStopOnArg`/`Prepare`/`BindArg`/帮助渲染等对外语义不动（见 D7）。
- 不改造任何消费方仓库，也不替它们拆分现有命令。
- 不做“默认子命令”注册机制（`app group` 不自动选出某个子命令执行）：组命令无子命令名时是否执行只由其自身 `Func` 决定（D6）。
- 不做命令名与别名的跨级冲突校验、不做分类/分组标题（category）、不做 shell 补全、不做三级以上的显式 API 承诺。
- 不改 `App` 既有导出符号的签名与语义。

## 已确认事实与规范

盘点基于 2026-10-13 只读阅读，`gookit2/goutil` HEAD `068984d`，工作区干净。

`capp` 现状：

| 事项 | 事实（位置） |
|---|---|
| 顶层命令集 | `App.names []string`、`App.cmds map[string]*Cmd`、`App.cmdAs maputil.Aliases`（`cflag/capp/app.go:31-33`），另用 `NameWidth` 控制帮助列宽（`app.go:39`、`app.go:149`） |
| 注册 | `App.Add` → `AddOrErr` → `addCmd`（`app.go:117-158`）：重名报错、追加 `names`、写 `cmds`、注册别名、更新 `NameWidth`、调用 `cmd.initCmd()`、触发 `OnAdd` 钩子 |
| 分发 | `RunWithArgs`（`app.go:175-221`）：`init()` → `preRun`（`Prepare` + 空参/帮助判断 + 全局选项 `DoParseStopOnArg`）→ `OnAppFlagParsed` → `args = RemainArgs()` → `args[0]` 当命令名 → `BeforeRun` → `cmd.Parse(cmdArgs)` → `AfterRun` |
| 查找 | `findCmd`（`app.go:255-264`）：`-` 开头不识别；否则别名解析 + 查表；未命中由 `RunWithArgs:203-206` 报 `input not exists command %q` |
| `Cmd` | `*cflag.CFlags` + `init/Name/Desc/Aliases/OnAdd/Func`（`cflag/capp/cmd.go:15-27`）；**没有** 父子关系与子命令集合，`RunWithArgs` 只为 `args[0]` 找到的那一个命令服务 |
| `Cmd` 初始化只跑一次 | `initCmd`（`cmd.go:85-110`）由 `init` 标记守卫，`App.addCmd` 注册时调用一次（`app.go:152`）；`Func == nil` 时提前返回（`cmd.go:92-94`），`CFlags` 的创建与 `(alias: ...)` 描述追加都在其后。所以“注册后再变成组命令”的时机无法再回到这里 |
| 别名读取时机 | 别名在注册时读取并写入 `cmdAs`（`app.go:145-147`），`Cmd.Aliases` 必须在 `Add` **之前**赋值；`(alias: ...)` 追加到 `CFlags.Desc`（`cmd.go:104-106`，仅 `Func != nil` 时执行到） |
| 命令执行 | `Cmd.Parse`（`cmd.go:79-83`）→ `cflag.CFlags.Parse`（`cflag/cflag.go:268-315`）：`Prepare` → `DoParse`（interspersed）→ `ErrHelp` 显示帮助 → `ErrStopRun` 静默返回 → `BeforeRun` → `Func` |
| 帮助 | 顶层 `App.showHelp`（`app.go:270-315`）排序平铺 `names`，列别名与描述、追加 `help` 行、`LongHelp`、`AfterHelpBuild` 钩子；命令帮助走 `CFlags.showHelp`（`cflag.go:618-672`），可被 `CFlags.HelpFunc`（`cflag.go:68`、`cflag.go:618-621`）整体替换 |

`cflag` 解析器中与本设计直接相关的既有行为：

| 事项 | 事实（位置） |
|---|---|
| `Prepare` 只在首次生效 | `CFlags.prepared` 守卫（`cflag.go:318-322`）：`VisitAll` 解析每个选项 usage 里的 `required`/短选项元信息，填充 `shortcuts` 与 `bindOpts`，并把 `FlagSet.Usage` 置空 |
| `DoParse*` 不调用 `Prepare` | `doParse`（`cflag.go:400-428`）直接做短选项替换、可选重排、`FlagSet.Parse`、`checkBindOpts`、`AfterFlagParse`、`bindParsedArgs`；`Prepare` 只由 `CFlags.Parse`（`cflag.go:283`）与 `App.preRun`（`app.go:225`）调用 |
| 停止式解析 | `DoParseStopOnArg`（`cflag.go:396-398`）→ `doParse(args, false)`：短选项替换带 `stopOnFirstArg` 回调，遇到第一个位置参数即停止改写，后续参数原样交给 `FlagSet.Parse` |
| 交错解析会“抢”选项 | interspersed 模式下 `interspersedArgs`（`cflag.go:438-476`）把位置参数之后能识别的长选项重新排到前面，并以 `--`（`ParseStopMark`，`util.go:145`）分隔；对组命令会把子命令自己的同名选项摘走 |
| 位置参数与剩余参数 | `FlagSet.Args()` 给出解析后的位置参数；`bindParsedArgs`（`cflag.go:535-574`）把其中前 N 个绑给 `AddArg` 声明的参数（`val := args[arg.Index]`，`cflag.go:552`），其余进 `remainArgs`（`cflag.go:566-572`）。所以**判定子命令名必须看 `Args()`**：看 `RemainArgs()` 会被 `AddArg` 吃掉子命令名 |
| 重复解析会改变结果 | 再次解析会重新 `Set` 全部选项；`cflag` 的 `Strings`/`Ints`/`Booleans` 的 `Set` 是 append 且 `IsRepeatable()==true`（`ext.go:251-253`/`224-229`/`273-278`、`ext.go:260`/`236`/`282`），`KVString` 是合并（`ext.go:391-401`），同一参数集解析两次会让它们翻倍 |
| `AfterFlagParse` 触发点 | `doParse` 在 `checkBindOpts` 之后调用（`cflag.go:422-425`）；返回 false 时以 `ErrStopRun` 结束并不再绑定位置参数。每次解析触发一次，所以不能对同一参数集解析两次 |
| 执行尾部 | `CFlags.Parse` 的尾部（`cflag.go:305-314`）：`BeforeRun` 返回 false 即 `return nil`，否则调用 `Func`；组命令复用这一段，不重新解析 |
| 别名表 | `maputil.Aliases`（`maputil/alias.go:10-69`）：`AddAliases(real, aliases)` 冲突即 panic，`ResolveAlias`、`GroupAliases` |
| 参照实现 | `gcli` 已有命令树：`Command.Add/AddCommand`（`gcli/cmd.go:235`、`:249`）、递归分发 `Command.innerDispatch`（`gcli/cmd.go:447-512`）、注册与别名 `base.addCommand`（`gcli/base.go:312`）；`capp` 定位是轻量替代，故不照搬其 category、`Category`、命令 ID 等机制（SR1403） |

适用规范：SR1403、SR1410、GR102、GR104、GR106。

## 总体方案

原则：**把“命令集合”抽成可复用的 `cmdSet`，让 `Cmd` 也能持有一个 `cmdSet`，分发从“查一次表”改成“逐级下钻”，叶子命令沿用今天的执行路径。**

1. **抽出 `cmdSet`（新文件 `capp/cmdset.go`）**：把 `App` 上的 `names/cmds/cmdAs/NameWidth` 四件套搬进 `cmdSet`，并搬入 `addCmd` 与 `findCmd` 的逻辑：

   ```go
   type cmdSet struct {
       names     []string
       cmds      map[string]*Cmd
       cmdAs     maputil.Aliases
       nameWidth int
   }
   func newCmdSet() *cmdSet
   func (s *cmdSet) add(c *Cmd) error          // 原 App.addCmd 逻辑（app.go:131-158）
   func (s *cmdSet) find(name string) (*Cmd, bool)  // 原 App.findCmd 逻辑（app.go:255-264）
   func (s *cmdSet) render(buf *strutil.Buffer, width int)      // 渲染 Commands: 列表（原 app.go:287-303）
   ```

   `App` 用 `cmdSet` 替换四个散字段；导出的 `App.NameWidth` **保持权威**：`App.showHelp` 仍读 `a.NameWidth` 并作为 `width` 传给 `render`（`app.go:297`、`app.go:301`），`cmdSet.nameWidth` 只服务于组命令的子命令列表。这样外部 `app.NameWidth = X` 的既有用法不受影响（评审 S5）。

2. **`Cmd` 增加父子关系与子命令注册（`capp/cmd.go`）**：

   ```go
   type Cmd struct {
       *cflag.CFlags
       init   bool
       app    *App   // 所属 app，可为 nil（独立使用 Cmd 时必须为 nil）
       parent *Cmd   // 所属组命令
       subs   *cmdSet // 惰性创建
       // ... 原字段不变
   }

   func (c *Cmd) AddOrErr(cmds ...*Cmd) error // 校验 + 注册 + 设 parent/subs
   func (c *Cmd) Add(cmds ...*Cmd)            // panic 版，对齐 App.Add
   func (c *Cmd) Sub(name string) (*Cmd, bool)
   func (c *Cmd) HasSubs() bool
   func (c *Cmd) Parent() *Cmd
   func (c *Cmd) Path() string                // "session list"；有 app 时前缀 app.Name
   ```

   `AddOrErr` 复用 `cmdSet.add`，并额外设置 `sub.parent = c`、`sub.app = c.app`；`App.addCmd` 相应设置 `c.app = a`。两种注册入口共用同一份校验、`OnAdd` 与 `initCmd` 逻辑，避免规则分叉。

   评审确认后必须同时处理的两处细节（A1、A4）：

   - **组命令不支持位置参数绑定**：`AddArg` 会绑走第一个位置参数（即子命令名），两者语义冲突。`Cmd` 覆盖同名方法 `AddArg`/`BindArg`（签名不变，转发给 `CFlags` 并记录声明）：声明位置参数时若 `HasSubs()` 为真则 panic；反过来，`Cmd.AddOrErr` 在把命令变成组命令时，若该命令已声明位置参数也 panic（“group command %q cannot declare positional args”）。直接调用 `c.CFlags.AddArg` 会绕过校验，文档与示例声明不支持（D10）。
   - **组帮助的 `HelpFunc` 在子命令注册时安装**，不在 `initCmd` 里按 `HasSubs()` 判断：`initCmd` 在 `App.addCmd` 时已跑过一次（`app.go:152`、`cmd.go:86-89`），那时组命令还没有子命令，`HasSubs()` 必为 false，“`app group -h`”会退回默认帮助。注册子命令是命令“变成组”的唯一时机，且两种注册顺序都能覆盖。

3. **分发改为逐级下钻（`capp/app.go`）**：`RunWithArgs` 里把 `err := cmd.Parse(cmdArgs)`（`app.go:214`）换成 `err := a.runCmd(cmd, cmdArgs)`：

   ```go
   func (a *App) runCmd(c *Cmd, args []string) error {
       if !c.HasSubs() {
           return c.Parse(args) // 叶子：与今天完全一致
       }

       // 组命令：只解析一次自身选项（stop-on-first-arg）
       if err := c.Prepare(); err != nil {
           return err
       }
       if err := c.DoParseStopOnArg(args); err != nil {
           if errors.Is(err, flag.ErrHelp) { c.ShowHelp(); return nil }
           if errors.Is(err, cflag.ErrStopRun) { return nil }
           return err
       }

       posArgs := c.Args() // 不用 RemainArgs()：AddArg 会吃掉子命令名
       if len(posArgs) > 0 && isHelp(posArgs[0]) { // help 是保留名
           c.ShowHelp()
           return nil
       }
       if len(posArgs) > 0 {
           sub, ok := c.subs.find(posArgs[0])
           if !ok { return fmt.Errorf("command %q not exists in %q", posArgs[0], c.Path()) }
           return a.runCmd(sub, posArgs[1:])
       }

       // 无位置参数：执行自身 Func 或显示组帮助
       if c.Func == nil {
           c.ShowHelp()
           return nil
       }
       // 等价于 CFlags.Parse 的执行尾部（cflag.go:305-314），不二次解析
       if c.CFlags.BeforeRun != nil && !c.CFlags.BeforeRun(c.CFlags) { return nil }
       return c.Func(c)
   }
   ```

   关键点：

   - 组命令必须用 `DoParseStopOnArg`（`cflag.go:396`），不能沿用叶子默认的 interspersed 解析：后者会把子命令名**之后**的、父子同名的选项摘到父级（`interspersedArgs`，`cflag.go:438-476`），语义不可预期。由此得到明确规则：**父级选项必须写在子命令名之前**，与 `App` 今天对全局选项的处理一致（`README.zh-CN.md:252-269`）。
   - 组命令分支必须显式 `c.Prepare()`：`DoParse*` 不调用它（`cflag.go:400-428`），而 `shortcuts`/`bindOpts` 是在 `Prepare` 里由 `parseFlagUsage` 填的（`cflag.go:318-340`）。漏掉它，父级短选项与 `required` 校验会静默失效。
   - **组命令只解析一次**（A2、A3）：无位置参数时用 `c.Args()` 与已解析状态直接跑执行尾部，不再调用 `c.Parse(args)`，否则 repeatable 选项会翻倍、`AfterFlagParse` 会触发两次。代价是组命令自身的选项同样要写在位置参数之前（与 D2 同一条规则），并且组命令不接受位置参数（D10）。
   - 子命令名从 `Args()` 取（A1）；`help` 在命令层是保留名，一律显示组帮助。
   - 第一位置参数不是已知子命令时一律报未知子命令，**不做**“回退成组命令位置参数”的处理，避免子命令拼写错误被静默当成参数执行。

   递归实现即天然支持三级；对外只承诺两级（D3）。

4. **组命令帮助**：在 `Cmd.AddOrErr`（命令第一次获得子命令时）设置 `c.CFlags.HelpFunc = c.renderGroupHelp`，且仅在用户未自行设置 `HelpFunc` 时写入（A4）。`CFlags.showHelp` 已经优先走 `HelpFunc`（`cflag.go:618-621`），所以 `--help`、`-h`、以及我们主动调用的 `ShowHelp()` 都路由到组帮助：
   - 描述取 `getDesc()`；
   - Usage 用 `Path()` 拼：`<root> session SUBCOMMAND [--Options...] [...Arguments]`；
   - 选项段复用 `c.RenderOptionsHelp(buf)`（`cflag.go:681`）加一行 `--help, -h`；
   - 子命令段复用 `cmdSet.render`，输出名字（含别名）+ 描述；
   - 末尾提示 `Use "<root> session SUBCOMMAND --help" for about a command`；
   - `LongHelp` 与 `Example` 按叶子同样渲染。

   不去改 `CFlags.Name()`/`FlagSet` 名来实现“全路径显示”：`CFlags.Name()` 取自 `FlagSet.Name()`（`cflag.go:591-594`），改它会牵动命令帮助模板里的 `{{cmd}}` 等替换（`cflag.go:624-631`），收益不抵风险。

5. **顶层帮助**：`App.showHelp` 复用 `cmdSet.render`；带子命令的条目保持一致外形，不加计数标注（已确认）。

6. **钩子语义固定**：`App.BeforeRun`/`AfterRun` 只在**实际执行的命令**上触发一次——叶子命令，或没有子命令名而执行自身 `Func` 的组命令——第二个参数是该命令实际收到的参数子串；`Cmd` 侧不加新钩子。这样单级应用的观察面与今天完全一致。

7. **兼容边界**：`App.Add`/`AddOrErr`/`findCmd` 的对外行为与错误文案不变（未命中仍是 `input not exists command %q`）；`cmd_test.go` 里 `NewCmd` + `MustParse` 的独立用法不变；`App.NameWidth` 仍可外部设置。

## 架构

不新增包，不新增导出类型（`cmdSet` 未导出）。改动集中在三个文件：

- `cflag/capp/cmdset.go`（新增）：`cmdSet` 定义、`add`/`find`/`render`。
- `cflag/capp/cmd.go`：`Cmd` 增加 `app/parent/subs` 字段与 `Add/AddOrErr/Sub/HasSubs/Parent/Path`；覆盖 `AddArg`/`BindArg` 记录位置参数声明（D10）；新增 `renderGroupHelp`，由 `AddOrErr` 在子命令注册时挂到 `HelpFunc`。
- `cflag/capp/app.go`：`App` 的四个散字段收敛为 `cmdSet`；`addCmd` 转调 `cmdSet.add`；新增 `runCmd`；`RunWithArgs` 改为调用 `runCmd`；`showHelp` 的命令列表段改为 `cmdSet.render`。

数据流：

```text
os.Args / RunWithArgs(args)
        │
        ├─ App.prepare：全局选项解析（stop-on-first-arg，App.preRun app.go:223-246）
        │
        └─ 逐级下钻 App.runCmd
             ├─ 叶子 Cmd → Cmd.Parse → cflag.CFlags.Parse → Cmd.Func
             └─ 组 Cmd  → Prepare + DoParseStopOnArg（只解析一次）
                          ├─ Args()[0] 是子命令名 → 递归 runCmd(sub, Args()[1:])
                          ├─ Args()[0] 是 help → renderGroupHelp
                          └─ 无位置参数 → Func 非空则执行 Func，否则 renderGroupHelp
```

对外 API 变化：只新增，不删除、不改签名。

| 符号 | 变化 | 说明 |
|---|---|---|
| `Cmd.Add`、`Cmd.AddOrErr` | 新增 | 与 `App.Add`/`AddOrErr` 同形，便于同一套注册代码复用 |
| `Cmd.Sub`、`Cmd.HasSubs`、`Cmd.Parent`、`Cmd.Path` | 新增 | 查询与帮助渲染所需；`Path` 用于 Usage 与错误信息 |
| `App`、`Cmd` 结构体字段 | 内部重组 | 新增的都是未导出字段；导出的 `App.NameWidth` 等不变 |

## 关键流程

### 分发与选项归属

```text
输入: app --workdir /tmp session -v list --json abc

1. App.preRun 解析全局选项（stop-on-first-arg）
   -> 全局: --workdir /tmp
   -> remainArgs: [session, -v, list, --json, abc]
2. runCmd(session, [-v, list, --json, abc])
   -> session 有子命令 -> Prepare + DoParseStopOnArg
   -> 组选项: -v（-v 若未定义则此处报 flag provided but not defined）
   -> Args(): [list, --json, abc]（不抵扣绑定参数，D10 下组命令也不会有绑定参数）
3. Args()[0]=list 命中子命令 -> runCmd(list, [--json, abc])
   -> list 是叶子 -> Cmd.Parse([--json, abc])（interspersed）
   -> 子命令选项: --json；位置参数: abc
4. App.BeforeRun / Cmd.Func / App.AfterRun 只在 list 上触发一次；session 自己的选项也只解析一次
```

边界与错误：

| 输入 | 结果 |
|---|---|
| `app session`（无位置参数） | `Func != nil` 执行 `Func`（用已解析状态，不二次解析），否则显示组帮助（D6） |
| `app session -h` / `--help` / `help` | 组帮助（前两者是 `FlagSet.Parse` 返回的 `flag.ErrHelp`，后者命中 `isHelp` 保留名） |
| `app session -v`（`-v` 是组选项） | 组选项解析完成且无位置参数：`Func` 非空则执行 `Func`，否则显示组帮助 |
| `app session -x`（`-x` 是组的 required 选项且未给） | 报 `flag option 'x' is required`（`checkBindOpts` 在 `doParse` 内先于帮助渲染，`cflag.go:417-419`）；只有 `-h`/`--help`/`help` 会先显示组帮助 |
| `app session nosuch` | `command "nosuch" not exists in "app session"`；**不会**回退成组命令的位置参数 |
| `app session --json`（`--json` 只属于子命令） | 报 `flag provided but not defined: -json`（父级选项必须在子命令名之前） |
| `app group sub -- a --json` | 叶子的 `--` 语义不变（`ParseStopMark`，`util.go:145-146`、`cflag.go:445-448`）；`app group -- sub` 也能强制把 `sub` 作为组的位置参数交给子命令名判定 |
| `app group1 group2 sub`（三级） | 递归天然支持；非承诺行为 |

### 组帮助输出

```text
Manage sessions

Usage: suag session SUBCOMMAND [--Options...] [...Arguments]
Options:
  --help, -h    Display command help
  -v, --verbose Show verbose output

Commands:
  list, ls    List sessions
  send        Send a message to a session

Use "suag session SUBCOMMAND --help" for about a command
```

（`Usage:` 行与末行提示用 `Path()` 拼接；选项段与子命令段分别复用 `RenderOptionsHelp` 与 `cmdSet.render`。）

### `cmdSet.render` 的呈现规则

沿用 `App.showHelp` 今天的规则（`app.go:287-303`）：名字排序、别名以 `name, alias1, alias2` 展示、描述首字母大写（`strutil.UpperFirst`）、右侧补齐。差异只有一处：宽度由调用方传入——`App.showHelp` 传 `a.NameWidth`（保持外部可覆盖，`app.go:297`、`app.go:301`），组帮助传该 `cmdSet` 自身的 `nameWidth`（初始值沿用 `App` 的 12，`app.go:76`）。

## 安全、数据、运维与回滚

- 无网络、无文件写入、无持久化数据、无 Secret；本改动是纯内存中的参数分发与文本渲染。
- 兼容性风险与对策（括号内为独立评审的发现编号）：
  - **短选项被误判**：组命令分支不做 interspersed 重排，子命令的短选项在组解析阶段不被改写（`ReplaceShorts` 的 `stopOnFirstArg` 回调，`util.go:155-215`）。测试覆盖 `app group -v sub -x val`。
  - **`Prepare` 漏调**：组分支显式调用，测试用父级 `required` 选项与短选项各覆盖一次。
  - **二次解析（A2、A3）**：组命令只解析一次；无位置参数且 `Func != nil` 时直接跑执行尾部。回归测试用组命令上的 repeatable 选项（如 `--names`，append 语义）断言只收到一次值，并用 `AfterFlagParse` 计数断言只触发一次。
  - **组命令吞掉子命令名（A1）**：子命令名从 `Args()` 取而非 `RemainArgs()`；位置参数声明与子命令共存时在注册期 panic（D10），并有回归测试。
  - **组帮助失效（A4）**：`HelpFunc` 在子命令注册时安装；两种注册顺序都测（先 `app.Add(group)` 再 `group.Add(sub)`，以及反向）。
  - **别名冲突**：`cmdSet.add` 保持现有行为——重名 panic（`app.go:137-139` 的等价物）与别名重复 panic（`maputil/alias.go:13-18`）。别名必须在 `Add` 之前设置（`app.go:145-147`），本设计沿用不改；跨级冲突校验不在本期（deferred）。
  - **`Cmd` 独立使用**：`Cmd.Sub` 不依赖 `app`；`Path()` 在 `app == nil` 时退化为从自身开始的名字链。
- 回滚：改动限于 `capp` 三个文件与文档，`git revert` 单个提交即可完全回退；无迁移、无外部状态。
- 运维：无。

## 决策

| 编号 | 决策 | 理由 |
|---|---|---|
| D1 | 采用“`Cmd` 挂子命令 + 逐级下钻”主方案；不采用“命令名写成 `group sub` 路径字符串”的轻量备选 | 备选无法给组命令自己的选项、别名与帮助（`app group -h` 只能特判合成），而组级选项正是消费方需要的能力；主方案新增代码约 150–200 行，仍属 SR1403 的最小充分实现 |
| D2 | 父级选项必须写在子命令名之前；组命令用 `DoParseStopOnArg`（`cflag.go:396`） | 与 `App` 现在对全局选项的规则一致（`README.zh-CN.md:252-269`），且避免 interspersed 重排把子命令同名选项摘到父级；规则单一、可解释 |
| D3 | 实现递归支持任意层级，只对两级做文档与示例承诺 | 递归不增加实现复杂度（`runCmd` 自调用）；承诺两级可避免深度带来的帮助与错误文案膨胀 |
| D4 | 新增导出 API：`Cmd.Add/AddOrErr/Sub/HasSubs/Parent/Path`；`App` 既有导出符号签名与语义不变 | 只增不改是外部 Interface 变化的低风险路径；“命令集能力”同时落在 `App` 与 `Cmd` 上才不至于规则分叉 |
| D5 | `App.BeforeRun`/`AfterRun` 只在实际执行的命令上触发一次，参数为它收到的参数 | 保持单级应用今天的观察面；组命令只有真正执行自身 `Func` 时才作为“被执行的命令”出现 |
| D6 | 组命令无位置参数时：`Func != nil` 执行 `Func`，`Func == nil` 显示组帮助；执行复用已解析状态，不再二次解析 | 用户 2026-10-13 决定；组命令既可作为纯组织节点，也可作为可执行命令，由应用侧自行选择，库不引入隐式执行；单次解析避免 repeatable 选项翻倍与钩子重复触发（A2、A3） |
| D7 | 不修改 `cflag` 包 | 组分发只需要 `Prepare`/`DoParseStopOnArg`/`Args`/`ShowHelp`/`HelpFunc`/`ErrStopRun`/`AfterFlagParse`，均已导出且语义如述（评审逐项核对通过）；改解析器会扩大回归面（SR1410） |
| D8 | 别名按级隔离，不做跨级冲突校验 | 各级 `cmdSet` 自带别名表；跨级校验需要新规则与新错误文案，交由待确认事项 |
| D9 | `cmdSet` 与 `renderGroupHelp` 均为未导出 | 本次目标是能力而非可扩展点；先不做抽象，待出现第二个真实复用场景再提升（SR1405.1） |
| D10 | 组命令不支持位置参数绑定：`Cmd` 覆盖 `AddArg`/`BindArg` 记录声明，与子命令共存时注册期 panic | 组命令的第一个位置参数就是子命令名，而 `bindParsedArgs` 会把它绑给第一个 `AddArg`（`cflag.go:552`），静默失效比报错更贵；直接调用 `c.CFlags.AddArg` 绕过校验，文档声明不支持 |
| D11 | 组帮助的 `HelpFunc` 在子命令注册时安装，且不覆盖用户已设的 `HelpFunc` | `initCmd` 只在 `App.addCmd` 时跑一次（`app.go:152`、`cmd.go:86-89`），那时还没有子命令，`HasSubs()` 必为 false（评审 A4） |

## 待确认事项

用户已于 2026-10-13 就初稿的四项核心选择给出决定，本版已按其结果改写，不再作为开放项：

1. **组命令无子名时的行为**：`Func != nil` 执行 `Func`，否则显示组帮助（D6）。
2. **`app help <cmd> [sub]` 形式**：本期不做。现有 `isHelp` 只识别首位 `help`/`--help`/`-h`（`app.go:317-319`），按路径显示任意命令帮助留待后续按需增加。
3. **顶层帮助的分组标注**：不加 `(N commands)` 之类标注，保持一致外形。
4. **层级承诺**：实现递归，README 与示例只承诺两级（D3）。

独立评审（2026-10-13，一轮合并轴）结论为“需修订”，4 条 `CORE_BLOCKING` 已在 0.3 全部处理：A1 组命令 `AddArg` 吞子命令名、A2 repeatable 选项二次解析翻倍、A3 `AfterFlagParse` 二次触发、A4 组帮助安装时机导致失效；处理方式见“总体方案”与“安全、数据、运维与回滚”。其余 5 条 `SUGGESTION` 已一并合入：S5 宽度同步、S6 required 与帮助的先后、S7 别名时序、S8 边界表措辞、S9 测试清单。

尚未决定、留待后续的事项：

- `Sub(name string) (*Cmd, bool)` 是否再加一个 panic 版 `MustSub`；按需在实施中发现。
- `_example/app.go` 是否拆一个独立的分组示例文件；按示例可读性决定。
- 跨级命令名/别名冲突校验，以及 `app help <cmd> <sub>` 帮助入口：前者影响注册期报错文案，后者影响帮助入口，两者都超出本次范围。
- 发现但不碰：`App.showHelp` 内 `sort.Strings(a.names)` 会就地排序内部切片（`app.go:288`），本次迁移到 `cmdSet.render` 时保持原行为，不额外改动；`cflag` 的 `ConfigOpt` 在选项不存在时 panic（`cflag.go:159-161`），与本次无关。

## 结论与人工计划 Gate

设计建议：把 `App` 的命令集合抽成 `cmdSet` 并让 `Cmd` 也能持有它；`Cmd` 新增 `Add/AddOrErr/Sub/HasSubs/Parent/Path`；`App.runCmd` 逐级下钻，组命令用 `DoParseStopOnArg` + 显式 `Prepare` **只解析一次**，子命令名从 `Args()` 取，无位置参数时按 `Func` 是否为空决定执行还是显示组帮助；组帮助的 `HelpFunc` 在子命令注册时安装；叶子命令执行路径完全不变；`cflag` 与消费方不动。对外 API 只增不改，回退为单提交 revert。

Gate：本设计（Draft 0.3，已过一轮独立评审）待用户批准后，才编写实施计划（`docs/plans/2026-10-13-capp-subcommands-plan.md` 级别）；计划经一轮评审与用户批准后才解除实施 Gate。批准设计不授权实施、不授权提交到远端、不授权改动消费方项目。
