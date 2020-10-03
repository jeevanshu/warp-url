package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeevanshu/warp-url/config"
	"github.com/jeevanshu/warp-url/pb"
	"github.com/jeevanshu/warp-url/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var conf *config.Config

func runGRPCServer(urlserver pb.ShortenServiceServer, healthserver pb.HealthServiceServer, listener net.Listener, done chan int) {

	grpcServer := grpc.NewServer()

	pb.RegisterHealthServiceServer(grpcServer, healthserver)
	pb.RegisterShortenServiceServer(grpcServer, urlserver)

	reflection.Register(grpcServer)
	log.Printf("Starting server at 0.0.0.0:8088")
	if err := grpcServer.Serve(listener); err != nil {
		log.Printf("An error occured while creating grpc server: %v", err)
	}

	done <- 1

}

func runHTTPServer(urlserver pb.ShortenServiceServer, done chan int) {

	m := mux.NewRouter()
	m.HandleFunc(
		"/{key}",
		func(rw http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			redirectURL, err := urlserver.FetchURL(context.Background(), &pb.FetchURLRequest{
				MinifyURL: vars["key"]})

			if err != nil {
				rw.Write([]byte(fmt.Sprintf("Could not find url %v, Error: %v", vars["key"], err.Error())))
				return
			}

			http.Redirect(rw, r, redirectURL.GetOriginalURL(), http.StatusSeeOther)
		},
	)

	log.Println("Starting http server")
	http.ListenAndServe(":"+conf.HTTPport, m)

	done <- 1
}

func main() {
	conf = config.New()
	listener, err := net.Listen("tcp", ":"+conf.GRPCport)

	done := make(chan int, 1)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
	defer listener.Close()

	healthserver := service.Service{}
	urlserver := service.URLService{}

	go runGRPCServer(&urlserver, &healthserver, listener, done)
	go runHTTPServer(&urlserver, done)
	<-done

}
