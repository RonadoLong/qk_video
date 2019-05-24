package lib

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node
const (
	regular = "^1([358][0-9]|4[579]|66|7[0135678]|9[89])[0-9]{8}$"
)

func init() {
	var err error
	snowflake.Epoch = time.Now().Unix()
	node, err = snowflake.NewNode(88)
	if err != nil {
		panic(err)
	}
}

// GetUUID 获取主键ID
func GetUUID() string {
	return fmt.Sprint(node.Generate())
}

// GetRandCode get rand number
func GetRandCode() string{
	rand.Seed(time.Now().Unix())
	code := strconv.Itoa(rand.Intn(1000000))
	return code
}

// CheckMobileNum 手机号码的验证
func CheckMobileNum(mobileNum string) bool {
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func StringJoinString(keys...string) string{
	builder := strings.Builder{}
	for _, val := range keys {
		builder.WriteString(val)
	}
	return  builder.String()
}
