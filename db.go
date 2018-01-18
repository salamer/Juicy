package Juicy

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	pb "github.com/salamer/Juicy/commandpb"
	rbt "github.com/salamer/RbTree"
	raft "github.com/salamer/naive_raft"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	SINGLE = iota
	DISTRIBUTED
)

type DB struct {
	port int
	host string

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

func NewDB(name string, mode int, conf RaftConf, host string, port int) *DB {
	if mode == SINGLE {
		return &DB{
			Tree: rbt.NewTree(),
			name: name,
			raft: nil,
			Mode: SINGLE,
			size: 0,

			host: host,
			port: port,
		}
	} else {
		return &DB{
			Tree: rbt.NewTree(),
			name: name,
			raft: raft.NewNode(conf.Name, conf.ID, host, port, conf.ConfPath),
			Mode: DISTRIBUTED,
			size: 0,

			host: host,
			port: port,
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
	fmt.Println("node", db.Tree.Find(Hash(key)), Hash(key))
	node, r := SafeString(db.Tree.Find(Hash(key)))
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
	node, r := SafeString(db.Tree.Find(Hash(key)))
	if r != nil {
		return nil, KeyError
	} else {
		return node, nil
	}
}

func (db *DB) SetValue(key string, value string) error {
	node, r := db.GetNode(key)
	if r != nil {
		node = NewNode(key, value)
		fmt.Println("hash", Hash(key))
		db.Tree.Insert(Hash(key), node)
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
			db.Tree.Delete(Hash(node.key))
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
		return true, nil
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

func (db *DB) CommandRPC(ctx context.Context, in *pb.CommandReq) (*pb.CommandResp, error) {
	switch in.Command {
	case pb.CommandReq_Set:

		err := db.SetValue(in.Key, in.Value)
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
		r, err := db.GetValue(in.Key)
		fmt.Println(r)
		if err != nil {
			return &pb.CommandResp{
				Success: false,
				Error:   err.Error(), //TODO:finish error
			}, nil
		} else {
			return &pb.CommandResp{
				Success: true,
				Value:   r,
			}, nil
		}

	case pb.CommandReq_Have:
		r, err := db.HaveKey(in.Key)
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
			Empty:   db.Empty(),
		}, nil

	case pb.CommandReq_Delete:
		err := db.Delete(in.Key)
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
		db.Persist(in.Key)
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

func (db *DB) Start() {
	var s *grpc.Server
	if db.Mode != SINGLE {
		s = db.raft.GetGRPCHandler()
	} else {
		s = grpc.NewServer()
	}
	lis, err := net.Listen("tcp", db.host+":"+strconv.Itoa(db.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	pb.RegisterDBCommandServer(s, db)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed  to serve: %v", err)
	}
}
