package main

import "week3/pkg"

func main() {
	//pkg.GetTitle("https://air.1688.com/kapp/channel-fe/cps-4c-pc/sytm?type=1&offerIds=660390230106,574965204819,949033739317", ".offer-title.ellipsis")
	//pkg.GetImage("https://air.1688.com/kapp/channel-fe/cps-4c-pc/sytm?type=1&offerIds=660390230106,574965204819,949033739317", ".offer-item__img")
	pkg.GetShow("https://air.1688.com/kapp/channel-fe/cps-4c-pc/sytm?type=1&offerIds=660390230106,574965204819,949033739317", ".offer-item__img")
}
