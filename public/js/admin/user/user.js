$(document).ready(function () {
    $('#list-user').DataTable({
    "ajax":"/admin/user/get",
    "columns":[
            {
                "data": null,
                "render": function(data, type, row){
                    return data.nama
                }
            },
            {
                "data": null,
                "render": function(data, type, row){
                    return data.hp
                }
            },
            {
                "data": null,
                "render": function(data, type, row){
                    return data.email
                }
            },
            {
                "data": null,
                "render": function(data, type, row){
                    return `${data.is_verified ? "terverifikasi" : "belum diverifikasi"}`
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
                            <a class="dropdown-item" href="/admin/user/profile/${data.id}">
                                <i class="si si-note mr-5"></i>Ubah
                            </a>
                        </div>
                    </div>
                    `
                }
            }
        ]
    })
})