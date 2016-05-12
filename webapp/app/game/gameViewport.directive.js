(function () {
    "use strict";

    angular
        .module("game")
        .directive("gameViewport", GameViewportDirective);

    GameViewportDirective.$inject = ["$document","GameService","Actions"];

    function GameViewportDirective($document, GameService, Actions) {

        return {
            restrict: "A",
            link: link
        };

        function link($scope, $element) {
            GameService.registerCanvas($element[0]);

            $document.on("keydown", keyDownHandler);
            $document.on("keyup", keyUpHandler);

            $scope.$on("$destroy",function () {
                $document.off(keyDownHandler);
                $document.off(keyUpHandler);
            })
        }

        function keyDownHandler(event) {

            var keyCode = event.keyCode;

            if (keyCode === 38) {
                GameService.setPlayerDirection(Actions.PlayerUp);
                event.preventDefault();
            }
            else if (keyCode === 40) {
                GameService.setPlayerDirection(Actions.PlayerDown);
                event.preventDefault();
            }

        }

        function keyUpHandler(event) {

            var keyCode = event.keyCode;

            if (keyCode === 38 || keyCode === 40) {
                GameService.setPlayerDirection(null);
                event.preventDefault();
            }

        }
    }
})();
