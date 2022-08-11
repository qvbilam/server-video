package api

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
	proto "video/api/v1"
)

var Client proto.VideoClient

func initClient() {
	port := 9802
	conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", port), grpc.WithInsecure())
	if err != nil {
		panic(any(err))
	}

	Client = proto.NewVideoClient(conn)
}

// 创建视频
func TestVideoServer_Create(t *testing.T) {
	initClient()

	response, err := Client.Create(context.Background(), &proto.UpdateVideoRequest{
		UserId:         1,
		RegionId:       1,
		CategoryId:     1,
		Name:           "测试视频",
		Introduction:   "测试视频简介",
		Icon:           "测试视频图标",
		HorizontalIcon: "测试视频纵版图标",
		Count:          1,
		TotalCount:     1,
	})

	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(response)
}

// 获取视频
func TestVideoServer_Get(t *testing.T) {
	initClient()
	response, err := Client.Get(context.Background(), &proto.GetVideoRequest{
		Search: nil,
		Page:   nil,
	})
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(response)
}
