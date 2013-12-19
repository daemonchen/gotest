

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
          alert(xhr.status);
          alert(thrownError);
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
          var out = ""
          var last_word_is_x = false
          console.log("bayes learn result:",result);
        },
        error: function (xhr, ajaxOptions, thrownError) {
        alert(xhr.status);
        alert(thrownError);
        }
    });
  }


  $("#segmentButton").click(function(){
    Segment()
  });

  $("#bayesButton").click(function(){
    bayesLearn()
  });

  $('.ui.radio.checkbox')
  .checkbox({
    onChange: function(a){
      Category = $(this).next().html()
    }
  })
;
});