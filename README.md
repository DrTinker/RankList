# go-redis 实现简易排行榜
功能不断完善中

**调用示例**
```go
func main() {
	rank := ranklist.InitRankList("127.0.0.1:6379", "", "salsry")
	cnt, err := rank.AddElementBatch(
		[]*models.RankElement{
			{
				Member: "jack",
				Score:  5000,
			},
			{
				Member: "tom",
				Score:  3200,
			},
			{
				Member: "lucy",
				Score:  7800,
			},
			{
				Member: "amy",
				Score:  2211,
			},
		},
	)
	fmt.Printf("cnt: %v, err: %v\n", cnt, err)

	res, _ := rank.GetTopRank(3)
	fmt.Printf("res1: %v\n", res)
	res, _ = rank.GetRankBetween(0, 1)
	fmt.Printf("res2: %v\n", res)
	rank.IncrScore(&models.RankElement{Member: "amy", Score: 1000})
	res, _ = rank.GetTopRank(3)
	fmt.Printf("res3: %v\n", res)
	cnt, err = rank.RemoveElementBatch([]*models.RankElement{
		{
			Member: "lucy",
		},
		{
			Member: "amy",
		},
	})
	fmt.Printf("cnt: %v, err: %v\n", cnt, err)
	res, _ = rank.GetTopRank(3)
	fmt.Printf("res4: %v\n", res)
}
```