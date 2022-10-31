import grpc
import os
from . import predict_pb2 as pb2, predict_pb2_grpc as pb2_grpc
from ml_model.client import PredictClient_ml

client = PredictClient_ml()
class PredictClient:
    def __init__(self):
        self.host = os.environ.get("PREDICT_SERVICE_HOST", "localhost")
        self.server_port = int(os.environ.get("PREDICT_SERVICE_PORT", "7000"))

        self.channel = grpc.insecure_channel(f"{self.host}:{self.server_port}")
        self.stub = pb2_grpc.PredictServiceStub(self.channel)

    def create_predict(self, sentence, user_id, intent):
        request = pb2.CreatePredictRequest(sentence=sentence,
                                        intent=intent,
                                        userID=user_id)
        return self.stub.CreatePredict(request)

    def get_predicts(self, user_id):
        request = pb2.GetPredictsRequest(userID=user_id)
        return self.stub.GetPredicts(request)

    def update_predict(self, id, sentence, intent):
        request = pb2.UpdatePredictRequest(id=id,
                                        sentence=sentence,
                                        intent=intent)
        return self.stub.UpdatePredict(request)

    def delete_predict(self, id):
        request = pb2.DeletePredictRequest(id=id)
        return self.stub.DeletePredict(request)

    def get_predict(self, item_id, user_id):
        request = pb2.GetPredictRequest(itemID=item_id, userID=user_id)
        return self.stub.GetPredict(request)