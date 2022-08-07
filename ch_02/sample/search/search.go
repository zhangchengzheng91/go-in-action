package search

import (
	"log"
	"sync"
)

// 注册用于搜索的匹配器的映射
var matchers = make(map[string]Matcher)

// Run 执行搜索逻辑
func Run(searchTerm string)  {
	// 获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个无缓冲的通道，接收匹配后的结果
	results := make(chan *Result)

	// 构造一个 waitGroup，以便处理所有的数据源
	var waitGroup sync.WaitGroup

	// 设置需要等待处理
	// 每个数据源的 goroutinne 的数量
	waitGroup.Add(len(feeds))

	// 为每一个数据源启动一个 goroutine 来查找结果
	for _, feed := range feeds {
		// 获取一个匹配器用户查找
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// 启动一个 gorutine 来执行搜索
		go func(matcher Matcher, feed * Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(match, feed)
	}

	// 启动一个 gorutine 来监控是否所有呃工作都做完了
	go func() {
		// 等候所有任务完成
		waitGroup.Wait()

		// 用关闭 gorutine 的方式，通知 Display 函数
		// 可以退出程序
		close(results)
	}()


	// 启动函数，显示返回的结果，并且
	// 在最后一个结果显示完后返回
	Display(results)
}
