
$(document).ready(function () {
    var daerahSelect = $('#field-daerah');
    daerahSelect.select2({
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

    $("#form-outlet").submit(function (e) {
        e.preventDefault();
        var formData = new FormData($('#form-outlet')[0]);
        $.ajax({
            url: laroute.route('mitra.outlet.simpan'),
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
                        // location.reload();
                        window.location.replace("http://www.w3schools.com");
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
            }
        });
    });

    $(document).on('change', '#field-daerah', function() {
        $.ajax({
            method: 'POST',
            url: laroute.route('getPos.json'),
            data: {
                daerah_id : $('#field-daerah').val(),
            },
            dataType: 'JSON',
            success: function(response) {
                $('#field-pos').val(response)
            },
        });
    });
});