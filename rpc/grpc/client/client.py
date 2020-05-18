import logging

import grpc

import arith_pb2
import arith_pb2_grpc


def run():
    # 注意(gRPC Python Team): .close()方法在channel上是可用的。
    # 并且应该在with语句不符合代码需求的情况下使用。
    req = arith_pb2.ArithRequest(a=9, b=2)
    with grpc.insecure_channel('localhost:2333') as channel:
        stub = arith_pb2_grpc.ArithServiceStub(channel)
        res = stub.Multiply(req)
        print("{0} * {1} = {2}\n".format(req.a, req.b, res.pro))

        res = stub.Divide(req)
        print("{0} / {1}, quo is {2}, rem is {3}\n".format(req.a,
                                                           req.b, res.quo, res.rem))


if __name__ == '__main__':
    logging.basicConfig()
    run()
