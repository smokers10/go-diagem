$(document).ready(function () {
    var croppie = null;
    var cropSampul = null;
    var el = document.getElementById('resizer');
    var formData = new FormData();
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
  
    $("#upload").on("click", function() {
        croppie.result({
            type: 'base64',
            size: 'original',
			format:'jpeg',
			size: { 
                width: 700, height: 700 
            }
        }).then(function(base64) {
            $("#cropModal").modal("hide");
            $("#logo_prev").attr("src", base64);
            $("#logo").val(base64);
        });
    });

    $(".rotate").on("click", function() {
        // console.log(sampul);
        if(cropSampul == null)
        {
            croppie.rotate(parseInt($(this).data('deg')));
        }else{
            cropSampul.rotate(parseInt($(this).data('deg')));
        }
    });

    $("#field-logo").on("change", function(event) {
        $("#cropModal").modal();
        croppie = new Croppie(el, {
                viewport: {
                    width: 400,
                    height: 400,
                    type: 'square'
                },
                original : {
                    width: 400,
                    height: 400,
                    type: 'square'
                },
                boundary: {
                    width: 450,
                    height: 450,
                },
                enableOrientation: true
            });
        $.getImage(event.target, croppie);
    });

    $("#field-sampul").on("change", function(event) {
        $("#sampulModal").modal();
        sampul = document.getElementById('cropSampul');
        cropSampul = new Croppie(sampul, {
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
        $.getImage(event.target, cropSampul);
    });

    $("#sampulSave").on("click", function() {
        cropSampul.result({
            type: 'base64',
            size: 'original',
			format:'jpeg',
			size: { 
                width: 1200, height: 400 
            }
        }).then(function(base64) {
            $("#sampulModal").modal("hide");
            $("#sampulPrev").attr("src", base64);
            $("#sampul").val(base64);
            cropSampul = null;
        });
    });
});

$("#form-toko_info").submit(function (e) {
    e.preventDefault();
    var formData = new FormData($('#form-toko_info')[0]);
    $.ajax({
        url: laroute.route('mitra.toko.update'),
        type: 'POST',
        data: formData,
        cache: false,
        contentType: false,
        processData: false,
        beforeSend: function(){
            Swal.fire({
                title: 'Tunggu Sebentar...',
                text: ' ',
                imageUrl: laroute.url('assets/img/loading.gif', ['']),
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
                    html: `Informasi Toko Berhasil Disimpan!`,
                    showCancelButton: false,
                    showConfirmButton: false,
                    allowOutsideClick: false,
                    timer: 3000
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
    });
});