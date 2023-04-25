(function (global, factory) {
    typeof exports === 'object' && typeof module !== 'undefined' ? factory(exports) :
    typeof define === 'function' && define.amd ? define(['exports'], factory) :
    (global = typeof globalThis !== 'undefined' ? globalThis : global || self, factory(global.MyModule = {}));
})(this, (function (exports) { 'use strict';

    function add(n1, n2) {
        return n1 + n2;
    }
    function mul(n1, n2) {
        return n1 * n2;
    }

    var User = /** @class */ (function () {
        function User(name) {
            this.name = name;
        }
        return User;
    }());
    function call() {
        console.log("caller is invoked");
    }

    function doOutput() {
        console.log(add(10, 20));
        console.log(mul(10, 20));
        var user = new User("rhie");
        console.log(user.name);
        call();
    }

    exports.doOutput = doOutput;

}));
