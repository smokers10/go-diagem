$(document).ready(function () {
	var alamatModal = $('#alamatPilih')
	$("#paket-pengiriman").hide()
	$("#hidden-summary").hide()
	$("#submit-checkout").hide()

	load_detail_cart()
	load_alamat()
	loadAlamatOrigin()

	$("input[name=courier]").change(function() {
		ongkir()
	})

	$(document).on('change', 'input[name=pilihPaket]:radio', function() {
		selected = $('input[name="pilihPaket"]:checked')
		value = JSON.parse(selected.val())
		$("#courier-paket").val(value.service)
		$("#ongkir").val(value.cost.value)

		// hitung final total
		var final = parseInt($("#sub_total").val()) + parseInt($("#ongkir").val())
		$("#final_total").val(final)

		// assign ke DOM
		$("#ongkir_show").text(toRupiah.format(value.cost.value))
		$("#total_belanja").text(toRupiah.format(final))

		// tampilkan button checkout
		$("#submit-checkout").show()
		$("#hidden-summary").show()

		// scroll to top
		window.scrollTo({ top: 0, behavior: 'smooth' })
	})

	$("#btn-next").on("click", function() {
		wizard.steps('next')
		$(this).addClass('hide')
		$('#btn-bayar').removeClass('hide')
	})

	$("form#checkout").submit(function (e) {
		e.preventDefault()
		var formData = new FormData($('form#checkout')[0])
		
	})
	
	$(".btn-pilih_alamat").on("click", function() {
		alamatModal.find('h5.modal-title').html('Pilih Alamat Pengiriman')
        alamatModal.modal({
            backdrop: 'static',
            keyboard: false
        })
	})

	$(document).on('change', 'input[name=alamatpilih]:radio', function () {
		selected = $('input[name="alamatpilih"]:checked')

		alamat = JSON.parse(selected.val())
		data = JSON.parse(selected.attr("data-detail"))

		// assigement
		$("#selected_alamat").val(JSON.stringify(data))
		$("#alamatData").html("")
		$("#alamatData").append(_createAlamatElement(data))
		$("#destination").val(data.kota_id)

		// hide
		alamatModal.modal("hide")
	
		ongkir()
	})

	$("#btn-finish-checkout").on('click', function (e) {
		var alamat = JSON.parse($("#selected_alamat").val())
		
		const data = {
			alamat_id : alamat.id,
			ongkir : parseInt($("#ongkir").val()),
			kurir : $("#courier").val(),
			paket_kurir : $("#courier-paket").val(),
		}

		console.log(data)

		alert("halo")
	})
})

// alamat
function load_alamat(){
	$.ajax({
        url: "/alamat/read",
        type: "GET",
        dataType: "JSON",
        success: function(response) {
            const alamatData = $("#alamatData")
			const alamat = response.data
			alamatData.html(``)
			alamatData.append(`
				<div class="text-center">
					<img class="empty-img" src="/img/placeholder/alamat.png" height="250" width="250">
					<div>
						<h3 class="font-size-24 font-weight-bold mt-1">Silahkan pilih alamat pengiriman</h3>
						<a href="/alamat" class="btn btn-primary btn-lg"><i class="fa fa-plus mr-1"></i>Tambah Alamat Jika Belum Ada</a>
					</div>
				</div>
			`)

			alamat.forEach(element => {
				_createAlamatForChoosing(element)
			})
        },
        error: function(jqXHR, textStatus, errorThrown) {
            alert('error get data')
        }
    })
}

function loadAlamatOrigin() {
	$.ajax({
		url:"/alamat-origin",
		success: function(res) {
			const data = res.data
			$("#origin").val(data.kota_id)
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

function _createAlamatForChoosing(v) {
	$("#pilihAlamat").append(`<div class="col-md-12">
			<label class="aiz-megabox d-block bg-white">
			<span class="d-flex p-3 aiz-megabox-elem">
				<label class="align-items-center as-radio-button as-radio-button mr-3">
					<input type="radio" name="alamatpilih" data-detail='${JSON.stringify(v)}' class="as-radio-button__input pilihAlamat" value='${JSON.stringify(v)}'>
					<span class="as-radio-button__check"></span>
				</label>
				<span class="flex-grow-1 pl-3">
					<div>
						<span class="font-weight-bold">${v.penerima}</span>
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
}

// cart element

function load_detail_cart(){
	$.ajax({
		url: "/cart/read",
		type: "GET",
        dataType: "JSON",
		success: function(response) {
			const data = response.data
			var totalBelanja = 0
			var totalBerat = 0
			for (let i = 0; i < response.data.length; i++) {
				const element = response.data[i]
				totalBelanja = totalBelanja + element.subtotal
				totalBerat = totalBerat + (element.produk.berat * element.quantity)
			}
			data.forEach(dataEl => {
				$("#detail-belanja").append(_createCartElement(dataEl))
			})
			$("#total-belanja").text(toRupiah.format(totalBelanja))
			$("#sub_total").val(totalBelanja)
			$("#weight").val(totalBerat)
		},
		error: function(jqXHR, textStatus, errorThrown) {
            alert('Error deleting data')
        }
	})
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

// ongkir
function ongkir() {
	const data = {
		origin : $("#origin").val(),
		destination : $("#destination").val(),
		weight : parseInt($("#weight").val()),
		courier : $("input[name=courier]:checked").val()
	}

	$("#courier").val($("input[name=courier]:checked").val())

	if (data.destination == "") {
		alert("silahkan pilih alamat tujuan")
		return
	}

	if (!data.courier) {
		alert("silahkan pilih kurir")
		return
	}

	load_ongkir(data)
}

function load_ongkir(data) {
	$.ajax({
		url:"/cart/ongkir",
		method: "post",
        data : JSON.stringify(data),
        dataType: "json",
        contentType: "application/json",
		success: function (res) {
			var dataRes = JSON.parse(res.data)
			var level1 = dataRes.rajaongkir
			var result = level1.results[0]
			var { costs, name } = result
			$("#penyedia-paket").text(`Paket Yang Tersedia di ${name}`)
			$("#data-paket").html("")
			costs.forEach(data => {
				$("#data-paket").append(createCostElement(data.cost[0], data.service, data.description)) 
			})
			$("#paket-pengiriman").show()
		}
	})
}

function createCostElement(cost, service, description) {
	value = {cost, service}
	return `
	<div class="col-md-12">
			<label class="aiz-megabox d-block bg-white">
			<span class="d-flex p-3 aiz-megabox-elem">
				<label class="align-items-center as-radio-button as-radio-button mr-3">
					<input type="radio" name="pilihPaket" class="as-radio-button__input" value='${JSON.stringify(value)}'>
					<span class="as-radio-button__check"></span>
				</label>
				<span class="flex-grow-1 pl-3">
					<div>
						<span>Paket ${service}</span>
						<span class="font-weight-bold">${description}</span>
						<span>(${cost.etd} hari)</span>
					</div>
					<div>
						<span> ${ toRupiah.format(cost.value)} </span>
					</div>
				</span>
			</span>
		</label>
	</div>
	`
}

// submit checkout
function submitCheckout(data) {
	$.ajax({
		url: "",
		type: 'POST',
		data: data,
		dataType: "json",
		contentType: "application/json",
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
			if (response.success) {
				Swal.fire({
					title: "Checkout Selesai",
					text: "Silahkan bayar tagihan Anda",
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
}