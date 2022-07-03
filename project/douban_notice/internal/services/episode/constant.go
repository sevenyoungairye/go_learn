package episode

const (
	TotalPage = 50
	PageLimit = 20
)

const (
	Last = 1
	Rank = 2
	Rmd  = 3
)

// 剧集类型
const (
	Movie = "movie"
	Tv    = "tv"
)

func ComputePageData(pageNo int) (pageStart int, pageLimit int) {
	if pageNo <= 0 || pageNo >= TotalPage {
		return 0, PageLimit
	}

	return (pageNo - 1) * PageLimit, PageLimit
}
