// jshint esversion: 6
$(function(){
    var modal = document.getElementById('myModal');

    $("a[href^='/download']").on("click", function(e) {
        if ($(this).data('ext') === "jpg" || $(this).data('ext') === "png" || $(this).data('ext') === "gif"){
            e.preventDefault();
            $('#myModal > .modal-content > .modal-header > h2').html($(this).data('name') + " " + $(this).data('size'));
            $('#myModal > .modal-content > .modal-body ').html("<img src='/files/" +  $(this).data('link') + "' alt='" + $(this).data('link') + "' />");
            $('#myModal').fadeIn();
            let href = $(this).attr('href');
            $('.download_img').on('click', function(){
                window.location.href = href;
            });
        }else if ($(this).data('ext') === "html" || $(this).data('ext') === "txt" || $(this).data('ext') === "log") {
            e.preventDefault();
            $('#myModal > .modal-content > .modal-header > h2').html($(this).data('link') + " " + $(this).data('size'));
            $('#myModal > .modal-content > .modal-body ').load($(this).attr('href'));
            $('#myModal').fadeIn();
            let href = $(this).attr('href');
            $('.download_img').on('click', function(){
                window.location.href = href;
            });
        }else if ($(this).data('ext') === "mp4" || $(this).data('ext') === "mkv" || $(this).data('ext') === "avi") {
            e.preventDefault();
            $('#myModal > .modal-content > .modal-header > h2').html($(this).data('link') + " " + $(this).data('size'));
            $('#myModal > .modal-content > .modal-body ').html('<video width="640" height="480" autoplay controls preload="auto"> <source src="'+ $(this).attr('href') + '" type="video/mp4"> Your Browser does not support the video tag</video>');
            $('#myModal').fadeIn();
            let href = $(this).attr('href');
            $('.download_img').on('click', function(){
                window.location.href = href;
            });
        }else{
            return confirm('Es-tu sûr de télécharger ce fichier?');
        }
    });


    $('.close').on('click', function(){
        $('#myModal').fadeOut();
        $('#myModal > .modal-content > .modal-body ').html("");
    });

    $(window).click(function(e){
        if (e.target === modal){
            $('#myModal').fadeOut();
            $('#myModal > .modal-content > .modal-body ').html("");
        }
    });
});
