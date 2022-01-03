const currentPath = location.pathname
const CPSplit = currentPath.split("/")  
const resellerID = CPSplit[4]

jQuery(document).ready(function(){
    $.ajax({
        url:`/admin/reseller/get/${resellerID}`,
        success:function(res){
            const data = res.data
            const {nama, email, kontak, alamat, seller_id, kota, kota_id} = data
            $("#field-id").val(resellerID)
            $("#field-nama").val(nama)
            $("#field-email").val(email)
            $("#field-kontak").val(kontak)
            $("#field-alamat").val(alamat)
            $("#kota").val(kota)
            $("#kota_id").val(kota_id)
            $("#field-id_seller").val(seller_id)

            $.ajax({
                url:"/rajaongkir/kota",
                success: function(res){
                    var data = JSON.parse(res.data)
                    var { results } = data.rajaongkir
                    results.forEach(element => {
                        $("#field-kota").append(`<option value='${JSON.stringify(element)}'"  ${element.city_id == kota_id ? "selected" : ""}>${element.type} ${element.city_name}</option>`)
                    })
                }
            })
        },
        error: function (jqXHR, textStatus, errorThrown) {
            Swal.close()
            alert('Error adding / update data')
        }
    })

    $('#form-mitra').submit(function(e){
        e.preventDefault()
        var data = new FormData($('#form-mitra')[0])
        $.ajax({
            url : "/admin/reseller/update",
            type: 'PUT',
            data,
            cache: false,
            contentType: false,
            processData: false,
            beforeSend: function(){
                Swal.fire({
                    title: 'Tunggu Sebentar...',
                    text: ' ',
                    imageUrl: '/img/loading.gif',
                    showConfirmButton: false,
                    allowOutsideClick: false,
                })
            },
            success: function (response) {
                $('.is-invalid').removeClass('is-invalid')
                if (response.success) {
                    Swal.fire({
                        title: `Berhasil!`,
                        showConfirmButton: false,
                        icon: 'success',
                        html: `Mitra Berhasil Diupdate!
                            <br><br>
                            <a href="/admin/reseller" class="btn btn-keluar btn-alt-danger"><i class="si si-close mr-1"></i>Keluar</a>`,
                        showCancelButton: false,
                        showConfirmButton: false,
                        allowOutsideClick: false,
                    })
                }else{
                    alert("terjadi kesalahan silahkan hubingi developer")
                }
            },
            error: function (jqXHR, textStatus, errorThrown) {
                Swal.close()
                alert('Error adding / update data')
            }
        })
    })
})

function chooseAlamtEvent(el) {
    var data = JSON.parse($(el).val())
    $("#kota").val(data.city_name)
    $("#kota_id").val(data.city_id)
}