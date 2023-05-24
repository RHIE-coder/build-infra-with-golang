from flask import Flask
from flask import url_for
from flask import render_template
from flask import request
from flask import make_response
from flask import jsonify
import os
import json

template_path = os.path.join(os.path.dirname(__file__), "..", "..", "public")
static_path = os.path.abspath(os.path.join(template_path))
app = Flask(__name__,**dict(
    template_folder=template_path,
    static_folder=static_path,
    static_url_path="/"
))
@app.route("/")
def a():
    return "<p>Hello, World!</p>"

@app.route("/<name>")
def aa(name):
    return name

# string(default) int, float, path(string but accepts slashs),uuid
@app.route("/a/<path:subsub>")
def aaa(subsub):
    # /a/b/c --> b/c
    # /a     --> a
    return subsub

@app.route("/aa/<int:num>")
def aaaa(num):
    return num

@app.route("/test")
def t():
    return "test1"

@app.route("/test/2")
def tt():
    return "test2"

@app.route("/test/3")
def ttt():
    return "test3"

with app.test_request_context():
    print(url_for("t"))
    print(url_for("tt"))
    print(url_for("ttt"))
    # print(url_for("tttt")) # werkzeug.routing.exceptions.BuildError: Could not build url for endpoint 'tttt'

@app.route("/b", methods=["GET"])
def b():
    return "b"

@app.route("/bb")
@app.route("/bbb")
def bb():
    print(dir(request))
    return render_template("index.html") 

with app.test_request_context("/bb", method="GET"):
    print(request.path)
    assert request.path=='/bb'


@app.route("/c")
def c():
    data = dict(
        a = 10,
        b = "hello",
        c = dict(
            score=3.14,
            age=30,
        )
    )
    return jsonify(data)

@app.route("/cc")
def cc():
    class Data:
        def __init__(self):
            self.a=10
            self.b="hhh"
    
    data = Data()
    return jsonify(data.__dict__)


@app.route("/ccc")
def ccc():
    class Data:
        def __init__(self):
            self.a=10
            self.b="hhh"
            self.c=dict(
                name="rhie",
                age=20,
            )
        def sum(self):
            return 10
    
    data = Data()
    print(data.__dict__)
    data = json.dumps(data, indent=4, default=lambda o:o.__dict__)
    print(data)
    return jsonify(data)

@app.route("/cccc")
def cccc():
    class Data:
        def __init__(self):
            self.a=10
            self.b="hhh"
            self.c=dict(
                name="rhie",
                age=20,
            )
        def sum(self):
            return 10
    
    data = Data()
    print(dir(data))
    print(data.__dict__)
    data = json.dumps(data, indent=4, default=lambda o:o.__dict__)
    resp = make_response(data)
    resp.content_type = 'application/json'
    return resp

@app.route("/ccccc")
def ccccc():
    class Data:
        def __init__(self):
            self.a=10
            self.b="hhh"
            self.c=dict(
                name="rhie",
                age=20,
            )
        def sum(self):
            return 10
    
    data = Data()
    print(data.__dict__) #{'a': 10, 'b': 'hhh', 'c': {'name': 'rhie', 'age': 20}}
    print(data.__class__) #<class 'server.flaskr.ccccc.<locals>.Data'>
    print(data.__dir__) #<built-in method __dir__ of Data object at 0x000001BFF019D070>
    print(data.__module__) #server.flaskr
    print(data.__format__) #<built-in method __format__ of Data object at 0x000001BFF019D070>
    print(data.__str__) #<method-wrapper '__str__' of Data object at 0x000001BFF019D070>
    resp = make_response([1, 2, 3, 4, 5])
    resp.headers['content-type'] = 'application/json'
    # resp = jsonify([1, 2, 3, 4, 5])
    return resp

if __name__ == '__main__':
    flaskr.app.run(**dict(
        port=3333
    ))