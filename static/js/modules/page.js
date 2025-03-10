/**
 * Copyright (c) 2018 phachon@163.com
 */

var Page = {

    /**
     * ajax Save
     * @param element
     * @returns {boolean}
     */
    ajaxSave: function (element, sendEmail, isAutoFollow) {

        /**
         * <LABEL_1095>
         * @param message
         * @param data
         */
        function successBox(message, data) {
            Layers.successMsg(message)
        }

        /**
         * <LABEL_1096>
         * @param message
         * @param data
         */
        function failed(message, data) {
            Layers.failedMsg(message)
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
                // remove storage
                var documentId = $("input[name='document_id']").val();
                var storageId = "mm_wiki_doc_"+documentId;
                Storage.remove(storageId);
                successBox(result.message, result.data);
            }
            if (result.redirect.url) {
                var sleepTime = result.redirect.sleep || 3000;
                setTimeout(function () {
                    parent.location.href = result.redirect.url;
                }, sleepTime);
            }
        }

        var containerHtml = '<div class="container-fluid" style="padding: 20px 20px 0 20px">';
        containerHtml += '<textarea name="edit_comment" class="form-control" rows="3" autofocus="autofocus" style="resize:none""></textarea>';
        containerHtml += '<div style="margin-top: 8px;text-align: left">';

        if (sendEmail === "1") {
            containerHtml += '<label>&nbsp;&nbsp;<LABEL_949>&nbsp;</label><input type="checkbox" name="is_notice_user" checked="checked" value="1">';
        }
        if (isAutoFollow !== "1") {
            containerHtml += '<label>&nbsp;&nbsp;<LABEL_1097>&nbsp;</label><input type="checkbox" name="is_follow_doc" checked="checked" value="1">';
        }
        containerHtml += "</div>";
        containerHtml += "</div>";

        layer.open({
            title: '<i class="fa fa-volume-up"></i> <LABEL_695>',
            type: 1,
            area: ['380px', '232px'],
            content: containerHtml,
            btn: ['<LABEL_1738>','<LABEL_1688>'],
            yes: function(index, layero){
                var commentText = $("textarea[name='edit_comment']").val().trim();
                var isNoticeUser = "0";
                var isFollowDoc = "0";
                if (sendEmail === "1") {
                    var isNoticeCheck = $("input[type='checkbox'][name='is_notice_user']").is(':checked');
                    if (isNoticeCheck) {
                        isNoticeUser = "1";
                    }
                }
                if (isAutoFollow !== "1") {
                    var followDocCheck = $("input[type='checkbox'][name='is_follow_doc']").is(':checked');
                    if (followDocCheck) {
                        isFollowDoc = "1";
                    }
                }
                if (commentText.length > 50 ) {
                    layer.tips("<LABEL_1817>50<LABEL_1595>！", $("textarea[name='edit_comment']"))
                } else {
                    layer.close(index);
                    var options = {
                        dataType: 'json',
                        success: response,
                        data: {'comment': commentText, 'is_notice_user': isNoticeUser, 'is_follow_doc': isFollowDoc}
                    };
                    $(element).ajaxSubmit(options);
                }
                // if (commentText && commentText.length > 0) {
                //
                // }else {
                //     $("textarea[name='edit_comment']").focus();
                // }
            },
            btn2: function(index, layero){
                layer.close(index);
            }
        });

        // layer.prompt({
        //     title: '<i class="fa fa-volume-up"></i> <LABEL_695>',
        //     formType: 2,
        //     maxlength: 150,
        //     value: '',
        //     area: ['340px', '80px']
        // }, function(comment, index, elem){
        //     if (comment.trim()) {
        //         layer.close(index);
        //         var options = {
        //             dataType: 'json',
        //             success: response,
        //             data: {'comment': comment}
        //         };
        //         $(element).ajaxSubmit(options);
        //     }else {
        //         elem.focus()
        //     }
        // });

        return false;
    },

    /**
     * cancel save
     */
    cancelSave: function (title, url) {
        title = '<i class="fa fa-volume-up"></i> '+title;
        layer.confirm(title, {
            btn: ['<LABEL_1838>','<LABEL_1839>'],
            skin: Layers.skin,
            btnAlign: 'c',
            title: "<i class='fa fa-warning'></i><strong> <LABEL_1689></strong>"
        }, function() {
            var documentId = $("input[name='document_id']").val();
            var storageId = "mm_wiki_doc_"+documentId;
            Storage.remove(storageId);
            parent.location = url
        }, function() {

        });
    },

    /**
     * upload attachment
     */
    attachment: function (documentId) {
        layer.open({
            type: 2,
            skin: Layers.skin,
            title: '<strong><LABEL_1620></strong>',
            shadeClose: true,
            shade : 0.1,
            resize: false,
            maxmin: false,
            area: ["900px", "500px"],
            content: "/attachment/page?document_id="+documentId,
            padding:"10px"
        });
    },

    /**
     * <LABEL_1476>
     * @param element
     * @param message
     */
    uploadErrorBox: function (element, message) {
        $(element).html('');
        $(element).removeClass('hide');
        $(element).addClass('alert alert-danger');
        $(element).append('<a class="close" href="#" onclick="$(this).parent().hide();">×</a>');
        $(element).append('<strong><i class="glyphicon glyphicon-remove-circle"></i> <LABEL_1477>：</strong>');
        $(element).append(message);
        $(element).show();
    },
};