$(document).ready(function () {
	var alamatForm = $('#modalAlamat');
	var alamatModal = $('#alamatPilih');

    var wizard = $("#wizard").steps({
		headerTag: "h4",
		bodyTag: "section",
		transitionEffect: "fade",
		autoFocus: true,
		enablePagination: false,
		enableAllSteps: false,
		enableKeyNavigation: false,
		onStepChanged: function (event, currentIndex, newIndex) {
		}
	});
	$("#btn-next").on("click", function() {
		wizard.steps('next');
		$(this).addClass('hide');
		$('#btn-bayar').removeClass('hide');
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

	$(".btn-add_alamat").on("click", function() {  
		alamatForm.find('h5.modal-title').html('Tambah Alamat Baru');
		
        alamatForm.modal({
            backdrop: 'static',
            keyboard: false
        });
        $('#form-alamat')[0].reset();
	});
	
	var daerah = $('#field-desa').select2({
        placeholder: 'Cari Kelurahan',
        language: 'id',
        ajax: {
            url: laroute.route('wilayah.jsonSelect'),
            type: 'POST',
            dataType: 'JSON',
            delay: 250,
            data: function (params) {
                return {
                    searchTerm: params.term
                };
            },
            processResults: function (response) {
                return {
                    results: response
                };
            },
            cache: true
        },
        
        minimumInputLength: 3,
        // templateResult: formatResult,
        templateResult: function(response) {
            if(response.loading)
            {
                return "Mencari...";
            }else{
                var selectionText = response.text.split(",");
                var $returnString = $('<span>'+selectionText[0] + ', ' + selectionText[1] + '</br>' + selectionText[2]+ ', ' + selectionText[3] +'</span>');
                return $returnString;
            }
        },
        templateSelection: function(response) {
            return response.text;
        },
    });
    
	$("#form-alamat").submit(function (e) {
		e.preventDefault();
		var formData = new FormData($('#form-alamat')[0]);
		$.ajax({
			url: laroute.route('user.alamat.simpan'),
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
						text: "Alamat Baru Berhasil Ditambahkan",
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
	
	$(".btn-pilih_alamat").on("click", function() {
		$.ajax({
			url: laroute.route('user.alamat.json'),
			type: "GET",
			dataType: "JSON",
			success: function(response) {
				$.each(response.data, function(k, v) {
					if(v.is_utama == 1)
					{
						utama = 'checked';
					}else{
						utama = '';
					}
					$('#pilihAlamat').html(
						`<div class="col-md-6">
							<label class="aiz-megabox d-block bg-white" for="alamat-`+ v.id +`">
							<span class="d-flex p-3 aiz-megabox-elem">
								<label class="align-items-center as-radio-button as-radio-button">
									<input type="radio" name="alamatpilih" id="alamat-`+ v.id +`" class="as-radio-button__input pilihAlamat" value="`+ v.id +`" `+ utama +`>
									<span class="as-radio-button__check"></span>
								</label>
								<span class="flex-grow-1 pl-3">
									<div>
										<span class="font-weight-bold">`+ v.penerima +`</span>
										<span>(`+ v.nama +`)</span>
									</div>
									<div>
										<span>`+ v.phone +`</span>
									</div>
									<div>
										`+ v.alamat +`
									</div>
									<div>
										`+ v.daerah +`
									</div>
								</span>
							</span>
						</label>
					</div>`
					);
				});
			},
			error: function(jqXHR, textStatus, errorThrown) {
				alert('Error deleting data');
			}
		});
		alamatModal.find('h5.modal-title').html('Pilih Alamat Pengiriman');
        alamatModal.modal({
            backdrop: 'static',
            keyboard: false
        });
	});

	$(document).on('change', 'input[name=alamatpilih]:radio', function () {
		alamat_sel = $('input[name="alamatpilih"]:checked');
		id_alamat = alamat_sel.val();
		$.ajax({
			url: laroute.route('user.alamat.edit', { id: id_alamat }),
			type: "GET",
			dataType: "JSON",
			success: function(response) {
				utama = '';
				if(response.is_utama === 1)
				{
					utama = '<span class="badge badge-success">Utama</span>';
				}

				$('#alamatData').html(`<input type="hidden" name="alamat['id']" value="`+ response.id +`">
				<input type="hidden" name="alamat['penerima']" value="`+ response.penerima +`">
				<input type="hidden" name="alamat['phone']" value="`+ response.phone +`">
				<input type="hidden" name="alamat['alamat']" value="`+ response.alamat +`">
				<input type="hidden" name="alamat['kelurahan_id']" value="`+ response.kelurahan_id +`">
				<input type="hidden" name="alamat['kd_pos']" value="`+ response.kd_pos +`">
				<div>
					<p class="mb-0">
						<b class="nama-penerima">`+ response.penerima +`</b>
						<span class="nama-alamat">(`+ response.nama +`)</span>
						`+ utama +`
					</p>
					<div class="phone">
						`+ response.phone +`
					</div>
					<div class="alamat-lengkap">
						`+ response.alamat +`<br>
						`+ response.daerah +`, `+ response.kd_pos +`
					</div>
				</div>
				`);
				alamatModal.modal('hide');
			},
			error: function(jqXHR, textStatus, errorThrown) {
				alert('error get data');
			}
		});
	});

});

function load_alamat()
{
	$.ajax({
        url: laroute.route('cart.data'),
        type: "GET",
        dataType: "JSON",
        success: function(response) {
            $('#cart-content').html(response.html);
        },
        error: function(jqXHR, textStatus, errorThrown) {
            alert('error get data');
        }
    });
}