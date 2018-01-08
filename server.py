#!/usr/bin/python
# encoding=utf-8

"""
fake REST Server
"""
import flask
import flask_restful
import time

APP = flask.Flask(__name__)
API = flask_restful.Api(APP, catch_all_404s=True)

class Fake(flask_restful.Resource):
    def get(self):
        """handle GET request
        """
        time.sleep(0.5)
        return {"server response": "it's working"}

    def post(self):
        """handle post request
        """
        data = flask.request.data
        return {"you posted": data.decode()}

def main():
    """ main function
    """

    API.add_resource(Fake, '/fakeapi')
    APP.run(host='0.0.0.0', port=8080, debug=True)

if __name__ == '__main__':
    try:
        main()
    except Exception as error:
        print(error)
