const currentPath = location.pathname
const CPSplit = currentPath.split("/")  
const slug = CPSplit[3]

$("#varian-section").hide()

$(document).ready(function(){
    $("#discount").hide()

    $("body").scrollspy({
        target: "#product-nav",
        offset: 583
    })

    $("#feedback-form").submit(function(e) {
        e.preventDefault()
        data = {
            comment : $("#comment").val(),
            rating : $("#rating").val(),
            produk_id : $("#id-produk").val()
        }

        submitFeedback(data)
    })

    $.ajax({
        url:`/produk/get/${slug}`,
        success: function(response) {
            const data = response.data
            const { 
                nama, kode, harga, berat, satuan_berat, stok, tinggi, lebar, panjang,
                kategori, produk_foto, deskripsi, spesifikasi, id, is_has_variant,
                variasi, discount
            } = data

            // set to DOM
            $("#nama-produk").text(nama)
            $("#deskripsi").html(deskripsi)
            $("#stok").text(stok)

            // set value
            $("#id-produk").val(id)
            $("#harga-original").val(harga)
            $("#discount-form").val(discount)
            $("#buy-quantity").attr("data-max",stok)
        
            // init price state
            if (is_has_variant) {
                _initialPriceStateWithVarian(variasi, discount)
            }else{
                _initialPriceState(discount, harga)
            }
            
            // produk foto
            produk_foto.forEach(data => {
                if (data.is_utama) {
                    $("#foto-produk-main").attr("src", data.path)
                }
                $("#foto-produk-nav").append(_createImageNav(data.path))
            })

            // ngisi table spek
            spesifikasi.forEach(el => {
                $("#table-spec > tbody").append(`
                    <tr>
                        <td><b>${el.nama}</b></td>
                        <td>${el.value}</td>
                    <tr>
                `)
            })

            // jika punya variasi
            if (is_has_variant) {
                $("#varian-section").show()
                variasi.forEach(data => {
                    $("#variasi").append(_createVariantElement(data))
                })
            }

            // ambil feedback
            LoadAllFeedback(id)
        }
    })
})

$(".btn-qty").on("click", function () {
    var tipe = $(this).attr("data-type")
    var subtotalForm = $("#subtotal-form")
    var bqty = parseInt($("#buy-quantity").val())
    var hargaOriginal = parseInt($("#harga-original").val())
    
    // buy quantity data attribute
    const quantityForm = $("#buy-quantity")
    const maxStok = parseInt(quantityForm.attr("data-max"))
    const minStok = parseInt(quantityForm.attr("data-min"))

    // perhitungan subtotal
    var besaranDiscount = parseInt($("#discount-form").val())
    var discount = besaranDiscount / 100
    var potongan = hargaOriginal - (hargaOriginal * discount)
    
    // processed subtotal
    var subtotal 
    var qty
    
    // pengurangan / penambahan stok yang dibeli
    if (tipe == "minus") {
        if (bqty <= minStok) return
        qty = bqty - 1
    } else {
        if (bqty >= maxStok) return
        qty = bqty + 1
    }

    // set buy quantity
    $("#buy-quantity").val(qty)

    // pengurangan / penambahan subtotal
    if (besaranDiscount > 0) {
        subtotal = potongan * qty
    }else {
        subtotal = hargaOriginal * qty
    }
    
    subtotalForm.val(subtotal)
    $("#subtotal").text(toRupiah.format(subtotal))
})

function _createVarianElement(data) {
    return `
        <label class="product-label">${data.variant}</label>
        <div class="product-variant-item">
            <label for="nilai_Varian" class="radio-chip b-outline b-secondary">
                <input type="radio" id="nilai_Varian" name="var1" value="nilai">
                <span>nilai</span>
            </label>
        </div>
    `
}

function fotoOnclick(el){
    var fotoProdukMain = $("#foto-produk-main");
    var fotoProdukNav = document.getElementById('foto-produk-nav');

    let src = $(el).attr("src")
    fotoProdukMain.attr("src", src)
}

function imgOnclick(el){
    var modal = document.getElementById("myModal");
    var img = $("#foto-produk-main");
    var modalImg = document.getElementById("mainImg");
    var captionText = document.getElementById("caption");
    var span = document.getElementsByClassName("close-img")[0];
    
    modal.style.display = "block";
    modalImg.src = $(el).attr("src");
    // captionText.innerHTML = this.alt;

    span.onclick = function() { 
        modal.style.display = "none";
    }
}

function _createImageNav(path) {
    return `
        <div class="co-4">
            <img src="${path}" class="slider-nav__item h-100" alt="navigasi foto" onclick="fotoOnclick(this)"/>
        </div>
    `
}

function _createVariantElement(data) {
    let {id, harga, stok, variant} = data
    return `
        <button href="#" class="btn btn-outline-secondary mr-1 mb-3" data-id="${id}" data-harga="${harga}" data-stok="${stok}" data-variant="${variant}" onclick="_chooseVarian(this)">
            ${variant}
        </button>
    `
}

function _initialPriceStateWithVarian(variasi, discount) {
    var {harga, stok, variant, id} = variasi[0]

    // set to display
    $("#stok").text(stok)
    $("#choosed-variant").text(`(${variant})`)
    
    // set to attribute
    $("#harga-original").val(harga)
    $("#buy-quantity").attr("data-max",stok)
    $("#harga-original").val(harga)
    $("#id-varian").val(id)
    
    _initialPriceState(discount, harga)
}

function _initialPriceState(discount, harga) {
    let discountToInt = parseInt(discount)
    let persentase = discountToInt / 100
    var subtot
    var potongan = harga - (harga * persentase)

    if(discountToInt > 0) {
        subtot = parseInt($("#buy-quantity").val()) * potongan
        $("#discount").show()
        $("#besaran-discount").text(`${discount}%`)
        $("#harga-produk").text(toRupiah.format(potongan))
        $("#harga-produk-original").text(toRupiah.format(harga))
        $("#subtotal").text(toRupiah.format(subtot)) 
    } else {
        subtot = parseInt($("#buy-quantity").val()) * harga
        $("#harga").text(toRupiah.format(harga))
        $("#harga-produk").text(toRupiah.format(harga))
        $("#subtotal").text(toRupiah.format(subtot)) 
    }

    $("#subtotal-form").val(subtot)
}

function _chooseVarian(data) {
    let el = $(data)
    let harga = el.attr("data-harga")
    let variantID = el.attr("data-id")
    let stok = el.attr("data-stok")
    let namaVariant = el.attr("data-variant")
    let discount = parseInt($("#discount-form").val())

    // set to display
    $("#stok").text(stok)
    $("#choosed-variant").text(`(${namaVariant})`)
    // set to attribute
    $("#buy-quantity").val(1)
    $("#buy-quantity").attr("data-max",stok)
    $("#harga-original").val(harga)
    $("#id-varian").val(variantID)
    
    _initialPriceState(discount, harga)
}

// feed back & rating zone

function submitFeedback(data) {
    console.log(data)
    $.ajax({
        url:"/feedback/create",
        type: "post",
        data: JSON.stringify(data),
        dataType: "json",
        contentType: "application/json",
        beforeSend: function(){
            Swal.fire({
                title: 'Tunggu Sebentar...',
                text: '',
                imageUrl: '/img/loading.gif',
                showConfirmButton: false,
                allowOutsideClick: false,
            })
        },
        success: function () {
            Swal.fire({
                title: `Berhasil!`,
                text: "Rating & Ulasan Anda tersimpan",
                icon: 'success',
                showCancelButton: false,
                showConfirmButton: false,
                timer: 3000
            })

            setInterval(() => {
                location.reload()
            }, 3000);
        },
        error: function (jqXHR, textStatus, errorThrown) {
            Swal.close()
            alert('Error adding / update feedback')
        }
    })
}

function createFeedbackElement(data) {
    return `
        <div class="media">
            <div class="media-body">
                <h5 class="mt-0 mb-0">${data.user.nama}</h5>
                <div class="mb-5">
                    ${createStarElement(data.rating)}
                </div>
                <h6>pada ${data.created_at}</h6>
                <p>${data.comment}</p>
                <hr>
            </div>
        </div>
    `
}

function createStarElement(starCount) {
    var starDump = ""
    for (let i = 0; i < 5; i++) {
        if (i < starCount) starDump += `<i class="fas fa-star fa-lg text-primary"></i>`; 
        else starDump += `<i class="far fa-star fa-lg"></i>`;
    }
    return starDump
}

// trigger pengambilan feedback by rating 
function getFeedbackByRating(el) {
    const target = $(el)
    const productID = $("#id-produk").val()
    const rating = target.attr("data-rating")

    rating == "all" ? LoadAllFeedback(productID) : LoadFeedbackByRating(productID, rating)
}

// untuk mengambil data feedback sesuai rating
function LoadFeedbackByRating(productID, rating) {
    $.ajax({
        url:`/feedback/${productID}/${rating}`,
        beforeSend: function(){
            Swal.fire({
                title: 'Tunggu Sebentar...',
                text: '',
                imageUrl: '/img/loading.gif',
                showConfirmButton: false,
                allowOutsideClick: false,
            })
        },
        success: function(res){
            const { data } = res
            const { feedback_list } = data
            $("#feedback-ulasan").html("")
            if (feedback_list) {
                feedback_list.forEach(element => {
                    $("#feedback-ulasan").append(createFeedbackElement(element))
                })
            }
            Swal.close()
        }
    })
}

// untuk mengambil semua data feedback 
function LoadAllFeedback(productID) {
    $.ajax({
        url:`/feedback/${productID}`,
        success: function(res){
            const { data } = res
            const { feedback_list, by_rate, average_rating } = data

            $("#1-star").text(`Bintang 1 (${by_rate.one_star})`)
            $("#2-star").text(`Bintang 2 (${by_rate.two_star})`)
            $("#3-star").text(`Bintang 3 (${by_rate.three_star})`)
            $("#4-star").text(`Bintang 4 (${by_rate.four_star})`)
            $("#5-star").text(`Bintang 5 (${by_rate.five_star})`)

            $("#rating-score").text(average_rating)
            $("#rating-star").append(createStarElement(parseInt(average_rating)))

            $("#feedback-ulasan").html("")
            feedback_list.forEach(element => {
                $("#feedback-ulasan").append(createFeedbackElement(element))
            })
        }
    })
}


