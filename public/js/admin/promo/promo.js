jQuery(document).ready(function () {


    table = $('#list-promo').DataTable({
        processing: true,
        serverSide: true,
        ordering : false,
        ajax: {
            url: laroute.route('admin.post'),
            data: function (d) {
                d.kategori = $("#field-kategori").val()
            }
          },
        columns: [
            {
                data: 'judul',
                name: 'judul'
            },
            {
                data: 'status',
                name: 'status'
            },
            {
                data: 'periode',
                name: 'periode'
            },
            {
                data: 'action',
                name: 'action',
                orderable: false,
                searchable: false
            },
        ],
        fnDrawCallback :function(){
            $('div.dataTables_wrapper div.dataTables_filter').css('text-align', 'left');
            $('div.dataTables_wrapper div.dataTables_filter input').css({'display':'block', 'width':'100%'}).removeClass("form-control-sm");
            $('div.dataTables_wrapper div.dataTables_filter label').css('display','block');
        }
    });
    
    $('#kategori').html('<select class="form-control" name="kategori" id="field-kategori"></select>');
    
    $("#field-kategori").select2({
        placeholder: 'Pilih Kategori',
        allowClear: true,
        ajax: {
            url: laroute.route('admin.postKategori.json'),
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
    }).on('select2:unselecting', function(e) {
        $(this).val(null).trigger('change');
        e.preventDefault();
    });

    $("#field-kategori").change(function(){
        table.draw();
    });
});
function hapus(id) {
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
            url: laroute.route('admin.post.hapus', { id: id }),
            type: "GET",
            dataType: "JSON",
            beforeSend: function(){
                Swal.fire({
                    title: 'Tunggu Sebentar...',
                    text: ' ',
                    imageUrl: laroute.url('assets/img/loading.gif', ['']),
                    showConfirmButton: false,
                    allowOutsideClick: false,
                });
            },
            success: function(data) {
                Swal.fire({
                    title: "Berhasil",
                    text: 'Berita Berhasil Dihapus!',
                    timer: 3000,
                    showConfirmButton: false,
                    icon: 'success'
                });
                $('#list-post').DataTable().ajax.reload();
            },
            error: function(jqXHR, textStatus, errorThrown) {
                Swal.close();
                alert('Error deleting data');
            }
        });
        }else{
            Swal.close();
        }
    });
}