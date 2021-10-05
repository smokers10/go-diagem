
jQuery(function() {
    var croppie = null;
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

    $("#uploadThumb").on("change", function(event) {
        readURL(this);  
    });

    function readURL(input) {    
        if (input.files && input.files[0]) {   
            var reader = new FileReader();
            var filename = $("#uploadThumb").val();
            filename = filename.substring(filename.lastIndexOf('\\')+1);
            reader.onload = function(e) {
            $('#thumbPrev').attr('src', e.target.result);
            $('.custom-file-label').text(filename);
            }
            reader.readAsDataURL(input.files[0]);    
        }
    }

    $(".rotate").on("click", function() {
        croppie.rotate(parseInt($(this).data('deg')));
    });

    $('#cropModal').on('hidden.bs.modal', function (e) {
        setTimeout(function() { croppie.destroy(); }, 100);
    });

    oTable = $('#list-rekening').DataTable({
        processing: true,
        serverSide: true,
        ajax: laroute.route('admin.rekening'),
        ordering: false,
        columns: [
            {
                data: 'bank',
                name: 'bank'
            },
            {
                data: 'kode',
                name: 'kode'
            },
            {
                data: 'rekening',
                name: 'rekening'
            },
            {
                data: 'nama',
                name: 'nama'
            },
            {
                data: 'aksi',
                name: 'aksi',
                orderable: false,
                searchable: false
            },
        ]
    });

    $('#cari_produk').keyup(function () {
        oTable.search($(this).val()).draw();
    });
    
    var modal = $('#modalRekening');

    $(document).on('click', '#btn-add_rekening', function () {  
        modal.find('h3.modal-title').html('Tambah Rekening Baru');
        modal.modal({
            backdrop: 'static',
            keyboard: false
        });
        $('#form-rekening')[0].reset();
        $('#form-rekening').find('.form-control').removeClass('is-invalid');
    
        $("#form-rekening").submit(function (e) {
            e.preventDefault();
            var formData = new FormData($('#form-rekening')[0]);
            $.ajax({
                url: laroute.route('admin.rekening.simpan'),
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
                            text: "Rekening Baru Berhasil Ditambahkan",
                            timer: 3000,
                            showConfirmButton: false,
                            icon: 'success'
                        });
                        $('#modalRekening').modal('hide');
                        $('#list-etlase').DataTable().ajax.reload();
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


    $(document).on('click', '.btn-edit', function () { 
        var id = $(this).attr('data-id');
        $.ajax({
            url: laroute.route('admin.rekening.edit', { id: id }),
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

                modal.find('h5.modal-title').html('Ubah Rekening');
                modal.find('input#field-id').val(data.id);
                modal.find('input#field-bank').val(data.bank);
                modal.find('input#field-kode').val(data.kode);
                modal.find('input#field-rekening').val(data.rekening_no);
                modal.find('input#field-nama').val(data.nama);

                if(data.icon === null)
                {
                    $("#thumbPrev").attr("src", "https://via.placeholder.com/128x64.png?text=ICON+BANK");
                }else{
                    $("#thumbPrev").attr("src", data.icon_url);
                }
            },
            error: function(jqXHR, textStatus, errorThrown) {
                Swal.close();
                alert('Error deleting data');
            }
        });

        $("#form-rekening").submit(function (e) {
            e.preventDefault();
            var formData = new FormData($('#form-rekening')[0]);
            $.ajax({
                url: laroute.route('admin.rekening.update'),
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
                            text: "Rekening Berhasil Diperbaharui!",
                            timer: 3000,
                            showConfirmButton: false,
                            icon: 'success'
                        });
                        $('#modalRekening').modal('hide');
                        oTable.ajax.reload();
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


    $(document).on('click', '.btn-hapus', function () { 
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
                url: laroute.route('admin.rekening.hapus', { id: id }),
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
                        text: 'Rekening Berhasil Dihapus!',
                        timer: 3000,
                        showConfirmButton: false,
                        icon: 'success'
                    });
                    $('#list-rekening').DataTable().ajax.reload();
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
