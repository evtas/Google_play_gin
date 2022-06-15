package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

const PAGESIZE = 36

type Game struct {
	Id            int
	Name          string
	Author        string
	StarRating    string
	DownloadTimes string
	ContentRating string
	Introduction  string
	UpdateTime    string
	Image         string
}

func cal(x, y int) (z int) {
	return x + y
}

func cal2(x, y int) (z int) {
	return x + y - 5
}

func cal3(x, y int) int {
	return x + y + 1 - 5
}

func calEq(x, y, z int) bool {
	return y+z-5 == x
}

func main() {
	//连接pg数据库
	db, err := sql.Open("postgres", "user=postgres password=binshao123 dbname=google_play_games sslmode=disable")
	if err != nil {
		fmt.Println("连接PG数据库失败！err：", err.Error())
		return
	}
	defer db.Close()
	fmt.Println("连接PG成功")

	//获取pg里的数据
	var rows *sql.Rows
	rows, err = db.Query("select id, name, author, star_rating, download_times, content_rating, introduction, update_time, image from games_games")
	if err != nil {
		fmt.Println("查询数据失败!err:", err.Error())
		return
	}
	defer rows.Close()
	fmt.Println("查询数据成功!")

	//保存数据到结构体中
	var games []*Game
	for rows.Next() {
		var id int
		var name string
		var author string
		var starRating string
		var downloadTimes string
		var contentRating string
		var introduction string
		var updateTime string
		var image string
		rows.Scan(&id, &name, &author, &starRating, &downloadTimes, &contentRating, &introduction, &updateTime, &image)
		game := &Game{
			Id:            id,
			Name:          name,
			Author:        author,
			StarRating:    starRating,
			DownloadTimes: downloadTimes,
			ContentRating: contentRating,
			Introduction:  introduction,
			UpdateTime:    updateTime,
			Image:         image,
		}
		games = append(games, game)
	}

	//for i := 0; i < len(games); i++ {
	//	fmt.Println(games[i].name)
	//}

	//设置路由器
	r := gin.Default()
	r.Static("/static", "./static")
	r.SetFuncMap(template.FuncMap{
		"cal":   cal,
		"cal2":  cal2,
		"calEq": calEq,
		"cal3":  cal3,
	})
	r.LoadHTMLGlob("templates/*")

	r.GET("/search", func(context *gin.Context) {
		page := context.Query("page")
		keywords := context.Query("keywords")
		pageId, _ := strconv.Atoi(page)

		pagedGames := games[pageId*PAGESIZE : (pageId+1)*PAGESIZE]
		maxPage := len(games) / PAGESIZE
		rangeTimes := [5]int{1, 2, 3, 4, 5}

		if keywords != "" {
			var searchGames []*Game
			for _, game := range games {
				if strings.Contains(strings.ToLower(game.Name), strings.ToLower(keywords)) {
					searchGames = append(searchGames, game)
				}
			}
			if len(searchGames) < PAGESIZE {
				pagedGames = searchGames
			} else {
				pagedGames = searchGames[pageId*PAGESIZE : (pageId+1)*PAGESIZE]
			}
		}
		context.HTML(http.StatusOK, "search.html", gin.H{
			"games":      pagedGames,
			"pageId":     pageId,
			"maxPage":    maxPage,
			"rangeTimes": rangeTimes,
		})
	})
	r.GET("/detail/:idx", func(context *gin.Context) {
		idx := context.Param("idx")
		index, _ := strconv.Atoi(idx)
		var flag int
		for i, g := range games {
			if g.Id == index {
				flag = i
				break
			}
		}
		game := games[flag]
		context.HTML(http.StatusOK, "detail.html", gin.H{
			"game": game,
		})
	})
	r.GET("/download/:index", func(context *gin.Context) {
		index := context.Param("index")
		context.File("/Users/Ben/download_apks/" + index + ".apk")
	})

	r.Run()
}
