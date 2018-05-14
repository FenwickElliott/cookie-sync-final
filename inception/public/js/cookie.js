let inceptionID = getCookie("inceptionID");

if (!inceptionID) {
    let date = new Date();
    inceptionID = sha1(window + date);
    date.setTime(date.getTime()+(365*24*60*60*1000));
    document.cookie = "inceptionID=" + inceptionID + ";expires="+ date.toUTCString();
}

function getCookie(name) {
    let value = "; " + document.cookie;
    let parts = value.split("; " + name + "=");
    if (parts.length == 2) return parts.pop().split(";").shift();
}