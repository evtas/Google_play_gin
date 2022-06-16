package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google_play_games/google_play_grpc/google_play"
	"os"
	"strconv"
)

func main() {
	var id int
	id, _ = strconv.Atoi(os.Args[1])

	fmt.Println(id)
	conn, err := grpc.Dial("127.0.0.1:9091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if nil != err {
		fmt.Println(err.Error())
	}
	defer conn.Close()

	c := pb.NewGooglePlayClient(conn)

	gameId := &pb.GameId{
		Id: int32(id),
	}

	response, err2 := c.GetGameDetail(context.Background(), gameId)
	if err2 != nil {
		fmt.Printf("get user error: %s \n", err2)
		return
	}

	fmt.Printf("response msg: %s \n", response)

}
