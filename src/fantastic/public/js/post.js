semantic.sidebar = {};
semantic.sidebar.ready = function(){
  $('.sidebar').first().sidebar('attach events', '#menu');
};
$(document)
  .ready(semantic.sidebar.ready)
;

$(function(){

  var stamp = $("#stamp").html()

  var formatTimeLabel = function(){
    var stampLabel = parseInt(stamp)
    $("#time").html(moment(stampLabel*1000).format("YYYY年MM月DD日  HH:mm:ss"))
  }

  formatTimeLabel()

  var doPostUpdate = function(url, stamp, articleContent){
    $.ajax({
        url: url,
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
  }
  var doPostComment = function(url, commentData){
    $.ajax({
        url: url,
        type: "POST",
        data: { commentData: JSON.stringify(commentData) },
        dataType: "json",
        success: function (result) {
          console.log("doPostComment result:",result);
          window.location.reload();
          // window.localStorage.clear();
        },
        error: function (xhr, ajaxOptions, thrownError) {
        console.log(xhr.status);
        console.log(thrownError);
        }
    });
  }

  var doDeleteUpdate = function(url){
    $.ajax({
        url: url,
        type: "DELETE",
        // data: { stamp: stamp},
        // dataType: "json",
        success: function (result) {
          console.log("todo: change url to detail","result:",result);
          window.location.reload();

        },
        error: function (xhr, ajaxOptions, thrownError) {
        console.log(xhr.status,thrownError);
        }
    });
  }

  //editable for content
  $("#save").click(function(){
    var contentEditable = CKEDITOR.instances.editor
    if(contentEditable == undefined){return}
    var articleContent = contentEditable.getData()
    if(articleContent.length == 0){
      return console.log("你在逗我么");
    }
    doPostUpdate("/post/update", stamp, articleContent)
  });
  $("#delete").click(function(){
    console.log(stamp)
    doDeleteUpdate("/post/index/?stamp="+stamp)
  });
  //comments

  var editorComment = CKEDITOR.replace( 'comment' );

  // localStorage user comment stuff
  $("#username").keyup(function(){
    window.localStorage.setItem("edittingUsername",$(this).val());
  });

  $("#email").keyup(function(){
    window.localStorage.setItem("edittingEmail",$(this).val());
  });

  editorComment.on("key",function(evt){
    console.log("data when click:",evt.editor.getData())
    window.localStorage.setItem("edittingCommentContent",evt.editor.getData());
  })

  var userName = window.localStorage.getItem("edittingUsername");
  var userEmail = window.localStorage.getItem("edittingEmail");
  var userComment = window.localStorage.getItem("edittingCommentContent");
  if(userName&&userName.length>0){
    $("#username").val(userName)
  }
  if(userEmail&&userEmail.length>0){
    $("#email").val(userEmail)
  }
  if(userComment&&userComment.length>0){
    editorComment.setData(userComment)
  }


  $(".submit").click(function(){
    var commentData = {
      relativeStamp: stamp,
      userName: $("#username").val() || "游客",
      userEmail: $("#email").val(),
      commentText: editorComment.getData(),
      commentTime: new Date().valueOf().toString()
    }
    if(commentData.commentText.length == 0){
      return console.log("你在逗我么");
    }
    doPostComment("/post/comment",commentData)
  });

  // transform comment result
  $(".comment .avatar img").each(function(index){
    var item = $(".comment .avatar img").eq(index);
    item.attr("src","http://www.gravatar.com/avatar/" + MD5(item.attr("data")))
  });

  $(".comment .content .metadata .date").each(function(index){
    var item = $(".comment .content .metadata .date").eq(index);
    var time = parseInt(item.html());
    item.html(moment(time).fromNow())
  });

});