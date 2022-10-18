package api

import (
	"context"
	"fmt"
	"testing"
	proto "video/api/qvbilam/video/v1"
)

func TestDramaServer_Get(t *testing.T) {
	initClient()

	response, err := DramaClient.Get(context.Background(), &proto.SearchDramaRequest{
		Keyword: "虚假",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v\n", response)
}

func TestDramaServer_Update(t *testing.T) {
	initClient()

	response, err := DramaClient.Update(context.Background(), &proto.UpdateDramaRequest{
		Id:        2,
		PlayCount: 1,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response)
}
