// get current URL path
const currentPath = location.pathname
const CPSplit = currentPath.split("/")  
const sliderID = CPSplit[4]

// get selected slider data
$.ajax({
    url: `/admin/slider/get/${sliderID}`,
    success: function(response) {
        const data = response.data
        const { name, description, type, is_publish, image, url } = data
        $("#field-name").val(name)
        $("#field-type").val(type)
        $("#field-link").val(url)
        $("#field-deskripsi").summernote("pasteHTML", description)
        image ? $("#thumbnail").attr("src", "/"+image.split("public/")[1]) : $("#thumbnail").attr("src","/img/slider.png")
        is_publish ? $("#status_publikasi").prop("checked", true) : $("#status_draft").prop("checked", true)
    },
    error: function (jqXHR, textStatus, errorThrown) {
        Swal.close()
        alert('Error adding / update data')
    }
})

jQuery(document).ready(function () {
    // upload thumbnails
    const thumbnails = $('#thumbnail-upload')
    thumbnails.change(function () {
        var file = thumbnails.prop('files')[0]
        var reader = new FileReader()
        reader.onloadend = function() {
            $("#thumbnail").attr('src', reader.result)
            let src = reader.result
            let base64 = src.split(";base64,")[1]
            let filetype = src.split(";base64,")[0]
            let format = filetype.split("data:image/")[1]
            const data = JSON.stringify({
                id: sliderID,
                sampul: {base64,format},
            })
            $.ajax({
                id: sliderID,
                url:"/admin/slider/update-cover",
                type:"PUT",
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
                    Swal.fire({
                        title: `Berhasil!`,
                        text: 'Sampul Slider Diperbaharui',
                        timer: 3000,
                        showConfirmButton: false,
                        icon: 'success',
                        showCancelButton: false,
                        showConfirmButton: false,
                    })
                },
                error: function (jqXHR, textStatus, errorThrown) {
                    Swal.close()
                    alert('Error adding / update data')
                }
            })
        }
        reader.readAsDataURL(file)
    })

    // inisiasi summernote
    $('#field-deskripsi').summernote({
        height: '250px',
        toolbar: [
            ['undo', ['undo',]],
            ['redo', ['redo',]],
            ['style', ['bold', 'italic', 'underline','clear']],
            ['font', ['strikethrough',]],
            ['fontsize', ['fontsize']],
            ['color', ['color']],
            ['para', ['ul', 'ol', 'paragraph']],
        ]
    })

    // submit
    $("#form-slide").on('submit', function (e) {
        e.preventDefault()
        var url, type;
        if($('#method').val() === 'update'){
            url = "/admin/slider/update"
            type = "PUT"
        }else{
            url = "/admin/slider/create"
            type = "POST"
        }

        let src = $("#thumbnail").attr("src")
        const base64 = src.split(";base64,")[1]
        const filetype = src.split(";base64,")[0]
        const format = filetype.split("data:image/")[1]

        const data = JSON.stringify({
            id: sliderID,
            name: $("#field-name").val(),
            description: $("#field-deskripsi").val(),
            type:parseInt($("#field-type").val()),
            is_publish:$("input[name=status]:checked").val() == "1",
            url:$("#field-link").val(),
            sampul: {base64,format},
        })

        // submit form
        $.ajax({
            id: sliderID,
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
                if (response.success) {
                    $('#modal_embed').modal('hide')
                    Swal.fire({
                        title: `Berhasil!`,
                        showConfirmButton: false,
                        icon: 'success',
                        html: `
                            ${ type == "PUT" ? "Slider Baru Berhasil Diupdate!" : "Slider Baru Berhasil Disimpan!"}
                            <br><br>
                            <a href="/admin/slider" class="btn btn-keluar btn-alt-danger"><i class="si si-close mr-1"></i>Keluar</a>
                            ${ type == "POST" ? `<a href="/admin/slider/tambah" class="btn btn-tambah_baru btn-alt-primary"><i class="si si-plus mr-1"></i>Tambah Slider Lain</a>` : ""}`,
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