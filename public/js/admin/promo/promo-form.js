// get current URL path
const currentPath = location.pathname
const CPSplit = currentPath.split("/")  
const promoID = CPSplit[4]

// get selected promo data
$.ajax({
    url: `/admin/promo/get/${promoID}`,
    success: function(response) {
        const data = response.data
        const {
            judul, deskripsi, image, seo_deskripsi, seo_keyword, 
            seo_tags, tgl_mulai, tgl_selesai, is_publish, is_featured
        } = data

        $("#field-judul").val(judul)
        $("#seo-keyword").val(seo_keyword)
        $("#seo-tags").val(seo_tags)
        $("#seo-deskripsi").val(seo_deskripsi)
        $("#field-tgl-mulai").val(tgl_mulai)
        $("#field-tgl-selesai").val(tgl_selesai)
        $("#field-deskripsi").summernote('pasteHTML', deskripsi)

        image ? $("#img_preview").attr("src", "/"+image.split("public/")[1]) : $("#img_preview").attr("src", "/img/poster.png")
        is_publish ? $("#is-publish-checked").prop("checked", true) : $("#is-publish-unchecked").prop("checked", true)
        is_featured ? $("#is-featured-checked").prop("checked", true) : $("#is-featured-unchecked").prop("checked", true)
    },
    error: function (jqXHR, textStatus, errorThrown) {
        Swal.close()
        alert('Error adding / update data')
    }
})

jQuery(document).ready(function () {
    var croppie = null
    var el = document.getElementById('resizer')
    moment.locale('id')

    $.getImage = function(input, croppie) {
        if (input.files && input.files[0]) {
            var reader = new FileReader()
            reader.onload = function(e) {
                croppie.bind({
                    url: e.target.result,
                })
            }
            reader.readAsDataURL(input.files[0])
        }
    }

    $("#file-upload").on("change", function(event) {
        $("#cropModal").modal()
        croppie = new Croppie(el, {
            viewport: {
                width: 700,
                height: 332,
                type: 'square'
            },
            original : {
                width: 700,
                height: 332,
                type: 'square'
            },
            boundary: {
                width: 730,
                height: 376,
            },
            enableOrientation: true
        })
        $.getImage(event.target, croppie)
    })

    $("#sampulSave").on("click", function() {
        croppie.result({
            type: 'base64',
            size: 'original',
			format:'jpeg',
			size: { width: 1376, height: 652 }
        }).then(function(base64result) {
            $("#cropModal").modal("hide")
            $("#img_preview").attr("src", base64result)
            $("#image").val(base64result)

            if($("#method").val() == "update") {
                let base64 = base64result.split(";base64,")[1]
                let filetype = base64result.split(";base64,")[0]
                let format = filetype.split("data:image/")[1]
                let data = JSON.stringify({
                    id: promoID,
                    sampul: {base64,format},
                })

                $.ajax({
                    url: "/admin/promo/update-cover",
                    type: "put",
                    data,
                    dataType: "json",
                    contentType: "application/json",
                    beforeSend: function(){
                        Swal.fire({
                            title: 'Tunggu Sebentar...',
                            text: 'Data Sedang Diproses!',
                            imageUrl:'/img/loading.gif',
                            showConfirmButton: false,
                            allowOutsideClick: false,
                        })
                    },
                    success: function (response) {
                        Swal.close()
                        $('.is-invalid').removeClass('is-invalid')
                        if (response.success) {
                            $('#modal_embed').modal('hide')
                            Swal.fire({
                                icon: 'success',
                                title: 'Sampul Berhasil Diupdate',
                                showConfirmButton: false,
                                timer: 1500
                              })
                        } else {
                            Swal.close()
                            for (control in response.errors) {
                                $('#field-' + control).addClass('is-invalid')
                                $('#error-' + control).html(response.errors[control])
                                $.notify({
                                    icon: 'fa fa-times',
                                    message: response.errors[control]
                                },{
                                    delay: 7000,
                                    type: 'danger'
                                })
                            }
                        }
                    },
                    error: function (jqXHR, textStatus, errorThrown) {
                        Swal.close()
                        alert('Error adding / update data')
                    }
                })
            }
        })
    })

    $(".rotate").on("click", function() {
        croppie.rotate(parseInt($(this).data('deg')))
    })

    $('#cropModal').on('hidden.bs.modal', function (e) {
        setTimeout(function() { croppie.destroy() }, 100)
    })

    $('#field-deskripsi').summernote({
        height: '250px'
    })

    $("#form-promo").on('submit', function (e) {
        e.preventDefault()
        var url, type;

        if($('#method').val() === 'update'){
            url = "/admin/promo/update"
            type = "PUT"
        }else{
            url = "/admin/promo/create"
            type = "POST"
        }

        let src = $("#img_preview").attr("src")
        const base64 = src.split(";base64,")[1]
        const filetype = src.split(";base64,")[0]
        const format = filetype.split("data:image/")[1]

        const data = JSON.stringify({
            id: promoID,
            judul: $("#field-judul").val(),
            deskripsi: $("#field-deskripsi").val(),
            seo_keyword: $("#seo-keyword").val(),
            seo_tags: $("#seo-tags").val(),
            seo_deskripsi: $("#seo-deskripsi").val(),
            is_publish: $("input[name=is-publish]:checked").val() == "1",
            is_featured: $("input[name=is-featured]:checked").val() == "1",
            tgl_mulai: $("#field-tgl-mulai").val(),
            tgl_selesai: $("#field-tgl-selesai").val(),
            sampul: {base64,format},
        })

        // submit form
        $.ajax({
            url,
            type,
            data,
            dataType: "json",
            contentType: "application/json",
            beforeSend: function(){
                Swal.fire({
                    title: 'Tunggu Sebentar...',
                    text: 'Data Sedang Diproses!',
                    imageUrl:'/img/loading.gif',
                    showConfirmButton: false,
                    allowOutsideClick: false,
                })
            },
            success: function (response) {
                Swal.close()
                $('.is-invalid').removeClass('is-invalid')
                if (response.success) {
                    $('#modal_embed').modal('hide')
                    Swal.fire({
                        title: `Berhasil!`,
                        showConfirmButton: false,
                        icon: 'success',
                        html: `
                            ${ type == "PUT" ? "Promo Baru Berhasil Diupdate!" : "Promo Baru Berhasil Disimpan!"}
                            <br><br>
                            <a href="/admin/promo" class="btn btn-keluar btn-alt-danger"><i class="si si-close mr-1"></i>Keluar</a>
                            ${ type == "POST" ? `<a href="/admin/promo/tambah" class="btn btn-tambah_baru btn-alt-primary"><i class="si si-plus mr-1"></i>Tambah Promo Lain</a>` : ""}`,
                        showCancelButton: false,
                        showConfirmButton: false,
                    })
                } else {
                    Swal.close()
                    for (control in response.errors) {
                        $('#field-' + control).addClass('is-invalid')
                        $('#error-' + control).html(response.errors[control])
                        $.notify({
                            icon: 'fa fa-times',
                            message: response.errors[control]
                        },{
                            delay: 7000,
                            type: 'danger'
                        })
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