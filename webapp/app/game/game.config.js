(function () {
    'use strict';

    angular
        .module("game")
        .config(ModuleConfig);

    ModuleConfig.$inject = ["$stateProvider"];
    
    function ModuleConfig($stateProvider) {
        
        $stateProvider
            .state('game', {
                url: "/game",
                templateUrl: "app/game/game.tmpl.html",
                controller: "GameController as vm"
            });
    }
})();