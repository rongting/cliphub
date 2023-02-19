$(document).ready(function(){

    $("#send").click(function() {
        var message = $("#message").val();
        $.post("/send",
            {
                message: message
            },
            function(data, status){
                $("#generated-token-message").text(data);
            });
    });

    $("#receive").click(function() {
        var token = $("#token-message").val();
        $.post("/receive",
            {
                token: token
            },
            function(data, status){
                $("#received-message").text(data);
            });
    });

    $("#upload").click(function() {
        var formData = new FormData();
        var fileData = $("#file").prop("files")[0];
        formData.append("fileName", fileData);
        $.ajax({
            url: '/upload',
            type: 'POST',
            async: false,
            data: formData,
            cache: false,
            contentType: false,
            processData: false,
            success: function (data) {
                $("#generated-token-file").text(data);
            }
        });
    });

    $("#download").click(function() {
        var token = $("#token-file").val();
        window.location.href = "/download?token=" + token;
    });
});