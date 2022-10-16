package api

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
	proto "video/api/qvbilam/video/v1"
)

var Client proto.VideoClient
var DramaClient proto.DramaClient

func initClient() {
	port := 9802
	conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", port), grpc.WithInsecure())
	if err != nil {
		panic(any(err))
	}

	Client = proto.NewVideoClient(conn)
	DramaClient = proto.NewDramaClient(conn)
}

// 创建视频
func TestVideoServer_Create(t *testing.T) {
	initClient()

	response, err := Client.Create(context.Background(), &proto.UpdateVideoRequest{
		UserId:         1,
		CategoryId:     1,
		Name:           "测试视频",
		Introduce:      "测试视频简介",
		Icon:           "测试视频图标",
		HorizontalIcon: "测试视频纵版图标",
	})

	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(response)
}

// 获取视频
func TestVideoServer_Get(t *testing.T) {
	initClient()
	response, err := Client.Get(context.Background(), &proto.SearchVideoRequest{
		Keyword: "进击人",
		Page:    1,
		PerPage: 10,
	})
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(response)
}

func TestVideoServer_Update(t *testing.T) {

	sorts := []string{
		"-play_count",
		"play_count",
	}

	for _, s := range sorts {
		if string(s[0]) == "-" {
			fmt.Printf("sort type: %s;", string(s[0]))
			fmt.Printf("sort field: %s\n", s[0:])
		} else {
			fmt.Printf("sort type: %s;", string(s[0]))
			fmt.Printf("sort field: %s\n", s[0:])
		}
	}

}
