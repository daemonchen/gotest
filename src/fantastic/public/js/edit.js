$(function(){
  var doClear = function(){
    $("#title").val("");
    $("#content").val("");
    window.localStorage.clear();
  }
  var title = window.localStorage.getItem("edittingArticleTitle");
  var content = window.localStorage.getItem("edittingArticleContent");
  if(title&&title.length>0){
    $("#title").val(title)
  }
  if(content&&content.length>0){
    $("#content").val(content)
  }
  $("#title").keyup(function(){
    window.localStorage.setItem("edittingArticleTitle",$(this).val());
  });
  $("#content").keyup(function(){
    window.localStorage.setItem("edittingArticleContent",$(this).val());
  });
  $("#submit").click(function(){
    var articleTitle = $("#title").val();
    var articleContent = $("#content").val();
    if(articleTitle.length == 0 || articleContent.length == 0){
      console.log("你在逗我么");
    }
    $.ajax({
        url: "/edit/post",
        type: "POST",
        data: { title: articleTitle, content: articleContent},
        dataType: "json",
        success: function (result) {
          console.log("todo: change url to detail","result:",result);
          doClear()
          // todo: change url to detail
        },
        error: function (xhr, ajaxOptions, thrownError) {
        console.log(xhr.status);
        console.log(thrownError);
        }
    });
  });
});