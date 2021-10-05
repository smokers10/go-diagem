$(document).ready(function () {
    load_variasi();
    var croppie = null;
    var el = document.getElementById('resizer');
    var formData = new FormData();
    $.base64ImageToBlob = function(str) {
        // extract content type and base64 payload from original string
        var pos = str.indexOf(';base64,');
        var type = str.substring(5, pos);
        var b64 = str.substr(pos + 8);

        // decode base64
        var imageContent = atob(b64);

        // create an ArrayBuffer and a view (as unsigned 8-bit)
        var buffer = new ArrayBuffer(imageContent.length);
        var view = new Uint8Array(buffer);

        // fill the view, using the decoded base64
        for (var n = 0; n < imageContent.length; n++) {
          view[n] = imageContent.charCodeAt(n);
        }

        // convert ArrayBuffer to Blob
        var blob = new Blob([buffer], { type: type });

        return blob;
    }

    $.getImage = function(input, croppie) {
        if (input.files && input.files[0]) {
            var reader = new FileReader();
            reader.onload = function(e) {
                croppie.bind({
                    url: e.target.result,
                });
            }
            reader.readAsDataURL(input.files[0]);
        }
    }

    $("#file-upload").on("change", function(event) {
        $("#cropModal").modal();
        croppie = new Croppie(el, {
                viewport: {
                    width: 350,
                    height: 350,
                    type: 'square'
                },
                original : {
                    width: 350,
                    height: 350,
                    type: 'square'
                },
                boundary: {
                    width: 360,
                    height: 360
                },
                enableOrientation: true
            });
        $.getImage(event.target, croppie);
    });

    $("#upload").on("click", function() {
        croppie.result({
            type: 'base64',
            size: 'original',
			format:'jpeg',
			size: { 
                width: 700, height: 700 
            }
        }).then(function(base64) {
            $("#cropModal").modal("hide");
            $("#img_preview").attr("src", base64);

            formData.append("foto", $.base64ImageToBlob(base64));
        });
    });

    $(".rotate").on("click", function() {
        croppie.rotate(parseInt($(this).data('deg')));
    });

    $('#cropModal').on('hidden.bs.modal', function (e) {
        setTimeout(function() { croppie.destroy(); }, 100);
    });

    merkSelect = $("#field-merk").select2({
        placeholder: 'Pilih Merk',
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
    });

    merkOption = new Option($('#field-merk').attr("data-name"), $('#field-merk').attr("data-id"), true, true);
    merkSelect.append(merkOption).trigger('change');

    kategoriSelect = $("#field-kategori").select2({
        placeholder: 'Pilih Kategori',
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
    });

    kategoriOption = new Option($('#field-kategori').attr("data-name"), $('#field-kategori').attr("data-id"), true, true);
    kategoriSelect.append(kategoriOption).trigger('change');

    unitSelect = $("#field-unit").select2({
        placeholder: 'Pilih Unit Satuan',
        ajax: {
            url: laroute.route('mitra.satuan.json'),
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
    });

    unitOption = new Option($('#field-unit').attr("data-text"), $('#field-unit').attr("data-id"), true, true);
    unitSelect.append(unitOption).trigger('change');

    $('#kelola_stok').click(function(){
        if($(this).prop("checked") == true){
            // console.log("Checkbox is checked.");
            // $('.inventaris').html('');
            $('.inventaris').show();

        }
        else if($(this).prop("checked") == false){
            // console.log("Checkbox is unchecked.");
            $('.inventaris').hide();
        }
    });

    $('#variasi_kelola_stok').click(function(){
        if($(this).prop("checked") == true){
            // console.log("Checkbox is checked.");
            // $('.inventaris').html('');
            $('.variasi_inventaris').show();

        }
        else if($(this).prop("checked") == false){
            // console.log("Checkbox is unchecked.");
            $('.variasi_inventaris').hide();
        }
    });

    $('#ubah_variasi_kelola_stok').click(function(){
        if($(this).prop("checked") == true){
            // console.log("Checkbox is checked.");
            // $('.inventaris').html('');
            $('.ubah_variasi_inventaris').show();

        }
        else if($(this).prop("checked") == false){
            // console.log("Checkbox is unchecked.");
            $('.ubah_variasi_inventaris').hide();
        }
    });

    $(document).on('click', '.btnVariasi', function() {
        $('#form-variasi')[0].reset();
        $('.variasi_inventaris').show();
        $('#modal_variasi').modal({
            backdrop: 'static',
            keyboard: false
        })
    });

    $("#form-variasi").submit(function (e){
        e.preventDefault();
        var formData = new FormData($('#form-variasi')[0]);
        if($('#has_variasi').val() == 1)
        {
            formData.append('old_sku', $('#field-sku').val());
            formData.append('old_hrg_modal', $('#field-hrg_modal').val());
            formData.append('old_hrg_jual', $('#field-hrg_jual').val());
            if($('#kelola_stok').prop("checked") == true){
                formData.append('old_min_stok', $('#field-min_stok').val());
                formData.append('old_satuan_id', $('#field-satuan_id').val());
                formData.append('old_satuan_nama', $('#field-satuan_id').attr("data-name"));
                formData.append('old_kelola_stok', 1);
            }
            else if($('#kelola_stok').prop("checked") == false){
                formData.append('old_kelola_stok', 0);
            }
        }
        $.ajax({
            url: laroute.route('mitra.variasi.tambah'),
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
            complete: function(){
                Swal.close();
            },
            success: function (response) {
                $('.is-invalid').removeClass('is-invalid');
                $('.text-danger').html('');
                if (response.fail == true) {
                    for (control in response.errors) {
                        $('#field-' + control).addClass('is-invalid');
                        $('#error-' + control).html(response.errors[control]);
                    }
                }else{
                    if ($('#count_variasi').val() == '0') {
                        $('#modal_variasi').modal('hide');
                        $('.tanpa_variasi').hide();
                        $(response).insertAfter(".info_produk");
                        $('#count_variasi').val(1);
                    } else {
                        $('#modal_variasi').modal('hide');
                        $('#tbl_variasi tbody').append(response);
                        $('#count_variasi').val(parseInt($('#count_variasi').val()) + 1);
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

    $('.btnAddUnit').click(function (e) {
        $('#form-satuan')[0].reset();
        $('#modal_addUnit').modal({
            backdrop: 'static',
            keyboard: false
        })
    });

    $("#form-kategori").submit(function (e) {
        e.preventDefault();
        var formData = new FormData($('#form-kategori')[0]);
        $.ajax({
            url: laroute.route('mitra.kategori.simpan'),
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
            complete: function(){
                Swal.close();
            },
            success: function (response) {
                $('.is-invalid').removeClass('is-invalid');
                if (response.fail == false) {
                    $('#modal_embed').modal('hide');
                    Swal.fire({
                        title: "Berhasil",
                        text: "Kategori Baru Berhasil Ditambahkan",
                        timer: 3000,
                        showConfirmButton: false,
                        icon: 'success'
                    });
                    window.setTimeout(function () {
                        location.reload();
                    }, 1500);
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

    $("#form-satuan").submit(function (e) {
        e.preventDefault();
        var formData = new FormData($('#form-satuan')[0]);
        $.ajax({
            url: laroute.route('mitra.satuan.simpan'),
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
            complete: function(){
                Swal.close();
            },
            success: function (response) {
                $('.is-invalid').removeClass('is-invalid');
                if (response.fail == false) {
                    $('#modal_addUnit').modal('hide');
                    Swal.fire({
                        title: "Berhasil",
                        text: "Satuan Unit Label Baru Berhasil Ditambahkan",
                        timer: 3000,
                        showConfirmButton: false,
                        icon: 'success'
                    });
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

    $(document).on('click', '.btn-hapus', function() {
        Swal.fire({
            title: "Anda Yakin?",
            text: "Data Yang Dihapus Tidak Akan Bisa Dikembalikan",
            icon: "warning",
            showCancelButton: true,
            confirmButtonText: 'Ya, Hapus!',
            cancelButtonText: 'Tidak, Batalkan!',
            reverseButtons: true,
            allowOutsideClick: false,
            confirmButtonColor: '#af1310',
            cancelButtonColor: '#fffff',
        })
        .then(value => {
            if (value) {
                // alert();
                dihapus = $(this).closest('tr').attr("data-variasi_id");
                $('#total_variasi').after('<input type="hidden" name="variasihapus[]" value="'+ dihapus +'">');
                $(this).closest('tr').remove();
                // $('#count_variasi').val($('#count_variasi').val() - 1);
                // if($('#count_variasi').val() == 0)
                // {
                //     $('.sh_variasi').remove();
                //     $('.tanpa_variasi').show();
                // }
            }
        });
    });

    $("#form-produk").submit(function (e) {
        e.preventDefault();
        // var formData = new FormData($('#form-produk')[0]);
        var poData = jQuery(document.forms['form-produk']).serializeArray();
        for (var i=0; i<poData.length; i++)
        {
            formData.append(poData[i].name, poData[i].value);
        }
        $.ajax({
            url: laroute.route('mitra.produk.update'),
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
                $('.is-invalid').removeClass('is-invalid');
                if (response.fail == false) {
                    Swal.fire({
                        title: "Berhasil",
                        text: "Data Produk Berhasil Diperbaharui!",
                        timer: 3000,
                        showConfirmButton: false,
                        icon: 'success'
                    });
                    window.setTimeout(function () {
                        location.reload();
                    }, 1500);
                } else {
                    Swal.close();
                    for (control in response.errors) {
                        $('#field-' + control).addClass('is-invalid');
                        $('#error-' + control).html(response.errors[control]);
                    }
                }
            },
            error: function (jqXHR, textStatus, errorThrown) {
                alert('Error adding / update data');
                Swal.close();
            }
        });
    });

    $("#form-ubah_variasi").submit(function (e) {
        e.preventDefault();
        var formData = new FormData($('#form-ubah_variasi')[0]);
        $.ajax({
            url: laroute.route('mitra.variasi.changeRow'),
            type: 'POST',
            data: formData,
            cache: false,
            contentType: false,
            processData: false,
            success: function (response) {
                $('.is-invalid').removeClass('is-invalid');
                if (response.fail == true) {
                    for (control in response.errors) {
                        $('#field-' + control).addClass('is-invalid');
                        $('#error-' + control).html(response.errors[control]);
                    }
                } else {
                    $('input[name^="variasi['+response.row+'][nama]"]').val(response.nama);
                    $('input[name^="variasi['+response.row+'][sku]"]').val(response.sku);
                    $('input[name^="variasi['+response.row+'][hrg_modal]"]').val(response.hrg_modal);
                    $('input[name^="variasi['+response.row+'][hrg_jual]"]').val(response.hrg_jual);
                    $('input[name^="variasi['+response.row+'][show_inventaris]"]').val(response.show_inventaris);
                    $('input[name^="variasi['+response.row+'][kelola_stok]"]').val(response.kelola_stok);
                    $('input[name^="variasi['+response.row+'][satuan_id]"]').val(response.satuan_id);
                    $('input[name^="variasi['+response.row+'][satuan_nama]"]').val(response.satuan_nama);
                    $('#modal_ubah_variasi').modal('hide');

                    Swal.fire({
                        title: "Berhasil",
                        text: "Variasi Produk Berhasil Diperbaharui",
                        timer: 3000,
                        showConfirmButton: false,
                        icon: 'success'
                    });
                }
            },
            error: function (jqXHR, textStatus, errorThrown) {
                alert('Error adding / update data');
                $('#btnSimpan').text('Simpan');
                $('#btnSimpan').attr('disabled', false);

            }
        });
    });

});

function load_variasi()
{
    if($("#has_variasi").val() > 1){
        $('.tanpa_variasi').hide();
    }
}

function ubah(id)
{
    $('#form-ubah_variasi')[0].reset();
    $('#row_ubah_variasi').val(id);
    if($('input[name^="variasi['+id+'][kelola_stok]"]').val() == 1)
    {
        $('#ubah_variasi_kelola_stok').prop("checked", true);
        $('.ubah_variasi_inventaris').show();
    }else{
        $('#ubah_variasi_kelola_stok').prop("checked", false);
        $('.ubah_variasi_inventaris').hide();
    }
    $('#field-ubah_variasi_nama').val($('input[name^="variasi['+id+'][nama]"]').val());
    $('#field-ubah_variasi_sku').val($('input[name^="variasi['+id+'][sku]"]').val());
    $('#field-ubah_variasi_hrg_modal').val($('input[name^="variasi['+id+'][hrg_modal]"]').val());
    $('#field-ubah_variasi_hrg_jual').val($('input[name^="variasi['+id+'][hrg_jual]"]').val());

    var satuanSelect = $('#field-ubah_variasi_unit');
    var option = new Option($('input[name^="variasi['+id+'][satuan_nama]"]').val(), $('input[name^="variasi['+id+'][satuan_id]"]').val(), true, true);
    satuanSelect.append(option).trigger('change');

    $('#modal_ubah_variasi').modal({
        backdrop: 'static',
        keyboard: false
    });
}