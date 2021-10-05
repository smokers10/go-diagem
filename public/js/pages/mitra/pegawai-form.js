
    $(document).ready(function () {

        $('#wilayah').select2({
            placeholder: 'Cari Kelurahan',
            ajax: {
                url: laroute.route('wilayah.json'),
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
            templateResult: function(response) {
                var selectionText = response.text.split(",");
                var $returnString = $('<span>'+selectionText[0] + '</br>' + selectionText[1] + ', ' + selectionText[2]+ ', ' + selectionText[3] +'</span>');
                return $returnString;
            },
            templateSelection: function(response) {
                if(response.id === '')
                {
                    return 'Cari Kelurahan';
                }else{
                    return response.text;
                }
            },
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

        $("#form-pegawai").submit(function (e) {
            e.preventDefault();
            var formData = new FormData($('#form-pegawai')[0]);
            $.ajax({
                url: laroute.route('mitra.pegawai.simpan'),
                type: 'POST',
                data: formData,
                cache: false,
                contentType: false,
                processData: false,
                beforeSend: function(){
                    Swal.fire({
                        title: 'Tunggu Sebentar...',
                        text: 'Data Sedang Diproses',
                        imageUrl: laroute.url('assets/img/loading.gif', ['']),
                        showConfirmButton: false,
                        allowOutsideClick: false,
                    });
                },
                success: function (response) {
                    Swal.close();
                    $('.is-invalid').removeClass('is-invalid');
                    if (response.fail == false) {
                        $('#modal_embed').modal('hide');
                        Swal.fire({
                            title: "Berhasil",
                            text: "Toko Baru Berhasil Ditambahkan",
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
                    Swal.close();
                    alert('Error adding / update data');
                    $('#btnSimpan').text('Simpan');
                    $('#btnSimpan').attr('disabled', false);
                }
            });
        });
    });