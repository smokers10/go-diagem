// get current URL path
const currentPath = location.pathname
const CPSplit = currentPath.split("/")  
const userID = CPSplit[4]

// get selected user-administratif data
$.ajax({
    url: `/admin/user-administratif/get/${userID}`,
    success: function(response) {
        const data = response.data
        const {nama, email, username, roles} = data
        $("#field-name").val(nama)
        $("#field-username").val(username)
        $("#field-email").val(email)
        $("#field-role").val(roles)
    },
    error: function (jqXHR, textStatus, errorThrown) {
        Swal.close()
        alert('Error adding / update data')
    }
})

jQuery(document).ready(function () {
    $("#form-user-admin").on('submit', function (e) {
        e.preventDefault()
        var url, type;
        var data = JSON.stringify({
            id: parseInt(userID),
            nama: $("#field-name").val(),
            username: $("#field-username").val(),
            email: $("#field-email").val(),
            password: $("#field-password").val(),
            roles: $("#field-role").val(),
        })

        if($('#method').val() === 'update'){
            url = "/admin/user-administratif/update"
            type = "PUT"
        }else{
            url = "/admin/user-administratif/create"
            type = "POST"
        }

        // submit form
        $.ajax({
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
                $('.is-invalid').removeClass('is-invalid')
                if (response.success) {
                    $('#modal_embed').modal('hide')
                    Swal.fire({
                        title: `Berhasil!`,
                        showConfirmButton: false,
                        icon: 'success',
                        html: `
                            ${ type == "PUT" ? "User Administratif Baru Berhasil Diupdate!" : "User Administratif Baru Berhasil Disimpan!"}
                            <br><br>
                            <a href="/admin/user-administratif" class="btn btn-keluar btn-alt-danger"><i class="si si-close mr-1"></i>Keluar</a>
                            ${ type == "POST" ? `<a href="/admin/user-administratif/tambah" class="btn btn-tambah_baru btn-alt-primary"><i class="si si-plus mr-1"></i>Tambah User Admin Lain</a>` : ""}`,
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
    
    $("#form-update-password").on("submit", function(e){
        e.preventDefault()
        var data = JSON.stringify({
            id: parseInt(userID),
            password: $("#field-password").val(),
        })
       
        // submit change password
        $.ajax({
            url:"/admin/user-administratif/update-password",
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
                Swal.close()
                $('.is-invalid').removeClass('is-invalid')
                if (response.success) {
                    $('#modal_embed').modal('hide')
                    Swal.fire({
                        icon: 'success',
                        title: 'Password Berhasil Diupdate',
                        showConfirmButton: false,
                        timer: 1500
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