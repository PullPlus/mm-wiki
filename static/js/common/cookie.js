
if(!window.localStorage){
    alert("<LABEL_1108>localstorage");
}else{
}

var Cookie = Cookies.withConverter({
    read: function (value, name) {
        return Base64.decode(value);
    },
    write: function (value, name) {
        return Base64.encode(value);
    }
});