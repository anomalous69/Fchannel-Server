package route

import (
	"github.com/anomalous69/fchannel/activitypub"
	"github.com/anomalous69/fchannel/db"
	"github.com/anomalous69/fchannel/util"
	"github.com/anomalous69/fchannel/webfinger"
)

type PageData struct {
	Title             string
	PreferredUsername string
	Board             webfinger.Board
	Pages             []int
	CurrentPage       int
	TotalPage         int
	Boards            []webfinger.Board
	Posts             []activitypub.ObjectBase
	Key               string
	PostId            string
	Instance          activitypub.Actor
	ReturnTo          string
	NewsItems         []db.NewsItem
	BoardRemainer     []int
	Meta              Meta
	PostType          string

	Themes      *[]string
	ThemeCookie string

	Referer string

	ServerVersion string
}

type AdminPage struct {
	Title         string
	Board         webfinger.Board
	Key           string
	Actor         string
	Boards        []webfinger.Board
	Following     []string
	Followers     []string
	Domain        string
	IsLocal       bool
	PostBlacklist []util.PostBlacklist
	AutoSubscribe bool
	BoardType	  string
	RecentPosts   []activitypub.ObjectBase
	Instance      activitypub.Actor
	Meta          Meta
	ServerVersion string

	Themes      *[]string
	ThemeCookie string
}

type Meta struct {
	Title       string
	Description string
	Url         string
	Preview     string
}

type BanInfo struct {
	Bans []db.Ban
	//Post         activitypub.ObjectBase
}

type errorData struct {
	Message string
	Error   error
}
