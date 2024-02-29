/**
 * Copyright (c) 2018 phachon@163.com
 * // var treeData = [
 //     { id:10, pId:0, name:"<LABEL_1640>", open:true},
 //     { id:11, pId:0, name:"<LABEL_1596>"},
 //     { id:12, pId:0, name:"<LABEL_1597>"},
 //     { id:13, pId:0, name:"<LABEL_1598>", isParent:true},
 //     { id:111, pId:11, name:"<LABEL_1098>", isParent:true},
 //     { id:112, pId:11, name:"UED<LABEL_1875>", isParent:true},
 //     { id:113, pId:11, name:"<LABEL_1099>", isParent:true},
 //     { id:114, pId:11, name:"<LABEL_1100>", isParent:true},
 //     { id:115, pId:11, name:"<LABEL_1101>", isParent:true},
 //
 //     { id:121, pId:111, name:"<LABEL_1818>CSS<LABEL_1819>"},
 //     { id:122, pId:112, name:"UED <LABEL_1478>"},
 //     { id:123, pId:113, name:"PHP <LABEL_1479>"},
 //     { id:124, pId:114, name:"<LABEL_950>"},
 //     { id:125, pId:115, name:"<LABEL_531>"},
 //
 //     { id:121, pId:12, name:"<LABEL_1599>", isParent:true},
 //     { id:122, pId:12, name:"<LABEL_1600>", isParent:true},
 //     { id:123, pId:12, name:"<LABEL_1601>", isParent:true},
 //     { id:124, pId:12, name:"<LABEL_1602>", isParent:true}
 // ];
 */
var Document = {

    /**
     * list nav
     */
    ListTree: function (element, treeData, defaultId, isEditor, isDelete) {

        //<LABEL_1480>
        var setting = {
            view: {
                showIcon: showIconForTree,
                addHoverDom: addHoverDom,
                removeHoverDom: removeHoverDom
            },
            edit: {
                enable: true,
                showRemoveBtn: true,
                showRenameBtn: false
                // removeTitle: '<LABEL_1664>',
                // renameTitle: '<LABEL_1663>'
            },
            data: {
                simpleData: {
                    enable: true
                }
            },
            callback: {
                beforeClick: beforeClick,
                onClick: onClick,

                beforeEditName: beforeEditName,

                beforeRemove: beforeRemove,
                onRemove: onRemove,

                beforeRename: beforeRename,
                onRename: onRename,

                beforeDrag: beforeDrag,
                onDrag: onDrag,
                beforeDrop: beforeDrop,
                onDrop: onDrop
            },
            drag: {
                isCopy: false,
                isMove: true,
                prev: true,
                inner: true,
                next: true
            }
        };

        if (isDelete == false) {
            setting.edit.showRemoveBtn = false;
        }

        function beforeClick(treeId, treeNode) {
            console.log("<LABEL_347>");
            // $("#mainFrame").attr("src", "/page/view?document_id="+treeNode.id);
            location.href = "/document/index?document_id=" + treeNode.id
        }

        function onClick() {
            console.log("<LABEL_348>");
        }

        function beforeEditName(treeId, treeNode) {
            console.log("<LABEL_349>");
            return true;
        }

        function beforeRename(treeId, treeNode, newName) {
            console.log("<LABEL_350>");
            return true;
        }

        function onRename(e, treeId, treeNode, isCancel) {
            console.log("<LABEL_351>");
            // setTimeout(function() {
            // 	location.href = listUrl;
            // }, 2000);
            return true;
        }

        function beforeRemove(treeId, treeNode) {
            console.log("<LABEL_532>");
            console.log(treeNode);
            if (treeNode.isParent) {
                if (treeNode.children && treeNode.children.length > 0) {
                    Layers.failedMsg("<LABEL_41>！");
                    return false;
                }
            }

            var title = '<i class="fa fa-volume-up"></i> <LABEL_533>？';
            layer.confirm(title, {
                btn: ['<LABEL_1838>', '<LABEL_1839>'],
                skin: Layers.skin,
                btnAlign: 'c',
                title: "<i class='fa fa-warning'></i><strong> <LABEL_1689></strong>"
            }, function () {
                Common.ajaxSubmit("/document/delete?document_id=" + treeNode.id);
                // location.href = "/document/index?document_id="+moveNode.id;
            }, function () {

            });

            return false;
        }

        function onRemove(e, treeId, treeNode) {
            console.log("<LABEL_696>");
            return false;
        }

        function showIconForTree(treeId, treeNode) {
            return true;
        }

        function beforeDrag(treeId, treeNodes) {
            console.log("<LABEL_534>");
            if (isEditor == false) {
                return false;
            }
            // if (treeNodes[0].isParent) {
            //     return false;
            // }
            return true;
        }

        function onDrag() {
            console.log("<LABEL_535>");
            return true;
        }

        function beforeDrop(treeId, treeNodes, targetNode, moveType) {
            console.log("<LABEL_1481>:", treeId, treeNodes, targetNode, moveType);
            console.log("<LABEL_352>");
            if (isEditor == false) {
                return false;
            }
            var moveNode = treeNodes[0];
            // <LABEL_536>
            if (moveType === "prev" || moveType === "next") {
                let moveUrl = "/document/move?move_type=" + moveType
                    + "&document_id=" + moveNode.id
                    + "&target_id=" + targetNode.id;
                Common.ajaxSubmit(moveUrl, moveUrl);
                return false;
            }

            // <LABEL_537>
            if (moveNode.isParent) {
                return false;
            }
            if (!targetNode.isParent) {
                return false;
            }
            
            var title = '<i class="fa fa-volume-up"></i> <LABEL_538>？';
            layer.confirm(title, {
                btn: ['<LABEL_1838>', '<LABEL_1839>'],
                skin: Layers.skin,
                btnAlign: 'c',
                title: "<i class='fa fa-warning'></i><strong> <LABEL_1689></strong>"
            }, function () {
                Common.ajaxSubmit("/document/move?document_id=" + moveNode.id + "&target_id=" + targetNode.id);
            }, function () {

            });

            return false;
        }

        function onDrop(treeId, treeNodes, targetNode, moveType) {
            console.log("<LABEL_247>");
            return false;
        }

        function addHoverDom(treeId, treeNode) {
            if (isEditor == false) {
                return false;
            }
            if (treeNode.isParent === false || treeNode.isParent === undefined) {
                return false
            }
            var sObj = $("#" + treeNode.tId + "_span");
            var addBtn = $("#addBtn_" + treeNode.tId);
            if (addBtn.length > 0) return;

            var spanHtml = "<span class='button add' id='addBtn_" + treeNode.tId + "' title='<LABEL_1482>' onfocus='this.blur();'></span>";
            sObj.append(spanHtml);

            // bind add
            var addBtn = $("#addBtn_" + treeNode.tId);
            if (addBtn) addBtn.bind("click", function () {
                var content = "/document/add?space_id=" + treeNode.spaceId + "&parent_id=" + treeNode.id;
                layer.open({
                    type: 2,
                    skin: Layers.skin,
                    title: '<strong><LABEL_1134></strong>',
                    shadeClose: true,
                    shade: 0.6,
                    maxmin: true,
                    area: ["800px", "345px"],
                    content: content,
                    padding: "10px"
                });
                return false;
            });
        }

        function removeHoverDom(treeId, treeNode) {
            $("#addBtn_" + treeNode.tId).unbind().remove();
        }

        $(document).ready(function () {
            $.fn.zTree.init($(element), setting, treeData);
            var zTreeMenu = $.fn.zTree.getZTreeObj("dir_tree");
            var node = zTreeMenu.getNodeByParam("id", defaultId);
            zTreeMenu.selectNode(node, true);
            zTreeMenu.expandNode(node, true, false);
            //initialize fuzzysearch function
            fuzzySearch("dir_tree", '#document_search', null, true);
        });
    }
};