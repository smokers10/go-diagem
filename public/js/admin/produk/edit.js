const fotoInput = $("#file-upload-input")
const fotoEditorPreview = $("#imagePreview")
const imagePreviewDefaultSrc = "/img/placeholder/product.png"
const selectedFotoRow = $("#selected-foto-row")
var photoEditor

// get current URL path
const currentPath = location.pathname
const CPSplit = currentPath.split("/")  
const produkID = CPSplit[4]

// once page is ready
$(document).ready(function(){
    // get data kategori
    $.ajax({
        url:"/admin/kategori/get",
        method: "get",
        success: function(res){
            const {data} = res
            data.forEach(kategori => {
                $("#field-kategori").append(`<option value="${kategori.id}">${kategori.nama}</option>`)
            })
        }
    })

    // get data
    $.ajax({
        url: `/admin/produk/get/${produkID}`,
        method: "get",
        success: function (res) {
            const data = res.data
            const { produk_foto, variasi, spesifikasi, nama, stok, deskripsi,
                 kode, harga, is_has_variant, kategori, berat, satuan_berat,
                 lebar, panjang, tinggi, discount } = data
            const { id } = kategori 

            // has variant or not
            if (is_has_variant) {
                $("#harga").hide()
                $("#variasi").show()
            } else {
                $("#harga").show()
                $("#variasi").hide()
            }

            // set to input html
            $("#field-nama").val(nama) 
            $("#field-deskripsi").html(deskripsi)
            $("#field-kategori").val(id)
            $("#field-berat").val(berat)
            $("#field-berat-satuan").val(satuan_berat)
            $("#field-kategori").val(id)
            $("#field-lebar").val(lebar)
            $("#field-panjang").val(panjang)
            $("#field-tinggi").val(tinggi)
            $("#field-harga").val(harga)
            $("#field-stok").val(stok)
            $("#field-kode").val(kode)
            $("#field-discount").val(discount)

            // set varian
            if (variasi) {
                variasi.forEach(v => {
                    _renderVarian(v)
                })
            }

            // set spesifikasi
            _renderSpek(spesifikasi)

            // set foto produk
            _renderFoto(produk_foto)

            // set deskripsi
            $("#field-deskripsi").summernote('pasteHTML', deskripsi)

            
        }
    })

    // hapus spesifikasi
    $(document).on('click', '#field-spek .btn-delete-spek', function () {
        $(this).closest('tr').remove();
    })

    // Deskripsi Summernote init
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

    // tambah spesifikasi 
    $("#btn-add-spek").click(function(){
        var spekEl = $('#field-spek')
        var spekCountEl = $('#field-spek-count') 
        var count = parseInt(spekCountEl.val())

        // append form
        spekEl.find('tbody').append(_createSpekEl(count))
        count += 1

        spekCountEl.val(count)
    })

    // upload foto produk
    // buka modal foto
    $("#foto-modal-btn").click(function() {
        fotoEditorPreview.croppie('destroy')
        fotoEditorPreview.attr("src", imagePreviewDefaultSrc)
    })

    // preview foto produk
    fotoInput.change(function(){readURL(this)})

    // add varian
    $("#btn-add-variasi").click(function(){
        let newVarian = {
            variant: $("#nama-varian").val(),
            sku: $("#kode-varian").val(),
            harga: parseFloat($("#harga-varian").val()),
            stok: parseInt($("#stok-varian").val()),
            produk_id: produkID,
        }

        $.ajax({
            url: "/admin/varian",
            method: "post",
            data: JSON.stringify(newVarian),
            dataType: "json",
            contentType: "application/json",
            beforeSend: function(){
                Swal.fire({
                    title: 'Tunggu Sebentar...',
                    text: '',
                    imageUrl: '/img/loading.gif',
                    showConfirmButton: false,
                    allowOutsideClick: false,
                })
            },
            success: function(){
                $("#nama-varian").val("")
                $("#kode-varian").val("")
                $("#harga-varian").val("")
                $("#stok-varian").val("")
                Swal.fire({
                    icon: 'success',
                    title: 'Varian produk berhasil disimpan',
                    showConfirmButton: false,
                    showCancelButton: false,
                    showConfirmButton: false,
                    timer: 1500
                })

                _renderVarian(newVarian)
            },
            error: function (jqXHR, textStatus, errorThrown) {
                Swal.close()
                alert('Error adding / update data')
            }
        })
    })

    // update info produk
    $("#submit-info-produk").click(function() {
        // mendapatkan data spesifikasi 
        const spek_count = parseInt($('#field-spek-count').val())
        var spesifikasi = []
        for (let i = 0; i < spek_count; i++) {
            var nama = $(`#nama${i+1}`).val()
            var value = $(`#value${i+1}`).val()
            var spek = {nama, value}
            if (!nama == "" && !value == "") {
                spesifikasi.push(spek)
            }
        }
        
        const data = {
            nama : $("#field-nama").val(),
            deskripsi: $("#field-deskripsi").val(),
            spesifikasi: JSON.stringify(spesifikasi),
            kategori_id: $("#field-kategori").val(),
            berat: parseInt($("#field-berat").val()),
            satuan_berat: $("#field-berat-satuan").val(),
            lebar: parseFloat($("#field-lebar").val()),
            panjang: parseFloat($("#field-panjang").val()),
            tinggi: parseFloat($("#field-tinggi").val()),
            kode: $("#field-kode").val(),
            harga: parseInt($("#field-harga").val()),
            stok: parseInt($("#field-stok").val()),
            discount: parseInt($("#field-discount").val()),
            id: produkID
        }

        $.ajax({
            url: "/admin/produk/update",
            type:"PUT",
            data: JSON.stringify(data),
            dataType: "json",
            contentType: "application/json",
            beforeSend: function(){
                Swal.fire({
                    title: 'Tunggu Sebentar...',
                    text: '',
                    imageUrl: '/img/loading.gif',
                    showConfirmButton: false,
                    allowOutsideClick: false,
                })
            },
            success: function (response) {
                $('.is-invalid').removeClass('is-invalid')
                Swal.fire({
                    title: `Berhasil!`,
                    showConfirmButton: false,
                    icon: 'success',
                    html: `Produk Baru Berhasil Diupdate!
                        <br>
                        <br>
                        <a href="/admin/produk" class="btn btn-keluar btn-alt-danger"><i class="si si-close mr-1"></i>Keluar</a>`,
                    showCancelButton: false,
                    showConfirmButton: false,
                    allowOutsideClick: false,
                    timer: 3000
                })
            },
            error: function (jqXHR, textStatus, errorThrown) {
                Swal.close()
                alert('Error adding / update data')
            }
        })
    })
})

// foto functionallity
function readURL(input) {
    if (input.files && input.files[0]) {
        var reader = new FileReader()

        reader.onload = function (e) {
            fotoEditorPreview.attr('src', e.target.result)
            photoEditor = fotoEditorPreview.croppie({
                enableExif: true,
                viewport: {
                    width: 350,
                    height: 350,
                    type: 'square'
                },
                original : {
                    width: 350,
                    height: 350,
                    type: 'square'
                },
                boundary: {
                    width: 360,
                    height: 360
                },
                enableOrientation: true
            })
        }

        reader.readAsDataURL(input.files[0])
    }
}

function _uploadSelectedFoto() {
    photoEditor.croppie('result', 'base64').then(function (res){
        let base64 = res.split(";base64,")[1]
        let type = res.split(";base64,")[0]
        let format = type.split("data:image/")[1]
        let data = {
            base64,
            format,
            produk_id: produkID
        }

        $.ajax({
            url: "/admin/produk/foto/upload",
            method: "post",
            data: JSON.stringify(data),
            dataType: "json",
            contentType: "application/json",
            beforeSend: function(){
                Swal.fire({
                    title: 'Tunggu Sebentar...',
                    text: '',
                    imageUrl: '/img/loading.gif',
                    showConfirmButton: false,
                    allowOutsideClick: false,
                })
            },
            success: function(res){
                let {data} = res
                selectedFotoRow.append(_createFotoElement(data))
                Swal.fire({
                    icon: 'success',
                    title: 'Foto Produk Baru Berhasil Disimpan',
                    showConfirmButton: false,
                    showCancelButton: false,
                    showConfirmButton: false,
                    timer: 3000
                })
                $("#foto-produk-modal").modal('hide')
            },
            error: function (jqXHR, textStatus, errorThrown) {
                Swal.close()
                alert('Error adding / update data')
            }
        })
    })
}

function _renderFoto(data){
    data.forEach(foto => {
        selectedFotoRow.append(_createFotoElement(foto))
    })
}

function _createFotoElement(data){
    return `
        <div class="col-md-4 mb-2">
            <center>
                <img src="${data.path}" class="selected-foto-item mb-3" width="80%" height="80%"/>
                <button class="btn btn-danger" data-id="${data.id}" data-src="${data.path}" onclick="_deleteFoto(this)">Hapus</button>
            </center>
        </div>
    `
}

function _deleteFoto(data) {
    var id = $(data).attr("data-id")
    var path = $(data).attr("data-src")
    Swal.fire({
        title: "Anda Yakin?",
        text: "Foto Yang Dihapus Tidak Akan Bisa Dikembalikan",
        icon: "warning",
        showCancelButton: true,
        confirmButtonText: 'Ya, Hapus!',
        cancelButtonText: 'Tidak, Batalkan!',
        reverseButtons: true,
        allowOutsideClick: false,
        confirmButtonColor: '#af1310',
        cancelButtonColor: '#fffff',
    }).then((result) => {
        if (result.value) {
            $.ajax({
                url: "/admin/produk/foto/delete",
                method: "delete",
                data: {id, path},
                beforeSend: function(){
                    Swal.fire({
                        title: 'Tunggu Sebentar...',
                        text: '',
                        imageUrl: '/img/loading.gif',
                        showConfirmButton: false,
                        allowOutsideClick: false,
                    })
                },
                success: function(){
                    Swal.fire({
                        icon: 'success',
                        title: 'Foto produk berhasil dihapus',
                        showConfirmButton: false,
                        showCancelButton: false,
                        showConfirmButton: false,
                        timer: 1500
                    })
                },
                error: function (jqXHR, textStatus, errorThrown) {
                    Swal.close()
                    alert('Error adding / update data')
                }
            })
        }else{
            Swal.close()
        }
    })
}

// varian functionallity
// render
function _renderVarian(data){
    $("#tbl_variasi").find("tbody").append(`
        <tr id="varian-${data.id}">
            <td>
                <input type="text" name="nama-varian" class="form-control" value="${data.variant}"></input> 
            </td>

            <td>
                <input type="number" name="harga-varian" class="form-control" min="0" value="${data.harga}"></input> 
            </td>

            <td>
                <input type="number" name="stok-varian" class="form-control" min="0" value="${data.stok}"></input> 
            </td>

            <td>
                <input type="text" name="sku-varian" class="form-control" value="${data.sku}"></input> 
            </td>

            <td>
                <button type="button" class="btn btn-success" data-id="${data.id}" onclick="_updateVarian(this)">
                    <i class="fas fa-save"></i>
                </button>
                <button type="button" class="btn btn-danger" data-id="${data.id}" onclick="_deleteVarian(this)">
                    <i class="si si-trash mr-1"></i>
                </button>
            </td>
        </tr>
    `)
}

// hapus
function _deleteVarian(data){
    var id = $(data).attr("data-id")
    Swal.fire({
        title: "Anda Yakin?",
        text: "Data Yang Dihapus Tidak Akan Bisa Dikembalikan",
        icon: "warning",
        showCancelButton: true,
        confirmButtonText: 'Ya, Hapus!',
        cancelButtonText: 'Tidak, Batalkan!',
        reverseButtons: true,
        allowOutsideClick: false,
        confirmButtonColor: '#af1310',
        cancelButtonColor: '#fffff',
    })
    .then((result) => {
        if (result.value) {
            $.ajax({
                url: "/admin/varian",
                method: "delete",
                data: {
                    produk_id: produkID,
                    id: id,
                },
                beforeSend: function(){
                    Swal.fire({
                        title: 'Tunggu Sebentar...',
                        text: '',
                        imageUrl: '/img/loading.gif',
                        showConfirmButton: false,
                        allowOutsideClick: false,
                    })
                },
                success: function(){
                    $("#nama-varian").val("")
                    $("#kode-varian").val("")
                    $("#harga-varian").val("")
                    $("#stok-varian").val("")
                    Swal.fire({
                        icon: 'success',
                        title: 'Varian produk berhasil dihapus',
                        showConfirmButton: false,
                        showCancelButton: false,
                        showConfirmButton: false,
                        timer: 1500
                    })

                    $(`#varian-${id}`).remove()
                },
                error: function (jqXHR, textStatus, errorThrown) {
                    Swal.close()
                    alert('Error adding / update data')
                }
            })
        }else{
            Swal.close()
        }
    })
}

// update
function _updateVarian(data){
    var id = $(data).attr("data-id")
    var trEl = $(`#varian-${id}`)
    const updatedVarian =  {
        variant: trEl.find('input[name="nama-varian"]').val(),
        sku: trEl.find('input[name="sku-varian"]').val(),
        harga: parseFloat(trEl.find('input[name="harga-varian"]').val()),
        stok: parseInt(trEl.find('input[name="stok-varian"]').val()),
        produk_id: produkID,
        id: id,
    }

    $.ajax({
        url: "/admin/varian",
        method: "put",
        data: JSON.stringify(updatedVarian),
        dataType: "json",
        contentType: "application/json",
        beforeSend: function(){
            Swal.fire({
                title: 'Tunggu Sebentar...',
                text: '',
                imageUrl: '/img/loading.gif',
                showConfirmButton: false,
                allowOutsideClick: false,
            })
        },
        success: function(){
            Swal.fire({
                icon: 'success',
                title: 'Varian produk berhasil disimpan',
                showConfirmButton: false,
                showCancelButton: false,
                showConfirmButton: false,
                timer: 1500
            })
        },
        error: function (jqXHR, textStatus, errorThrown) {
            Swal.close()
            alert('Error adding / update data')
        }
    })
}

// fungsionalitas informasi produk
function _renderSpek(spek_array) {
    var spekEl = $('#field-spek')
    var spekCountEl = $('#field-spek-count') 

    // append form
    for (let i = 0; i < spek_array.length; i++) {
        spekEl.find('tbody').append(_createSpekElWithData(i+1, spek_array[i].nama, spek_array[i].value))
    }
    
    spekCountEl.val(spek_array.length)
}

function _createSpekEl(count) {
    return `
        <tr>
            <td>
                <input type="text" class="form-control" id="nama${count}" placeholder="Nama" autocomplete="off">
            </td>

            <td>
                <input type="text" class="form-control" id="value${count}" placeholder="Nilai Spesifikasi" autocomplete="off">
            </td>

            <td width="5%">
                <button type="button" class="btn btn-sm btn-alt-danger btn-delete-spek"><i class="si si-trash"></i></button>
            </td>
        </tr>
    `
}

function _createSpekElWithData(count, name, value) {
    return `
        <tr>
            <td>
                <input type="text" class="form-control" id="nama${count}" value="${name}" placeholder="Nama" autocomplete="off">
            </td>

            <td>
                <input type="text" class="form-control" id="value${count}" value="${value}" placeholder="Nilai Spesifikasi" autocomplete="off">
            </td>

            <td width="5%">
                <button type="button" class="btn btn-sm btn-alt-danger btn-delete-spek"><i class="si si-trash"></i></button>
            </td>
        </tr>
    `
}