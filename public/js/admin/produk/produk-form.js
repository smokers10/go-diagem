jQuery(document).ready(function () {
    var inc_variasi = 0;
    var var2_status = $('#var2_status');
    var category_name = "";
    var subcategory_name = "";
    var subsubcategory_name = "";
    var category_id = null;
    var subcategory_id = null;
    var subsubcategory_id = null;
    var kategori_id = null;
    var el = document.getElementById('resizer');
    var croppie = null;

    $('#subcategory_list').hide();
    $('#subsubcategory_list').hide();

    var kategoriSelect = $("#field-kategori").select2({
        placeholder: 'Pilih Kategori',
        theme : 'bootstrap4',
        ajax: {
            url: laroute.route('admin.kategori.json'),
            type: 'POST',
            dataType: 'JSON',
            delay: 250,
            data: function (params) {
                return {
                    searchTerm: params.term
                };
            },
            processResults: function (response) {
                return {
                    results: response
                };
            },
            cache: true
        }
    });


    $(document).on('click', '#btn-add-variasi', function () {
        $(this).addClass('d-none');
        $('#btn-hapus-variasi').removeClass('d-none');
        $('#is_variasi').val(1);
        $('.non-variasi').last().after(
            `<div class="form-group row mb-0 variasi">
                <label class="col-lg-3 col-form-label" for="field-harga">Variasi 1</label>
                <div class="col-lg-7">
                    <div class="form-group row">
                        <label class="col-lg-2 col-form-label" for="field-var1_nama">Nama</label>
                        <div class="col-lg-10">
                            <input type="text" class="form-control variasi-1_nama" id="field-var1_nama" name="var1_nama" placeholder="Masukkan Nama Variasi, contoh: Warna, dll.">
                            <div class="text-danger font-size-sm" id="error-var1_nama"></div>
                        </div>
                    </div>
                    <div class="form-group row">
                        <label class="col-lg-2 col-form-label" for="field-var1_pilihan">Pilihan</label>
                        <div class="col-lg-10">
                            <input type="text" class="form-control" id="field-var1_pilihan" name="var1_pilihan" placeholder="Masukkan Pilihan Variasi, contoh: Hijau, dll.">
                            <div class="text-danger font-size-sm" id="error-var1_pilihan"></div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="form-group row variasi">
                <label class="col-lg-3 col-form-label" for="field-harga">Variasi 2</label>
                <div class="col-lg-7">
                    <button type="button" class="btn btn-outline-primary btn-block" id="btn-add-var2">
                        <i class="si si-plus mr-3"></i> Tambah Variasi 2
                    </button>
                    <button type="button" class="btn btn-outline-danger btn-block mb-3 mt-0 d-none var2-form" id="btn-hapus-var2">
                        <i class="si si-trash mr-3"></i> Hapus Variasi 2
                    </button>
                    <div class="form-group row var2-form d-none">
                        <label class="col-lg-2 col-form-label" for="field-var2_nama">Nama</label>
                        <div class="col-lg-10">
                            <input type="text" class="form-control var2-nama" id="field-var2_nama" name="var2_nama" placeholder="Masukkan Nama Variasi, contoh: Ukuran, dll.">
                            <div class="text-danger font-size-sm" id="error-var2_nama"></div>
                        </div>
                    </div>
                    <div class="form-group row var2-form d-none">
                        <label class="col-lg-2 col-form-label" for="field-var2_pilihan">Pilihan</label>
                        <div class="col-lg-10">
                            <input type="text" class="" id="field-var2_pilihan" name="var2_pilihan" placeholder="Masukkan Pilihan Variasi, contoh: Hijau, dll.">
                            <div class="text-danger font-size-sm" id="error-var2_pilihan"></div>
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
                                <th class="var2-nama d-none" width="18%">Nama</th>
                                <th width="30%">Harga</th>
                                <th width="14%">Stok</th>
                                <th width="20%">Kode Variasi</th>
                            </tr>
                        </thead>
                        <tbody>
                            
                        </tbody>
                    </table>
                </div>
            </div>`
        );

        $('.non-variasi').not(":last").addClass('d-none');

        $('#field-var1_pilihan').tagsinput({
            tagClass: 'badge badge-primary',
            maxTags: 20
        });

        $('#field-var1_pilihan').tagsinput({
            tagClass: 'badge badge-primary',
            maxTags: 20
        });
    
        $('#field-var1_pilihan').on('itemAdded', function(event) {
            update_variasi();
        });
    
        $('#field-var1_pilihan').on('beforeItemRemove', function(event) {
            $('#tbl_variasi tbody').find('tr').each(function(){
                $(this).find('td input[value="'+ event.item + '"]').eq(0).closest('tr').remove();
            });
        });

        $(document).on('click', '#btn-add-var2', function () {
            var2_status.val(1);
            $(this).addClass('d-none');
            $('.var2-form').removeClass('d-none');
            $('.var2-nama').removeClass('d-none');

            $('#tbl_variasi tbody').find('tr').each(function(){
                $(this).find('td').eq(0).after('<td>Pilihan</td>');
            });

            $(document).on('keyup', '.var2-nama', function () {
                $('th.var2-nama').html($(this).val());
            });
        
            $('#field-var2_pilihan').tagsinput({
                tagClass: 'badge badge-primary',
                maxTags: 20
            });
        
            $('#field-var2_pilihan').on('itemAdded', function(event) {
                update_variasi();
            });
        
            $('#field-var2_pilihan').on('beforeItemRemove', function(event) {
                $('#tbl_variasi tbody').find('tr').each(function(){
                    $(this).find('td input[value="'+ event.item + '"]').eq(0).closest('tr').remove();
                });
            });

        });

        $(document).on('click', '#btn-hapus-var2', function () {
            $(this).addClass('d-none');
            $('#btn-add-var2').removeClass('d-none');
            $('.var2-form').addClass('d-none');
            $('.var2-nama').addClass('d-none');
            $('#tbl_variasi tbody').find('tr').each(function(){
                $(this).find('td').eq(1).remove();
            });
            var2_status.val(0);
        });

        $(document).on('click', '#btn-hapus-variasi', function () {
            $(this).addClass('d-none');
            $('#btn-add-variasi').removeClass('d-none');
            $('.non-variasi').removeClass('d-none');
            $('.variasi').remove();
            $('#is_variasi').val(0);
        });
    });

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
    });

    $("input.file-upload").on("change", function(event) {
        $("#cropModal").modal();
        $('#image_id').val($(this).data('id'));

        croppie = new Croppie(el, {
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
            });
        $.getImage(event.target, croppie);
    });

    $("#upload").on("click", function() {
        id = $('#image_id').val();
        croppie.result({
            type: 'base64',
            size: 'original',
			format:'jpeg',
			size: { 
                width: 700, height: 700 
            }
        }).then(function(base64) {
            $("#cropModal").modal("hide");
            $("#previewImg-"+ id ).attr("src", base64);
            $("#foto-"+ id ).val(base64);
        });
    });

    $(".rotate").on("click", function() {
        croppie.rotate(parseInt($(this).data('deg')));
    });

    $('#cropModal').on('hidden.bs.modal', function (e) {
        setTimeout(function() { croppie.destroy(); }, 100);
    });

    $("#form-produk").on("submit", function (e) {
        e.preventDefault();
        var formData = new FormData($('#form-produk')[0]);
        $.ajax({
            url: laroute.route('admin.produk.simpan'),
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
                $('.is-invalid').removeClass('is-invalid');
                if (response.fail == false) { 
                    Swal.fire({
                        title: `Berhasil!`,
                        showConfirmButton: false,
                        icon: 'success',
                        html: `Produk Baru Berhasil Disimpan!
                            <br><br>
                            <a href="`+ laroute.route('admin.produk') +`" class="btn btn-keluar btn-alt-danger"><i class="si si-close mr-1"></i>Keluar</a> 
                            <a href="`+ laroute.route('admin.produk.tambah') +`" class="btn btn-tambah_baru btn-alt-primary"><i class="si si-plus mr-1"></i>Tambah Produk Lain</a>`,
                        showCancelButton: false,
                        showConfirmButton: false,
                        allowOutsideClick: false,
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

    $(document).on('click', '#field-spek #btn-add-spek', function () {
        var spek = $('#field-spek');
        var spek_count = parseInt($('#field-spek_count').val());
        spek.find('tbody').append(`
        <tr>
            <td>
                <input type="text" class="form-control" name="spek[${ spek_count }][nama]" placeholder="Nama" autocomplete="off">
            </td>
            <td>
                <input type="text" class="form-control" name="spek[${ spek_count }][value]" placeholder="Nilai" autocomplete="off">
            </td>
            <td width="5%">
                <button type="button" class="btn btn-sm btn-alt-danger btn-delete-spek"><i class="si si-trash"></i></button>
            </td>
        </tr>`);
        spek_count += 1;
        $('#field-spek_count').val(spek_count);
    });

    $(document).on('click', '#field-spek .btn-delete-spek', function () {
        $(this).closest('tr').remove();
        var spek_count = parseInt($('#field-spek_count').val());
        spek_count -= 1;
        $('#field-spek_count').val(spek_count);
    });
});


function list_item_highlight(el){
    $(el).parent().children().each(function(){
        $(this).removeClass('selected');
    });
    $(el).addClass('selected');
}

function get_subcategories_by_category(el, cat_id){
    list_item_highlight(el);
    category_id = cat_id;
    kategori_id = cat_id;
    subcategory_id = null;
    subsubcategory_id = null;
    category_name = $(el).html();
    $('#subcategories').html(null);
    $('#subsubcategory_list').hide();
    $.ajax({
        type:"POST",
        url: laroute.route('kategori.sub_kategori_json'),
        data: {
            category_id:category_id
        },
        success: function(data){
            for (var i = 0; i < data.length; i++) {
                $('#subcategories').append('<li onclick="get_subsubcategories_by_subcategory(this, '+data[i].id+')">'+data[i].nama+'</li>');
            }
            $('#subcategory_list').show();
        }
    });
}

function get_subsubcategories_by_subcategory(el, subcat_id){
    list_item_highlight(el);
    subcategory_id = subcat_id;
    kategori_id = subcat_id;
    subsubcategory_id = null;
    subcategory_name = $(el).html();
    $('#subsubcategories').html(null);
    $.ajax({
        type:"POST",
        url: laroute.route('kategori.sub_kategori_json'),
        data: {
            category_id:subcategory_id
        },
        success: function(data){
            for (var i = 0; i < data.length; i++) {
                $('#subsubcategories').append('<li onclick="confirm_subsubcategory(this, '+data[i].id+')">'+data[i].nama+'</li>');
            }
            $('#subsubcategory_list').show();
        }
    });
}

function confirm_subsubcategory(el, subsubcat_id){
    list_item_highlight(el);
    subsubcategory_id = subsubcat_id;
    kategori_id = subsubcat_id;
    subsubcategory_name = $(el).html();
}

function filterListItems(el, list){
    filter = el.value.toUpperCase();
    li = $('#'+list).children();
    for (i = 0; i < li.length; i++) {
        if ($(li[i]).html().toUpperCase().indexOf(filter) > -1) {
            $(li[i]).show();
        } else {
            $(li[i]).hide();
        }
    }
}

function closeModal(){
    if(category_id > 0){
        $('#kategori_id').val(kategori_id);
        $('#category_id').val(category_id);
        $('#subcategory_id').val(subcategory_id);
        $('#subsubcategory_id').val(subsubcategory_id);
        if(category_id != null)
        {
            kategorinya = category_name;
        }
        if(subcategory_id != null)
        {
            kategorinya += ' > '+ subcategory_name;
        }
        if(subsubcategory_id != null)
        {
            kategorinya += ' > '+ subsubcategory_name;
        }
        $('#field-kategori').html(kategorinya);
        $('#categorySelectModal').modal('hide');
    }
    else{
        alert('Please choose categories...');
    }
}

function update_variasi(){
    $.ajax({
       type:"POST",
       url: laroute.route('admin.variasi_update'),
       data: {
           pil1 :  $('#field-var1_pilihan').val(),
           var2_status : $('#var2_status').val(),
           pil2 : $('#field-var2_pilihan').val(),
       },
       success: function(response){
            $('#tbl_variasi tbody').html(response.html);

            $("#tbl_variasi").rowspanizer({
                vertical_align: 'middle'
            });
       }
   });
}