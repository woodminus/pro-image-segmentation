
function changeImage(activeImage, hiddenImage) {
  $('#show-' + activeImage).addClass('btn-primary');
  $('#show-' + activeImage).removeClass('btn-default');
  $('#show-' + hiddenImage).addClass('btn-default');
  $('#show-' + hiddenImage).removeClass('btn-primary');
  $('#' + hiddenImage + '-image').hide();
  $('#' + activeImage + '-image').show();
}

$(function() {
  $.material.init();

  var currentOriginal = '/img/original.jpg';
  var currentResult = '/img/result.png';

  $('#sigma-slider').noUiSlider({
    start: 0.8,
    step: 0.05,
    connect: 'lower',
    range: {
      min: 0.0,
      max: 1.0
    }
  });

  $('#k-slider').noUiSlider({
    start: 300,
    step: 10,
    connect: 'lower',
    range: {
      min: 0,
      max: 3000
    }
  });

  $('#minsize-slider').noUiSlider({