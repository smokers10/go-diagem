$(document).ready(function () {
    var select = $('.selectpicker').selectpicker();
	$(document).on('change', 'input[name=payment_method]:radio', function () {
        radio = $('input[name=payment_method]:checked');
        parent_label = radio.parent().parent().parent();
        parent_label.find('.form-group').removeClass('hide');

        $('.as-payment-cat').each(function () {
            if($(this).find('input[name=payment_method]').is(":checked")){
                $(this).find('.form-group').removeClass('hide');
            }else{
                $(this).find('.form-group').addClass('hide');
                $(this).find('.selectpicker').val('');
                $('.selectpicker').selectpicker('refresh');
            }
        });
    });
    
    $("form#checkout").submit(function (e) {
		e.preventDefault();
		var formData = new FormData($('form#checkout')[0]);
		$.ajax({
			url: laroute.route('checkout.simpan'),
			type: 'POST',
			data: formData,
			cache: false,
			contentType: false,
			processData: false,
			beforeSend: function(){
				Swal.fire({
					title: 'Tunggu Sebentar...',
					text: ' ',
					imageUrl: laroute.url('public/img/loading.gif', ['']),
					showConfirmButton: false,
					allowOutsideClick: false,
				});
			},
			success: function (response) {
				$('.is-invalid').removeClass('is-invalid');
				if (response.fail == false) {
					Swal.fire({
						title: "Berhasil",
						text: "Alamat Berhasil Diperbaharui!",
						timer: 3000,
						showConfirmButton: false,
						icon: 'success'
					});
				} else {
					Swal.close();
					for (control in response.errors) {
						$('#field-' + control).addClass('is-invalid');
						$('#error-' + control).html(response.errors[control]);
					}
				}
			},
			error: function (jqXHR, textStatus, errorThrown) {
				Swal.close();
				alert('Error adding / update data');
			}
		});
	});

});