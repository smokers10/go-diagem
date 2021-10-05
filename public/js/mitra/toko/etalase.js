
$(document).ready(function () {
    oTable = $('#list-etalase').DataTable({
        processing: true,
        serverSide: true,
        ajax: laroute.route('mitra.etalase'),
        ordering: false,
        columns: [
            {
                data: 'nama',
                name: 'nama'
            },
            {
                data: 'produk',
                name: 'produk'
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
    
    var modal = $('#modalEtalase');

    $(document).on('click', '#btn-add_etalase', function () {  
        modal.find('h3.modal-title').html('Tambah Etalase Baru');
        modal.modal({
            backdrop: 'static',
            keyboard: false
        });
        $('#form-etalase')[0].reset();
    
        $("#form-etalase").submit(function (e) {
            e.preventDefault();
            var formData = new FormData($('#form-etalase')[0]);
            $.ajax({
                url: laroute.route('mitra.etalase.simpan'),
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
                    $('.is-invalid').removeClass('is-invalid');
                    if (response.fail == false) {
                        Swal.fire({
                            title: "Berhasil",
                            text: "Etalase Baru Berhasil Ditambahkan",
                            timer: 3000,
                            showConfirmButton: false,
                            icon: 'success'
                        });
                        $('#modalEtalase').modal('hide');
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
            url: laroute.route('mitra.etalase.edit', { id: id }),
            type: "GET",
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
            success: function(data) {
                Swal.close();
                modal.modal({
                    backdrop: 'static',
                    keyboard: false
                });

                modal.find('h5.modal-title').html('Ubah Etalase');
                modal.find('input#field-id').val(data.id);
                modal.find('input#field-nama').val(data.nama);
            },
            error: function(jqXHR, textStatus, errorThrown) {
                Swal.close();
                alert('Error deleting data');
            }
        });

        $("#form-etalase").submit(function (e) {
            e.preventDefault();
            var formData = new FormData($('#form-etalase')[0]);
            $.ajax({
                url: laroute.route('mitra.etalase.update'),
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
                    $('.is-invalid').removeClass('is-invalid');
                    if (response.fail == false) {
                        Swal.fire({
                            title: "Berhasil",
                            text: "Etalase Berhasil Diperbaharui!",
                            timer: 3000,
                            showConfirmButton: false,
                            icon: 'success'
                        });
                        $('#modalEtalase').modal('hide');
                        $('#list-etalase').DataTable().ajax.reload();
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
                url: laroute.route('mitra.etalase.hapus', { id: id }),
                type: "GET",
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
                success: function() {
                    Swal.fire({
                        title: "Berhasil",
                        text: 'Etalase Berhasil Dihapus!',
                        timer: 3000,
                        showConfirmButton: false,
                        icon: 'success'
                    });
                    $('#list-etalase').DataTable().ajax.reload();
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
