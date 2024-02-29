/**
 * Form.js <LABEL_1111>
 * <LABEL_1826> jquery.form.js
 */

var Form = {

    /**
     * <LABEL_1827> div
     */
    failedBox: '#failedBox',

    /**
     * <LABEL_953>
     */
    inPopup: false,

    /**
     * ajax submit
     * @param element
     * @param inPopup
     * @returns {boolean}
     */
    ajaxSubmit: function(element, inPopup) {

        if (inPopup) {
            Form.inPopup = true;
        }

        /**
         * <LABEL_1095>
         * @param message
         * @param data
         */
        function successBox(message, data) {
            Common.successBox(Form.failedBox, message)
        }

        /**
         * <LABEL_1096>
         * @param message
         * @param data
         */
        function failed(message, data) {
            Common.errorBox(Form.failedBox, message)
        }

        /**
         * request success
         * @param result
         */
        function response(result) {
            //console.log(result)
            if (result.code == 0) {
                failed(result.message, result.data);
            }
            if (result.code == 1) {
                successBox(result.message, result.data);
            }
            $("body,html").animate({scrollTop:0},300);
            if (result.redirect.url) {
                var sleepTime = result.redirect.sleep || 3000;
                setTimeout(function() {
                    if (Form.inPopup) {
                        parent.location.href = result.redirect.url;
                    } else {
                        location.href = result.redirect.url;
                    }
                }, sleepTime);
            }
        }

        var options = {
            dataType: 'json',
            success: response
        };

        $(element).ajaxSubmit(options);

        return false;
    }
};