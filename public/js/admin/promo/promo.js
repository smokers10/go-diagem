jQuery(document).ready(function () {
    $('#list-promo').DataTable({
        "ajax":"/admin/promo/get",
        "columns":[
            {
                "data": null,
                "render": function(data, type, row){
                    return data.judul
                }
            },
            {
                "data": null,
                "render": function(data, type, row){
                    return `${ data.is_publish ? "Publish" : "Draft" }`
                }
            },
            {
                "data": null,
                "render": function(data, type, row){
                    return `${ data.tgl_mulai } s/d ${ data.tgl_selesai }`
                }
            },
            {
                "data":null,
                "orderable":false,
                "render":function(data, type, row){
                    return `
                    <div class="btn-group" role="group">
                        <button type="button" class="btn btn-secondary btn-sm dropdown-toggle" id="actionButton-${data.id}" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                            OPSI
                        </button>
    
                        <div class="dropdown-menu" aria-labelledby="actionButton-${data.id}">
                            <a class="dropdown-item" href="/admin/slider/edit/${data.id}">
                                <i class="si si-note mr-5"></i>Ubah
                            </a>
                            <a class="dropdown-item btn-hapus" href="javascript:void(0)" onClick="hapus('`+ data.id +`')">
                                <i class="si si-trash mr-5"></i>Hapus
                            </a>
                        </div>
                    </div>
                    `
                }
            }
        ],
    })
})

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
            url: "/admin/promo/delete",
            data : {id},
            type: "DELETE",
            dataType: "JSON",
            beforeSend: function(){
                Swal.fire({
                    title: 'Tunggu Sebentar...',
                    text: ' ',
                    imageUrl: '/img/loading.gif',
                    showConfirmButton: false,
                    allowOutsideClick: false,
                })
            },
            success: function(response) {
                if (response.success) {
                    Swal.fire({
                        title: "Berhasil",
                        text: 'Promo Berhasil Dihapus!',
                        timer: 3000,
                        showConfirmButton: false,
                        icon: 'success'
                    })
                    $('#list-promo').DataTable().ajax.reload()
                }else {
                    alert("kesalahan terjadi hubungi developer")
                }
            },
            error: function(jqXHR, textStatus, errorThrown) {
                Swal.close()
                alert('Error deleting data')
            }
        })
        }else{
            Swal.close()
        }
    })
}