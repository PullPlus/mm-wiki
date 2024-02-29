(function(){
    var factory = function (exports) {
        var lang = {
            name : "zh-tw",
            description : "<LABEL_1410>Markdown<LABEL_1577><br/>Open source online Markdown editor.",
            tocTitle    : "<LABEL_1775>",
            toolbar     : {
                undo             : "<LABEL_1776>（Ctrl+Z）",
                redo             : "<LABEL_1727>（Ctrl+Y）",
                bold             : "<LABEL_1777>",
                del              : "<LABEL_1578>",
                italic           : "<LABEL_1778>",
                quote            : "<LABEL_1730>",
                ucwords          : "<LABEL_36>",
                uppercase        : "<LABEL_338>",
                lowercase        : "<LABEL_339>",
                h1               : "<LABEL_1779>1",
                h2               : "<LABEL_1779>2",
                h3               : "<LABEL_1779>3",
                h4               : "<LABEL_1779>4",
                h5               : "<LABEL_1779>5",
                h6               : "<LABEL_1779>6",
                "list-ul"        : "<LABEL_1411>",
                "list-ol"        : "<LABEL_1358>",
                hr               : "<LABEL_1732>",
                link             : "<LABEL_1733>",
                "reference-link" : "<LABEL_1412>",
                image            : "<LABEL_1780>",
                code             : "<LABEL_1413>",
                "preformatted-text" : "<LABEL_1075> / <LABEL_1579>（<LABEL_1414>）",
                "code-block"     : "<LABEL_1579>（<LABEL_1076>）",
                table            : "<LABEL_1363>",
                datetime         : "<LABEL_1415>",
                emoji            : "Emoji <LABEL_1734>",
                "html-entities"  : "HTML <LABEL_1416>",
                pagebreak        : "<LABEL_1077>",
                watch            : "<LABEL_925>",
                unwatch          : "<LABEL_926>",
                preview          : "<LABEL_1078>HTML（<LABEL_1859> Shift + ESC <LABEL_1781>）",
                fullscreen       : "<LABEL_1736>（<LABEL_1859> ESC <LABEL_1781>）",
                clear            : "<LABEL_1737>",
                search           : "<LABEL_1782>",
                help             : "<LABEL_1417>",
                info             : "<LABEL_1783>" + exports.title
            },
            buttons : {
                enter  : "<LABEL_1784>",
                cancel : "<LABEL_1688>",
                close  : "<LABEL_1785>"
            },
            dialog : {
                link   : {
                    title    : "<LABEL_1418>",
                    url      : "<LABEL_1419>",
                    urlTitle : "<LABEL_1420>",
                    urlEmpty : "<LABEL_1786>：<LABEL_673>。"
                },
                referenceLink : {
                    title    : "<LABEL_927>",
                    name     : "<LABEL_1421>",
                    url      : "<LABEL_1419>",
                    urlId    : "<LABEL_1787>ID",
                    urlTitle : "<LABEL_1420>",
                    nameEmpty: "<LABEL_1786>：<LABEL_175>。",
                    idEmpty  : "<LABEL_1786>：<LABEL_518>ID。",
                    urlEmpty : "<LABEL_1786>：<LABEL_518>URL<LABEL_1714>。"
                },
                image  : {
                    title    : "<LABEL_1422>",
                    url      : "<LABEL_1423>",
                    link     : "<LABEL_1424>",
                    alt      : "<LABEL_1425>",
                    uploadButton     : "<LABEL_1426>",
                    imageURLEmpty    : "<LABEL_1786>：<LABEL_519>。",
                    uploadFileEmpty  : "<LABEL_1786>：<LABEL_340>！",
                    formatNotAllowed : "<LABEL_1786>：<LABEL_341>，<LABEL_121>："
                },
                preformattedText : {
                    title             : "<LABEL_176>", 
                    emptyAlert        : "<LABEL_1786>：<LABEL_56>。"
                },
                codeBlock : {
                    title             : "<LABEL_1079>",                 
                    selectLabel       : "<LABEL_1427>：",
                    selectDefaultText : "<LABEL_674>",
                    otherLanguage     : "<LABEL_1428>",
                    unselectedLanguageAlert : "<LABEL_1786>：<LABEL_122>。",
                    codeEmptyAlert    : "<LABEL_1786>：<LABEL_675>。"
                },
                htmlEntities : {
                    title : "HTML<LABEL_1416>"
                },
                help : {
                    title : "<LABEL_1417>"
                }
            }
        };
        
        exports.defaults.lang = lang;
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
                var editormd = require("../editormd");
                factory(editormd);
            });
		}
	} 
	else
	{
        factory(window.editormd);
	}
    
})();