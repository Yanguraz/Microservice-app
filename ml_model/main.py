from ml_model import ml_model_pb2_grpc as pb2_grpc
from ml_model.ml_model import Ml_Service
from concurrent.futures import ThreadPoolExecutor
import grpc

def serve():
    server = grpc.server(ThreadPoolExecutor(max_workers=10))
    pb2_grpc.add_Ml_ServiceServicer_to_server(Ml_Service(), server)
    print('server ready on port 8080')
    server.add_insecure_port('[::]:8080')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()