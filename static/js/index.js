// jshint esversion: 6
$(function(){
    var modal = document.getElementById('myModal');
    // $('.modal').load('http://localhost:9000/download/js/appar-multiple.html')

    // $('a').on("click", function(){
    //     confirm('Es-tu sûr de télécharger ce fichier?')
    // })
    // $('.cat').hide().addClass('hidden');
    // $('.cat + .files').prev().show().removeClass('hidden');
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
        }else{
            return confirm('Es-tu sûr de télécharger ce fichier?');
        }
    });


    $('.close').on('click', function(){
        $('#myModal').fadeOut();
    });

    $(window).click(function(e){
        if (e.target === modal){
            $('#myModal').fadeOut();
        }
    });
});
