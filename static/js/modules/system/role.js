/**
 * <LABEL_1634>
 * Copyright (c) 2018 phachon@163.com
 */
var Role = {

	isChecked : true,

	defaults : function(defaults, disableds) {
		console.log(defaults);
		console.log(disableds);
		$('[name="privilege_id"]').each(function() {
			var checked = $.inArray(parseInt(this.value), defaults) > -1 ? true : false;
			this.checked = checked;
            var isDisabled = $.inArray(parseInt(this.value), disableds) > -1 ? true : false;
            if (isDisabled) {
            	this.checked = true;
                this.disabled = true;
            }
		});
	},

	privilege : function(element) {

		var privilegeId = $(element).val();
		var checked = $(element).is(':checked');
		var type = $(element).attr("data-type");
		//console.log(checked);

		if(checked == false) {
			Role.isChecked = false;
		}else {
			Role.isChecked = true;
		}
		if(type == 'navigator') {
			Role.navigator(privilegeId);
		}
		if(type == 'menu') {
			Role.menu(privilegeId);
		}
		if(type == 'controller') {
			Role.controller(privilegeId);
		}
		
	},

	navigator : function(privilegeId) {

		var navigatorIds = [];
		var menuIds = [];
		var controllerIds = [];

		navigatorIds.push(privilegeId);
		//<LABEL_1487>
		for (var i = 0; i < menus.length; i++) {
			if(menus[i].parentId == privilegeId) {
				menuIds.push(menus[i].privilegeId);
			}
		}

		for (var i = 0; i < menuIds.length; i++) {
			//<LABEL_1103>
			for (var y = 0; y < controllers.length; y++) {
				if(controllers[y].parentId == menuIds[i]) {
					controllerIds.push(controllers[y].privilegeId);
				}
			}
		}

		var allIds = [].concat(navigatorIds, menuIds, controllerIds);
		Role.downLevel(allIds);
	},

	menu : function(privilegeId) {
		var navigatorIds = [];
		var menuIds = [];
		var controllerIds = [];

		menuIds.push(privilegeId);
		//<LABEL_1104>
		for (var i = 0; i < controllers.length; i++) {
			if(controllers[i].parentId == privilegeId) {
				controllerIds.push(controllers[i].privilegeId);
			}
		}

		//<LABEL_1488>
		for (var i = 0; i < menus.length; i++) {
			if(menus[i].privilegeId == privilegeId) {
				for (var y = 0; y < navigators.length; y++) {
					if(navigators[y].privilegeId == menus[i].parentId) {
						navigatorIds.push(navigators[y].privilegeId);
						continue
					}
				}
				continue
			}
		}

		var allIds = [].concat(menuIds, controllerIds);
		//<LABEL_1822>
		Role.downLevel(allIds);
		//<LABEL_1653>
		Role.upLevel(navigatorIds);
	},

	controller : function(privilegeId) {
		var navigatorIds = [];
		var menuIds = [];
		var controllerIds = [];
		controllerIds.push(privilegeId);

		//<LABEL_1105>
		for (var i = 0; i < controllers.length; i++) {
			if(controllers[i].privilegeId == privilegeId) {
				for (var y = 0; y < menus.length; y++) {
					if(menus[y].privilegeId == controllers[i].parentId) {
						menuIds.push(menus[y].privilegeId);
						for (var z = 0; z < navigators.length; z++) {
							if(navigators[z].privilegeId == menus[y].parentId) {
								navigatorIds.push(navigators[z].privilegeId);
								continue;
							}
						}
						continue;
					}
				}
				continue;
			}
		}

		//<LABEL_1822>
		Role.downLevel(controllerIds);
		//<LABEL_1653>
		var allIds = [].concat(navigatorIds, menuIds);
		Role.upLevel(allIds);
	},

	/**
	 * <LABEL_1106>
	 */
	downLevel : function(privilegeIds) {
		for (var i = 0; i < privilegeIds.length; i++) {
			if(Role.isChecked == true) {
				$('input:checkbox[value="'+ privilegeIds[i] +'"]').prop('checked', true);
				$('input:checkbox[value="'+ privilegeIds[i] +'"]').attr('checked', true);
			} else {
				$('input:checkbox[value="'+ privilegeIds[i] +'"]').attr('checked', false);
				$('input:checkbox[value="'+ privilegeIds[i] +'"]').prop('checked', false);
				$('input:checkbox[value="'+ privilegeIds[i] +'"]').removeAttr('checked');
			}
		}
	},

	/**
	 * <LABEL_1107>
	 */
	upLevel : function(privilegeIds) {
		for (var i = 0; i < privilegeIds.length; i++) {
			if(Role.isChecked == true) {
				$('input:checkbox[value="'+ privilegeIds[i] +'"]').prop('checked', true);
				$('input:checkbox[value="'+ privilegeIds[i] +'"]').attr('checked', true);
			}
		}
	}

};