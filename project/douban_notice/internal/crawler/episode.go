package crawler

// SubjectAbstract 豆瓣剧集信息异步接口
// https://movie.douban.com/j/subject_abstract?subject_id=35235813
type SubjectAbstract struct {
	R       int     `fmt:"r"`
	Subject Subject `fmt:"subject"`
}

type Subject struct {
	Actors           []string     `bson:"actors" fmt:"actors"`
	Blacklisted      string       `bson:"blacklisted" fmt:"blacklisted"`
	CollectionStatus string       `bson:"collection_status" fmt:"collection_status"`
	Directors        []string     `bson:"directors" fmt:"directors"`
	Duration         string       `bson:"duration" fmt:"duration"`
	EpisodesCount    string       `bson:"episodes_count" fmt:"episodes_count"`
	ID               string       `fmt:"id"`
	IsTv             bool         `bson:"is_tv" fmt:"is_tv"`
	Playable         bool         `bson:"playable" fmt:"playable"`
	Rate             string       `bson:"rate" fmt:"rate"`
	Region           string       `bson:"region" fmt:"region"`
	ReleaseYear      string       `bson:"release_year" fmt:"release_year"`
	ShortComment     ShortComment `bson:"short_comment" fmt:"short_comment"`
	Star             int          `bson:"star" fmt:"star"`
	Subtype          string       `bson:"subtype" fmt:"subtype"`
	Title            string       `bson:"title" fmt:"title"`
	Types            []string     `bson:"types" fmt:"types"`
	URL              string       `fmt:"url"`
}

type ShortComment struct {
	Author  string `fmt:"author"`
	Content string `fmt:"content"`
}

type Episode struct {
	Cover        string `json:"cover" fmt:"cover"`
	CoverX       int64  `json:"cover_x" fmt:"cover_x"`
	CoverY       int64  `json:"cover_y" fmt:"cover_y"`
	EpisodesInfo string `json:"episodes_info" fmt:"episodes_info"`
	ID           string `json:"id" fmt:"id"`
	IsNew        bool   `json:"is_new" fmt:"is_new"`
	Playable     bool   `json:"playable" fmt:"playable"`
	Rate         string `json:"rate" fmt:"rate"`
	Title        string `json:"title" fmt:"title"`
	URL          string `json:"url" fmt:"url"`
}

type BaseSubject struct {
	Subjects []Episode `json:"subjects" fmt:"subjects"`
}

/*
func (b *BaseSubject) convert() *[]mongodb.MovieInfo {
	data := make([]mongodb.MovieInfo, 0)
	for _, item := range b.Subjects {
		data = append(data, *item.covert())
	}
	return &data
}

func (e *Episode) covert() *mongodb.MovieInfo {
	return &mongodb.MovieInfo{
		Id:           "",
		MvTagId:      "",
		MvTagInfo:    "",
		DbId:         "",
		Title:        "",
		Url:          "",
		Cover:        "",
		Rate:         "",
		IsNew:        false,
		Playable:     false,
		CoverX:       0,
		CoverY:       0,
		EpisodesInfo: "",
		PublicDate:   time.Time{},
		Torrent:      "",
		DriverUrl:    "",
		Created:      time.Time{},
		Creator:      "",
		Updater:      "",
		UpdateTime:   time.Time{},
	}
}
*/
