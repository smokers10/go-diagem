$.ajax({
    url:"/alamat-origin/footer",
    type: "GET",
    dataType: "JSON",
    success : function (response) {
        const { alamat_lengkap, kd_pos, nama_kota, nama_provinsi } = response.data
        const literalAddress = `${alamat_lengkap}, ${nama_kota}, ${nama_provinsi}, ${kd_pos}`
        $("#alamat-origin").text(literalAddress)
    }
})