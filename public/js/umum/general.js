const toRupiah = new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
})

jQuery(function() {
    const logged = JSON.parse(localStorage.getItem("logged")) 
    if (logged.nama != null) {
        $("#header-user-logged-name").text(logged.nama)
    }
    cart_count()

    $(document).on('click', '#detail-kerangjang', function (){
        console.log("detail kerangjang SAT!")
    })
})

// Log out
$("#logout-button").on("click", function(){
    localStorage.removeItem("logged")
})

// Cart
const createCartItemHeader = (data) => {
    var item
    var { quantity, subtotal, produk, produk_single_foto, variasi } = data 
    var { harga, is_has_variant, nama, slug  } = produk
    var {variant, harga: hargaVariant} = variasi 

    if (!is_has_variant) {
        item = `<div class="cart-item">
            <div class="cart-item-prodImg">
                <img class="img-fluid" src="${ produk_single_foto.path }" alt="Image Description">
            </div>
            <div class="cart-item-prodInfo">
                <div class="cart-item-prodName">
                    <a href="/produk/detail/${slug}">${ nama }</a>
                </div>
                <span class="cart-item-prodQty">${ toRupiah.format(harga) } x (${ quantity }) Barang</span>
            </div>
            <div class="cart-item-prodTotal">
                ${ toRupiah.format(subtotal) }
            </div>
        </div>`
    }else {
        item = `<div class="cart-item">
            <div class="cart-item-prodImg">
                <img class="img-fluid" src="${ produk_single_foto.path }" alt="Image Description">
            </div>
            <div class="cart-item-prodInfo">
                <div class="cart-item-prodName">
                    <a href="/produk/detail/${slug}">${ nama } (${variant})</a>
                </div>
                <span class="cart-item-prodQty">${ toRupiah.format(hargaVariant) } x (${ quantity }) Barang</span>
            </div>
            <div class="cart-item-prodTotal">
                ${ toRupiah.format(subtotal) }
            </div>
        </div>`
    }

    return item
}

$(document).on('mouseenter','div.btn-cart', function () {
    $(this).dropdown("toggle")
    var content = $(this).find('.cart-content')
    var cart_header = $(this).find('.cart-header')
    $.ajax({
        url: "/cart/read",
        type: "GET",
        dataType: "JSON",
        beforeSend : function(){
            content.html('')
            content.append(`
            <div class="py-6 text-center" id="data-loading">
                <div class="height-50 spinner-grow text-primary width-50" role="status">
                    <span class="sr-only">Loading...</span>
                </div>
            </div>`)
        },
        success: function(response) {
            cart_header.find('span.cart-count').html(response.data.length)
            content.html('')
            var dataList = ''
            if(response.data.length !== 0){
                $.each(response.data, function(k, item) {
                    dataList += createCartItemHeader(item)
                })
            }else{
                dataList += `<div class="text-center py-5">
                    <img src="/img/placeholder/cart_is_empty.png" alt="empty-basket" class="max-width-200">
                    <div class="font-size-16 font-weight-bold mt-3">
                        Yah keranjang belanjaanmu masih kosong!
                    </div>
                </div>`
            }
            content.append(dataList)
        },
        error: function(httpObj, textStatus, errorThrown) {
            content.html(`<div class="text-center py-5">
                <img src="/img/placeholder/cart_is_empty.png" alt="empty-basket" class="max-width-200">
                <div class="font-size-16 font-weight-bold mt-3">
                    Yah keranjang belanjaanmu masih kosong!
                </div>
            </div>`)
        }
    })
})

// Product Detail
function checkAddToCartValidity() {
    has_var = $('input[name="has_variasi"]').val()
    if (has_var == '1') {
        var1 = $('input[name=var1]:checked', '#option-choice-form')
        if ($('input[name="has_variasi"]').attr("data-var2") == '1') {
            var2 = $('input:radio[name=var2]:checked').val()
            if (var1.length > 0 && var2.length > 0) {
                return true
            } else {
                return false
            }
        } else {
            if (var1.length > 0) {
                return true
            } else {
                return false
            }
        }
    } else {
        return true
    }
}

$(document).on('change', '#option-choice-form input', function () {
    if (checkAddToCartValidity()) {
        getVariantPrice()
        $('#error_cart').addClass('hide')
    } else {
        alert('asdas')
        $('#error_cart').removeClass('hide')
    }
})

$(document).on('change', '#option-choice-form input', function () {
    if (checkAddToCartValidity()) {
        getVariantPrice()
        $('#error_cart').addClass('hide')
    } else {
        $('#error_cart').removeClass('hide')
    }
})

function getVariantPrice() {
    if ($('#option-choice-form input[name=quantity]').val() > 0 && checkAddToCartValidity()) {
        $.ajax({
            type: "POST",
            url: laroute.route('variant_price'),
            data: $('#option-choice-form').serializeArray(),
            success: function (response) {
                if (response.fail == false) {
                    $('.product-harga').html(response.harga)
                    $('.total_harga').html(response.total)
                    $('.total-field').removeClass('hide')
                }
            }
        })
    }
}

$(document).on('click', '#btn-add-cart', function () {
    if (checkAddToCartValidity()) {
        $.ajax({
            type: "POST",
            url: "/cart/add-to-cart",
            data: $('#option-choice-form').serializeArray(),
            beforeSend: function () {
                Swal.fire({
                    title: 'Tunggu Sebentar...',
                    text: 'Data Sedang Di Proses!',
                    imageUrl: '/img/loading.gif',
                    showConfirmButton: false,
                    allowOutsideClick: false,
                })
            },
            success: function (response) {
                Swal.close()
                cart_count()
                $('#addToCart').modal()
            },
            error: function (httpObj, textStatus, errorThrown) {
                Swal.close()
            }
        })
    } else {
        var variant =[]
        $(".product-variant").find('.product-label').each(function(index, obj){
            variant.push($(this).text().toLocaleLowerCase())
        })

        var alertText = ''
        if(variant.length > 1){
            alertText = `Pilih ${ variant[0] } & ${ variant[1] } dulu, ya`
        }else{
            alertText = `Pilih ${ variant[0] } dulu, ya`
        }

        Swal.fire({
            title: "Oops!",
            text: alertText,
            timer: 3000,
            showConfirmButton: false,
            icon: 'warning'
        })
    }
})

function cart_count(){
    $.ajax({
        url: "/cart/total-carts",
        type: "GET",
        dataType: "JSON",
        success: function(response) {
            $("div.btn-cart").find('.cart-notification').html(response.data.total_carts)
        },
        error: function(httpObj, textStatus, errorThrown) {
            $("div.btn-cart").find('.cart-notification').html(0)
        }
    })
}