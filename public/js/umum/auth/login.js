$(document).ready(function() {
    $("#form-login").submit(function (e) {
        e.preventDefault()
        var formData = new FormData($('#form-login')[0])
        $.ajax({
            url: "/login",
            type: 'POST',
            data: formData,
            cache: false,
            contentType: false,
            processData: false,
            beforeSend: function(){
                Swal.fire({
                    title: 'Tunggu Sebentar...',
                    text: 'Data Sedang Diproses!',
                    imageUrl: '/img/loading.gif',
                    showConfirmButton: false,
                    allowOutsideClick: false,
                })
            },
            success: function (res) {
                if (res.success) {
                    Swal.fire({
                        title: "Berhasil",
                        text: "Proses Masuk Ke Akun Kamu Berhasil!",
                        timer: 3000,
                        showConfirmButton: false,
                        icon: 'success'
                    })
                    window.setTimeout(function () {
                        location.reload()
                    }, 1500)
                    var user = JSON.stringify(res.data)
                    localStorage.setItem("logged", user)
                }else {
                    Swal.close()
                }
            },
            error: function (jqXHR, textStatus, errorThrown) {
                Swal.close()
                alert('Error adding / update data')
            }
        })
    })
})
