package service

import (
	"log"
	"net/url"

	"github.com/jeevanshu/warp-url/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal/status"
	"k8s.io/apimachinery/pkg/util/rand"
)

// URLService struct
type URLService struct {
}

// Shorten function to handle urls
func (svc *URLService) Shorten(ctx context.Context, request *pb.ShortenRequest) (*pb.ShortenResponse, error) {
	log.Printf("Recieved URL shorten request")
	u := request.GetURL()
	domainName, err := url.ParseRequestURI(u)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Only absolute url allowed, please check url: %v", request.URL)
	}
	var urlKey string
	if request.Key != "" {
		urlKey = request.Key
	} else {
		urlKey = rand.String(5)
	}

	var uc URLConfiguration
	uc.Key = urlKey
	uc.OrignalURL = domainName
	storeValue(uc)

	return &pb.ShortenResponse{
		MinifyURL: urlKey,
	}, nil
}
