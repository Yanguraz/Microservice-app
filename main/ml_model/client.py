import grpc
from ml_model import ml_model_pb2 as pb2, ml_model_pb2_grpc as pb2_grpc
import os


class PredictClient_ml():
    def __init__(self):
        self.host = os.environ.get("ML_SERVICE_HOST", "localhost")
        self.server_port = int(os.environ.get("ML_SERVICE_PORT", "8080"))

        self.channel = grpc.insecure_channel(f'{self.host}:{self.server_port}')
        self.stub = pb2_grpc.Ml_ServiceStub(self.channel)

    def make_predict(self, sentence):
        request = pb2.MakePredictRequest(sentence=sentence)
        return self.stub.MakePredict(request)