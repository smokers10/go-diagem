jQuery(document).ready(function () {
    moment.locale('id');
    $('#tgl_range').daterangepicker({
        locale: {
            format: 'DD MMMM YYYY'
        }
    }, function(start, end, label) {
        $('input#tgl_mulai').val(start.format('YYYY-MM-DD'));
        $('input#tgl_selesai').val(end.format('YYYY-MM-DD'));
    });

    var croppie = null;
    var el = document.getElementById('resizer');

    $.getImage = function(input, croppie) {
        if (input.files && input.files[0]) {
            var reader = new FileReader();
            reader.onload = function(e) {
                croppie.bind({
                    url: e.target.result,
                });
            }
            reader.readAsDataURL(input.files[0]);
        }
    }

    $("#file-upload").on("change", function(event) {
        $("#cropModal").modal();
        croppie = new Croppie(el, {
            viewport: {
                width: 700,
                height: 233.33,
                type: 'square'
            },
            original : {
                width: 700,
                height: 233.33,
                type: 'square'
            },
            boundary: {
                width: 730,
                height: 263.33,
            },
            enableOrientation: true
        });
        $.getImage(event.target, croppie);
    });

    $("#sampulSave").on("click", function() {
        croppie.result({
            type: 'base64',
            size: 'original',
			format:'jpeg',
			size: { 
                width: 1200, height: 400 
            }
        }).then(function(base64) {
            $("#cropModal").modal("hide");
            $("#img_preview").attr("src", base64);
            $("#image").val(base64);
        });
    });

    $(".rotate").on("click", function() {
        croppie.rotate(parseInt($(this).data('deg')));
    });

    $('#cropModal').on('hidden.bs.modal', function (e) {
        setTimeout(function() { croppie.destroy(); }, 100);
    });

    $('#field-deskripsi').summernote({
        height: '250px'
    });

    $("#form-promo").submit(function (e) {
        e.preventDefault();
        var formData = new FormData($('#form-promo')[0]);
        $.ajax({
            url: laroute.route('admin.promo.update'),
            type: 'POST',
            data: formData,
            cache: false,
            contentType: false,
            processData: false,
            beforeSend: function(){
                Swal.fire({
                    title: 'Tunggu Sebentar...',
                    text: 'Data Sedang Diproses!',
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
                        title: `Berhasil!`,
                        showConfirmButton: false,
                        icon: 'success',
                        html: `Promo Berhasil Berhasil Diperbaharui!
                            <br><br>
                            <a href="`+ laroute.route('admin.promo') +`" class="btn btn-keluar btn-alt-danger"><i class="si si-close mr-1"></i>Keluar</a> 
                            <a href="`+ laroute.route('admin.promo.tambah') +`" class="btn btn-tambah_baru btn-alt-primary"><i class="si si-plus mr-1"></i>Tambah Promo Lain</a>`,
                        showCancelButton: false,
                        showConfirmButton: false,
                    });
                } else {
                    Swal.close();
                    for (control in response.errors) {
                        $('#field-' + control).addClass('is-invalid');
                        $('#error-' + control).html(response.errors[control]);
                        $.notify({
                            // options
                            icon: 'fa fa-times',
                            message: response.errors[control]
                        },{
                            // settings
                            delay: 7000,
                            type: 'danger'
                        });
                    }
                }
            },
            error: function (jqXHR, textStatus, errorThrown) {
                Swal.close();
                alert('Error adding / update data');
            }
        });
    });
});