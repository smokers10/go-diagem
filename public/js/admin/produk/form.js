jQuery(function() {
    $("#form-product").on("submit", function (e) {
        e.preventDefault();
        var formData = new FormData($('#form-product')[0]);

        if($('#metode').val() === 'update'){
            link = laroute.route('admin.produk.update')
        }else{
            link = laroute.route('admin.produk.simpan')
        }

        $.ajax({
            url: link,
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
                        text: "Kategori Baru Berhasil Ditambahkan",
                        timer: 1500,
                        showConfirmButton: false,
                        icon: 'success'
                    });
                    $('#kategori-tree').jstree().refresh();
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
    
});