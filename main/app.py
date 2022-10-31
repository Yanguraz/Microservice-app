from flask import Flask, render_template
from auth.auth import auth_blueprint
from predict.predict import predict_blueprint
import os

app = Flask(__name__)
app.secret_key = os.urandom(24)


app.register_blueprint(auth_blueprint, url_prefix='/auth')
app.register_blueprint(predict_blueprint, url_prefix='/predict')


@app.route('/', methods=['GET'])
def root_route():
    return render_template('home.html')

if __name__ == '__main__':
    app.run(debug=True)