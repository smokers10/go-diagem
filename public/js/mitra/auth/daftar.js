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

    var daerah = $('#field-wilayah').select2({
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

    $(document).on('change', '#daftar-terms', function() {
        if($(this).is(':checked'))
            $("#finish_btn").prop("disabled", false);
        else
            $("#finish_btn").prop("disabled", true);
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
            var formData = new FormData($('#form-pendaftaran')[0]);
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
                        // console.log()
                        // $.each(response.data, function(control) {
                        // });
                        $.each(response.data, function(index) {
                            // alert(index);
                            $('#preview-' + index).find('p').html(': '+response.data[index]);
                        });
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