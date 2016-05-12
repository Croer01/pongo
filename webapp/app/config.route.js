(function () {
    'use strict';

    angular
        .module("app")
        .config(ModuleConfig);

    ModuleConfig.$inject = ["$stateProvider", "$urlRouterProvider"];
    
    function ModuleConfig($stateProvider, $urlRouterProvider) {
        
        $urlRouterProvider.otherwise("/home");
        
        $stateProvider
            .state('home', {
                url: "/home",
                templateUrl: "app/home.tmpl.html",
                controller: "HomeController as vm"
            });
    }
})();