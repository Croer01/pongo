(function () {
    'use strict';

    angular
        .module("game")
        .service("GameResponseHandler", GameResponseHandler);

    GameResponseHandler.$inject = ["Actions"];

    function GameResponseHandler(Actions) {
        var service = this;

        service.handle = handleResponse;
        /////////////////

        function handleResponse(serverResponse, gameData) {
            switch (serverResponse.action) {
                case Actions.GameStart:
                    gameStartHandler(serverResponse, gameData);
                    break;
                case Actions.GameEnd:
                    if (gameData.stop)
                        gameData.stop();
                    break;
                case Actions.PlayerDown:
                case Actions.PlayerUp:
                    var player = gameData.players[0];
                    if (player.idPlayer != serverResponse.idPlayer) {
                        player = gameData.players[1];
                    }

                    player.position = -serverResponse.result;

                    break;
                case Actions.MoveBall:
                    moveBallHandler(serverResponse, gameData);
                    break;
            }
        }

        function gameStartHandler(serverResponse, gameData) {
            gameData.ball = serverResponse.result.ball;
            gameData.players = serverResponse.result.players;
            gameData.start();
        }

        function moveBallHandler(serverResponse, gameData) {
            var ball = serverResponse.result;
            ball.Y = -ball.Y;
            gameData.ball = ball;
        }
    }
})();