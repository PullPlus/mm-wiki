/**
 * main
 * Copyright (c) 2018
 */

var sliderLayout = null;

function mainRightHeight() {
    var mainHeight = $(window).height() - $('header').height() - 55;
    $('#mainFrame').height(mainHeight);
}

// webui-popover
function initPopover() {
    $("[data-toggle='web-popover']").webuiPopover({animation: 'pop',autoHide:3000});
    $('[data-toggle="tooltip"]').tooltip();
}

function iniMenu() {
    $('.menu-nav > li > a').click(function() {
        $(".menu-nav > li").each(function () {
            $(this).removeClass('active');
        });
        $(this).parent('li').addClass('active');
    });
}
$(window).resize(function () {
    mainRightHeight();
});

$(window).load(function() {
    mainRightHeight();
    initPopover();
    iniMenu();
});

$(document).ready(function () {
    sliderLayout = $('body').layout({
        west__size:                 230,
        west__spacing_open:		    4,
        west__spacing_closed:		4,
        west__togglerTip_closed:	"<LABEL_1483>",
        west__togglerTip_open:	    "<LABEL_1484>",
        west__resizerTip:	        "<LABEL_1485>",
        west__resizerCursor :       "col-resize",
        west__sliderTip:	        "<LABEL_1483>",
        west__slideTrigger_open:	"click",
        west__slideTrigger_close:	"click",
        center__maskContents:		 true
    });
});

function layoutOpen() {
    sliderLayout.open("west");
}

function layoutClose() {
    sliderLayout.close("west");
}

function layoutToggle() {
    sliderLayout.toggle("west");
}

function hiddenScrollY() {
    // $("#mainFrame").attr("overflow-y", "hidden");
    $("#mainFrame").attr("scrolling", "no");
}

function windowResized() {
    if ($(window).width() < 768) {
        layoutClose()
    }else{
        layoutOpen()
    }
}

$(windowResized);
$(window).resize(windowResized);
