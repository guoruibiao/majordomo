package scripts

import (
	"github.com/guoruibiao/majordomo/library/clipboard"
	"github.com/guoruibiao/majordomo/models/data/elasticsearch"
	"time"
)

const (
	CLIPBOARD_LISTEN_INTERVAL = time.Second * 2
)

var (
	clipboardStr = ""
	darwinClipboard = clipboard.NewDarwin()
)

/**
 * 监听系统剪切板变化，并发送到 elasticsearch
 */
func StartClipboardListening() {
	go clipboardListener(elasticsearch.AddLog)
}

func clipboardListener(callback func(content string) error) {
	ticker := time.NewTicker(CLIPBOARD_LISTEN_INTERVAL)
	for {
		select {
		case <- ticker.C:
			if newClipboardStr, _ := darwinClipboard.Get(); clipboardStr != newClipboardStr {
				// 回调
				callback(newClipboardStr)
				clipboardStr = newClipboardStr
			}
		}
	}
}