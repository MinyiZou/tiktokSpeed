package rpc

import (
	"log"

	"github.com/cloudwego/kitex/client"
	"github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/video/videoservice"
)

var VideoClient videoservice.Client

// Video RPC 客户端初始化
func InitVideo() {

	c, err := videoservice.NewClient("Video", client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		panic(err)
	}
	//initialize.GlobalLogger.Println("初始化Video", err)
	log.Println(err)
	VideoClient = c
	log.Println(VideoClient)
}

//在这里调用kitex_gen里面真正业务逻辑的接口

// func GetUserFeed(ctx context.Context, req *video.DouyinFeedRequest) (resp *video.DouyinFeedRequest, err error) {
// 	resp, err = videoClient.Feed(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if resp.StatusCode != 0 {
// 		// return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
// 		initialize.GlobalLogger.Println("获取视频流失败")
// 		return nil, err
// 	}
// 	return resp, nil
// }
