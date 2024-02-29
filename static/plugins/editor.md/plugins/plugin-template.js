/*!
 * Link dialog plugin for Editor.md
 *
 * @file        link-dialog.js
 * @author      pandao
 * @version     1.2.0
 * @updateTime  2015-03-07
 * {@link       https://github.com/pandao/editor.md}
 * @license     MIT
 */

(function() {

    var factory = function (exports) {

		var $            = jQuery;           // if using module loader(Require.js/Sea.js).

		var langs = {
			"zh-cn" : {
				toolbar : {
					table : "<LABEL_1788>"
				},
				dialog : {
					table : {
						title      : "<LABEL_1363>",
						cellsLabel : "<LABEL_1429>",
						alignLabel : "<LABEL_1430>",
						rows       : "<LABEL_1789>",
						cols       : "<LABEL_1790>",
						aligns     : ["<LABEL_1613>", "<LABEL_1580>", "<LABEL_1431>", "<LABEL_1581>"]
					}
				}
			},
			"zh-tw" : {
				toolbar : {
					table : "<LABEL_1363>"
				},
				dialog : {
					table : {
						title      : "<LABEL_1363>",
						cellsLabel : "<LABEL_1432>",
						alignLabel : "<LABEL_1433>",
						rows       : "<LABEL_1791>",
						cols       : "<LABEL_1792>",
						aligns     : ["<LABEL_1793>", "<LABEL_1582>", "<LABEL_1434>", "<LABEL_1583>"]
					}
				}
			},
			"en" : {
				toolbar : {
					table : "Tables"
				},
				dialog : {
					table : {
						title      : "Tables",
						cellsLabel : "Cells",
						alignLabel : "Align",
						rows       : "Rows",
						cols       : "Cols",
						aligns     : ["Default", "Left align", "Center align", "Right align"]
					}
				}
			}
		};

		exports.fn.htmlEntities = function() {
			/*
			var _this       = this; // this == the current instance object of Editor.md
			var lang        = _this.lang;
			var settings    = _this.settings;
			var editor      = this.editor;
			var cursor      = cm.getCursor();
			var selection   = cm.getSelection();
			var classPrefix = this.classPrefix;

			$.extend(true, this.lang, langs[this.lang.name]); // l18n
			this.setToolbar();

			cm.focus();
			*/
			//....
		};

	};
    
	// CommonJS/Node.js
	if (typeof require === "function" && typeof exports === "object" && typeof module === "object")
    { 
        module.exports = factory;
    }
	else if (typeof define === "function")  // AMD/CMD/Sea.js
    {
		if (define.amd) { // for Require.js

			define(["editormd"], function(editormd) {
                factory(editormd);
            });

		} else { // for Sea.js
			define(function(require) {
                var editormd = require("./../../editormd");
                factory(editormd);
            });
		}
	} 
	else
	{
        factory(window.editormd);
	}

})();
