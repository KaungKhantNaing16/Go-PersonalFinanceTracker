$(document).ready(function () {
	$('#example').DataTable();
});

document.addEventListener("DOMContentLoaded", function () {
	flatpickr("#datepicker", {
		dateFormat: "Y-m-d", // Change this to match the date format you want to use (YYYY-MM-DD in this case)
	});
});

(function ($) {

	"use strict";
	var fullHeight = function () {

		$('.js-fullheight').css('height', $(window).height());
		$(window).resize(function () {
			$('.js-fullheight').css('height', $(window).height());
		});

	};
	fullHeight();

	$('#sidebarCollapse').on('click', function () {
		$('#sidebar').toggleClass('active');
	});

})(jQuery);

$(document).ready(function () {
	ImgUpload();
});

function ImgUpload() {
	var imgWrap = "";
	var imgArray = [];

	$('.upload__inputfile').each(function () {
		$(this).on('change', function (e) {
			imgWrap = $(this).closest('.upload__box').find('.upload__img-wrap');
			imgWrap.css('border', '2px dashed var(--main-color)')
			$('.submit-btn').css('display', 'block')
			var maxLength = $(this).attr('data-max_length');

			var files = e.target.files;
			var filesArr = Array.prototype.slice.call(files);
			var iterator = 0;
			filesArr.forEach(function (f, index) {

				if (!f.type.match('image.*')) {
					return;
				}

				if (imgArray.length > maxLength) {
					return false
				} else {
					var len = 0;
					for (var i = 0; i < imgArray.length; i++) {
						if (imgArray[i] !== undefined) {
							len++;
						}
					}
					if (len > maxLength) {
						return false;
					} else {
						imgArray.push(f);

						var reader = new FileReader();
						reader.onload = function (e) {
							var html = "<div class='upload__img-box'><div style='background-image: url(" + e.target.result + ")' data-number='" + $(".upload__img-close").length + "' data-file='" + f.name + "' class='img-bg'><div class='upload__img-close'></div></div></div>";
							imgWrap.append(html);
							iterator++;
						}
						reader.readAsDataURL(f);
					}
				}
			});
		});
	});

	$('body').on('click', ".upload__img-close", function (e) {
		var file = $(this).parent().data("file");
		for (let i = 0; i < imgArray.length; i++) {
			if (imgArray[i].name === file) {
				imgArray.splice(i, 1);
				break;
			}
		}
		if (imgArray.length === 0 ) {
			imgWrap.css('border', 'transparent')
			$('.submit-btn').css('display', 'none')
		}
		$(this).parent().parent().remove();
	});
}