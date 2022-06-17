package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "google_play_games/pb"
	"log"
	"net"
)

type GameService struct {
	pb.UnimplementedGamerServer
}

func (*GameService) GetGameDetail(ctx context.Context, req *pb.GameRequest) (*pb.GameResponse, error) {
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

	s := "select id, name, author, star_rating, download_times, content_rating, introduction, update_time, image from games_games where id = " + fmt.Sprint(req.Id)
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
		var image string

		rows.Scan(&id, &name, &author, &starRating, &downloadTimes, &contentRating, &introduction, &updateTime, &image)

		return &pb.GameResponse{
			Id:            int32(id),
			Name:          name,
			Author:        author,
			StarRating:    starRating,
			DownloadTimes: downloadTimes,
			ContentRating: contentRating,
			Introduction:  introduction,
			UpdateTime:    updateTime,
			Image:         image,
		}, nil
	}
	return nil, nil
}

func NewServer() *GameService {
	return &GameService{}
}

func StartServer() {
	fmt.Println("start...")
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if nil != err {
		fmt.Println(err.Error())
		return
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGamerServer(grpcServer, NewServer())

	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}

func main() {
	StartServer()
}
