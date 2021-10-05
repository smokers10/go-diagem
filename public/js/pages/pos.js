jQuery(document).ready(function () {
    get_product_suggestion_list('semua', 'semua', null);
    get_cart();

    $('select#produk_kategori, select#produk_merk').on('change', function(e) {
        $('input#suggestion_page').val(1);
        get_product_suggestion_list(
            $('select#produk_kategori').val(),
            $('select#produk_merk').val(),
            null
        );
    });

    $('#form-produkvariasi').submit(function(e) { 
        e.preventDefault();
        var formData = new FormData($('#form-produkvariasi')[0]);
        $.ajax({
            url: laroute.route('mitra.pos.addCart'),
            type: 'POST',
            data: formData,
            cache: false,
            contentType: false,
            processData: false,
            // beforeSend: function(){
            //     Swal.fire({
            //         title: 'Tunggu Sebentar...',
            //         text: ' ',
            //         imageUrl: laroute.url('assets/img/loading.gif', ['']),
            //         showConfirmButton: false,
            //         allowOutsideClick: false,
            //     });
            // },
            success: function (response) {
                if (response.fail == false) {
                    $('#variasiModal').modal('hide');
                    row_cart = parseInt($('input#row_cart').val());
                    response.html.forEach(function(entry){
                        if(row_cart > 0)
                        {
                            $('#pos_table tbody').find('tr').each(function() {
                            var row_v_id = $(this).find('.row_variasi_id').val();
                                if (row_v_id == entry.variasi_id)
                                {
                                    alert('sudah ada');
                                }else{
                                    $('table#pos_table tbody').append(entry.html);
                                }
                            });
                        }else{
                            $('table#pos_table tbody').append(entry.html);
                        }
                    });
                    $('input#row_cart').val(row_cart+response.row);
                    // foreach(response.html as h
                    // $('#pos_table tbody').find('tr').each(function() {
                    //     var row_v_id = $(this)
                    //         .find('.row_variasi_id')
                    //         .val();
                    //     if (row_v_id == variation_id)
                    //     {
                            //Increment product quantity
                            // qty_element = $(this).find('.pos_quantity-'+);
                            // var qty = __read_number(qty_element);
                            // __write_number(qty_element, qty + 1);
                            // qty_element.change();
                    //     }
                    // });
                    // if(row_cart > 0)
                    // {
                    //     $('table#pos_table tbody')
                    //         .append(response.html_content);
                    // }

                } else {
                    // for (control in response.errors) {
                    //     $('#field-' + control).addClass('is-invalid');
                    //     $('#error-' + control).html(response.errors[control]);
                    // }
                }
            },
            error: function (jqXHR, textStatus, errorThrown) {
                alert('Error adding / update data');
                $('#btnSimpan').text('Simpan');
                $('#btnSimpan').attr('disabled', false);

            }
        });
    });
   
    $(document).on('click', 'div.product_box', function() {
       
        $.ajax({
            method: 'GET',
            url: laroute.route('mitra.variasi.jsonModal'),
            data: {
                produk_id: $(this).data('produk_id'),
                tipe: 'penjualan'
            },
            dataType: 'JSON',
            success: function(result) {
                $('#modal_produkNama').html(result.produk_nama);
                $('#modal_produkKategori').html(result.produk_kategori);
                 $('table#variasi_table tbody').html('');
                result.variasi.forEach(function(variasi, index){
                    $('table#variasi_table tbody').append(variasi);
                });
                // $('#form-produkvariasi').html('');
                // $('#form-produkvariasi').append(result);
                row_cart = $('input#row_cart').val();
                if(row_cart > 0)
                {
                    $('#variasi_table tbody').find('tr').each(function() {
                        var variasi_id = $(this).find('.variasi_id_modal').val();
                        var qty_modal = $(this).find('.kuantias_modal');
                        $('#pos_table tbody').find('tr').each(function() {
                            var row_v_id = $(this).find('.row_variasi_id').val();
                            if (row_v_id == variasi_id) {
                                var qty_cart = $(this).find('.row_kuantitas').val();
                                qty_modal.val(qty_cart);
                            }
                        });
                    });
                }
                $('#variasiModal').modal({
                    backdrop: 'static',
                    keyboard: false
                })
            },
        });
    });
});

function get_product_suggestion_list(kategori_id, merk_id, url = null) {
    // if($('div#product_list_body').length == 0) {
    //     return false;
    // }
    if (url == null) {
        url = laroute.route('mitra.pos.get_product', []);
    }
    var page = $('input#suggestion_page').val();
    if (page == 1) {
        $('div#product_list_body').html('');
    }
    $.ajax({
        method: 'GET',
        url: url,
        data: {
            kategori_id: kategori_id,
            merk_id: merk_id,
            page: page,
        },
        dataType: 'html',
        success: function(result) {
            $('div#product_list_body').append(result);
        },
    });
}

// function addRow