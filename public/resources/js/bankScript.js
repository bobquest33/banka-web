$(function(){
    $("main").load("/bank/accounts");

    $(".accounts").click(function(){
        $(".nav ul li a").each(function(){
            $(this).removeClass("selected")
        });

        $(this).addClass("selected")
        $("main").html("");
        $("main").load("/bank/accounts");

        return false;
    });

    $(".loans").click(function(){
        $(".nav ul li a").each(function(){
            $(this).removeClass("selected")
        });

        $(this).addClass("selected")
        $("main").html("");
        $("main").load("/bank/loans");

        return false;
    });

    $(".transactions").click(function(){
        $(".nav ul li a").each(function(){
            $(this).removeClass("selected")
        });

        $(this).addClass("selected")
        $("main").html("");
        $("main").load("/bank/transactions");

        return false;
    });
});
