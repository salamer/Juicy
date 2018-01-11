package main

import (
	"context"
	"log"

	"github.com/abiosoft/ishell"
	"github.com/fatih/color"
	Juicy "github.com/salamer/Juicy"
	pb "github.com/salamer/Juicy/commandpb"
	"google.golang.org/grpc"
)

const (
	Set = iota
	Get
	Have
	Empty
	Clear
	Persist
)

func sendRPC(command int, arg1 string, arg2 string) (*pb.CommandResp, error) {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDBCommandClient(conn)

	switch command {
	case Set:
		r, err := c.CommandRPC(context.Background(), &pb.CommandReq{
			Command: pb.CommandReq_Set,
			Arg1:    arg1,
			Arg2:    arg2,
		})
		return r, err

	case Get:
		r, err := c.CommandRPC(context.Background(), &pb.CommandReq{
			Command: pb.CommandReq_Get,
			Arg1:    arg1,
		})
		return r, err
	case Have:
		r, err := c.CommandRPC(context.Background(), &pb.CommandReq{
			Command: pb.CommandReq_Have,
			Arg1:    arg1,
		})
		return r, err
	case Empty:
		r, err := c.CommandRPC(context.Background(), &pb.CommandReq{
			Command: pb.CommandReq_Empty,
		})
		return r, err
	case Clear:
		r, err := c.CommandRPC(context.Background(), &pb.CommandReq{
			Command: pb.CommandReq_Clear,
		})
		return r, err
	case Persist:
		r, err := c.CommandRPC(context.Background(), &pb.CommandReq{
			Command: pb.CommandReq_Persist,
		})
		return r, err
	}
	return nil, Juicy.MissCommandError
}

func main() {

	shell := ishell.New()

	shell.Println("Juicy Interactive Shell")

	red := color.New(color.FgRed).SprintFunc()
	ArgWarning := red("arg error,please check 'help' first.")

	shell.AddCmd(&ishell.Cmd{
		Name: "Set",
		Help: "Set {key} {value}",
		Func: func(c *ishell.Context) {
			if len(c.Args) == 2 {
				c.Println(sendRPC(Set, c.Args[0], c.Args[1]))
			} else {
				c.Println(ArgWarning)
			}
		},
	})

	// run shell
	shell.Run()
}
