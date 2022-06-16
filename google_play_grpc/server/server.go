package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	pb "google_play_games/google_play_grpc/google_play"
	"net"
)

type GameService struct {
	pb.UnimplementedGooglePlayServer
}

func (*GameService) GetGameDetail(ctx context.Context, req *pb.GameId) (*pb.GameDetail, error) {
	fmt.Println(req.Id)
	fmt.Printf("get game %d detail from server\n", req.Id)

	db, err := sql.Open("postgres", "user=postgres password=binshao123 dbname=google_play_games sslmode=disable")
	if err != nil {
		fmt.Println("连接PG失败!err: ", err.Error())
		return nil, nil
	}
	defer db.Close()
	fmt.Println("连接PG成功！")

	var rows *sql.Rows

	s := "select id, name, author, star_rating, download_times, content_rating, introduction, update_time, genre from games_games where id = " + fmt.Sprint(req.Id)
	fmt.Println(s)
	rows, err = db.Query(s)
	if err != nil {
		fmt.Println("查询失败!err:", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var author string
		var starRating string
		var downloadTimes string
		var contentRating string
		var introduction string
		var updateTime string
		var genre string

		rows.Scan(&id, &name, &author, &starRating, &downloadTimes, &contentRating, &introduction, &updateTime, &genre)

		return &pb.GameDetail{
			Id:            int32(id),
			Name:          name,
			Author:        author,
			StarRating:    starRating,
			DownloadTimes: downloadTimes,
			ContentRating: contentRating,
			Introduction:  introduction,
			UpdateTime:    updateTime,
			Genre:         genre,
		}, nil
	}
	return nil, nil
}

func NewServer() *GameService {
	return &GameService{}
}

func StartServer() {
	fmt.Println("start...")
	lis, err := net.Listen("tcp", "127.0.0.1:9091")
	if nil != err {
		fmt.Println(err.Error())
		return
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGooglePlayServer(grpcServer, NewServer())

	grpcServer.Serve(lis)
}

func main() {
	StartServer()
}
