jQuery(function() {
    load_cart()
})

$(document).on('change', '#select-all', function () {
    $('.cart-item').find('input[type=checkbox]').not(this).prop('checked', this.checked)
    updateCart()
})

$(document).on('change', '.cart-store input[type=checkbox]', function () {
    parent = $(this).closest('.cart-store').parent()
    parent.find('.cart-product input[type=checkbox]').not(this).prop('checked', this.checked)
    updateCart()
})

$(document).on('change', '.cart-group input[type=checkbox]', function () {
    updateCart()
})

$(document).on('change', '.cart-quantity input.input-number', function(event) {
    qty = parseInt($(this).val())
    min = parseInt($(this).attr('min'))
    max = parseInt($(this).attr('max'))
    id = $(this).attr('data-id')
    if (qty < min){
        $(this).val(min)
    } else if (qty > max){
        $(this).val(max)
    }
    updateCartQtyAPI(qty, id)
})

function _cartQuantityBtn(data) {
    var isNegative = $(data).attr("data-type") == "minus" ? true : false,
    input = $(data).closest('.cart-quantity').find('input'),
    qty = parseInt(input.val()),
    min = parseInt(input.attr('min')),
    max = parseInt(input.attr('max')),
    cartID = $(data).attr("data-id")

    if(isNegative){
        qty = qty - 1
    }else{
        qty = qty + 1
    }

    if (qty < min){
        input.val(min)
    }else if (qty > max){
        input.val(max)
    }else{
        input.val(qty)
    }

    updateCartQtyAPI(qty, cartID)
}

function updateCartQtyAPI(qty, id){
    $.ajax({
        type:"PUT",
        url: "/cart/update-quantity",
        dataType: "json",
        contentType: "application/json",
        data: JSON.stringify({
            id : parseInt(id),
            Quantity : parseInt(qty)
        }),
        success: function(data){
            $('#cartTopHover span').html(parseInt($('#cartTopHover span').text(), 10) - 1)
            updateCart()
        }
    })

    load_cart()
}

function updateCart(){
    location.reload()
}

function load_cart(){
    $.ajax({
        url: "/cart/read",
        type: "GET",
        dataType: "JSON",
        beforeSend : function(){
            $("#data-list").after(`<div class="py-6 text-center" id="data-loading">
            <div class="height-50 spinner-grow text-primary width-50" role="status">
                <span class="sr-only">Loading...</span>
            </div>
        </div>`)
        },
        success: function(response) {
            var dataList = ''
            var totalBelanja = 0
            $("#data-list").html('')
            if(response.data.length !== 0){
                $.each(response.data, function(k, item) {
                    dataList += _createElement(item)
                })
                for (let i = 0; i < response.data.length; i++) {
                    const element = response.data[i]
                    totalBelanja = totalBelanja + element.subtotal
                }
                $("#total-belanja").text(toRupiah.format(totalBelanja))
                $("#total").val(totalBelanja)
            }else{
                dataList += `<div class="height-380 py-5 text-center">
                    <img class="empty-img" src="/img/placeholder/cart_empty.png" width="200px">
                    <div>
                        <h3 class="font-size-24 font-weight-bold mt-5">Cart kamu masih kosong</h3>
                        <p class="font-size-16"></p>
                        <a href="/produk" class="btn btn-primary btn-lg" id="btn-pilih-barang">
                            <i class="fa fa-plus mr-1"></i>Pilih Barang</a>
                    </div>
                </div>`
            }
            $('#data-list').append(dataList)
        },
        complete: function(){
            $("#data-loading").remove()
        },
        error: function(jqXHR, textStatus, errorThrown) {
            alert('Error deleting data')
        }
    })
}

function deleteFromCart(el) {
    const id = $(el).attr("data-id")
    Swal.fire({
        title: "Anda Yakin?",
        text: "Barang yang kamu pilih akan dihapus dari keranjangmu.",
        icon: "warning",
        showCancelButton: true,
        confirmButtonText: 'Ya, Hapus!',
        cancelButtonText: 'Tidak, Batalkan!',
        reverseButtons: true,
        allowOutsideClick: false,
        confirmButtonColor: '#af1310',
        cancelButtonColor: '#fffff',
    })
    .then((result) => {
        if (result.value) {
            $.ajax({
                url: "/cart/delete",
                type: "DELETE",
                data: {
                    id : parseInt(id)
                },
                dataType: "JSON",
                beforeSend: function(){
                    Swal.fire({
                        title: 'Tunggu Sebentar...',
                        text: ' ',
                        imageUrl: '/img/loading.gif',
                        showConfirmButton: false,
                        allowOutsideClick: false,
                    })
                },
                success: function() {
                    Swal.fire({
                        title: "Berhasil",
                        text: 'Produk berhasil dihapus dari keranjang!',
                        timer: 3000,
                        showConfirmButton: false,
                        icon: 'success'
                    })
                    load_cart()
                },
                error: function(jqXHR, textStatus, errorThrown) {
                    alert('Error deleting data')
                }
            })
        }else{
            Swal.close()
        }
    })
}

function _createElement(data){
    const { id, quantity, subtotal, variasi, produk, produk_single_foto } = data
    const { harga, id: produkID, nama, slug, discount, stok, is_has_variant  } = produk
    const { variant, harga: hargaVariant, stok: stokVariant } = variasi
    const { path } = produk_single_foto
    const pengaliDiscount = discount / 100
    var item
    var denganPotongan

    if (discount > 0) {
        denganPotongan = subtotal - (subtotal * pengaliDiscount)
    }

    if (!is_has_variant) {
        item = `<div class="cart-item">
            <div class="cart-item-product">
                <input type="hidden" class="harga" value="${ subtotal }" id="subtotal-${id}">
                <div class="cart-item-product_img">
                    <img src="${ path }" data-src="" alt="" class="img-fluid lazyImage" data-loaded="true">
                </div>
                <div class="cart-item-product_info">
                    <div class="cart-item-product_title">
                        <a href="">${ nama }</a>
                    </div>
                    <span>${ toRupiah.format(harga) } x (${quantity}) </span>
                    <div class="cart-item-product_price" id="view-cart-subtotal-${id}">
                        ${ discount > 0 ? toRupiah.format(denganPotongan) : toRupiah.format(subtotal) }
                    </div>
                    ${ discount > 0 ? `<span> sudah dipotong discount ${discount}% </span>` : "" }
                </div>
            </div>
            <div class="cart-item-action my-auto">
                <div class="d-flex">
                    <button class="btn btn-secondary btn-sm mr-2" data-id="${id}" onclick="deleteFromCart(this)">
                        <i class="si si-trash"></i>
                    </button>

                    <div class="cart-id">
                        <input type="hidden" value="${id}">
                    </div>

                    <div class="cart-quantity">
                        <button class="btn-decrement" type="button" data-type="minus" onclick="_cartQuantityBtn(this)" data-id=${id}>
                            <i class="fa fa-minus"></i>
                        </button>
                        <input type="number" name="quantity" class="input-number" data-id=${id} value="${ quantity }" min="1" max="${ stok }">
                        <button class="btn-increment" type="button" data-type="plus" onclick="_cartQuantityBtn(this)" data-id=${id}>
                            <i class="fa fa-plus"></i>
                        </button>
                    </div>
                </div>
            </div>
        </div>`
    } else {
        item = `<div class="cart-item">
            <div class="cart-item-product">
                <input type="hidden" class="harga" value="${ subtotal }" id="subtotal-${id}">
                <div class="cart-item-product_img">
                    <img src="${ path }" data-src="" alt="" class="img-fluid lazyImage" data-loaded="true">
                </div>
                <div class="cart-item-product_info">
                    <div class="cart-item-product_title">
                        <a href="">${ nama }</a>
                    </div>
                    <div class="cart-item-product_variant">
                        <b>
                        ${ variant}
                        </b>
                    </div>
                    <span>${ toRupiah.format(hargaVariant) } x (${quantity}) </span>
                    <div class="cart-item-product_price" id="view-cart-subtotal-${id}">
                        ${ discount > 0 ? toRupiah.format(denganPotongan) : toRupiah.format(subtotal) }
                    </div>
                    ${ discount > 0 ? `<span>dipotong discount ${discount}% </span>` : "" }
                </div>
            </div>
            <div class="cart-item-action my-auto">
                <div class="d-flex">
                    <button class="btn btn-secondary btn-sm mr-2" data-id="${id}" onclick="deleteFromCart(this)">
                        <i class="si si-trash"></i>
                    </button>

                    <div class="cart-id">
                        <input type="hidden" value="${id}">
                    </div>

                    <div class="cart-quantity">
                        <button class="btn-decrement" type="button" data-type="minus" onclick="_cartQuantityBtn(this)" data-id=${id}>
                            <i class="fa fa-minus"></i>
                        </button>
                        <input type="number" name="quantity" class="input-number" data-id=${id} value="${ quantity }" min="1" max="${ stokVariant }">
                        <button class="btn-increment" type="button" data-type="plus" onclick="_cartQuantityBtn(this)" data-id=${id}>
                            <i class="fa fa-plus"></i>
                        </button>
                    </div>
                </div>
            </div>
        </div>`
    }

    return item
}
