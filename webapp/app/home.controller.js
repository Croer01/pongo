(function () {
    'use strict';

    angular
        .module("app")
        .controller("HomeController", HomeController);

    HomeController.$inject = ["$state","$cookies","Restangular"];

    function HomeController($state, $cookies, Restangular) {
        var vm = this;
        
        vm.nick="";
        vm.login = login;
        
        /////////////////

        function login(){
            Restangular.one("login").post(null,{
                nick :vm.nick
            }).then(function (login) {
                $cookies.put("nick",login);
                $state.go("game");
            })
        }
        
    }
})();