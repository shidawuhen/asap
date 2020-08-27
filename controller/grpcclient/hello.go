package grpcclient

import (
	"asap/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"time"
	pb "asap/lib/helloworld"
	"log"
	"context"
)
const (
	address     = "localhost:50051"
	defaultName = "world"
	GROUP = "b2c"
	TEAM =  "i18n"
)

func Hello(contextGin *gin.Context)  {
	// Set up a connection to the server.
	serviceName := "/"+GROUP+ "/" + TEAM + "/"
	address := global.GetService(serviceName)
	fmt.Println(address)
	if len(address) == 0 {
		return
	}
	useAddress := ""
	for _, ad := range address {
		useAddress = ad
	}
	fmt.Println(useAddress)

	conn, err := grpc.Dial(useAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
	contextGin.String(http.StatusOK, r.GetMessage())
}
