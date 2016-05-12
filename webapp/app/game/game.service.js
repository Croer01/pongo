(function () {
    'use strict';

    angular
        .module("game")
        .service("GameService", GameService);

    GameService.$inject = ["$interval", "GameResponseHandler","$window"];

    function GameService($interval, GameResponseHandler, $window) {
        var service = this;

        var stopRefresh;
        var viewport;
        var width;
        var height;
        var playerDirection;
        var connection;

        service.registerCanvas = registerCanvas;
        service.setPlayerDirection = setPlayerDirection;
        service.gameData = {
            players: [],
            ball: {X: 0, Y: 0},
            start: startGame
        };
        service.gameData.start = service.gameData.start.bind(service.gameData);

        init();
        /////////////////

        function init() {
            var loc = $window.location;
            var new_uri = "ws://" + loc.host + loc.pathname + "ws";
            connection = new WebSocket(new_uri);

            connection.onmessage = function (event) {
                GameResponseHandler.handle(JSON.parse(event.data), service.gameData)
            };
        }

        function render() {
            viewport.clearRect(0, 0, width, height);
            renderPlayer(service.gameData.players[0], 0);
            renderPlayer(service.gameData.players[1], width - 10);

            renderBall(service.gameData.ball);
        }

        function renderPlayer(player, xPosition) {
            viewport.fillStyle = "#0F0";
            viewport.fillRect(xPosition, player.position, 10, 50);
        }

        function renderBall(ballPosition) {
            viewport.beginPath();
            viewport.arc(ballPosition.X, ballPosition.Y, 10, 0, 2 * Math.PI);
            viewport.fillStyle = '#F00';
            viewport.fill();
        }

        function update() {
            updatePlayer();
        }

        function updatePlayer() {
            if (playerDirection != null) {
                var request = {
                    "action": playerDirection
                };

                connection.send(JSON.stringify(request));
            }
        }

        function registerCanvas(canvasElement) {
            viewport = canvasElement.getContext("2d");
            width = canvasElement.width;
            height = canvasElement.height;

            viewport.textAlign = "center";
            viewport.font = "40px arial";
            viewport.fillText("Wait another user ...", canvasElement.width / 2, canvasElement.height / 2);

        }

        function startGame() {
            if (this.stop)
                this.stop();

            this.stop = $interval(function () {
                if (viewport) {
                    update();
                    render();
                }
            }, 1000 / 60);
        }

        function setPlayerDirection(direction) {
            playerDirection = direction;
        }
    }
})();