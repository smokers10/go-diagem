jQuery(document).ready(function () {
    getProduk(null, null, null);
    get_cart();

    pelangganOption = new Option("Pembeli Langsung", 1);
    var selectPelanggan = $("#field-pelanggan").select2({
        ajax: {
            url: laroute.route('mitra.pelanggan.json'),
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
    })
    selectPelanggan.append(pelangganOption).trigger('change');

    $("#field-merk").select2({
        placeholder: 'Pilih Merk',
        allowClear: true,
        ajax: {
            url: laroute.route('mitra.merk.json'),
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

    $('select#produk_kategori, select#produk_merk').on('change', function(e) {
        $('input#suggestion_page').val(1);
        get_product_suggestion_list(
            $('select#produk_kategori').val(),
            $('select#produk_merk').val(),
            null
        );
    });
   
    $(document).on('click', 'div.product_box', function(){
        $.ajax({
            method: 'GET',
            url: laroute.route('mitra.variasi.jsonModal'),
            data: {
                produk_id: $(this).data('produk_id'),
                tipe : 'penjualan',
            },
            dataType: 'JSON',
            success: function(result) {
                $('#modal_produkNama').html(result.produk_nama);
                $('#modal_produkKategori').html(result.produk_kategori);
                $('table#variasi_table tbody').html('');
                result.variasi.forEach(function(variasi, index){
                    $('table#variasi_table tbody').append(variasi);
                });
                $('#variasi_table tbody').find('tr').each(function() {
                    var variasi_id = $(this).find('.variasi_id_modal').val();
                    var qty_modal = $(this).find('.kuantias_modal');
                    $(this).find('.kuantias_modal').attr({"max" : result.max_stok});
                    $('#pos_table tbody').find('tr').each(function() {
                        var row_v_id = $(this).find('.row_variasi_id').val();
                        if (row_v_id == variasi_id) {
                            var qty_cart = $(this).find('.row_kuantitas').val();
                            qty_modal.val(qty_cart);
                        }
                    });
                });
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
        formData.append('row_cart', $('#row_cart').val());
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
            //         text: '',
            //         imageUrl: laroute.url('assets/img/loading.gif', ['']),
            //         showConfirmButton: false,
            //         allowOutsideClick: false,
            //     });
            // },
            success: function (response) {
                if (response.fail == false) {
                    $('#variasiModal').modal('hide');

                    $('#pos_table tbody').html(response.html);
                    
                    $('#total_subtotal').text(__currency_trans_from_en(response.sub_total, true, true));
                    __write_number($('input#total_subtotal_input'), response.sub_total, true);

                    $('#grand_total').text(__currency_trans_from_en(response.total, true, true));
                    __write_number($('input#grand_total_hidden'), response.total, true);

                    __write_number($('input#row_cart'), response.row_cart, true);
                    
                } else {
                }
            },
            error: function (jqXHR, textStatus, errorThrown) {
                alert('Error adding / update data');
                $('#btnSimpan').text('Simpan');
                $('#btnSimpan').attr('disabled', false);

            }
        });
    });

    $(document).on('click', 'button.bayar', function() {
        if($("#row_cart").val() > 0)
        {
            var total = $('input#grand_total_hidden').val();
            $('.total_tagihan').html(__currency_trans_from_en(total, false));
            $('input#field-uang_diterima').val(total, false);
            $('input#field-uang_diterima').attr("placeholder", "Rp " + __currency_trans_from_en(total, false));
            $('#payModal').modal({
                backdrop: 'static',
                keyboard: false
            });
        }else{
            Swal.fire({
                title: "Oops!",
                text: "Keranjang Penjualan Masih Kosong",
                timer: 3000,
                showConfirmButton: false,
                icon: 'error'
            });
        }
    });

    $(document).on('input', 'input#field-uang_diterima', function() {
        if($(this).val() !== '')
        {
            var kembali  =  $(this).val() - $('#pos_table tfoot').find('input.total_cart').val();
            $('input#field-uang_kembali').val(kembali);
            $('div.kembalian').show();
            $('button#bayar_pas').hide();
            $('button#bayar_lebih').show();
        }else{
            $('div.kembalian').hide();
            $('button#bayar_lebih').hide();
            $('button#bayar_pas').show();
        }
    });

    $(document).on('click', 'button#bayar_pas, button#bayar_lebih', function() {
        var formData = new FormData($('#form-penjualan')[0]);
        formData.append('uang_diterima', $('#field-uang_diterima').val());
        $.ajax({
            method: 'POST',
            url: laroute.route('mitra.pos.checkout'),
            data: formData,
            cache: false,
            contentType: false,
            processData: false,
            dataType: "JSON",
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
                    $('#payModal').modal('hide');
                    Swal.fire({
                        title: `Berhasil!`,
                        icon: 'success',
                        width: 480,
                        html: `Transaksi Berhasil Diproses!<br><br>
                            <button type="button" class="btn btn-alt-danger btn-cancel"><i class="si si-close mr-1"></i>Batal</button>
                            <button type="button" class="btn btn-alt-primary btn-add_another"><i class="si si-check mr-1"></i>Jual Produk Lagi!</button> 
                            <button type="button" class="btn btn-alt-primary btn-add_stok_awal"><i class="si si-check mr-1"></i>Print Faktur</button>`,
                        showCancelButton: false,
                        showConfirmButton: false,
                        onBeforeOpen: () => {
                            const add_another = document.querySelector('.btn-add_another')
                            const add_stok_awal = document.querySelector('.btn-add_stok_awal')
                            const batal = document.querySelector('.btn-cancel')

                            add_another.addEventListener('click', () => {
                                document.location.href = laroute.route('mitra.pos');
                            })

                            add_stok_awal.addEventListener('click', () => {
                                window.setTimeout(function () {
                                    location.reload();
                                }, 1500);
                            })
                            batal.addEventListener('click', () => {
                                document.location.href = laroute.route('mitra.penjualan');
                            })
                        }
                    });
                    // window.setTimeout(function () {
                    //     location.reload();
                    // }, 1500);
                } else {
                    for (control in response.errors) {
                        $('#field-' + control).addClass('is-invalid');
                        $('#error-' + control).html(response.errors[control]);
                    }
                }
            },
            error: function (jqXHR, textStatus, errorThrown) {
                alert('Error adding / update data');
                $('#btnSimpan').text('Simpan');
                $('#btnSimpan').attr('disabled', false);

            }
        });
    });
    
    $('table#pos_table tbody').on('click', 'button.ubah_cart', function() {
        var variasi_id = $(this).parents('tr').find('.row_variasi_id').val();
        var qty = $(this).parents('tr').find('.row_kuantitas').val();
        $.ajax({
            method: 'GET',
            url: laroute.route('mitra.variasi.jsonFind'),
            data: {
                variasi_id: variasi_id,
                qty: qty,
            },
            dataType: 'JSON',
            success: function(result) {
                $('#form-ubahcart').find('.modal_produkNama').html(result.produk_nama);
                $('#form-ubahcart').find('.modal_variasiNama').html(result.variasi_nama);
                $('#form-ubahcart').find('.hrg_net').html(result.hrg_jual);
                $('#form-ubahcart').find('.modal_hrgaxqty').html(result.sub_tot);
                $('#form-ubahcart').find('.input_quantity').val(result.qty);
                $('#ubah_variasi_id').val(result.variasi_id);
                $('#ubahCartModal').modal({
                    backdrop: 'static',
                    keyboard: false
                })
            },
        });
        //     .remove();
        // pos_total_row();
        // alert(variasi_id);
    });

    $('#form-ubahcart').submit(function(e) { 
        e.preventDefault();
        var formData = new FormData($('#form-ubahcart')[0]);
        $.ajax({
            url: laroute.route('mitra.pos.updateCart'),
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
                    $('#ubahCartModal').modal('hide');

                    $('#pos_table tbody').html(response.html);
                    
                    $('#total_subtotal').text(__currency_trans_from_en(response.sub_total, true, true));
                    __write_number($('input#total_subtotal_input'), response.sub_total, true);

                    $('#grand_total').text(__currency_trans_from_en(response.total, true, true));
                    __write_number($('input#grand_total_hidden'), response.total, true);

                    __write_number($('input#row_cart'), response.row_cart, true);
                    
                } else {
                }
            },
            error: function (jqXHR, textStatus, errorThrown) {
                alert('Error adding / update data');
                $('#btnSimpan').text('Simpan');
                $('#btnSimpan').attr('disabled', false);

            }
        });
    });

    $('table#pos_table tbody').on('click', 'button.hapus_cart', function() {
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
                url: laroute.route('mitra.pos.deleteCart'),
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

                    $('#pos_table tbody').html(response.html);
                    
                    $('#total_subtotal').text(__currency_trans_from_en(response.sub_total, true, true));
                    __write_number($('input#total_subtotal_input'), response.sub_total, true);
        
                    $('#grand_total').text(__currency_trans_from_en(response.total, true, true));
                    __write_number($('input#grand_total_hidden'), response.total, true);
        
                    __write_number($('input#row_cart'), response.row_cart, true);
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
    if(keyword !== null)
    {
        keyword = $('#field-keyword').val();
    }
    if(kategori_id !== null)
    {
        kategori_id = $('#field-kategori').val();
    }
    if( merk_id !== null)
    {
        merk_id = $('#field-merk').val();
    }
    var page = $('input#produk_page').val();
    if (page == 1) {
        $('div#product_list_body').html('');
    }
    $.ajax({
        method: 'GET',
        url: laroute.route('mitra.produk.json'),
        data: {
            keyword: keyword,
            kategori_id: kategori_id,
            merk_id: merk_id,
            page: page,
        },
        dataType: 'JSON',
        beforeSend: function(){
        },
        success: function(response) {
            $('div#product_list_body').html('');
            $('div#product_list_body').append(response.html);
        },
    });
}

function get_cart()
{
    $.ajax({
        method: 'GET',
        url: laroute.route('mitra.pos.getCart'),
        dataType: 'JSON',
        success: function(response) {
            $('#pos_table tbody').html(response.html);
            
            $('#total_subtotal').text(__currency_trans_from_en(response.sub_total, true, true));
            __write_number($('input#total_subtotal_input'), response.sub_total, true);

            $('#grand_total').text(__currency_trans_from_en(response.total, true, true));
            __write_number($('input#grand_total_hidden'), response.total, true);

            __write_number($('input#row_cart'), response.row_cart, true);
            
        },
    });
}

// function addRow