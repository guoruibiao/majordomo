package clipboard

import (
	"fmt"
	"github.com/guoruibiao/commands"
)

type Darwin struct {
}

/**
 * macOS 剪切板工具
 *
 * return: *Darwin
 */
func NewDarwin() *Darwin {
	return &Darwin{}
}

/**
 * 获取 macOS 系统 剪切板内容
 *
 * return: string
 * return: error
 */
func (d *Darwin) Get() (content string, err error) {
	// check pbpaste or pbcopy valid or not.
	if status, output := commands.New().GetStatusOutput(`pbpaste`); !status {
		err = fmt.Errorf("[Darwin] get content from clipboard with %s", err.Error())
	}else{
		content = output
	}
	return
}

/**
 * 将内容写入系统剪切板
 *
 * param: string content
 * return: error
 */
func (d *Darwin) Set(content string) (err error) {
	// commands 设计的遗漏，err 被吞掉了 -_-||
	commands.New().Run(`echo`, ` ` + content + ` | pbcopy`)
	return
}