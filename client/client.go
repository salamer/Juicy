package client

import (
	"context"
	"log"
	"strconv"

	pb "github.com/salamer/Juicy/commandpb"
	"google.golang.org/grpc"
)

type JuicyClient struct {
	address string
	port    int
	host    string
}

func NewJuicyClient(host string, port int) *JuicyClient {
	return &JuicyClient{
		port:    port,
		host:    host,
		address: host + ":" + strconv.Itoa(port),
	}
}

func (jc *JuicyClient) Set(key string, value string) {

	conn, err := grpc.Dial(jc.address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDBCommandClient(conn)
	r, err := c.CommandRPC(context.Background(), &pb.CommandReq{
		Command: pb.CommandReq_Set,
		Key:     key,
		Value:   value,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("set: %+v\n", r)
}

func (jc *JuicyClient) Get(key string) string {

	conn, err := grpc.Dial(jc.address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDBCommandClient(conn)
	r, err := c.CommandRPC(context.Background(), &pb.CommandReq{
		Command: pb.CommandReq_Get,
		Key:     key,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("get: %+v\n", r)
	return r.Value
}

func (jc *JuicyClient) Persist(filename string) {

	conn, err := grpc.Dial(jc.address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDBCommandClient(conn)
	r, err := c.CommandRPC(context.Background(), &pb.CommandReq{
		Command:  pb.CommandReq_Persist,
		Filename: filename,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("persist: %+v\n", r)
}

func (jc *JuicyClient) Have(key string) bool {

	conn, err := grpc.Dial(jc.address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDBCommandClient(conn)
	r, err := c.CommandRPC(context.Background(), &pb.CommandReq{
		Command: pb.CommandReq_Have,
		Key:     key,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("have: %+v\n", r)
	return r.Have
}

func (jc *JuicyClient) Clear() {

	conn, err := grpc.Dial(jc.address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDBCommandClient(conn)
	r, err := c.CommandRPC(context.Background(), &pb.CommandReq{
		Command: pb.CommandReq_Clear,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("clear: %+v\n", r)
}

func (jc *JuicyClient) Empty() bool {

	conn, err := grpc.Dial(jc.address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDBCommandClient(conn)
	r, err := c.CommandRPC(context.Background(), &pb.CommandReq{
		Command: pb.CommandReq_Empty,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("empty: %+v\n", r)
	return r.Empty
}

func (jc *JuicyClient) Delete(key string) bool {

	conn, err := grpc.Dial(jc.address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDBCommandClient(conn)
	r, err := c.CommandRPC(context.Background(), &pb.CommandReq{
		Command: pb.CommandReq_Delete,
		Key:     key,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("delete: %+v\n", r)
	return r.Success
}
