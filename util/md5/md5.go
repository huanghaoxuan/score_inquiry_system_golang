package md5

import (
	"crypto/md5"
	"encoding/hex"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/6/25 22:19
 * @Version 1.0
	生成md5
	工具类
*/

func GeneratedMD5(string string) string {
	ctx := md5.New()
	ctx.Write([]byte(string))
	return hex.EncodeToString(ctx.Sum(nil))
}
