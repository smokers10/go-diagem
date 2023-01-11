var toRupiah = new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    maximumFractionDigits: 0,
})

$('#data-barang').DataTable({
    "ajax":"/admin/produk/get",
    "columns":[
        {
            "data":null,
            "render": function(data, type, row) {
                return `
                <div class="media">
                <img class="media-object" src="${ data.produk_single_foto.path ? data.produk_single_foto.path : "/img/placeholder.png" }" style="width:45px">
                    <div class="media-body ml-3">
                        <div class="font-size-h6 font-w700 mb-0">
                            <a href="/admin/produk/edit/${data.id}" target="_blank">${ data.nama }</a>
                        </div>
                        <span class="text-primary">${ data.kategori.nama }</span>
                    </div>
                </div>
                `
            }
        },
        {
            "data":null,
            "render":function(data, type, row){
                let { variasi } = data

                if (typeof variasi == "object" || typeof variasi == Object) {
                    let firstItem = variasi[0]
                    let lastItem = variasi[variasi.length - 1]
                    return `${toRupiah.format(firstItem.harga)} - ${toRupiah.format(lastItem.harga)}`
                } 

                return toRupiah.format(data.harga)
            }
        },
        {
            "data":null,
            "render":function(data, type, row){
                return data.stok ? data.stok : `<a href="/admin/produk/edit/${data.id}">lihat detail</a>`
            }
        },
        {
            "data":null,
            "orderable":false,
            "render":function(data, type, row){
                return `
                <div class="btn-group" role="group">
                    <button type="button" class="btn btn-secondary btn-sm dropdown-toggle" id="actionButton-${ data.id }" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                        OPSI
                    </button>

                    <div class="dropdown-menu" aria-labelledby="actionButton-${ data.id }">
                        <a class="dropdown-item" href="/admin/produk/edit/${ data.id }">
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
            url: '/admin/produk/delete',
            type: "DELETE",
            data: {id},
            beforeSend: function(){
                Swal.fire({
                    title: 'Tunggu Sebentar...',
                    text: ' ',
                    imageUrl: '/img/loading.gif',
                    showConfirmButton: false,
                    allowOutsideClick: false,
                })
            },
            success: function(data) {
                Swal.fire({
                    title: "Berhasil",
                    text: 'Produk Berhasil Dihapus!',
                    timer: 3000,
                    showConfirmButton: false,
                    icon: 'success'
                })
                setTimeout(() => {
                    location.reload()
                }, 3000)
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