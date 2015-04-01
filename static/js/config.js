/**
 * Created by haroldmiao on 2015/3/29.
 */

$(document).ready(function() {
    $.ajax({
        type:"GET",
        url: "/data?action=get_total_status",
        dateType : "json",
        success: function(result) {
            console.log(result);
            console.log(result.mac);
            $("#mac_addr").text(result.mac);
            $("#ip_addr").text(result.ip);

            $("#used_storage").text(result.usedStorage);
            $("#all_storage").text(result.allStorage);

        },
        error: function(){
            alert('获取服务器信息失败', '错误');
        }
    });
});

//function resetContent() {
//    var today = new Date();
//    var year = today.getFullYear();
//    var month = today.getMonth() < 9 ? "0" + (today.getMonth() + 1) : today.getMonth() + 1;
//    var date = today.getDate() < 10 ? "0" + today.getDate() : today.getDate();
//    var page = year.toString() + month.toString() + date.toString() + "/index_vod.html";
//    document.getElementById('configContent').src = "config/network.html";
//}
//window.onload = resetContent;

//
//function openNetwork() {
//    document.getElementById('configContent').src = "config/network.html";
//}


