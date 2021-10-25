$(document).ready(function(){
    // get logged user
    const logged = JSON.parse(localStorage.getItem("logged")) 
    $("#logged-user-name").text(logged.nama)

    // resend OTP
    $("#resend-button").click(()=>{
        $.ajax({
            url: "/verifikasi/create",
            method: "post",
            cache: false,
            contentType: false,
            processData: false,
            beforeSend: function(){
                Swal.fire({
                    title: 'Tunggu Sebentar...',
                    text: 'Kode Verifikasi Sedang Diproses!',
                    imageUrl: '/img/loading.gif',
                    showConfirmButton: false,
                    allowOutsideClick: false,
                })
            },
            success: function (res) {
                if (res.success) {
                    Swal.fire({
                        title: "Berhasil",
                        text: "Kami Berhasil Mengirim Kode Verifikasi Baru, Silahkan Check Inbox Email Kamu",
                        timer: 3000,
                        showConfirmButton: false,
                        icon: 'success'
                    })
                }else {
                    Swal.close()
                }
            },
            error: function (jqXHR, textStatus, errorThrown) {
                Swal.close()
                alert('Error Create Verification Code')
            }
        })
    })

    // verifikasi
    $("#verificate-button").click(function(e){
        var formData = new FormData($('#form-verifikasi')[0])
        $.ajax({
            url: "/verifikasi/verificate",
            method: "post",
            data: formData,
            cache: false,
            contentType: false,
            processData: false,
            beforeSend: function(){
                Swal.fire({
                    title: 'Tunggu Sebentar...',
                    text: 'Proses Verifikasi Akun!',
                    imageUrl: '/img/loading.gif',
                    showConfirmButton: false,
                    allowOutsideClick: false,
                })
            },
            success: function (res) {
                if (res.success) {
                    Swal.fire({
                        title: "Berhasil",
                        text: "Akun Kamu Berhasil Diverifikasi!",
                        timer: 3000,
                        showConfirmButton: false,
                        icon: 'success'
                    })
                    setTimeout(() => {
                        location.reload()
                    }, 3000)
                }

                else if (!res.success && res.status != 500) {
                    Swal.fire({
                        title: "Kode Verifikasi Salah",
                        text: "",
                        timer: 3000,
                        showConfirmButton: false,
                        icon: 'error'
                    })
                }

                else {
                    Swal.close()
                }
            },
            error: function (jqXHR, textStatus, errorThrown) {
                Swal.close()
                alert('Error Create Verification Code')
            }
        })
    })
})