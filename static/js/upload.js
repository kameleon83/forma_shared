// jshint esversion: 6
$(function() {
    let submit = $('.submit');
    let folder = $('.folder');
    let file = $('.inputfile');

    // submit.attr('disabled',true);

    $('.inputfile').on('change', function(){
        // console.log(file.val());
        if ( file.val() !== ""){
            let res = file.val().split('\\');
            $('.name_file').html("Vous avez choisi ce fichier : <b>" + res[2] + "</b>");
            file = res[2];
        }
    });

    $('body').on('change', function(){
        let folder = $('.folder');
        let file = $('.inputfile');
            if (folder.val() !== "" && file.val() !== "") {
                submit.attr('disabled', false);
            }else{
                submit.attr('disabled', true);
            }
    });

});
