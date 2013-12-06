Segment = function() {
    var input = $('textarea#text').val();
    $.ajax({
        url: "/api/segment",
        type: "POST",
        data: { text: input },
        dataType: "json",
  success: function (result) {
    var out = ""
    var last_word_is_x = false
    console.log(">>>>>",result)
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