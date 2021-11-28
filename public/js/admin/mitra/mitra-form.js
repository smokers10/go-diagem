jQuery(document).ready(function(){
    $('#form-mitra').submit(function(e){
        e.preventDefault()
        var data = new FormData($('#form-mitra')[0])
        $.ajax({
            url : "/admin/reseller/create",
            type: 'POST',
            data,
            cache: false,
            contentType: false,
            processData: false,
            beforeSend: function(){
                Swal.fire({
                    title: 'Tunggu Sebentar...',
                    text: ' ',
                    imageUrl: '/img/loading.gif',
                    showConfirmButton: false,
                    allowOutsideClick: false,
                });
            },
            success: function (res) {
                $('.is-invalid').removeClass('is-invalid')
                if (res.success) {
                    Swal.fire({
                        title: `Berhasil!`,
                        showConfirmButton: false,
                        icon: 'success',
                        html: `Reseller Baru Berhasil Disimpan!
                            <br><br>
                            <a href="`+ "/admin/reseller" +`"class="btn btn-keluar btn-alt-danger"><i class="si si-close mr-1"></i>Keluar</a> 
                            <a href="`+ "/admin/reseller/tambah" +`"class="btn btn-tambah_baru btn-alt-primary"><i class="si si-plus mr-1"></i>Tambah Reseller Lain</a>`,
                        showCancelButton: false,
                        showConfirmButton: false,
                        allowOutsideClick: false,
                    })
                } else {
                    alert("terjadi kesalahan silahkan hubungi developer")
                }
            },
            error: function (jqXHR, textStatus, errorThrown) {
                Swal.close();
                alert('Error adding / update data');
            }
        })
    })
})