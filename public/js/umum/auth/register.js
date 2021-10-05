$(document).ready(function() {
    $("#form-register").submit(function (e) {
        e.preventDefault();
        var formData = new FormData($('#form-register')[0]);
        $.ajax({
            url: laroute.route('registerPost'),
            type: 'POST',
            data: formData,
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
                if (response.fail == false) {
                    Swal.fire({
                        title: "Berhasil",
                        text: "Panduan Verifikasi Akun Sudah Dikirim Ke Email",
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
                        $('#reg-' + control).addClass('is-invalid');
                        $('#reg_error-' + control).html(response.errors[control]);
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