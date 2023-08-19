// Code generated by Kitex v0.7.0. DO NOT EDIT.
package videoservice

import (
	server "github.com/cloudwego/kitex/server"
	video "github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/video"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler video.VideoService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
