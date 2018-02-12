/**
 * Created 2018/02/12 19:10 By lvmingyin
 */

$(function(){

    $.ajax({
        type: 'GET',
        url: '/measures',
        success: function (result) {
            if (result.code === 1) {
                var data = result.data;
                data.forEach(function(n){
                    $('tbody').append('<tr>' +
                        '<td>'+ n.Name +'</td>' +
                        '<td>'+ n.Company +'</td>' +
                        '<td>'+n.Phone+'</td>' +
                        '<td>'+n.Address+'</td>' +
                        '<td>' +
                        '<a class="a" href="/edit/'+n.Id+'">修改</a>&nbsp;' +
                        '<span class="a down" data-href="/downloadMeasure/'+n.Id+'">下载</span>&nbsp;' +
                        '<span class="a delete" data-id="'+n.Id+'">删除</span></td>' +
                        '</tr>');
                });
            }
        }
    });

    $('.list').on('click','.down', function(){
       window.open($(this).attr('data-href'));
    });

    $('.list').on('click','.delete', function(){
        var self = $(this);
        $.ajax({
            type: 'GET',
            url: '/deleteMeasure/'+self.attr('data-id'),
            success: function(result){
                if(result.code === 1){
                    alert("删除成功");
                    self.parent().parent().remove();
                }else{
                    alert("删除失败");
                }
            },
            error: function(){
                alert("删除失败");
            }
        });
    });

});
