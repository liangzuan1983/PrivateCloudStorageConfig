/**
 * Created by haroldmiao on 2015/4/7.
 */

$(document).ready(function() {
    $("#btnWireConfig").click(function(){
        $.ajax({
            type: "POST",
            url: "/config/network",
            dataType: "JSON",
            data: $("#formConfigWire").serialize(),
            timeout: 5000,
            success: function (result) {
                alert(result.ObjectId);
                console.log(result);
            },
            error: function () {
                alert('获取服务器信息失败', '错误');
            }
        });
    })
})
