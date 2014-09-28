/**
 * Created by CrossRun on 2014/9/28 0028.
 */
$(document).ready(function (){
    /*$("li").onmouseover(function(){
        this.css('color','#000')
    })*/
    $("li.titleB").click(function(){
//        this.css("color","#000")
//        alert(this)
        $("li.titleB").css("color","#b3b3b3")
        this.style.color="#fff"
    })
})