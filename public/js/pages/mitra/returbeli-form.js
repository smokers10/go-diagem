jQuery(document).ready(function () {
 
    get_cart();
    $("#field-supplier").select2({
        placeholder: 'Cari Supplier',
        language: 'id',
        ajax: {
            url: laroute.route('mitra.supplier.json'),
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
        }
    });

    $("#field-outlet").select2({
        placeholder: 'Pilih Outlet',
        allowClear: true,
        ajax: {
            url: laroute.route('mitra.outlet.json'),
            type: 'POST',
            dataType: 'JSON',
            delay: 250,
            data: function (params) {
                return {
                    searchTerm: params.term // search term
                };
            },
            processResults: function (response) {
                return {
                    results: response
                };
            },
            cache: true
        }
    }).on('select2:unselecting', function(e) {
        $(this).val(null).trigger('change');
        e.preventDefault();
    });

    $("#field-merk").change(function(){
        getProduk(null, null, $(this).val() ,null);
    });

    $("#field-kategori").select2({
        placeholder: 'Pilih Kategori',
        allowClear: true,
        ajax: {
            url: laroute.route('mitra.kategori.json'),
            type: 'POST',
            dataType: 'JSON',
            delay: 250,
            data: function (params) {
                return {
                    searchTerm: params.term // search term
                };
            },
            processResults: function (response) {
                return {
                    results: response
                };
            },
            cache: true
        }
    }).on('select2:unselecting', function(e) {
        $(this).val(null).trigger('change');
        e.preventDefault();
    });

    $(document).on('click', 'div.product_box', function() {
        $.ajax({
            method: 'GET',
            url: laroute.route('mitra.variasi.jsonModal'),
            data: {
                produk_id: $(this).data('produk_id'),
                tipe : 'returbeli',
            },
            dataType: 'JSON',
            success: function(result) {
                $('#modal_produkNama').html(result.produk_nama);
                $('#modal_produkKategori').html(result.produk_kategori);
                $('table#variasi_table tbody').html('');
                result.variasi.forEach(function(variasi, index){
                    $('table#variasi_table tbody').append(variasi);
                });
                if($('input#row_count').val() >= 1)
                {
                    $('#variasi_table tbody').find('tr').each(function() {
                        var variasi_id = $(this).find('.variasi_id_modal').val();
                        var qty_modal = $(this).find('.kuantias_modal');
                        $('#retur_table tbody').find('tr').each(function() {
                            var row_v_id = $(this).find('.row_variasi_id').val();
                            if (row_v_id == variasi_id) {
                                var qty_cart = $(this).find('.jumlah_beli').val();
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

    $('#form-produkvariasi').submit(function(e) { 
        e.preventDefault();
        var formData = new FormData($('#form-produkvariasi')[0]);
        formData.append('row_count', $('#row_count').val());
        $.ajax({
            url: laroute.route('mitra.returbeli.addCart'),
            type: 'POST',
            data: formData,
            cache: false,
            contentType: false,
            processData: false,
            beforeSend: function(){
                Swal.fire({
                    title: 'Tunggu Sebentar...',
                    text: ' ',
                    imageUrl: laroute.url('assets/img/loading.gif', ['']),
                    showConfirmButton: false,
                    allowOutsideClick: false,
                });
            },
            success: function (response) {
                if (response.fail == false) {
                    Swal.fire({
                        title: "Berhasil",
                        text: 'Produk Berhasil Ditambahkan!',
                        timer: 3000,
                        showConfirmButton: false,
                        icon: 'success'
                    });
                    // Tutup Modal Produk
                    $('#variasiModal').modal('hide');
                    $('#produkModal').modal('hide');
                    // Tampilkan Keranjang Pembelian

                    $('#retur_table tbody').html(response.html);
                    $('#retur_table tfoot').find('.total_subtotal').text(__currency_trans_from_en(response.sub_total, true, true));
                    __write_number($('input#total_subtotal_input'), response.sub_total, true);
        
                    $('#grand_total').text(__currency_trans_from_en(response.total, true, true));
                    __write_number($('input#grand_total_hidden'), response.total, true);

                    __write_number($('input#row_count'), response.total_item, true);
                }
            },
            error: function (jqXHR, textStatus, errorThrown) {
                Swal.close();
                alert('Error adding / update data');

            }
        });
    });

    $('#form-returbeli').submit(function(e) { 
        e.preventDefault();
        var formData = new FormData($('#form-returbeli')[0]);
        $.ajax({
            url: laroute.route('mitra.returbeli.simpan'),
            type: 'POST',
            data: formData,
            cache: false,
            contentType: false,
            processData: false,
            beforeSend: function(){
                Swal.fire({
                    title: 'Tunggu Sebentar...',
                    text: 'Data Sedang Diproses!',
                    imageUrl: laroute.url('assets/img/loading.gif', ['']),
                    showConfirmButton: false,
                    allowOutsideClick: false,
                });
            },
            success: function (response) {
                if (response.fail == false) {
                    Swal.fire({
                        title: "Berhasil",
                        text: 'Retur Produk berhasil Disimpan!',
                        timer: 3000,
                        showConfirmButton: false,
                        icon: 'success'
                    });
                    window.setTimeout(function () {
                        location.href = laroute.route('mitra.returbeli')
                    }, 1000);
                }else{
                    Swal.close();
                    if(response.pesan)
                    {
                        Swal.fire({
                            title: "Gagal",
                            text: response.pesan,
                            timer: 3000,
                            showConfirmButton: false,
                            icon: 'error'
                        });
                    }else{
                        for (control in response.errors) {
                            $('#field-' + control).addClass('is-invalid');
                            $('#error-' + control).html(response.errors[control]);
                        }
                    }
                }
            },
            error: function (jqXHR, textStatus, errorThrown) {
                Swal.fire({
                    title: "Gagal",
                    text: 'Data Gagal Diproses!',
                    timer: 3000,
                    showConfirmButton: false,
                    icon: 'error'
                });
            }
        });
    });

    $(document).on('change', '.jumlah_beli, .hrg_pokok', function() {
        var row = $(this).closest('tr');
        var variasi_id = row.find('.row_variasi_id').val();
        var quantity =  __read_number(row.find('input.jumlah_beli'), true);
        var hrg_pokok = __read_number(row.find('input.hrg_pokok'), true);

        updateCart(variasi_id, quantity, hrg_pokok)
    });

    var produkPage = 1;

    $(document).on('click', 'button#btn-cari_produk', function() {
        
        $('div#list-produk-body').html('');
        getProduk(null, null, null);
        $('#produkModal').modal();
    });

    var list_produk = $('#list-produk');
    list_produk.slimScroll({
        size: '5px',
        position: 'right',
        start: 'top',
        railVisible: true,
        railColor: '#222',
        railOpacity: 0.3,
        wheelStep: 10,
        allowPageScroll: false,
        disableFadeOut: false
    });

    list_produk.slimScroll().bind('slimscroll', function (e, pos) {
        // SOME CODE...
        if(pos == 'bottom')
        {
            produkPage++;
	        getProduk('semua', 'semua',produkPage);
        }
        // console.log(pos);
    });

    $('table#retur_table tbody').on('click', 'button.hapus_cart', function() {
        var variasi_id = $(this).parents('tr').find('.row_variasi_id').val();
        Swal.fire({
            title: "Anda Yakin?",
            text: "Produk Akan Dihapus Dari Keranjang Belanja!",
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
                url: laroute.route('mitra.returbeli.deleteCart'),
                type: "POST",
                dataType: "JSON",
                data: {
                    variasi_id: variasi_id,
                },
                success: function(response) {
                    Swal.fire({
                        title: "Berhasil",
                        text: 'Produk Berhasil Dihapus Dari Keranjang Belanja!',
                        timer: 3000,
                        showConfirmButton: false,
                        icon: 'success'
                    });
                    $('#retur_table tbody').html(response.html);
                    $('#retur_table tfoot').find('.total_subtotal').text(__currency_trans_from_en(response.sub_total, true, true));
                    __write_number($('input#total_subtotal_input'), response.sub_total, true);
        
                    $('#grand_total').text(__currency_trans_from_en(response.total, true, true));
                    __write_number($('input#grand_total_hidden'), response.total, true);
        
                    __write_number($('input#row_count'), response.total_item, true);
                },
                error: function(jqXHR, textStatus, errorThrown) {
                    alert('Error deleting data');
                }
            });
            } else {
                window.setTimeout(function(){
                    location.reload();
                } ,1500);
            }
        });
    });
});

function getProduk(keyword, kategori_id, merk_id, page) {
    if(keyword == null)
    {
        keyword = $('#field-keyword').val();
    }
    if(kategori_id == null)
    {
        kategori_id = $('#field-kategori').val();
    }
    if( merk_id== null)
    {
        merk_id = $('#field-merk').val();
    }
    $.ajax({
        method: 'GET',
        url: laroute.route('mitra.returbeli.getProduk'),
        data: {
            keyword: keyword,
            kategori_id: kategori_id,
            merk_id: merk_id,
            page: page,
        },
        dataType: 'JSON',
        beforeSend: function(){
        //    $('div.loading_ajax').show();
        },
        success: function(response) {
            // $('div.loading_ajax').hide();
            $('div#list-produk-body').html('');
            $('div#list-produk-body').append(response.html);
        },
    });
}

function get_cart()
{
    $.ajax({
        method: 'GET',
        url: laroute.route('mitra.returbeli.getCart'),
        dataType: 'JSON',
        success: function(response) {
            $('#retur_table tbody').html(response.html);
            $('#retur_table tfoot').find('.total_subtotal').text(__currency_trans_from_en(response.sub_total, true, true));
            __write_number($('input#total_subtotal_input'), response.sub_total, false);

            $('#grand_total').text(__currency_trans_from_en(response.total, true, true));
            __write_number($('input#grand_total_hidden'), response.total, false);

            __write_number($('input#row_count'), response.total_item, false);

            __write_number($('input#field-jml_bayar'), response.total, false);
        },
    });
}

function updateCart(cart_id, qty, hrg)
{
    $.ajax({
        url: laroute.route('mitra.returbeli.updateCart'),
        type: "POST",
        dataType: "JSON",
        data: {
            variasi_id : cart_id,
            qty:qty,
            hrg:hrg,
        },
        success: function (response) {
            if (response.fail == false) {
                $('#retur_table tbody').html(response.html);
                $('#retur_table tfoot').find('.total_subtotal').text(__currency_trans_from_en(response.sub_total, true, true));
                __write_number($('input#total_subtotal_input'), response.sub_total, true);
    
                $('#grand_total').text(__currency_trans_from_en(response.total, true, true));
                __write_number($('input#grand_total_hidden'), response.total, true);
            }
        },
        error: function (jqXHR, textStatus, errorThrown) {
            alert('Error adding / update data');
            $('#btnSimpan').text('Simpan');
            $('#btnSimpan').attr('disabled', false);

        }
    });
}
