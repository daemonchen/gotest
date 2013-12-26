$(function(){
  var editorContent = CKEDITOR.replace( 'content' );
  var doClear = function(){
    $("#title").val("");
    editorContent.setData("",function(){
      console.log("clear content when init");
    })
    window.localStorage.clear();
  }
  var title = window.localStorage.getItem("edittingArticleTitle");
  var content = window.localStorage.getItem("edittingArticleContent");
  if(title&&title.length>0){
    $("#title").val(title)
  }
  if(content&&content.length>0){
    editorContent.setData(content)
  }
  $("#title").keyup(function(){
    window.localStorage.setItem("edittingArticleTitle",$(this).val());
  });
  editorContent.on("key",function(evt){
    console.log("data when click:",evt.editor.getData())
    window.localStorage.setItem("edittingArticleContent",evt.editor.getData());
  })

  $("#submit").click(function(){
    var articleTitle = $("#title").val();
    var articleContent = editorContent.getData()
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