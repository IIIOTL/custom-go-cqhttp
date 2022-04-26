// Package silk Silk编码核心模块
package silk

import (
	"github.com/IIIOTL/custom-go-cqhttp/utils/base"
)

func init() {
	base.EncodeSilk = encode
	base.ResampleSilk = resample
}
