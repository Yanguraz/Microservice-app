package usecase

import (
	"encoding/json"
	"log"
	"predict/models"
	"predict/repository"
)

type predictUsecase struct {
	predictRepository repository.PredictRepository
}

type PredictUsecase interface {
	CreatePredict(predict models.Predict) (bool, error)
	GetPredicts(userID int64) (string, error)
	GetPredict(userID, itemID int64) (string, error)
	UpdatePredict(predict models.Predict) (bool, error)
	DeletePredict(id int64) (bool, error)
}

func InitUserUsecase(predictRepository repository.PredictRepository) PredictUsecase {
	return &predictUsecase{
		predictRepository,
	}
}

func (predictUsecase *predictUsecase) CreatePredict(predict models.Predict) (bool, error) {
	_, err := predictUsecase.predictRepository.CreatePredict(predict)
	if err != nil {
		log.Println("Error create predict", err)
		return false, err
	}

	return true, nil
}

func (predictUsecase *predictUsecase) GetPredicts(userID int64) (string, error) {
	predicts, err := predictUsecase.predictRepository.GetPredicts(userID)
	if err != nil {
		log.Println("Error get predicts", err)
		return "", err
	}

	result, err := json.Marshal(predicts)

	return string(result), nil
}

func (predictUsecase *predictUsecase) UpdatePredict(predict models.Predict) (bool, error) {
	_, err := predictUsecase.predictRepository.UpdatePredict(predict)
	if err != nil {
		log.Println("Error update predict", err)
		return false, err
	}

	return true, nil
}

func (predictUsecase *predictUsecase) DeletePredict(id int64) (bool, error) {
	_, err := predictUsecase.predictRepository.DeletePredict(id)
	if err != nil {
		log.Println("Error delete predict", err)
		return false, err
	}

	return true, nil
}

func (predictUsecase *predictUsecase) GetPredict(itemID, userID int64) (string, error) {
	predict, err := predictUsecase.predictRepository.GetPredict(itemID, userID)
	if err != nil {
		log.Println("Error get predicts", err)
		return "", err
	}

	result, err := json.Marshal(predict)

	return string(result), nil
}
