semantic.validateForm = {};

// ready event
semantic.validateForm.ready = function() {

  // selector cache
  var

    $form       = $('.ui.form'),
    // alias
    handler
  ;

  // event handlers
  handler = {

  };

  $.fn.form.settings.defaults = {
    firstName: {
      identifier  : 'first-name',
      rules: [
        {
          type   : 'empty',
          prompt : 'Please enter your first name'
        }
      ]
    },
    lastName: {
      identifier  : 'last-name',
      rules: [
        {
          type   : 'empty',
          prompt : 'Please enter your last name'
        }
      ]
    },
    username: {
      identifier : 'username',
      rules: [
        {
          type   : 'empty',
          prompt : 'Please enter a username'
        }
      ]
    },
    email: {
      identifier : 'email',
      rules: [
        {
          type   : 'empty',
          prompt : 'Please enter your email'
        },
        {
          type   : 'email',
          prompt : 'Please enter a valid email'
        }
      ]
    },
    password: {
      identifier : 'password',
      rules: [
        {
          type   : 'empty',
          prompt : 'Please enter a password'
        },
        {
          type   : 'length[6]',
          prompt : 'Your password must be at least 6 characters'
        }
      ]
    },
    passwordConfirm: {
      identifier : 'password-confirm',
      rules: [
        {
          type   : 'empty',
          prompt : 'Please confirm your password'
        },
        {
          identifier : 'password-confirm',
          type       : 'match[password]',
          prompt     : 'Please verify password matches'
        }
      ]
    },
    terms: {
      identifier : 'terms',
      rules: [
        {
          type   : 'checked',
          prompt : 'You must agree to the terms and conditions'
        }
      ]
    }
  };

  var login = function(data){
      $.ajax({
        type:"GET",
        url: '/login/login',
        data: data,//{timestamp:timestamp},
        success: function(xhr,result,obj){

          console.log('login success')
          window.location.href="/edit/index"
          console.log("back");
        },
        error: function(obj,err,xhr){
          alert('username or password is wrong')
        }
      });
  }
  $form
    .form({},{onSuccess: function(e){
      var username = $("#username").val()
      var password = $("#password").val()
      login({username:username,password:MD5(password),timestamp:e.timeStamp})

    }})
  ;

};


// attach ready event
$(document)
  .ready(semantic.validateForm.ready)
;