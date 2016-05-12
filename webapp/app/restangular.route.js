(function () {
    'use strict';

    angular
        .module("app")
        .config(ModuleConfig);

    ModuleConfig.$inject = ["RestangularProvider"];
    
    function ModuleConfig(RestangularProvider) {
        RestangularProvider.setBaseUrl('/api');
    }
})();