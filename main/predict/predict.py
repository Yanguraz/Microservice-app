import json
from flask import Blueprint, request, g , jsonify, session, render_template, redirect, url_for

from predict.client import PredictClient
from middleware.middleware import token_required
from ml_model.client import PredictClient_ml

predict_blueprint = Blueprint('predict_blueprint', __name__)

predict_blueprint.secret_key = "superman"

@predict_blueprint.route('/create', methods=['GET', 'POST'])
@token_required
def create_predict():
    if request.method == 'POST':
        try:
            sentence = request.form['sentence']
            client1 = PredictClient_ml()
            intent = client1.make_predict(sentence)
            user_data = session['user_data']
            if sentence == '':
                raise ValueError('Please fill the sentence')
            client = PredictClient()
            result = client.create_predict(sentence=sentence,
                                        intent=str(intent),
                                        user_id=user_data.get('ID'))
            return redirect(url_for('predict_blueprint.get_predicts'))
        except Exception as error:
            return render_template('predict/create.html', error=str(error), output = str(intent))
    return render_template('predict/create.html')

@predict_blueprint.route('/get', methods=['GET'])
@token_required
def get_predicts():
    try:
        user_data = session['user_data']
        client = PredictClient()
        result = client.get_predicts(user_id=user_data.get('ID'))
        data = json.loads(result.data)
        return render_template('predict/list.html', predicts=data)
    except Exception as error:
        return render_template('predict/list.html', error=str(error))

@predict_blueprint.route('/update/<int:predict_id>', methods=['POST'])
@token_required
def update_predict(predict_id):
    try:
        user_data = session['user_data']
        client = PredictClient()
        client1 = PredictClient_ml()
        sentence = request.form['sentence']
        intent = client1.make_predict(sentence=sentence)
        result_update = client.update_predict(predict_id, sentence, str(intent))
        result_get = client.get_predict(predict_id, user_data.get('ID'))
        data = json.loads(result_get.data)
        return render_template('predict/item.html', predict=data)
    except Exception as error:
        return render_template('predict/item.html', error=str(error), output = str(intent))


@predict_blueprint.route('/delete/<int:predict_id>', methods=['GET'])
@token_required
def delete_predict(predict_id):
    try:
        client = PredictClient()
        result = client.delete_predict(predict_id)
        return redirect(url_for('predict_blueprint.get_predicts'))
    except Exception as error:
        return render_template('predict/item.html', error=str(error))

@predict_blueprint.route('/get/<int:predict_id>', methods=['GET'])
@token_required
def get_predict(predict_id):
    try:
        user_data = session['user_data']
        client = PredictClient()
        result = client.get_predict(item_id=predict_id, user_id=user_data.get('ID'))
        data = json.loads(result.data)
        return render_template('predict/item.html', predict=data)
    except Exception as error:
        return render_template('predict/item.html', error=str(error))