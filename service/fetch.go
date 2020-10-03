package service

import (
	"fmt"

	"github.com/jeevanshu/warp-url/pb"
	"golang.org/x/net/context"
)

// FetchURL function to retrieve url from requested key
func (svc *URLService) FetchURL(ctx context.Context, request *pb.FetchURLRequest) (*pb.FetchURLResponse, error) {
	val := fetchValue(request.MinifyURL)
	fmt.Println(val)
	return &pb.FetchURLResponse{
		OriginalURL: val,
	}, nil
}
