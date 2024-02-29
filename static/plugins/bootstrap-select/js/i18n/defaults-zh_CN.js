/*!
 * Bootstrap-select v1.12.4 (http://silviomoreto.github.io/bootstrap-select)
 *
 * Copyright 2013-2017 bootstrap-select
 * Licensed under MIT (https://github.com/silviomoreto/bootstrap-select/blob/master/LICENSE)
 */

(function (root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module unless amdModuleId is set
    define(["jquery"], function (a0) {
      return (factory(a0));
    });
  } else if (typeof module === 'object' && module.exports) {
    // Node. Does not work with strict CommonJS, but
    // only CommonJS-like environments that support module.exports,
    // like Node.
    module.exports = factory(require("jquery"));
  } else {
    factory(root["jQuery"]);
  }
}(this, function (jQuery) {

(function ($) {
  $.fn.selectpicker.defaults = {
    noneSelectedText: '<LABEL_676>',
    noneResultsText: '<LABEL_677>',
    countSelectedText: '<LABEL_1794>{1}<LABEL_1795>{0}<LABEL_1874>',
    maxOptionsText: ['<LABEL_1436> (<LABEL_1437>{n}<LABEL_1874>)', '<LABEL_678>(<LABEL_1437>{n}<LABEL_1875>)'],
    multipleSeparator: ', ',
    selectAllText: '<LABEL_1796>',
    deselectAllText: '<LABEL_1438>'
  };
})(jQuery);


}));
