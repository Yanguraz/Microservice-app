from ml_model import ml_model_pb2 as pb2, ml_model_pb2_grpc as pb2_grpc
import logging
from processing import PredictSentence

class Ml_Service(pb2_grpc.Ml_ServiceServicer):
    def __init__(self):
        pass
    def MakePredict(self, request, context):
        logging.info(request.sentence)

        s = PredictSentence()
        predicted = s.predict_model(request.sentence)

        return pb2.MakePredictResponse(intent = predicted)