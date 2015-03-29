/**
 * Created by haroldmiao on 2015/3/29.
 */


//function resetContent() {
//    var today = new Date();
//    var year = today.getFullYear();
//    var month = today.getMonth() < 9 ? "0" + (today.getMonth() + 1) : today.getMonth() + 1;
//    var date = today.getDate() < 10 ? "0" + today.getDate() : today.getDate();
//    var page = year.toString() + month.toString() + date.toString() + "/index_vod.html";
//    document.getElementById('configContent').src = "config/network.html";
//}
//window.onload = resetContent;


function openNetwork() {
    document.getElementById('configContent').src = "config/network.html";
}
