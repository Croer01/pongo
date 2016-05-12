(function () {
    'use strict';

    //create autoincrement for easy develop
    function AutoIncrement(base) {
        this.base = base == undefined? 0 : base;
        this.i = undefined;
    }

    AutoIncrement.prototype.next = function () {
        this.i = (this.i == undefined ? this.base : this.i + 1);
        return this.i;
    };

    var clientautoIncrement = new AutoIncrement(0);
    var serverAutoIncrement = new AutoIncrement(1000);
    
    angular
        .module("game")
        .constant("Actions", {
            "PlayerUp": clientautoIncrement.next(),
            "PlayerDown": clientautoIncrement.next(),
            "GameStart": serverAutoIncrement.next(),
            "GameEnd": serverAutoIncrement.next(),
            "MoveBall": serverAutoIncrement.next()
        });
})();