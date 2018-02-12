/**
 * Created 2018/02/12 11:46 By lvmingyin
 */

$(function () {
    // $('input[type="text"]').each(function () {
    //     $(this).val($(this).prev().text());
    // });

    function strFirstLower(str) {
        var first = str.substring(0, 1);
        var other = str.substring(1);
        return first.toLowerCase() + other;
    }

    $('.saveBtn, .save').on('click', function () {
        if(!/^\/edit\/\d+$/.test(location.pathname)){
            $.ajax({
                type: "POST",//方法类型
                url: "/addMeasure",//url
                data: $('form').serialize(),
                success: function (result) {
                    if (result.code === 1) {
                        alert("添加成功");
                    } else {
                        alert("添加失败");
                    }
                },
                error: function () {
                    alert("异常！");
                }
            });
        }else{
            var paths = location.pathname.split("/");
            $.ajax({
                type: "POST",//方法类型
                url: "/editMeasure/" + paths[paths.length - 1],
                data: $('form').serialize(),
                success: function (result) {
                    if (result.code === 1) {
                        alert("修改成功");
                    } else {
                        alert("修改失败");
                    }
                },
                error: function () {
                    alert("异常！");
                }
            });
        }
    });

    if (/^\/edit\/\d+$/.test(location.pathname)) {
        var paths = location.pathname.split("/");
        $.ajax({
            type: "GET",
            url: '/measure/' + paths[paths.length - 1],
            success: function (result) {
                var data = result.data;
                for (var a in data) {
                    var ip = $('input[name="' + strFirstLower(a) + '"]');
                    if(ip.prop('type') === 'text'){
                        ip.val(data[a]);
                    }else if(ip.prop('type') === 'radio'){
                        ip.each(function(){
                            if(this.value === data[a]){
                                this.checked = true;
                            }
                        });
                    }
                }
            }
        });
    }
});
