$(document).ready(function() {
    $("#form-register").submit(function (e) {
        e.preventDefault()
        var formData = new FormData($('#form-register')[0])
        $.ajax({
            url: "/register",
            type: 'POST',
            data: formData,
            cache: false,
            contentType: false,
            processData: false,
            beforeSend: function(){
                Swal.fire({
                    title: 'Tunggu Sebentar...',
                    text: '',
                    imageUrl: '/img/loading.gif',
                    showConfirmButton: false,
                    allowOutsideClick: false,
                })
            },
            success: function (response) {
                if (response.success) {
                    Swal.fire({
                        title: "Berhasil",
                        text: "Panduan Verifikasi Akun Sudah Dikirim Ke Email",
                        timer: 3000,
                        showConfirmButton: false,
                        icon: 'success'
                    })
                    if (response.success) {
                        var user = JSON.stringify(response.data)
                        localStorage.setItem("logged", user)
                        window.setTimeout(function () {
                            location.reload()
                        }, 1500)
                    }
                } else {
                    Swal.close()
                    for (control in response.errors) {
                        $('#reg-' + control).addClass('is-invalid')
                        $('#reg_error-' + control).html(response.errors[control])
                    }
                }
            },
            error: function (jqXHR, textStatus, errorThrown) {
                Swal.close()
                alert('Error adding / update data')
            }
        })
    })
})