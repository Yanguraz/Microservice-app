from flask import Blueprint, request, render_template, redirect, url_for, session

from auth.client import AuthClient

auth_blueprint = Blueprint('auth_blueprint', __name__)

auth_blueprint.secret_key = "superman"

@auth_blueprint.route('/register', methods=['GET','POST'])
def register():
    if request.method == 'POST':
        try:
            if request.form['username'] == '' or request.form['password'] == '':
                raise ValueError('Please provide your username and password')
            auth_client = AuthClient()
            result = auth_client.register(username=request.form['username'],
                                        password=request.form['password'])
            return redirect(url_for('auth_blueprint.login'))
        except Exception as error:
            return render_template('auth/register.html', error=str(error))
    return render_template('auth/register.html')


@auth_blueprint.route('/login', methods=['GET', 'POST'])
def login():
    if request.method == 'POST':
        try:
            if request.form['username'] == '' or request.form['password'] == '':
                raise ValueError('Please provide your username and password')
            auth_client = AuthClient()
            result = auth_client.login(username=request.form['username'],
                                       password=request.form['password'])
            session['token'] = result.token
            return redirect(url_for('predict_blueprint.get_predicts'))
        except Exception as error:
            return render_template('auth/login.html', error=str(error))
    return render_template('auth/login.html')

@auth_blueprint.route('/logout', methods=['GET'])
def logout():
    session.clear()
    return redirect(url_for('auth_blueprint.login'))
