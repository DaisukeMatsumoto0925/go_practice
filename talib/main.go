package main

/*
quoteで株価取得、
talibで株価分析をする

インストール 済
go get "github.com/markcheno/go-quote"
go get "github.com/markcheno/go-talib"
*/

import (
	"fmt"
	//株価取得
	"github.com/markcheno/go-quote"
	//株価分析
	"github.com/markcheno/go-talib"
)

func main() {
	//Quote
	//yahooからSPY株価を取得
	spy, _ := quote.NewQuoteFromYahoo(
		"spy", "2018-04-01", "2020-01-01", quote.Daily, true)
	//"spy", "2018-04-01", "2020-01-01", quote.Min60, true)
	//"spy", "2018-04", "2020-01", quote.Monthly, true)

	fmt.Println(spy.CSV())

	//Talib
	//RSI取得
	//株価の買われすぎ、売られすぎの指標
	// spyの終値過去2日間を基にしたRSIを取得

	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)

	//移動平均線
	mva := talib.Sma(spy.Close, 14)
	fmt.Println(mva)

	//指数平滑移動平均
	ema := talib.Ema(spy.Close, 14)
	fmt.Println(ema)

	//この他、様々なインジケーターを取得できる.(ドキュメント)

}
