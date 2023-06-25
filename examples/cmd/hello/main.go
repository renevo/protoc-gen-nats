package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/nats-io/nats.go"

	helloworld "github.com/renevo/protoc-gen-nats/examples"
)

func main() {
	nc, err := nats.Connect("127.0.0.1:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	svc, err := helloworld.NewHelloWorldServer(context.Background(), "0.1.0", nc, &helloImpl{})
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = svc.Stop() }()

	log.Printf("%+v\n", svc)

	client := helloworld.NewHelloWorldClient(nc)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := client.Hello(ctx, &helloworld.HelloRequest{Subject: "World"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Text)

	if _, err := client.Err(ctx, &helloworld.EmptyRequest{}); err != nil {
		log.Println(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

type helloImpl struct{}

func (helloImpl) Hello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	return &helloworld.HelloResponse{Text: fmt.Sprintf("Hello, %s!", req.Subject)}, nil
}

func (helloImpl) Echo(ctx context.Context, req *helloworld.EchoRequest) (*helloworld.EchoResponse, error) {
	return &helloworld.EchoResponse{Data: req.Data}, nil
}

func (helloImpl) Err(ctx context.Context, req *helloworld.EmptyRequest) (*helloworld.EmptyResponse, error) {
	return nil, errors.New("broken")
}
