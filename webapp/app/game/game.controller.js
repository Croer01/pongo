(function () {
    'use strict';

    angular
        .module("game")
        .controller("GameController", HomeController);

    HomeController.$inject = ["$scope", "GameService"];

    function HomeController($scope, GameService) {
        var vm = this;
        // var connection;
        //
        // vm.sendRequest = sendRequest;
        //
        //
        // init();
        // /////////////////
        // function sendRequest() {
        //     var request = {
        //         "IdPlayer": "1",
        //         "Action": 1
        //     };
        //
        //     connection.send(JSON.stringify(request));
        // }
        // function init() {
        //     connection = new WebSocket('ws://localhost:8080/echo');
        //
        //     connection.onmessage = function (event) {
        //         var serverResponse = JSON.parse(event.data);
        //         GameService.gameData.player = -serverResponse.Result;
        //
        //         $scope.$applyAsync();
        //     };
        // }
    }
})();