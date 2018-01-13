
from __future__ import print_function

import grpc

import command_pb2
import command_pb2_grpc


class Client(object):

    def __init__(self,host,port):
        self.host = host
        self.port = port

        self.channel = grpc.insecure_channel(host+":"+str(port))
        self.stub = command_pb2_grpc.DBCommandStub(self.channel)

    def __contains__(self, key):
        if not isinstance(key,str):
            raise ValueError
        response = self.stub.CommandRPC(command_pb2.CommandReq(
            command=command_pb2.CommandReq.Have,
            arg1=key,
            ))
        return response.success
    
    def set(self,key,value):
        if not isinstance(key,str) or not isinstance(value,str):
            raise ValueError
        response = self.stub.CommandRPC(command_pb2.CommandReq(
            command=command_pb2.CommandReq.Set,
            arg1=key,
            arg2=value,
            ))
        return response
    
    def get(self,key):
        if not isinstance(key,str) :
            raise ValueError
        response = self.stub.CommandRPC(command_pb2.CommandReq(
                command=command_pb2.CommandReq.Get,
                arg1=key,
            ))
        return response

    def empty(self):
        response = self.stub.CommandRPC(command_pb2.CommandReq(
                command=command_pb2.CommandReq.Empty,
            ))
        return response

    def clear(self):
        response = self.stub.CommandRPC(command_pb2.CommandReq(
                command=command_pb2.CommandReq.Clear,
            ))
        return response

    def delete(self,key):
        if not isinstance(key,str) :
            raise ValueError
        response = self.stub.CommandRPC(command_pb2.CommandReq(
                command=command_pb2.CommandReq.Delete,
                arg1=key,
            ))
        return response

    def persist(self):
        response = self.stub.CommandRPC(command_pb2.CommandReq(
                command=command_pb2.CommandReq.Persist,
                arg1=key,
            ))
        return response
