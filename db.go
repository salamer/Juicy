package Juicy

import (
	"context"
	"fmt"
	"log"
	"net"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
	pb "github.com/salamer/Juicy/commandpb"
	raft "github.com/salamer/naive_raft"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	SINGLE = iota
	DISTRIBUTED
)

type DB struct {
	Tree *rbt.Tree
	name string
	raft *raft.Node
	Mode int
	size int
}

type Node struct {
	next  *Node
	key   string
	value string
}

type RaftConf struct {
	ID       int
	Name     string
	ConfPath string
	Port     int
	Host     string
}

func GetDBFromFile(filename string) *DB {
	//TODO: read from persist file
	return &DB{}
}

func NewDB(name string, mode int, conf RaftConf) *DB {
	if mode == SINGLE {
		return &DB{
			Tree: rbt.NewWithIntComparator(),
			name: name,
			raft: nil,
			Mode: SINGLE,
			size: 0,
		}
	} else {
		return &DB{
			Tree: rbt.NewWithIntComparator(),
			name: name,
			raft: raft.NewNode(conf.Name, conf.ID, conf.Host, conf.Port, conf.ConfPath),
			Mode: DISTRIBUTED,
			size: 0,
		}
	}
}

func NewNode(key string, value string) *Node {
	return &Node{
		key:   key,
		value: value,
	}
}

func (db *DB) GetValue(key string) (string, error) {
	node, r := SafeString(db.Tree.Get(Hash(key)))
	if r != nil {
		return "", KeyError
	} else {
		for node != nil {
			if node.key == key {
				return node.value, nil
			}
			node = node.next
		}
		return node.value, nil
	}
}

func (db *DB) GetNode(key string) (*Node, error) {
	node, r := SafeString(db.Tree.Get(Hash(key)))
	if r != nil {
		return nil, KeyError
	} else {
		return node, nil
	}
}

func (db *DB) SetValue(key string, value string) error {
	node, r := db.GetNode(key)
	fmt.Println(db.Tree)
	if r != nil {
		node = NewNode(key, value)
		db.Tree.Put(Hash(key), node)
		db.size += 1
		return nil
	} else {
		haveKey := false
		for node.next != nil {
			// for the same key
			if node.key == key {
				node.value = value
				haveKey = true
			}
			node = node.next
		}
		if !haveKey {
			node.next = NewNode(key, value)
		}
		db.size += 1
		return nil
	}

}

func (db *DB) Delete(key string) error {
	node, r := db.GetNode(key)
	if r != nil {
		return r
	} else {
		if node.key == key {
			db.Tree.Remove(Hash(node.key))
		} else {
			_node := node.next
			for _node != nil {
				if _node.key == key {
					node.next = _node.next
					return nil
				}
				_node = _node.next
				node = node.next
			}
		}
		db.size -= 1
		return nil
	}
}

func (db *DB) Size() int {
	return db.size
}

func (db *DB) HaveKey(key string) (bool, error) {
	node, r := db.GetNode(key)
	if r != nil && node != nil {
		return false, KeyError
	} else {
		for node != nil {
			if node.key == key {
				return true, nil
			}
		}
		return false, KeyError
	}
}

func (db *DB) Clear() {
	db.Tree.Clear()
}

func (db *DB) Empty() bool {
	if db.Tree.Empty() {
		return true
	} else {
		return false
	}
}

func (db *DB) Start() error {
	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDBCommandServer(s, db)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	if db.Mode == SINGLE {
		return nil
	} else {
		db.raft.Run()
		return nil
	}
}

func (db *DB) CommandRPC(ctx context.Context, in *pb.CommandReq) (*pb.CommandResp, error) {
	switch in.Command {
	case pb.CommandReq_Set:

		err := db.SetValue(in.Arg1, in.Arg2)
		if err != nil {
			return &pb.CommandResp{
				Success: false,
				Error:   err.Error(), //TODO:finish error
			}, nil
		}
		return &pb.CommandResp{
			Success: true,
			Error:   "", //TODO:finish error
		}, nil

	case pb.CommandReq_Get:
		r, err := db.GetValue(in.Arg1)
		if err != nil {
			return &pb.CommandResp{
				Success: false,
				Error:   err.Error(), //TODO:finish error
			}, nil
		} else {
			return &pb.CommandResp{
				Success: true,
				Res2:    r,
			}, nil
		}

	case pb.CommandReq_Have:
		r, err := db.HaveKey(in.Arg1)
		if r && err != nil {
			return &pb.CommandResp{
				Success: false,
				Error:   err.Error(), //TODO:finish error
			}, nil
		} else {
			return &pb.CommandResp{
				Success: true,
				Error:   "",
			}, nil
		}

	case pb.CommandReq_Clear:
		db.Clear()
		return &pb.CommandResp{
			Success: true,
			Error:   "", //TODO:finish error
		}, nil

	case pb.CommandReq_Empty:
		return &pb.CommandResp{
			Success: db.Empty(),
			Error:   "", //TODO:finish error
		}, nil

	case pb.CommandReq_Delete:
		err := db.Delete(in.Arg1)
		if err != nil {
			return &pb.CommandResp{
				Success: false,
				Error:   err.Error(), //TODO:finish error
			}, nil
		} else {
			return &pb.CommandResp{
				Success: true,
				Error:   "", //TODO:finish error
			}, nil
		}

	case pb.CommandReq_Persist:
		// TODB : db.Persist
		return &pb.CommandResp{
			Success: true,
			Error:   "", //TODO:finish error
		}, nil
	}
	return &pb.CommandResp{
		Success: false,
		Error:   MissCommandError.Error(), //TODO:finish error
	}, nil
}
