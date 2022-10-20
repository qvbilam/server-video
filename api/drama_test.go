package api

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	proto "video/api/qvbilam/video/v1"
)

const defaultUserId = 1
const categoryCartoon = 1
const categoryMovie = 3

const videoTypeAnime = "anime"

func TestDramaServer_Create(t *testing.T) {
	initClient()

	response, err := DramaClient.Create(context.Background(), &proto.UpdateDramaRequest{
		CategoryId: categoryCartoon,
		Name:       "fate stay night",
		Introduce:  "fate第二部哈哈哈哈哈",
		TotalCount: 24,
	})
	if err != nil {
		panic(any(err))
	}
	fmt.Println(response)
}

func TestDramaServer_Update(t *testing.T) {
	initClient()

	response, err := DramaClient.Update(context.Background(), &proto.UpdateDramaRequest{
		Id:        2,
		PlayCount: 1,
		Name:      "fate/zero",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response)
}

func TestDramaServer_Get(t *testing.T) {
	initClient()
	k := "传说中"

	response, err := DramaClient.Get(context.Background(), &proto.SearchDramaRequest{
		Keyword: k,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("struct: %+v\n", response)
	j, _ := json.Marshal(response)
	fmt.Printf("json: %s\n", j)
}
