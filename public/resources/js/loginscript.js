$(function(){
    var error = $(".error");
    var errorContent = error.html();

    if(errorContent != "") {
        setTimeout(function(){
            error.fadeOut();
        },2000);
    }
});
