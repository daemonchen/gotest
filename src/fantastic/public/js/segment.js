

$(function(){

  var Category = "Good"
  var Segment = function() {
      var input = $('textarea#text').val();
      $.ajax({
          url: "/api/segment",
          type: "POST",
          data: { text: input },
          dataType: "json",
          success: function (result) {
            var out = ""
            var last_word_is_x = false
            console.log("segment result:",result)
            for (var i = 0; i < result.segments.length; i++) {
              var segment = result.segments[i]
              if (segment.pos == "x") {
                if (segment.text != "\n") {
                  last_word_is_x = true
                } else {
                  last_word_is_x = false
                }

                out += segment.text
              } else {
                if (last_word_is_x) {
                  out += " "
                }
                out += segment.text + "/" + segment.pos + " "
              }
            }
            $('textarea#output').html(out);
          },
          error: function (xhr, ajaxOptions, thrownError) {
          console.log(xhr.status);
          console.log(thrownError);
          }
      });
  };

  var bayesLearn = function(){
    rowString = $('textarea#output').html();
    $.ajax({
        url: "/api/bayes",
        type: "POST",
        data: { category: Category, text: rowString},
        dataType: "json",
        success: function (result) {
          console.log("bayes learn result:",result);
        },
        error: function (xhr, ajaxOptions, thrownError) {
        console.log(xhr.status);
        console.log(thrownError);
        }
    });
  }

  var getBayesResult = function(text){
    $.ajax({
        url: "/api/logScore",
        type: "POST",
        data: { text: text},
        dataType: "json",
        success: function (result) {
          console.log("score:",result);
          $(".ui.info.message").show()
        },
        error: function (xhr, ajaxOptions, thrownError) {
        console.log(xhr.status);
        console.log(thrownError);
        }
    });
  }
  $("#segmentButton").click(function(){
    Segment()
  });

  $("#bayesButton").click(function(){
    bayesLearn()
  });

  $("#patternButton").click(function(){
    var input = $("#patternText").val()
    if(input.length == 0){
      console.log("please input your text for test")
      return
    }
    getBayesResult(input)
  });

  $('.ui.radio.checkbox')
  .checkbox({
    onChange: function(a){
      Category = $(this).next().html()
    }
  });
});