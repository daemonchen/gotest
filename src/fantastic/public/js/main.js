semantic.sidebar = {};
semantic.sidebar.ready = function(){
  $('.sidebar').first().sidebar('attach events', '#menu');
};
$(document)
  .ready(semantic.sidebar.ready)
;