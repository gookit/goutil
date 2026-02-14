package strutil

import (
	"hash/crc32"
	"math/rand" // TODO use v2 on 1.22+
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/gookit/goutil/x/basefn"
)

// global id:
//
//	https://github.com/rs/xid
//	https://github.com/satori/go.uuid
var (
	DefMinInt = 1000
	DefMaxInt = 9999
)

// MicroTimeID generate.
//   - return like: 16074145697981929446(len: 20)
//
// Conv Base:
//
//	mtId := MicroTimeID() // eg: 16935349145643425047 len: 20
//	b16id := Base10Conv(mtId, 16) // eg: eb067252154a9d17 len: 16
//	b32id := Base10Conv(mtId, 32) // eg: em1jia8akl78n len: 13
//	b36id := Base10Conv(mtId, 36) // eg: 3ko088phiuoev len: 13
//	b62id := Base10Conv(mtId, 62) // eg: kb24SKgsQ9V len: 11
func MicroTimeID() string { return MTimeBaseID(10) }

// MicroTimeHexID micro time HEX ID generate.
//
// return like: 643d4cec7db9e(len: 13)
func MicroTimeHexID() string { return MTimeHexID() }

// MTimeHexID micro time HEX ID generate.
//
// return like: 643d4cec7db9e(len: 13)
func MTimeHexID() string { return MTimeBaseID(16) }

// MTimeBase36 micro time BASE36 id generate.
func MTimeBase36() string { return MTimeBaseID(36) }

// MTimeBaseID micro time BASE id generate. toBase: 2-36
//
// Examples:
//   - toBase=16: 643d4cec7db9e(len: 13)
//   - toBase=36: hd312z9ka2(len: 10)
func MTimeBaseID(toBase int) string {
	// eg: 1763431181849557
	ms := time.Now().UnixMicro()
	// rand 1000 - 9999
	// ri := mathutil.RandomInt(DefMinInt, DefMaxInt)
	ri := 1000 + rand.Int63n(8999)

	if toBase > 36 {
		return BaseConvInt(uint64(ms)+uint64(ri), toBase)
	}
	return strconv.FormatInt(ms+ri, toBase)
}

// DatetimeNo generate. can use for order-no.
//
//   - No prefix, return like: 2023041410484904074285478388(len: 28)
//   - With prefix, return like: prefix2023041410484904074285478388(len: 28 + len(prefix))
func DatetimeNo(prefix string) string { return DateSN(prefix) }

// DateSN generate date serial number. PREFIX + yyyyMMddHHmmss + ext(微秒+随机数)
func DateSN(prefix string) string {
	nt := time.Now()
	pl := len(prefix)
	bs := make([]byte, 0, 28+pl)
	if pl > 0 {
		bs = append(bs, prefix...)
	}

	// micro datetime
	bs = nt.AppendFormat(bs, "20060102150405.000000")
	bs[14+pl] = '0'

	// host
	name, err := os.Hostname()
	if err != nil {
		name = "default"
	}
	c32 := crc32.ChecksumIEEE([]byte(name)) // eg: 4006367001
	bs = strconv.AppendUint(bs, uint64(c32%99), 10)

	// rand 1000 - 9999
	// rs := rand.New(rand.NewSource(nt.UnixNano()))
	bs = strconv.AppendInt(bs, 1000+rand.Int63n(8999), 10)

	return string(bs)
}

// DateSNOpt 基于时间生成唯一编号
type DateSNOpt struct {
	Layout    string // time layout
	RandMax   int    // rand max
	DateLen   int    // 时间格式长度，后面部分将会进行进制转换 默认 8(yyyyMMdd)
	ConvBase  int    // DateLen 之后的转换 base 2-64. default 36
	EnableSeq bool   // 需要高并发生成时可以启用自增序号。默认不启用
	SeqMaxVal int    // 自增序号最大值，之后后自动重置
	globalSeq int64  // 自增，确保同一时刻生成的编号不重复. EnableSeq=true 时启用
}

// default setting: {时间年到秒14位}
var defOpt = NewDateSNOpt()

// ConfigSNOpt config default date sn option
func ConfigSNOpt(fn func(opt *DateSNOpt)) {
	fn(defOpt)
}

// NewDateSNOpt create a new DateSNOpt instance.
func NewDateSNOpt() *DateSNOpt {
	return &DateSNOpt{
		Layout:    "20060102150405.000000",
		RandMax:   999,
		DateLen:   8,
		ConvBase:  36,
		EnableSeq: false,
		SeqMaxVal: 9999,
	}
}

// prepare for generate
func (do *DateSNOpt) prepare() {
	if do.DateLen <= 0 {
		do.DateLen = 8 // default 8 for yyyyMMdd
	}
	if do.ConvBase <= 0 {
		do.ConvBase = 36
	}
	if do.Layout == "" {
		do.Layout = "20060102150405.000000"
	}
}

// GenSN generate date serial number.
func (do *DateSNOpt) GenSN(prefix string) string {
	pl := len(prefix)
	bs := make([]byte, 0, 22+pl)
	if pl > 0 {
		bs = append(bs, prefix...)
	}
	do.prepare()

	// get time and format
	nt := time.Now()
	bs = nt.AppendFormat(bs, do.Layout)

	// remove the dot separator if exists
	dotIdx := -1
	for i := pl; i < len(bs); i++ {
		if bs[i] == '.' {
			dotIdx = i
			break
		}
	}
	if dotIdx > 0 {
		bs = append(bs[:dotIdx], bs[dotIdx+1:]...)
	}

	// determine date length (default 8 for yyyyMMdd)
	dateLen := do.DateLen
	// build extension part using UnixNano for maximum precision
	extNano := nt.UnixNano()

	// get sequence max value (default 9999)
	seqMax := do.SeqMaxVal
	if seqMax <= 0 {
		seqMax = 9999
	}

	// use atomic sequence for guaranteed uniqueness (even without EnableSeq)
	// this ensures no collisions in tight loops
	seq := atomic.AddInt64(&do.globalSeq, 1)

	// auto reset when seq exceeds SeqMaxVal (thread-safe using CAS)
	if seq > int64(seqMax) {
		// try to reset to 1, other goroutines may have already done it
		atomic.CompareAndSwapInt64(&do.globalSeq, seq, 1)
		seq = seq % int64(seqMax)
		if seq == 0 {
			seq = 1
		}
	}

	var ext int64
	if do.EnableSeq {
		// high concurrency mode: sequence is the main differentiator
		ext = extNano + seq
	} else {
		// normal mode: nanosecond timestamp + sequence offset + random
		// sequence ensures uniqueness, random adds entropy
		randMax := do.RandMax
		if randMax <= 0 {
			randMax = 9999
		}
		ext = extNano + seq*int64(randMax+1) + rand.Int63n(int64(randMax)+1)
	}

	// convert extension to target base
	if do.ConvBase > 36 {
		bs = append(bs[:dateLen+pl], BaseConvInt(uint64(ext), do.ConvBase)...)
	} else {
		bs = append(bs[:dateLen+pl], strconv.FormatInt(ext, do.ConvBase)...)
	}
	return string(bs)
}

// 确保同一时刻生成的编号不重复
// var globalSeqSnV2 int64 = 0

// DateSNv2 generate date serial number.
//   - 2 < extBase <= 36
//   - return: PREFIX + yyyyMMddHHmmss + extBase(6bit micro + 5bit random number)
//
// Example:
//   - prefix=P, extBase=16, return: P2023091414361354b4490(len=22)
//   - prefix=P, extBase=36, return: P202309141436131gw3jg(len=21)
func DateSNv2(prefix string, extBase ...int) string {
	pl := len(prefix)
	bs := make([]byte, 0, 22+pl)
	if pl > 0 {
		bs = append(bs, prefix...)
	}

	// micro datetime
	nt := time.Now()
	bs = nt.AppendFormat(bs, "20060102150405.000000")

	// rand 10000 - 99999
	// rs := rand.New(rand.NewSource(nt.UnixNano()))
	// 6bit micro + 5bit rand
	ext := strconv.AppendInt(bs[16+pl:], 10000+rand.Int63n(89999), 10)

	base := basefn.FirstOr(extBase, 36)
	// prefix + yyyyMMddHHmmss + ext(convert to base)
	bs = append(bs[:14+pl], strconv.FormatInt(SafeInt64(string(ext)), base)...)

	return string(bs)
}

// DateSNv3 generate date serial number.
//   - 2 < extBase <= 64
//   - return: PREFIX + DATETIME(yyyyMMddHHmmss).dateLen + extBase(DATETIME.after+6bit micro + 5bit random number)
//   - dateLen: 为 DATETIME(yyyyMMddHHmmss) 保留的长度，默认为 8(yyyyMMdd) 后面的给 extBase 使用
//
// Example:
//   - prefix=P, dateLen=8, extBase=16, return: P202511139vs99gbifnj len: 20
//   - prefix=P, dateLen=6, extBase=36, return: P2025119yn52qhefati len: 19
//   - prefix=P, dateLen=6, extBase=48, return: P202511k9ksgD1fe6x len: 18
//   - prefix=P, dateLen=4, extBase=62, return: P2025aZl8N0y58M7 len: 16
func DateSNv3(prefix string, dateLen int, extBase ...int) string {
	dso := NewDateSNOpt()
	dso.DateLen = dateLen
	dso.ConvBase = basefn.FirstOr(extBase, 36)
	return dso.GenSN(prefix)
}
