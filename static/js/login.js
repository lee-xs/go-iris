$(document).ready(function () {
    layui.use('form,layer', function () {
        var layer = layui.layer
            ,form = layui.form

        form.on('submit(submitForm)',function (data) {
            console.log(data)

        })
    })
})