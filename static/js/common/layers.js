var Layers = {

	/**
	 * <LABEL_1823>
	 */
	skin : 'default',

	/**
	 * success <LABEL_1109>
	 * @param title
	 */
	success : function (title) {
		layer.alert(title+"<br/>", {
			title: "<LABEL_1489>",
			icon: 1,
			skin: Layers.skin,
			closeBtn: 0
		})
	},

	/**
	 * error <LABEL_1109>
	 * @param title
	 */
	error : function (title) {
		layer.alert(title, {
			title: "<LABEL_1152>",
			icon: 2,
			skin: Layers.skin,
			closeBtn: 0
		})
	},
	
	failedMsg: function (info, callback) {
		var content = '<i class="fa fa-frown-o"></i> ';
		content += info;
		layer.msg(content, callback);
	},

	successMsg: function (info, callback) {
		var content = '<i class="fa fa-smile-o"></i> ';
		content += info;
		layer.msg(content, callback);
	},

	/**
	 * confirm <LABEL_1605>， post <LABEL_1698>
	 * @param title
	 * @param url
	 * @param data
	 */
	confirm: function (title, url, data) {
        title = '<i class="fa fa-volume-up"></i> '+title;
		layer.confirm(title, {
			btn: ['<LABEL_1838>','<LABEL_1839>'],
			skin: Layers.skin,
            btnAlign: 'c',
			title: "<i class='fa fa-warning'></i><strong> <LABEL_1689></strong>"
		}, function() {
			Common.ajaxSubmit(url, data)
		}, function() {

		});
	},

	/**
	 * confirm <LABEL_1605>， post <LABEL_1698>
	 * @param title
	 * @param confirm
	 * @param cancel
	 */
	confirmCallback: function (title, confirm, cancel) {
		title = '<i class="fa fa-volume-up"></i> '+title;
		layer.confirm(title, {
			btn: ['<LABEL_1838>','<LABEL_1839>'],
			skin: Layers.skin,
			btnAlign: 'c',
			title: "<i class='fa fa-warning'></i><strong> <LABEL_1689></strong>"
		}, function() {
			confirm();
		}, function() {
			cancel();
		});
	},

	/**
	 * bind iframe <LABEL_1876>
	 */
	bindIframe: function (element, title, height, width, url) {
		$(element).each(function () {
			height = height||"500px";
			width = width||"1000px";
			$(this).bind('click', function () {
				var content = url || $(this).attr("data-link");
				layer.open({
					type: 2,
					skin: Layers.skin,
					title: '<strong>'+title+'</strong>',
					shadeClose: true,
					shade : 0.6,
					maxmin: true,
					area: [width, height],
					content: content,
					padding:"10px"
				});
			})
		})
	},

	bindPage: function (element, title, height, width, html) {
        $(element).each(function () {
            height = height||"500px";
            width = width||"1000px";
            $(this).bind('click', function () {
                layer.open({
                    type: 1,
                    skin: Layers.skin,
                    title: '<strong>'+title+'</strong>',
                    shadeClose: true,
                    shade : 0.6,
                    maxmin: true,
                    area: [width, height],
                    content: html,
                    padding: "10px"
                });
            })
        })
    }
};