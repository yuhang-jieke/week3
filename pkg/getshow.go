package pkg

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func GetShow(address string, selector string) {
	url := launcher.New().
		Headless(false).
		Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0").
		MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()
	page := browser.MustPage(address)
	page.MustWaitLoad()
	if err := page.Timeout(10 * time.Second).MustElement(selector).WaitVisible(); err != nil {
		fmt.Printf("元素等待超时")
		return
	}
	time.Sleep(2 * time.Second)
	element := page.MustElements(selector)
	fmt.Printf("获取到%d个元素\n", len(element))
	fmt.Printf("正在点击元素...\n")
	show := element[0]
	show.MustScrollIntoView()
	wait := page.MustWaitNavigation()
	show.MustClick()
	wait()
}
