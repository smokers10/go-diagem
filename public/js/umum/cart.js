jQuery(function() {
    load_cart();
});

$(document).on('change', '#select-all', function () {
    $('.cart-item').find('input[type=checkbox]').not(this).prop('checked', this.checked);
    updateCart();
});

$(document).on('change', '.cart-store input[type=checkbox]', function () {
    parent = $(this).closest('.cart-store').parent();
    parent.find('.cart-product input[type=checkbox]').not(this).prop('checked', this.checked);
    updateCart();
});

$(document).on('change', '.cart-group input[type=checkbox]', function () {
    updateCart();
});

// $(document).on('change', '.product input.input-number', function () {
//     $.ajax({
//         type:"POST",
//         url: laroute.route('cart.updateQuantity'),
//         data: {
//             cart_id : $('.product input.cart_id').val(),
//             qty : parseInt($(this).val())
//         },
//         success: function(data){
//             $('#cartTopHover span').html(parseInt($('#cartTopHover span').text(), 10) - 1);
//             updateCart();
//         }
//     });
// });
$(document).on('change', '.cart-quantity input.input-number', function(event) {
    qty = parseInt($(this).val());
    min = parseInt($(this).attr('min'));
    max = parseInt($(this).attr('max'));
    if (qty < min){
        $(this).val(min);
    } else if (qty > max){
        $(this).val(max);
    }
    updateCartQty(qty);
});

$(document).on('click', '.btn-increment, .btn-decrement', function (e) {
  var isNegative = $(e.target).closest('.btn-decrement').is('.btn-decrement');
  var input = $(e.target).closest('.cart-quantity').find('input');
  min = parseInt(input.attr('min'));
  max = parseInt(input.attr('max'));
    var qty = parseInt(input.val());

    if(isNegative){
        qty = qty - 1;
    }else{
        qty = qty + 1;
    }

    if (qty < min){
        input.val(min);
    } else if (qty > max){
        input.val(max);
    }else{
        input.val(qty);
    }
    updateCartQty(qty);
});

function updateCartQty(qty){

    $.ajax({
        type:"POST",
        url: laroute.route('cart.updateQuantity'),
        data: {
            cart_id : $('.cart-item input.cart_id').val(),
            qty : qty
        },
        success: function(data){
            $('#cartTopHover span').html(parseInt($('#cartTopHover span').text(), 10) - 1);
            updateCart();
        }
    });
}


function updateCart()
{
    var total = 0;
    var produk = 0;
    var i = 0;
    var parent = $('form#cart-checkout');
    parent.find('input[type=hidden].ck').remove();
    $('.cart-item').each(function () {
        if($(this).find('input[type=checkbox]').is(":checked")){
            harga = $(this).find('input.harga').val();
            qty = parseInt($(this).find('input.input-number').val());
            cart_id = $(this).find('input.cart_id').val();
            total += harga*qty;
            produk += qty;
            i += 1;
            parent.append(`<input type="hidden" class="ck" name="c_id[]" value="`+cart_id+`">`);
        }
    });
    if($('.cart-item').find('input[type=checkbox]:checked').length > 0)
    {
        $('#hapus-all').removeClass('d-none');
        $('.btn-checkout').prop('disabled', false);
    }else{
        $('#hapus-all').addClass('d-none');
        $('.btn-checkout').prop('disabled', true);
    }
    $('.total_title').html('Total belanja ('+produk+' produk)');
    $('.total_belanja').text(total);
}

function load_cart()
{
    $.ajax({
        url: laroute.route('cart.data'),
        type: "GET",
        dataType: "JSON",
        beforeSend : function(){
            $("#data-list").after(`<div class="py-6 text-center" id="data-loading">
            <div class="height-50 spinner-grow text-primary width-50" role="status">
                <span class="sr-only">Loading...</span>
            </div>
        </div>`);
        },
        success: function(response) {
            var dataList = '';
            if(response.data.length !== 0){
                $('#cart-content').prepend(__createCartHeader());
                $.each(response.data, function(k, item) {
                    dataList += _createElement(item)
                });
            }else{
                dataList += `<div class="height-380 py-5 text-center">
                    <img class="empty-img" src="${ laroute.url('public/img/placeholder/alamat.png', []) }" width="200px">
                    <div>
                        <h3 class="font-size-24 font-weight-bold mt-5">Alamat Pengiriman Belum Ditambahkan</h3>
                        <p class="font-size-16"></p>
                        <button type="button" class="btn btn-primary btn-lg" id="btn-add_alamat">
                            <i class="fa fa-plus mr-1"></i>Tambah Alamat</button>
                    </div>
                </div>`;
            }
            $('#data-list').append(dataList);
        },
        complete: function(){
            $("#data-loading").remove();
        },
        error: function(jqXHR, textStatus, errorThrown) {
            alert('Error deleting data');
        }
    });
}

$(document).on('click', '#hapus-all', function () {
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
            var ck = [];
            $('.product').each(function () {
                if($(this).find('input[type=checkbox]').is(":checked")){
                    cart_id = $(this).find('input.cart_id').val();
                    ck.push(cart_id);
                }
            });
            $.ajax({
                url: laroute.route('cart.hapus'),
                type: "POST",
                data: {
                    c_id :ck
                },
                dataType: "JSON",
                beforeSend: function(){
                    Swal.fire({
                        title: 'Tunggu Sebentar...',
                        text: ' ',
                        imageUrl: laroute.url('public/img/loading.gif', ['']),
                        showConfirmButton: false,
                        allowOutsideClick: false,
                    });
                },
                success: function() {
                    Swal.fire({
                        title: "Berhasil",
                        text: 'Produk berhasil dihapus dari keranjang!',
                        timer: 3000,
                        showConfirmButton: false,
                        icon: 'success'
                    });
                    load_cart();
                },
                error: function(jqXHR, textStatus, errorThrown) {
                    alert('Error deleting data');
                }
            });
        }else{
            Swal.close();
        }
    });
});

function __createCartHeader(){
    var header = `<div class="block block-shadow block-bordered block-rounded mb-2">
    <div class="block-content block-content-full py-2">
        <div class="no-gutters row">
            <div class="col-6 my-auto ">
                <div class="custom-control custom-checkbox my-auto">
                    <input class="custom-control-input" type="checkbox" id="select-all">
                    <label class="custom-control-label" for="select-all">Pilih Semua Produk</label>
                </div>
            </div>
            <div class="col-6 my-auto text-right">
                <button type="button" class="btn btn-alt-danger btn-sm d-none" id="hapus-all">
                    <i class="fa fa-trash mr-1"></i>
                    Hapus
                </button>
            </div>
        </div>
    </div>
</div>`;
return header;
}

function _createElement(item){
    var item = `
    <div class="cart-item">
        <div class="d-flex">
            <div class="custom-control custom-checkbox my-auto">
                <input class="custom-control-input" type="checkbox" id="${ item.id }">
                <label class="custom-control-label" for="${ item.id }"></label>
            </div>
        </div>
        <div class="cart-item-product">
            <input type="hidden" class="cart_id" value="${ item.id }">
            <input type="hidden" class="harga" value="${ item.harga }">
            <div class="cart-item-product_img">
                <img src="${ item.produk.fotoUtama }" data-src="" alt="" class="img-fluid lazyImage" data-loaded="true">
            </div>
            <div class="cart-item-product_info">
                <div class="cart-item-product_title">
                    <a href="">${ item.produk.nama }</a>
                </div>
                <div class="cart-item-product_variant">
                   ${ item.variasi.variant }
                </div>
                <div class="cart-item-product_price">
                    ${ item.harga_frm }
                </div>
            </div>
        </div>
        <div class="cart-item-action my-auto">
            <div class="d-flex">
                <button class="btn btn-secondary btn-sm mr-2">
                    <i class="si si-trash"></i>
                </button>
                <div class="cart-quantity">
                    <button class="btn-decrement" type="button" data-type="minus">
                        <i class="fa fa-minus"></i>
                    </button>
                    <input type="number" name="quantity" class="input-number" value="${ item.qty }" min="1" max="${ item.variasi.stok }">
                    <button class="btn-increment" type="button" data-type="plus">
                        <i class="fa fa-plus"></i>
                    </button>
                </div>
            </div>
        </div>
    </div>`;

return item;
}