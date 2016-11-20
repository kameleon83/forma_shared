// jshint esversion: 6
$(function(){
    let formateur = $('tbody tr.m2i').length;
    let total = $('tbody tr').length;
    let eleve = total - formateur;
    let tab_tr = new Array(total);

    $('tbody tr').each(function(i, v){
        tab_tr.push(v);
    });

    // $('a[href="/annuaire"]').append('<span class="nbr_annuaire">' + total + '</span>');
    $('section').prepend('<span class="number_annuaire">Nombres de contacts.&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Total : <b class="total">' + total + '</b>&nbsp;&nbsp;&nbsp;-&nbsp;&nbsp;&nbsp;Formateur : <b class="formateur">' + formateur + '</b>&nbsp;&nbsp;&nbsp;-&nbsp;&nbsp;&nbsp;Elèves : <b class="eleves">' + eleve + '</b></span>');

    // $('section table thead tr th').each(function(){
    // 	$(this).prepend('<a href="#" class="before ' + $(this).attr('class') + '"><img src="/images/down-arrow.png"></a>')
    // 	$(this).append('<a href="#" class="after ' + $(this).attr('class') + '"><img src="/images/up-arrow.png"></a>')
    // })

    // console.log(tab_tr);

    // $('.before, .after').on('click', function(){
    //     for (var v in tab_tr) {
    //         if (tab_tr.hasOwnProperty(v)) {
    //             console.log($(tab_tr[v]).children().first());
    //         }
    //     }
    // });

});
