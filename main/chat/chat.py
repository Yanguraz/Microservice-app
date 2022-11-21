import json
from flask import Blueprint, request, render_template

from middleware.middleware import token_required
from ml_model.client import PredictClient_ml

chat_blueprint = Blueprint('chat_blueprint', __name__)

chat_blueprint.secret_key = "superman"


@chat_blueprint.route('/', methods=["GET", "POST"])
@token_required
def home():
    return render_template('index.html')

@chat_blueprint.route('/chatbot', methods=['POST'])
@token_required
def chatbot_response():
    message = request.form['msg']

    ml_client = PredictClient_ml()
    response = ml_client.make_response(sentence = message)
        
    return str(response)[10:][1:-2]
