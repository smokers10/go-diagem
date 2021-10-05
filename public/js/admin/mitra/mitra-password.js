jQuery(document).ready(function () {
    $('#mitra-password').submit(function(e){
        e.preventDefault()
        var data = new FormData($('#mitra-password')[0])
        $.ajax({
            url : "http://localhost/re-diagem/admin/reseller/update-password",
            type: 'POST',
            data,
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
                        title: `Berhasil!`,
                        showConfirmButton: false,
                        icon: 'success',
                        html: `Password Reseller Berhasil Diupdate!
                            <br><br>
                            <a href="`+ "http://localhost/re-diagem/admin/reseller" +`" class="btn btn-keluar btn-alt-danger"><i class="si si-close mr-1"></i>Keluar</a>`,
                        showCancelButton: false,
                        showConfirmButton: false,
                        allowOutsideClick: false,
                    });
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
        })
    })
})