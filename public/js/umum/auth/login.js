$(document).ready(function() {
    
    $("#form-login").submit(function (e) {
        e.preventDefault();
        var formData = new FormData($('#form-login')[0]);
        $.ajax({
            url: laroute.route('loginPost'),
            type: 'POST',
            data: formData,
            cache: false,
            contentType: false,
            processData: false,
            beforeSend: function(){
                Swal.fire({
                    title: 'Tunggu Sebentar...',
                    text: 'Data Sedang Diproses!',
                    imageUrl: laroute.url('public/img/loading.gif', ['']),
                    showConfirmButton: false,
                    allowOutsideClick: false,
                });
            },
            success: function (response) {
                if (response.fail == false) {
                    if(response.access == true)
                    {
                        Swal.fire({
                            title: "Berhasil",
                            text: "Proses Masuk Ke Akun Kamu Berhasil!",
                            timer: 3000,
                            showConfirmButton: false,
                            icon: 'success'
                        });
                        window.setTimeout(function () {
                            location.reload();
                        }, 1500);
                    }else{
                        window.setTimeout(function () {
                            location.reload();
                        }, 1500);
                    }
                }else {
                    Swal.close();
                    for (control in response.errors) {
                        $('#field_login-' + control).addClass('is-invalid');
                        $('#error_login-' + control).html(response.errors[control]);
                    }
                }
            },
            error: function (jqXHR, textStatus, errorThrown) {
                Swal.close();
                alert('Error adding / update data');
            }
        });
    });
    // $("form#form-register").validate({
    //     errorClass: "invalid-feedback",
    //     errorElement: "div",
    //     errorPlacement: function (e, r) {
    //         jQuery(r).parents(".form-group > div").append(e);
    //     },
    //     highlight: function (e) {
    //         jQuery(e).closest(".form-group").removeClass("is-invalid").addClass("is-invalid");
    //     },
    //     success: function (e) {
    //         jQuery(e).closest(".form-group").removeClass("is-invalid"), jQuery(e).remove();
    //     },
    //   rules: {
    //     nama_lengkap: "required",
    //     email: {
    //       required: true,
    //       email: true
    //     },
    //     password: {
    //       required: true,
    //       minlength: 5
    //     }
    //   },
    //   messages: {
    //     nama_lengkap: "Masukan Nama Lengkap",
    //     password: {
    //       required: "Please provide a password",
    //       minlength: "Your password must be at least 5 characters long"
    //     },
    //     email: "Please enter a valid email address"
    //   },
    //   submitHandler: function(form) {
    //     form.submit();
    //   }
    // });

});
