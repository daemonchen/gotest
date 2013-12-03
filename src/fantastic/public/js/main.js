semantic.validateForm = {};
semantic.validateForm.ready = function(){
  $('.ui.sidebar').sidebar('attach events', '.attached.button');
};
$(function(){
  semantic.validateForm.ready();
});