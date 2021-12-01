// get current URL path
const currentPath = location.pathname
const CPSplit = currentPath.split("/")  
const userID = CPSplit[4]

$(document).ready(function() {
    $.ajax({
        url:`/admin/user/get/${userID}`,
        success: function(res) {
            const data = res.data
            const { nama, email, hp, tahun, bulan, tanggal, alamat, kd_pos  } = data
            $("#field-nama").val(nama)
            $("#field-kontak").val(hp)
            $("#field-email").val(email)
            if ( tahun && bulan && tanggal) {
                $("#field-tgl-lahir").val(`${tahun}/${bulan}/${tanggal}`)
            }

            if (alamat) {
                for (let i = 0; i < alamat.length; i++) {
                    $("#alamat-row").append(`
                        <div class="col-md-6">
                            <h2>Alamat Ke-${i+1} ${ alamat[i].is_utama ? "(Utama)" : "" }</h2>
                            <div class="form-group">
                                <label for="field-nama-alamat">Nama Alamat</label>
                                <input type="text" class="form-control" id="field-nama-alamat" name="nama-alamat" autocomplete="off" value="${alamat[i].nama}" disabled>
                            </div>

                            <div class="form-group">
                                <label for="field-alamat-lengkap">Alamat Lengkap</label>
                                <input type="text" class="form-control" id="field-alamat-lengkap" name="alamat-lengkap" autocomplete="off" value="${alamat[i].alamat}" disabled>
                            </div>

                            <div class="form-group">
                                <label for="field-kode-pos">Kode Pos</label>
                                <input type="text" class="form-control" id="field-kode-pos" name="kode-pos" autocomplete="off" value="${alamat[i].kd_pos}" disabled>
                            </div>

                            <div class="form-group">
                                <label for="field-nama-penerima">Nama Penerima</label>
                                <input type="text" class="form-control" id="field-nama-penerima" name="nama-penerima" autocomplete="off" value="${alamat[i].penerima}" disabled>
                            </div>

                            <div class="form-group">
                                <label for="field-kontak-penerima">Kontak Penerima</label>
                                <input type="text" class="form-control" id="field-kontak-penerima" name="kontak-penerima" autocomplete="off" value="${alamat[i].phone}" disabled>
                            </div>                                    
                        </div>
                    `)
                }
            }
        }
    })
})