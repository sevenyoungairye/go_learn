package crawler

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
