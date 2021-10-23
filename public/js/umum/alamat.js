jQuery(function() {
    _loadContent()

    // triger tambah alamat modal
    $("#btn-add_alamat").click(()=>{
        $("#alamatModal").modal('show')
        $("#alamatModal-title").text("Tambah Alamat Baru")
        $("#alamatModalFrm").trigger("reset")
        $("#field-operation").val("create")
    })

    // submit add
    $("#alamatModalFrm").submit(function(e){
        e.preventDefault()
        const formData = new FormData($('#alamatModalFrm')[0])
        var operation = $("#field-operation").val()
        if (operation == "create") {
            _insert(formData)
        }
        
        if (operation == "update") {
            _update(formData)
        }
    })
})

function _loadContent(){
    $.ajax({
        url: "/alamat/read",
        type: "GET",
        dataType: "JSON",
        success: function(response) {
            var dataList = ''
            if(response.data.length !== 0){
                $.each(response.data, function(k, item) {
                    dataList += _createElement(item)
                })
            }else{
                dataList += `<div class="height-380 py-5 text-center">
                    <img class="empty-img" src="/img/placeholder/alamat.png" width="200px">
                    <div>
                        <h3 class="font-size-24 font-weight-bold mt-5">Alamat Pengiriman Belum Ditambahkan</h3>
                        <p class="font-size-16"></p>
                        <button type="button" class="btn btn-primary btn-lg" id="btn-add_alamat">
                            <i class="fa fa-plus mr-1"></i>Tambah Alamat</button>
                    </div>
                </div>`
            }
            $('#data-list').append(dataList)
        },
        complete :function(){
            $('#dataku-loading').remove()
        },
        error: function(jqXHR, textStatus, errorThrown) {
            alert('Error deleting data')
        }
    })
}

function _createElement(item){
    var row = `
    <div id="alamat-${item.id}" class="block block-bordered block-rounded border-2x mb-0 mb-1 `+ (item.is_utama == 1 ? `border-primary` : ``) +` ">
        <div class="block-header pb-0">
            <h3 class="block-title font-size-h6">
                <span class="font-w700">${ item.nama }</span>
                ${ item.is_utama == 1 ? `<div class="badge badge-secondary">Utama</div>` : "" }
            </h3>
            <div class="block-options">
                <button type="button" class="btn btn-sm btn-secondary btn-edit_alamat" data-id="${ item.id }" data-kode-pos=${item.kd_pos} data-nama="${item.nama}" data-penerima="${item.penerima}" data-phone="${item.phone}" data-alamat="${item.alamat}" onclick="_editAlamat(this)">Ubah</button>
                <button type="button" class="btn btn-sm btn-secondary" data-id="${item.id}" onclick="_hapusAlamat(this)">Hapus</button>
                ${item.is_utama != 1 ? `<button class="btn btn-secondary btn-sm">Jadikan Alamat Utama</button>` : ``}
            </div>
        </div>
        <div class="block-content block-content-full">
            <div class="row justify-content-space-between">
                <div class="col-lg-8">
                    <h4 class="font-size-18 font-w700 mb-0 nice-copy">${ item.penerima }</h4>
                    <p class="mb-0 nice-copy">${ item.phone }</p>
                    <p class="mb-0 nice-copy">${ item.alamat }</p>
                </div>
            </div>
        </div>
    </div>`

    return row
}

function _hapusAlamat(e) {
    const selected = $(e)
    var selectedID = selected.attr("data-id")
    Swal.fire({
        title: 'Hapus Alamat Ini?',
        text: 'Kamu Yakin Ingin Menghapus Alamat Ini?',
        showDenyButton: true,
        showCancelButton: false,
        confirmButtonText: 'Hapus',
        denyButtonText: `Batal`,
    }).then((result) => {
        if (result.isConfirmed) {
            $.ajax({
                url: "/alamat/delete",
                method:"delete",
                data: {
                    id : selectedID,
                },
                beforeSend: function(){
                    Swal.fire({
                        title: 'Tunggu Sebentar...',
                        text: 'Sedang Menghapus Alamat Baru',
                        imageUrl: '/img/loading.gif',
                        showConfirmButton: false,
                        allowOutsideClick: false,
                    })
                },
                success: function (response) {
                    if (response.success) {
                        Swal.fire({
                            title: "Berhasil",
                            text: "Alamat Baru Berhasil Disimpan",
                            timer: 1500,
                            showConfirmButton: false,
                            icon: 'success'
                        })
                        
                        $(`#alamat-${selectedID}`).remove()
                    }else {
                        Swal.close()
                    }
                },
            })
        }
    })
}

function _editAlamat(e) {
    const selected = $(e)
    let id = selected.attr("data-id")
    let nama = selected.attr("data-nama")
    let penerima = selected.attr("data-penerima")
    let phone = selected.attr("data-phone")
    let alamat = selected.attr("data-alamat")
    let kdPos = selected.attr("data-kode-pos")
    $("#alamatModal-title").text("Edit Alamat")
    $("#field-nama").val(nama)
    $("#field-penerima").val(penerima)
    $("#field-alamat").val(alamat)
    $("#field-kode-pos").val(kdPos)
    $("#field-phone").val(phone)
    $("#field-id").val(id)
    $("#alamatModal").modal('show')
    $("#field-operation").val("update")
}

// ajax request - simpan alamat baru
function _insert(formData) {
    $.ajax({
        url: "/alamat/create",
        method: "post",
        data : formData,
        cache: false,
        contentType: false,
        processData: false,
        beforeSend: function(){
            Swal.fire({
                title: 'Tunggu Sebentar...',
                text: 'Sedang Menyimpan Alamat Baru',
                imageUrl: '/img/loading.gif',
                showConfirmButton: false,
                allowOutsideClick: false,
            })
            $("#alamatModal").modal('hide')
        },
        success: function (response) {
            if (response.success) {
                Swal.fire({
                    title: "Berhasil",
                    text: "Alamat Baru Berhasil Disimpan",
                    timer: 3000,
                    showConfirmButton: false,
                    icon: 'success'
                })

                // appending alamat baru
                let newElement = _createElement(response.data)
                $('#data-list').append(newElement)
            }else {
                Swal.close()
            }
        },
        error: function (jqXHR, textStatus, errorThrown) {
            Swal.close()
            alert('Error adding data')
        }
    })
} 

// ajax request - update alamat
function _update(formData) {
    $.ajax({
        url: "/alamat/update",
        method: "put",
        data : formData,
        cache: false,
        contentType: false,
        processData: false,
        beforeSend: function(){
            Swal.fire({
                title: 'Tunggu Sebentar...',
                text: 'Sedang Memperbaharui Alamat',
                imageUrl: '/img/loading.gif',
                showConfirmButton: false,
                allowOutsideClick: false,
            })
            $("#alamatModal").modal('hide')
        },
        success: function (response) {
            if (response.success) {
                Swal.fire({
                    title: "Berhasil",
                    text: "Alamat Berhasil Diperbaharui",
                    timer: 3000,
                    showConfirmButton: false,
                    icon: 'success'
                })

                // update element alamat
                let data = response.data
                let updateElement = _createElement(data)
                $(`#alamat-${data.id}`).html(updateElement)
            }else {
                Swal.close()
            }
        },
        error: function (jqXHR, textStatus, errorThrown) {
            Swal.close()
            alert('Error adding data')
        }
    })
}

