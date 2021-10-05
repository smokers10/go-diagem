$(document).ready(function () {
    
    $.ajaxSetup({
        headers: {
            'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
        }
    });

    $("#field-bisnis_kategori").select2({
        placeholder: 'Pilih Kategori Bisnis',
        allowClear: true,
        ajax: {
            url: laroute.route('bisnisKategori.json'),
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

    var daerah = $('#field-bisnis_daerah').select2({
        placeholder: 'Cari Kelurahan',
        language: 'id',
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
        
        minimumInputLength: 3,
        // templateResult: formatResult,
        templateResult: function(response) {
            if(response.loading)
            {
                return "Mencari...";
            }else{
                var selectionText = response.text.split(",");
                var $returnString = $('<span>'+selectionText[0] + '</br>' + selectionText[1] + ', ' + selectionText[2]+ ', ' + selectionText[3] +'</span>');
                return $returnString;
            }
        },
        templateSelection: function(response) {
            return response.text;
        },
    });
    // Get Kode POS
    $(document).on('change', '#field-bisnis_daerah', function() {
        $.ajax({
            method: 'POST',
            url: laroute.route('getPos.json'),
            data: {
                daerah_id : $('#field-bisnis_daerah').val(),
            },
            dataType: 'JSON',
            success: function(response) {
                $('#field-pos').val(response)
            },
        });
    });

    $('#pendaftaran').bootstrapWizard({
        tabClass: 'nav nav-tabs',
        nextSelector: '[data-wizard="next"]',
        previousSelector: '[data-wizard="prev"]',
        firstSelector : '[data-wizard="first"]',
        lastSelector : '[data-wizard="lsat"]',
        finishSelector : '[data-wizard="finish"]',
        backSelector : '[data-wizard="back"]',
		onTabClick: function(tab, navigation, index) {
            return false;
        },
		onNext: function(tab, navigation, index) {
            formData = new FormData();
            formData.append('bisnis_nama',$('#field-bisnis_nama').val());
            if($('#field-bisnis_kategori').val() !== null)
            {
                formData.append('bisnis_kategori',$('#field-bisnis_kategori').val());
            }
            if($('#field-bisnis_daerah').val() !== null)
            {
                formData.append('bisnis_daerah',$('#field-bisnis_daerah').val());
            }
            formData.append('pos',$('#field-pos').val());
            formData.append('bisnis_alamat',$('#field-bisnis_alamat').val());
            var status = false;
            $.ajax({
                url: laroute.route('mitra.daftarStep1'),
                type: 'POST',
                data: formData,
                async: false,
                cache: false,
                contentType: false,
                processData: false,
                success: function (response) {
                    $('.is-invalid').removeClass('is-invalid');
                    if (response.fail == false) {
                        status = true;
                    } else {
                        status = false;
                        for (control in response.errors) {
                            $('#field-' + control).addClass('is-invalid');
                            $('#error-' + control).html(response.errors[control]);
                        }
                    }
                }
            });
            return status;
        },
        onFinish: function(tab, navigation, index) {
  		}
  });
  
    $("#form-pendaftaran").submit(function (e) {
        e.preventDefault();
        var formData = new FormData($('#form-pendaftaran')[0]);
        $.ajax({
            url: laroute.route('mitra.daftarStep2'),
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
                        text: "Supplier Baru Berhasil Ditambahkan",
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
                Swal.close();
                alert('Error adding / update data');
                $('#btnSimpan').text('Simpan');
                $('#btnSimpan').attr('disabled', false);
            }
        });
    });
});