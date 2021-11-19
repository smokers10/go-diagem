var isHasVarian = false
var varianArr = []
const hargaProdukFormElem = $("#harga-produk-form")
const variasiToggleBtn = $("#btn-toggle-variasi")

const fotoInput = $("#file-upload-input")
const fotoEditorPreview = $("#imagePreview")
const imagePreviewDefaultSrc = "/img/placeholder/product.png"
const selectedFotoRow = $("#selected-foto-row")
var photoEditor

// once page is ready
$(document).ready(function(){
    // kategori select 2
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

    $("#field-kategori").select2({
        placeholder: 'Pilih Kategori',
        theme : 'bootstrap4',
    })

    // tambah spesifikasi 
    $("#btn-add-spek").click(function(){
        // append form
        var spek = $('#field-spek')
        var spek_count = parseInt($('#field-spek_count').val());
        spek.find('tbody').append(`
        <tr>
            <td>
                <input type="text" class="form-control" id="nama${spek_count}" placeholder="Nama" autocomplete="off">
            </td>
            <td>
                <input type="text" class="form-control" id="value${spek_count}" placeholder="Nilai Spesifikasi" autocomplete="off">
            </td>
            <td width="5%">
                <button type="button" class="btn btn-sm btn-alt-danger btn-delete-spek"><i class="si si-trash"></i></button>
            </td>
        </tr>`)
        spek_count += 1
        $('#field-spek_count').val(spek_count)
        console.log(spek_count)
    })

    // hapus spesifikasi
    $(document).on('click', '#field-spek .btn-delete-spek', function () {
        $(this).closest('tr').remove();
    })

    // varian
    hargaProdukFormElem.html(_nonVariasi())

    variasiToggleBtn.click(function(){
        isHasVarian ? isHasVarian = false : isHasVarian = true
        console.log(isHasVarian)
        
        if (isHasVarian) {
            $(this).removeClass("btn-outline-primary")
            $(this).addClass("btn-outline-danger")
            $(this).html(`Nonaktifkan Variasi Produk`)
            hargaProdukFormElem.html(_variasi())
            if (varianArr.length) {
                let table = $("#tbl_variasi")
                varianArr.forEach(varian => {
                    table.find("tbody").append(`
                        <tr>
                            <td>${varian.variant}</td>
                            <td>${varian.harga}</td>
                            <td>${varian.stok}</td>
                            <td>${varian.sku}</td>
                            <td>
                                <button type="button" class="btn btn-danger">
                                    <i class="si si-trash mr-1"></i>
                                </button>
                            </td>
                        </tr>
                    `)
                })
            }
        }else {
            hargaProdukFormElem.html(_nonVariasi())
            $(this).removeClass("btn-outline-danger")
            $(this).addClass("btn-outline-primary")
            $(this).html(`<i class="si si-plus mr-1"></i> Aktifkan Variasi Produk`)
        }
    })

    $(document).on('click', '#btn-add-variasi', function(){
        let namaVarian = $("#nama-varian")
        let kodeVarian = $("#kode-varian")
        let stokVarian = $("#stok-varian")
        let hargaVarian = $("#harga-varian")
        let table = $("#tbl_variasi")

        var varian = {
            variant : namaVarian.val(),
            sku : kodeVarian.val(),
            stok : parseInt(stokVarian.val()),
            harga : parseInt(hargaVarian.val()),
        }
        
        varianArr.push(varian)

        namaVarian.val("")
        kodeVarian.val("")
        stokVarian.val("")
        hargaVarian.val("")

        table.find("tbody").append(`
            <tr>
                <td>${varian.variant}</td>
                <td>${varian.harga}</td>
                <td>${varian.stok}</td>
                <td>${varian.sku}</td>
                <td>
                    <button type="button" class="btn btn-danger">
                        <i class="si si-trash mr-1"></i>
                    </button>
                </td>
            </tr>
        `)
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

    // upload foto produk
    // buka modal foto
    $("#foto-modal-btn").click(function() {
        fotoEditorPreview.croppie('destroy')
        fotoEditorPreview.attr("src", imagePreviewDefaultSrc)
    })

    // preview foto produk
    fotoInput.change(function(){readURL(this)})

    // submit
    $("#submit-tbh-barang").click(function(e){
        e.preventDefault()
        // mendapatkan data spesifikasi 
        const spek_count = parseInt($('#field-spek_count').val())
        var spesifikasi = []
        for (let i = 0; i < spek_count; i++) {
            var nama = $(`#nama${i}`).val()
            var value = $(`#value${i}`).val()
            var spek = {nama, value}
            if (!nama == "" && !value == "") {
                spesifikasi.push(spek)
            }
        }

        // mendapatkan foto
        var selectedFotoSrc = []
        const selectedFotoItem = $(".selected-foto-item")
        for (let i = 0; i < selectedFotoItem.length; i++) {
            let src = selectedFotoItem[i].currentSrc
            const base64 = src.split(";base64,")[1]
            const type = src.split(";base64,")[0]
            const format = type.split("data:image/")[1]
            selectedFotoSrc.push({ base64, format })
        }

        const data = {
            nama : $("#field-nama").val(),
            deskripsi: $("#field-deskripsi").val(),
            spesifikasi: JSON.stringify(spesifikasi),
            produk_foto: selectedFotoSrc,
            kategori_id: parseInt($("#field-kategori").val()),
            berat: parseInt($("#field-berat").val()),
            satuan_berat: $("#field-berat-satuan").val(),
            lebar: parseFloat($("#field-lebar").val()),
            panjang: parseFloat($("#field-panjang").val()),
            tinggi: parseFloat($("#field-tinggi").val()),
            kode: $("#field-kode").val(),
            harga: parseInt($("#field-harga").val()),
            variasi: varianArr,
            stok: parseInt($("#field-stok").val()),
            is_has_variant: isHasVarian,
        }

        $.ajax({
            url: "/admin/produk/tambah",
            type:"post",
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
                    html: `Produk Baru Berhasil Disimpan!
                        <br><br>
                        <a href="/admin/produk" class="btn btn-keluar btn-alt-danger"><i class="si si-close mr-1"></i>Keluar</a> 
                        <a href="/admin/produk/tambah" class="btn btn-tambah_baru btn-alt-primary"><i class="si si-plus mr-1"></i>Tambah Produk Lain</a>`,
                    showCancelButton: false,
                    showConfirmButton: false,
                    allowOutsideClick: false,
                })
            },
            error: function (jqXHR, textStatus, errorThrown) {
                Swal.close()
                alert('Error adding / update data')
            }
        })
    })
})

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
        selectedFotoRow.append(`
            <div class="col-md-4 mb-2">
                <center>
                    <img src="${res}" class="selected-foto-item mb-3" width="80%" height="80%"/>
                    <button class="btn btn-danger">Hapus</button>
                </center>
            </div>
        `)
    })
    $("#foto-produk-modal").modal('hide')
}

function _nonVariasi() {
    var element = `
        <div class="form-group row non-variasi">
            <label class="col-lg-3 col-form-label" for="field-harga">Harga Produk</label>
            <div class="col-lg-9">
                <div class="input-group">
                    <div class="input-group-prepend">
                        <span class="input-group-text">
                            Rp
                        </span>
                    </div>
                    <input type="number" class="form-control" id="field-harga" name="harga" placeholder="Masukan Harga Produk" autocomplete="off">
                </div>
                <div class="text-danger font-size-sm" id="error-harga"></div>
            </div>
        </div>

        <div class="form-group row non-variasi">
            <label class="col-lg-3 col-form-label" for="field-stok">Stok Produk</label>
            <div class="col-lg-9">
                <input type="number" class="form-control" id="field-stok" name="stok" placeholder="Masukan Stok Produk">
                <div class="invalid-feedback font-size-sm" id="error-stok">Invalid feedback</div>
            </div>
        </div>
    `

    return element
}

function _variasi() {
    return `
        <div class="form-group row mb-0 variasi">
            <div class="col-lg-12">
                <div class="form-group row">
                    <label class="col-lg-2 col-form-label" for="field-nama">Nama</label>
                    <div class="col-lg-10">
                        <input type="text" class="form-control variasi-nama" id="nama-varian" name="var_nama" placeholder="Masukkan Nama Variasi, contoh: Warna, Jenis Batu, dll.">
                        <div class="text-danger font-size-sm" id="error-var1_nama"></div>
                    </div>
                </div>
                
                <div class="form-group row">
                    <label class="col-lg-2 col-form-label" for="field-kode">Kode</label>
                    <div class="col-lg-10">
                        <input type="text" class="form-control variasi-kode" id="kode-varian" name="var_kode" placeholder="Masukan Kode/SKU Variasi Ini">
                        <div class="text-danger font-size-sm" id="error-kode"></div>
                    </div>
                </div>

                <div class="form-group row">
                    <label class="col-lg-2 col-form-label" for="field-stok">Stok Variasi</label>
                    <div class="col-lg-10">
                        <input type="number" class="form-control variasi-stok" id="stok-varian" name="var_stok" placeholder="Masukan Stok Variasi Ini">
                        <div class="text-danger font-size-sm" id="error-stok"></div>
                    </div>
                </div>

                <div class="form-group row">
                    <label class="col-lg-2 col-form-label" for="field-harga">Harga</label>
                    <div class="col-lg-10">
                        <input type="number" class="form-control variasi-harga" id="harga-varian" name="var_harga" placeholder="Masukkan Nama Harga Untuk Variasi Ini">
                        <div class="text-danger font-size-sm" id="error-var_harga"></div>
                    </div>
                </div>

                <div class="form-group row">
                    <label class="col-lg-2 col-form-label"></label>
                    <div class="col-lg-10">
                        <button type="button" class="btn btn-outline-primary btn-block" id="btn-add-variasi">
                            <i class="si si-plus mr-1"></i> Tambahkan Variasi
                        </button>
                    </div>
                </div>
            </div>
        </div> 

        <div class="form-group row variasi">
            <div class="col-lg-12">
                <table class="table table-bordered table-vcenter text-center" id="tbl_variasi">
                    <thead>
                        <tr>
                            <th class="tb-variasi-1_nama" width="18%">Nama</th>
                            <th width="30%">Harga</th>
                            <th width="14%">Stok</th>
                            <th width="20%">Kode Variasi</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        
                    </tbody>
                </table>
            </div>
        </div>
    `
}