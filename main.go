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

const PAGESIZE = 36

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

func getPGitemsAmount() int {
	var amount int
	db, err := sql.Open("postgres", "user=postgres password=binshao123 dbname=google_play_games sslmode=disable")
	if err != nil {
		fmt.Println("连接数据库失败!err: ", err.Error())
		return -1
	}
	defer db.Close()
	var rows *sql.Rows
	rows, err = db.Query("select count(*) from games_games")
	if err != nil {
		fmt.Println("获取数据库总记录数失败!err: ", err.Error())
	}

	for rows.Next() {
		rows.Scan(&amount)
		fmt.Println(amount)
		return amount
	}
	return -1
}

func getPagedGames(pageId int) []*Game {
	//连接pg数据库
	db, err := sql.Open("postgres", "user=postgres password=binshao123 dbname=google_play_games sslmode=disable")
	if err != nil {
		fmt.Println("连接PG数据库失败！err：", err.Error())
		return nil
	}
	defer db.Close()
	fmt.Println("连接PG成功")

	//获取pg里的数据
	var rows *sql.Rows
	s := "select id, name, author, star_rating, download_times, content_rating, introduction, update_time, image from games_games limit " + fmt.Sprint(PAGESIZE) + " offset " + fmt.Sprint(pageId*PAGESIZE)
	rows, err = db.Query(s)
	if err != nil {
		fmt.Println("查询数据失败!err:", err.Error())
		return nil
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
	return games
}

func searchGame(keywords string) []*Game {
	fmt.Println(keywords)
	//连接pg数据库
	db, err := sql.Open("postgres", "user=postgres password=binshao123 dbname=google_play_games sslmode=disable")
	if err != nil {
		fmt.Println("连接PG数据库失败！err：", err.Error())
		return nil
	}
	defer db.Close()
	fmt.Println("连接PG成功")

	//获取pg里的数据
	var rows *sql.Rows
	s := "select id, name, author, star_rating, download_times, content_rating, introduction, update_time, image from games_games where lower(name) ilike regexp_replace(concat('%','" + strings.ToLower(keywords) + "','%'),'\\\\','\\\\\\','g')"
	rows, err = db.Query(s)
	if err != nil {
		fmt.Println("查询数据失败!err:", err.Error())
		return nil
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
	fmt.Println(games)
	return games
}

func getgameDetail(id int) Game {
	//连接pg数据库
	db, err := sql.Open("postgres", "user=postgres password=binshao123 dbname=google_play_games sslmode=disable")
	if err != nil {
		fmt.Println("连接PG数据库失败！err：", err.Error())
		return Game{}
	}
	defer db.Close()
	fmt.Println("连接PG成功")

	game := Game{}
	//获取pg里的数据
	db.QueryRow("select id, name, author, star_rating, download_times, content_rating, introduction, update_time, image from games_games where id = "+fmt.Sprint(id)).Scan(&game.Id, &game.Name, &game.Author, &game.StarRating, &game.DownloadTimes, &game.ContentRating, &game.Introduction, &game.UpdateTime, &game.Image)
	return game
}

func main() {
	amount := getPGitemsAmount()

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

	r.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/games/search")
	})

	r.GET("/games/search", func(context *gin.Context) {
		//获取当前页id
		page := context.Query("page")
		keywords := context.Query("keywords")
		pageId, _ := strconv.Atoi(page)

		//获取分好页的数据
		pagedGames := getPagedGames(pageId)

		//获得搜索关键字得到的游戏
		if keywords != "" {
			pagedGames = searchGame(keywords)
		}

		maxPage := amount / PAGESIZE
		rangeTimes := [5]int{1, 2, 3, 4, 5}

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
		game := getgameDetail(index)
		context.HTML(http.StatusOK, "detail.html", gin.H{
			"game": game,
		})
	})
	r.GET("/download/:index", func(context *gin.Context) {
		index := context.Param("index")
		context.File("/Users/Ben/download_apks/" + index + ".apk")
	})
	r.GET("/nginx-test", func(context *gin.Context) {
		context.String(http.StatusOK, "port：%s", context.GetHeader("X-real-IP"))
	})

	r.Run(":8001")
}
