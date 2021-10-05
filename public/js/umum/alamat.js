jQuery(function() {
    var modal = $('#modalAlamat');
    _loadContent();

    var daerah = $('#field-desa').select2({
        placeholder: 'Cari Kelurahan',
        language: 'id',
        ajax: {
            url: laroute.route('wilayah.jsonSelect'),
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
        },
        
        minimumInputLength: 3,
        // templateResult: formatResult,
        templateResult: function(response) {
            if(response.loading)
            {
                return "Mencari...";
            }else{
                var selectionText = response.text.split(",");
                var $returnString = $('<span>'+selectionText[0] + ', ' + selectionText[1] + '</br>' + selectionText[2]+ ', ' + selectionText[3] +'</span>');
                return $returnString;
            }
        },
        templateSelection: function(response) {
            return response.text;
        },
    });

    $(document).on('click', '#btn-add_alamat', function () {  
        modal.find('h5.modal-title').html('Tambah Alamat Baru');
        modal.modal({
            backdrop: 'static',
            keyboard: false
        });
        $('#form-alamat')[0].reset();
    
        $("#form-alamat").submit(function (e) {
            e.preventDefault();
            var formData = new FormData($('#form-alamat')[0]);
            $.ajax({
                url: laroute.route('user.alamat.simpan'),
                type: 'POST',
                data: formData,
                cache: false,
                contentType: false,
                processData: false,
                beforeSend: function(){
                    Swal.fire({
                        title: 'Tunggu Sebentar...',
                        text: ' ',
                        imageUrl: laroute.url('public/img/loading.gif', ['']),
                        showConfirmButton: false,
                        allowOutsideClick: false,
                    });
                },
                success: function (response) {
                    $('.is-invalid').removeClass('is-invalid');
                    if (response.fail == false) {
                        Swal.fire({
                            title: "Berhasil",
                            text: "Alamat Baru Berhasil Ditambahkan",
                            timer: 3000,
                            showConfirmButton: false,
                            icon: 'success'
                        });
                        $('#modalAlamat').modal('hide');
                        $('#list-alamat').DataTable().ajax.reload();
                    } else {
                        Swal.close();
                        for (control in response.errors) {
                            $('#field-' + control).addClass('is-invalid');
                            $('#error-' + control).html(response.errors[control]);
                        }
                    }
                },
                error: function (jqXHR, textStatus, errorThrown) {
                    Swal.close();
                    alert('Error adding / update data');
                }
            });
        });
    });

    $(document).on('click', '.btn-edit_alamat', function () { 
        var id = $(this).attr('data-id');
        $.ajax({
            url: laroute.route('user.alamat.edit', { id: id }),
            type: "GET",
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
            success: function(data) {
                Swal.close();
                modal.modal({
                    backdrop: 'static',
                    keyboard: false
                });

                modal.find('h5.modal-title').html('Ubah Alamat');
                modal.find('input#field-id').val(data.id);
                modal.find('input#field-user_id').val(data.user_id);
                modal.find('input#field-nama').val(data.nama);
                modal.find('input#field-penerima').val(data.penerima);
                modal.find('input#field-phone').val(data.phone);
                modal.find('#field-alamat').val(data.alamat);
                modal.find('input#field-pos').val(data.kd_pos);
                daerah;
                sel_option = new Option(data.daerah, data.kelurahan_id, true, true);
                daerah.append(sel_option).trigger('change');
            },
            error: function(jqXHR, textStatus, errorThrown) {
                Swal.close();
                alert('Error deleting data');
            }
        });

        $("#form-alamat").on('submit', function (e) {
            e.preventDefault();
            var formData = new FormData($('#form-alamat')[0]);
            $.ajax({
                url: laroute.route('user.alamat.update'),
                type: 'POST',
                data: formData,
                cache: false,
                contentType: false,
                processData: false,
                beforeSend: function(){
                    Swal.fire({
                        title: 'Tunggu Sebentar...',
                        text: ' ',
                        imageUrl: laroute.url('public/img/loading.gif', ['']),
                        showConfirmButton: false,
                        allowOutsideClick: false,
                    });
                },
                success: function (response) {
                    $('.is-invalid').removeClass('is-invalid');
                    if (response.fail == false) {
                        Swal.fire({
                            title: "Berhasil",
                            text: "Alamat Berhasil Diperbaharui!",
                            timer: 3000,
                            showConfirmButton: false,
                            icon: 'success'
                        });
                        $('#modalAlamat').modal('hide');
                        $('#list-alamat').DataTable().ajax.reload();
                    } else {
                        Swal.close();
                        for (control in response.errors) {
                            $('#field-' + control).addClass('is-invalid');
                            $('#error-' + control).html(response.errors[control]);
                        }
                    }
                },
                error: function (jqXHR, textStatus, errorThrown) {
                    Swal.close();
                    alert('Error adding / update data');
                }
            });
        });
    });

    $(document).on('click', '.btn-hapus_alamat', function () { 
        id = $(this).attr('data-id');
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
        .then((result) => {
            if (result.value) {
            $.ajax({
                url: laroute.route('user.alamat.hapus', { id: id }),
                type: "GET",
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
                        text: 'Alamat Berhasil Dihapus!',
                        timer: 3000,
                        showConfirmButton: false,
                        icon: 'success'
                    });
                    $('#list-alamat').DataTable().ajax.reload();
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

function _loadContent(){

    var kategori = $('#kategori').val();
    var keyword = $('#search-data-list').val();
    var page = $('#current_page').val();

    var navPrev = $('#dataku-prev');
    var navNext = $('#dataku-next');
    
    $.ajax({
        url: laroute.route('user.alamat'),
        type: "GET",
        dataType: "JSON",
        data: {
            keyword: keyword,
            kategori_id : kategori,
            page: page,
        },
        beforeSend: function(){
            navNext.prop('disabled', true);
            navPrev.prop('disabled', true);
        },
        success: function(response) {
            var dataList = '';
            if(response.data.length !== 0){
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

            // Table Navigation
            response.next_page_url !== null ? navNext.prop('disabled', false) : navNext.prop('disabled', true);
            response.prev_page_url !== null ? navPrev.prop('disabled', false) : navPrev.prop('disabled', true);
            if(response.total === 0){
                var navigasi = '0 - 0 / 0';
            }else{
                var navigasi = response.from +' - '+ response.to +' / '+ response.total;
            }
            $('.dataku-nav .dataku-nav_info span').html(navigasi);
        },
        complete :function(){
            $('#dataku-loading').remove();
        },
        error: function(jqXHR, textStatus, errorThrown) {
            alert('Error deleting data');
        }
    });
};

function _createElement(item){
    var row = `
    <div class="block block-bordered block-rounded border-2x mb-0 `+ (item.is_utama == 1 ? `border-primary` : ``) +` ">
        <div class="block-header pb-0">
            <h3 class="block-title font-size-h6">
                <span class="font-w700">${ item.nama }</span>
                <div class="badge badge-secondary">Utama</div>
            </h3>
            <div class="block-options">
                <button type="button" class="btn btn-sm btn-secondary btn-edit_alamat" data-id="${ item.id }">Ubah</button>
                <button type="button" class="btn btn-sm btn-secondary">Hapus</button>
            </div>
        </div>
        <div class="block-content block-content-full">
            <div class="row justify-content-space-between">
                <div class="col-lg-8">
                    <h4 class="font-size-18 font-w700 mb-0 nice-copy">${ item.penerima }</h4>
                    <p class="mb-0 nice-copy">${ item.phone }</p>
                    <p class="mb-0 nice-copy">${ item.alamat }</p>
                </div>
                <div class="col-lg-4 my-auto text-right">`+
                    (item.is_utama == 0 ? `<button class="btn btn-secondary btn-sm">Jadikan Alamat Utama</button>` : ``)
                +`</div>
            </div>
        </div>
    </div>`;

    return row;
}
