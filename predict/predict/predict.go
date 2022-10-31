package predict

import (
	context "context"
	"fmt"
	"log"
	"predict/models"
	"predict/usecase"
)

type Server struct {
	predictUsecase usecase.PredictUsecase
	UnimplementedPredictServiceServer
}

func InitServer(predictUsecase usecase.PredictUsecase) Server {
	return Server{
		predictUsecase,
		UnimplementedPredictServiceServer{},
	}
}

func (s *Server) CreatePredict(ctx context.Context, request *CreatePredictRequest) (*CreatePredictResponse, error) {
	log.Println("Create predict")
	predict := models.Predict{
		Sentence: request.Sentence,
		Intent:   request.Intent,
		UserID:   request.UserID,
	}
	_, err := s.predictUsecase.CreatePredict(predict)
	if err != nil {
		log.Println("Error create predict", err)
		return &CreatePredictResponse{Success: false, Message: fmt.Sprintf("%v", err)}, err
	}
	return &CreatePredictResponse{Success: true, Message: "Success to create predict"}, nil
}

func (s *Server) GetPredicts(ctx context.Context, request *GetPredictsRequest) (*GetPredictsResponse, error) {
	log.Println("Get predict")
	predicts, err := s.predictUsecase.GetPredicts(request.UserID)
	if err != nil {
		log.Println("Error create predict", err)
		return &GetPredictsResponse{Success: false, Message: fmt.Sprintf("%v", err)}, err
	}
	return &GetPredictsResponse{Success: true, Message: "Success to get predicts", Data: predicts}, nil
}

func (s *Server) UpdatePredict(ctx context.Context, request *UpdatePredictRequest) (*UpdatePredictResponse, error) {
	log.Println("Update predict")
	predict := models.Predict{
		Sentence: request.Sentence,
		Intent:   request.Intent,
		ID:       request.Id,
	}
	_, err := s.predictUsecase.UpdatePredict(predict)
	if err != nil {
		log.Println("Error update predict", err)
		return &UpdatePredictResponse{Success: false, Message: "Failed to update"}, err
	}

	return &UpdatePredictResponse{Success: true, Message: "Success to update predict"}, err
}

func (s *Server) DeletePredict(ctx context.Context, request *DeletePredictRequest) (*DeletePredictResponse, error) {
	log.Println("Delete predict")
	_, err := s.predictUsecase.DeletePredict(request.Id)
	if err != nil {
		log.Println("Error update predict", err)
		return &DeletePredictResponse{Success: false, Message: "Failed to delete"}, err
	}

	return &DeletePredictResponse{Success: true, Message: "Success to delete predict"}, err
}

func (s *Server) GetPredict(ctx context.Context, request *GetPredictRequest) (*GetPredictResponse, error) {
	log.Println("Get predict")
	predict, err := s.predictUsecase.GetPredict(request.ItemID, request.UserID)
	if err != nil {
		log.Println("Error create predict", err)
		return &GetPredictResponse{Success: false, Message: fmt.Sprintf("%v", err)}, err
	}
	return &GetPredictResponse{Success: true, Message: "Success to get predicts", Data: predict}, nil
}
