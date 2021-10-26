$(function () {
    $("#loginForm").validate({
        onfocusout: function(element) {
            $(element).valid()
            if ($(element).valid()) {
                $('#loginForm').find('button:submit').prop('disabled', false)  
            } else {
                $('#loginForm').find('button:submit').prop('disabled', 'disabled')
            }
        },    
        errorClass: "invalid-feedback font-size-sm animated fadeInDown",
        errorElement: "div",
        errorPlacement: function (e, n) {
            jQuery(n).parents(".form-group").find('div.invalid-feedback').html(e)
        },
        highlight: function (e) {
            jQuery(e).closest(".form-group").removeClass("is-invalid").addClass("is-invalid")
        },
        success: function (e) {
            jQuery(e).closest(".form-group").removeClass("is-invalid").addClass("is-valid")
            jQuery(e).closest(".form-group").removeClass("is-invalid"), jQuery(e).remove()
        },
        rules: {
            email: {
                required: true,
                email: true
            },
            password: {
                required: true,
                minlength:6
            },
        },
        messages: {
            email: {
                required: "Alamat Email Wajib Diisi!",
                email: "Format Alamat Email Salah!"
            },
            password: {
                required: 'Password Wajib Diisi!',
                minlength: 'Password Kurang Dari 6 Karakter!'
            },
        },
        submitHandler: function (form) {
            var form = $('#loginForm')[0]
            var formData = new FormData(form)
            $.ajax({
                type: 'POST',
                url: '/admin/login',
                data: formData,
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
                    })
                },
                success: function (response) {
                    if (response.success) {
                        Swal.fire({
                            title: "Berhasil",
                            text: "Anda Akan Segera Dialihkan!",
                            timer: 3000,
                            showConfirmButton: false,
                            icon: 'success'
                        })
                        var user = JSON.stringify(response.data)
                        localStorage.setItem("logged", user)
                        window.setTimeout(function () {
                            location.reload()
                        }, 1500)
                    }else {
                        Swal.fire({
                            title: "Gagal",
                            text: "Periksa Form Input!",
                            timer: 3000,
                            showConfirmButton: false,
                            icon: 'error'
                        })
                        for (control in response.errors) {
                            $('#login-' + control).addClass('is-invalid')
                            $('#error-' + control).html(response.errors[control])
                        }
                    }
                },
                error : function(xhr, ajaxOptions, thrownError) {
                    var res = xhr.responseJSON
                    Swal.fire({
                        title: "Gagal",
                        text: res.message,
                        timer: 3000,
                        showConfirmButton: false,
                        icon: "error"
                    })
                }
            })
            return false
        }
    })
})

$(document).ready(function() {
    $("#show_hide_password a").on('click', function(event) {
        event.preventDefault()
        if($('#show_hide_password input').attr("type") == "text"){
            $('#show_hide_password input').attr('type', 'password')
            $('#show_hide_password i').addClass( "fa-eye-slash" )
            $('#show_hide_password i').removeClass( "fa-eye" )
        }else if($('#show_hide_password input').attr("type") == "password"){
            $('#show_hide_password input').attr('type', 'text')
            $('#show_hide_password i').removeClass( "fa-eye-slash" )
            $('#show_hide_password i').addClass( "fa-eye" )
        }
    })
})