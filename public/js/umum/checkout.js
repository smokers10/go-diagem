$(document).ready(function () {
	load_detail_cart()
	load_alamat()
	loadAlamatOrigin()
	var alamatModal = $('#alamatPilih')

    var wizard = $("#wizard").steps({
		headerTag: "h4",
		bodyTag: "section",
		transitionEffect: "fade",
		autoFocus: true,
		enablePagination: false,
		enableAllSteps: false,
		enableKeyNavigation: false,
		onStepChanged: function (event, currentIndex, newIndex) {}
	})
	
	$("#btn-next").on("click", function() {
		wizard.steps('next')
		$(this).addClass('hide')
		$('#btn-bayar').removeClass('hide')
	})

	$("form#checkout").submit(function (e) {
		e.preventDefault()
		var formData = new FormData($('form#checkout')[0])
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
					imageUrl: '/img/loading.gif',
					showConfirmButton: false,
					allowOutsideClick: false,
				})
			},
			success: function (response) {
				$('.is-invalid').removeClass('is-invalid')
				if (response.fail == false) {
					Swal.fire({
						title: "Berhasil",
						text: "Alamat Berhasil Diperbaharui!",
						timer: 3000,
						showConfirmButton: false,
						icon: 'success'
					})
				} else {
					Swal.close()
					for (control in response.errors) {
						$('#field-' + control).addClass('is-invalid')
						$('#error-' + control).html(response.errors[control])
					}
				}
			},
			error: function (jqXHR, textStatus, errorThrown) {
				Swal.close()
				alert('Error adding / update data')
			}
		})
	})
	
	$(".btn-pilih_alamat").on("click", function() {
		$.ajax({
			url: '/alamat/read',
			type: "GET",
			dataType: "JSON",
			success: function(response) {
				$.each(response.data, function(k, v) {
					if(v.is_utama == 1){
						utama = 'checked';
					}else{
						utama = '';
					}

					$("#pilihAlamat").append(`<div class="col-md-12">
							<label class="aiz-megabox d-block bg-white" for="alamat-`+ v.id +`">
							<span class="d-flex p-3 aiz-megabox-elem">
								<label class="align-items-center as-radio-button as-radio-button mr-3">
									<input type="radio" name="alamatpilih" id="alamat-${v.id} " class="as-radio-button__input pilihAlamat" value='${JSON.stringify(v)}' ${utama}>
									<span class="as-radio-button__check"></span>
								</label>
								<span class="flex-grow-1 pl-3">
									<div>
										<span class="font-weight-bold">`+ v.penerima +`</span>
										<span>( ${v.nama} )</span>
									</div>
									<div>
										<span> ${v.phone} </span>
									</div>
									<div>
										<span> ${v.alamat} </span> 
									</div>
									<div>
										<span>${v.nama_provinsi} - ${v.nama_kota} </span>
									</div>
								</span>
							</span>
						</label>
					</div>`)
				})
			},
			error: function(jqXHR, textStatus, errorThrown) {
				alert('Error deleting data')
			}
		})
		alamatModal.find('h5.modal-title').html('Pilih Alamat Pengiriman')
        alamatModal.modal({
            backdrop: 'static',
            keyboard: false
        })
	})

	$(document).on('change', 'input[name=alamatpilih]:radio', function () {
		selected = $('input[name="alamatpilih"]:checked')
		alamat = JSON.parse(selected.val())
		console.log(alamat)
	})
})

function load_alamat(){
	$.ajax({
        url: "/alamat/read",
        type: "GET",
        dataType: "JSON",
        success: function(response) {
            const alamatData = $("#alamatData")
			const alamat = response.data
			var loopIteration
			alamatData.html(``)
			if (alamat.length == 0) {
				alamatData.append(`
					<div class="text-center">
						<img class="empty-img" src="/img/placeholder/alamat.png">
						<div>
							<h3 class="font-size-24 font-weight-bold mt-1">Alamat Pengiriman Belum Ada</h3>
							<a href="/alamat" class="btn btn-primary btn-lg"><i class="fa fa-plus mr-1"></i>Tambah Alamat</a>
						</div>
					</div>
				`)
				return
			}

			for (let i = 0; i < alamat.length; i++) {
				const el = alamat[i]
				loopIteration = i
				if (el.is_utama) {
					break	
				}
			}

			$("#selected_alamat").val(JSON.stringify(alamat[loopIteration]))
			$("#alamatData").append(_createAlamatElement(alamat[loopIteration]))
        },
        error: function(jqXHR, textStatus, errorThrown) {
            alert('error get data')
        }
    })
}

function load_detail_cart(){
	$.ajax({
		url: "/cart/read",
		type: "GET",
        dataType: "JSON",
		success: function(response) {
			const data = response.data
			var totalBelanja = 0
			for (let i = 0; i < response.data.length; i++) {
				const element = response.data[i]
				totalBelanja = totalBelanja + element.subtotal
			}
			data.forEach(dataEl => {
				$("#detail-belanja").append(_createCartElement(dataEl))
			})
			$("#total-belanja").text(toRupiah.format(totalBelanja))
		},
		error: function(jqXHR, textStatus, errorThrown) {
            alert('Error deleting data')
        }
	})
}

function _createAlamatElement(data) {
	const { alamat, is_utama, nama_kota, nama_provinsi, penerima, phone, kd_pos, nama } = data
	return `
		<div>
			<p class="mb-0">
				<b class="nama-penerima">${penerima}</b>
				<span class="nama-alamat">${nama}</span>
				${ is_utama ? `<span class="badge badge-success">Utama</span>` : "" }
			</p>

			<div class="phone">
				${phone}
			</div>

			<div class="alamat-lengkap">
				${alamat} ${nama_kota} ${nama_provinsi}<br>
				${kd_pos}
			</div>
		</div>
	`
}

function _createCartElement(data) {
	const { id, quantity, subtotal, variasi, produk, produk_single_foto } = data
    const { harga, id: produkID, nama, slug, discount, stok, is_has_variant  } = produk
    const { variant, harga: hargaVariant, stok: stokVariant } = variasi
    const { path } = produk_single_foto
    const pengaliDiscount = discount / 100
    var denganPotongan

    if (discount > 0) {
        denganPotongan = subtotal - (subtotal * pengaliDiscount)
    }

	return `
		<div class="cart-item">
			<div class="cart-item-product">
				<div class="cart-item-product_img">
					<img src="${path}" data-src="${path}" alt="foto produk" class="img-fluid lazyImage" data-loaded="true">
				</div>
				<div class="cart-item-product_info">
					<div class="cart-item-product_title">
						<a href="">
							${nama} ${is_has_variant ? `(${variant})` : ""}
						</a>
					</div>
					<div class="cart-item-product_variant">
						Ukuran 30
					</div>
					<input type="hidden" class="cart_id" value="${id}">
					<input type="hidden" class="harga" value="17900.00">
					<div class="cart-item-product_price">
						${ is_has_variant ? toRupiah.format(hargaVariant) : toRupiah.format(harga) } x ${quantity} ${ discount > 0 ? ` dengan discount ${discount}%` : "" }
					</div>
				</div>
			</div>
			<div class="my-auto">
				<div class="float-right font-size-20 font-weight-bold product__price py-5 text-orange">
					${ discount > 0 ? toRupiah.format(denganPotongan) : toRupiah.format(subtotal) }
				</div>
			</div>
		</div>
	`
}

function load_ongkir(params) {
	$.ajax({
		url:"/cart/ongkir",
		method: "post",
        data : JSON.stringify(formData),
        dataType: "json",
        contentType: "application/json",
		success: function (params) {
			console.log(params)
		}
	})
}

function loadAlamatOrigin() {
	$.ajax({
		url:"/alamat-origin",
		success: function(res) {
			
		}
	})
}