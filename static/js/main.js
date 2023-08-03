$(document).ready(function () {
    $('#example').DataTable();
});

document.addEventListener("DOMContentLoaded", function () {
	flatpickr("#datepicker", {
		dateFormat: "Y-m-d", // Change this to match the date format you want to use (YYYY-MM-DD in this case)
	});
});

(function($) {

	"use strict";
	var fullHeight = function() {

		$('.js-fullheight').css('height', $(window).height());
		$(window).resize(function(){
			$('.js-fullheight').css('height', $(window).height());
		});

	};
	fullHeight();

	$('#sidebarCollapse').on('click', function () {
      $('#sidebar').toggleClass('active');
  	});

})(jQuery);

