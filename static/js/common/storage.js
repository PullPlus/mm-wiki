
if(!window.localStorage){
    alert("<LABEL_1108>localstorage");
}

var Storage = {

    get: function (key) {
        var v = localStorage.getItem(key);
        if (v == null || v == undefined) {
            return ""
        }
        return Base64.decode(v);
    },

    set: function (key, value) {
        localStorage.setItem(key, Base64.encode(value));
    },

    remove: function (key) {
        localStorage.removeItem(key)
    }
};