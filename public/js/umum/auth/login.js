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
                var titleText 
                var iconType
                res.success ? titleText = "Berhasil" : titleText = "Gagal"
                res.success ? iconType = "success" : iconType = "error"

                Swal.fire({
                    title: titleText,
                    text: res.message,
                    timer: 3000,
                    showConfirmButton: false,
                    icon: iconType
                })

                if (res.success) {
                    var user = JSON.stringify(res.data)
                    localStorage.setItem("logged", user)
                    window.setTimeout(function () {
                        location.reload()
                    }, 1500)
                }
            },
            error : function(xhr, ajaxOptions, thrownError) {
                var res = xhr.responseJSON
                Swal.fire({
                    title: "Opps",
                    text: res.message,
                    timer: 3000,
                    showConfirmButton: false,
                    icon: "error"
                })
            }
        })
    })
})
