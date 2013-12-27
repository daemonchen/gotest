semantic.sidebar = {};
semantic.sidebar.ready = function(){
  $('.sidebar').first().sidebar('attach events', '#menu');
};
$(document)
  .ready(semantic.sidebar.ready)
;

var formatTimeLabel = function(){
  var stamp = parseInt($("#stamp").html())
  // moment(stamp*1000).format("YYYY:MM:DD  hh:mm:ss")
  $("#time").html(moment(stamp*1000).format("YYYY年MM月DD日  hh:mm:ss"))
}
$(function(){
  formatTimeLabel()
  $("#save").click(function(){
    var contentEditable = CKEDITOR.instances.editor
    if(contentEditable == undefined){return}
    var stamp = $("#stamp").html()
    var articleContent = contentEditable.getData()
    if(articleContent.length == 0){
      console.log("你在逗我么");
    }
    $.ajax({
        url: "/post/update",
        type: "POST",
        data: { stamp: stamp, content: articleContent},
        dataType: "json",
        success: function (result) {
          console.log("todo: change url to detail","result:",result);
        },
        error: function (xhr, ajaxOptions, thrownError) {
        console.log(xhr.status);
        console.log(thrownError);
        }
    });
  })
});