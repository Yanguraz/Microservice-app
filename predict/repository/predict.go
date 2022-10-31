package repository

import (
	"database/sql"
	"log"
	"predict/models"

	"github.com/jmoiron/sqlx"
)

type predictRepository struct {
	db *sqlx.DB
}

type PredictRepository interface {
	CreatePredict(predict models.Predict) (bool, error)
	GetPredicts(userID int64) ([]models.Predict, error)
	GetPredict(itemID, userID int64) (models.Predict, error)
	UpdatePredict(predict models.Predict) (bool, error)
	DeletePredict(id int64) (bool, error)
}

func InitPredictRepository(db *sqlx.DB) PredictRepository {
	return &predictRepository{
		db,
	}
}

func (predictRepository *predictRepository) CreatePredict(predict models.Predict) (bool, error) {
	var err error
	var result bool

	tx, errTx := predictRepository.db.Begin()
	if errTx != nil {
		log.Println("Error create predict: ", errTx)
	} else {
		err = insertPredict(tx, predict)
		if err != nil {
			log.Println("Error create predict: ", err)
		}
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		log.Println("Error create predict: ", err)
	}

	return result, err
}

func insertPredict(tx *sql.Tx, predict models.Predict) error {
	_, err := tx.Exec(`
	INSERT INTO predicts (
		sentence,
		intent,
		user_id
	)
	VALUES(
		$1,
		$2,
		$3
	);
	`,
		predict.Sentence,
		predict.Intent,
		predict.UserID,
	)

	return err
}

func (predictRepository *predictRepository) GetPredicts(userID int64) ([]models.Predict, error) {
	var predicts []models.Predict
	rows, err := predictRepository.db.Query(`
		SELECT id, sentence, intent FROM predicts WHERE user_id=$1;
	`, userID)

	if err != nil {
		log.Println("Error get predicts", err)
		return []models.Predict{}, err
	}

	defer rows.Close()
	for rows.Next() {
		var predict models.Predict
		err = rows.Scan(&(predict.ID), &(predict.Sentence), &(predict.Intent))
		if err != nil {
			log.Println("Error get predicts", err)
			return []models.Predict{}, err
		}

		predicts = append(predicts, predict)
	}

	return predicts, nil
}

func (predictRepository *predictRepository) UpdatePredict(predict models.Predict) (bool, error) {
	var err error
	var result bool

	tx, errTx := predictRepository.db.Begin()
	if errTx != nil {
		log.Println("Error update predict: ", errTx)
	} else {
		err = updatePredict(tx, predict)
		if err != nil {
			log.Println("Error update predict: ", err)
		}
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		log.Println("Error update predict: ", err)
	}

	return result, err
}

func updatePredict(tx *sql.Tx, predict models.Predict) error {
	_, err := tx.Exec(`
	UPDATE predicts
	SET
		sentence=$1,
		intent=$2
	WHERE id=$3
	
	`,
		predict.Sentence,
		predict.Intent,
		predict.ID,
	)

	return err
}

func (predictRepository *predictRepository) DeletePredict(id int64) (bool, error) {
	var err error
	var result bool

	tx, errTx := predictRepository.db.Begin()
	if errTx != nil {
		log.Println("Error to delete predict", err)
	} else {
		err = deletePredict(tx, id)
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		log.Println("Error to delete predict", err)
	}

	return result, err
}

func deletePredict(tx *sql.Tx, id int64) error {
	_, err := tx.Exec(`
		DELETE FROM predicts
		WHERE id=$1
	`, id,
	)

	return err
}

func (predictRepository *predictRepository) GetPredict(itemID, userID int64) (models.Predict, error) {
	var predict models.Predict
	row := predictRepository.db.QueryRow(`
		SELECT id, sentence, intent FROM predicts WHERE user_id=$1 AND id=$2;
	`, userID, itemID)

	err := row.Scan(&(predict.ID), &(predict.Sentence), &(predict.Intent))
	if err != nil {
		log.Println("Error get predict", err)
		return models.Predict{}, err
	}

	return predict, nil
}
