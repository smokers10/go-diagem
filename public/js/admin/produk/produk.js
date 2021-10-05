jQuery(function() {

    load_content();

    // Filter Table
    $('#search-data-list').on('input', function(){
        clearTimeout(this.delay);
        this.delay = setTimeout(function(){
           $(this).trigger('search');
        }.bind(this), 800);
    }).on('search', function(){
        load_content();
        $('#current_page').val(1);
    });

    // Navigation Table
    $('#next-data-list').on('click', function(){
        old = parseInt($('#current_page').val());
        old += 1;
        $('#current_page').val(old);
        load_content();
    });

    $('#prev-data-list').on('click', function(){
        old = parseInt($('#current_page').val());
        old -= 1;
        $('#current_page').val(old);
        load_content();
    });

    $('a.btn-hapus').on('click', function(){
        alert('asdsa');
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
            url: laroute.route('admin.produk.hapus', { id: id }),
            type: "GET",
            dataType: "JSON",
            beforeSend: function(){
                Swal.fire({
                    title: 'Tunggu Sebentar...',
                    text: ' ',
                    imageUrl: laroute.url('public/img/loading.gif', ['']),
                    showConfirmButton: false,
                    allowOutsideClick: false,
                });
            },
            success: function(data) {
                Swal.fire({
                    title: "Berhasil",
                    text: 'Produk Berhasil Dihapus!',
                    timer: 3000,
                    showConfirmButton: false,
                    icon: 'success'
                });
                load_content();
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


function load_content(){
    var keyword = $('#search-data-list').val();
    var page = $('#current_page').val();

    var navNext = $('#next-data-list');
    var navPrev = $('#prev-data-list');
    
    $.ajax({
        url: laroute.route('anggota'),
        type: "GET",
        dataType: "JSON",
        data: {
            keyword: keyword,
            page: page,
        },
        beforeSend: function(){
            $('#data-list tbody tr#loading').removeClass('d-none');
            navNext.prop('disabled', true);
            navPrev.prop('disabled', true);
        },
        success: function(response) {
            $('#data-list tbody tr').not('#data-list tbody tr#loading').remove();

            var row = "";
            if(response.data.length !== 0){
                $.each(response.data, function(k, item) {
                    row += create_element(item);
                });
            }else{
               row = `<tr>
                    <td colspan="5">
                        <div class="text-center">
                            <img class="img-fluid" src="`+ laroute.url('public/img/icon/not_found.png', ['']) +`">
                            <div>
                                <h3 class="font-size-24 font-w600 mt-3">Data Tidak Ditemukan</h3>
                            </div>
                        </div>
                    </td>
                </tr>`;
            }
            $('#data-list tbody').append(row);

            // Table Navigation
            response.next_page_url !== null ? navNext.prop('disabled', false) : navNext.prop('disabled', true);
            response.prev_page_url !== null ? navPrev.prop('disabled', false) : navPrev.prop('disabled', true);
            if(response.total === 0){
                var navigasi = '0 - 0 / 0';
            }else{
                var navigasi = response.from +' - '+ response.to +' / '+ response.total;
            }
            $('#content-nav span').html(navigasi);
        },
        complete: function(){
            $('#data-list tbody tr#loading').addClass('d-none');
        },

        error: function(jqXHR, textStatus, errorThrown) {
            alert('Error deleting data');
        }
    });
}

function create_element(item){
    let row = `<tr>
        <td>#</td>
        <td>
            <div class="media">
            <img class="media-object" src="${ item.fotoUtama }" style="width:45px">
                <div class="media-body ml-3">
                    <div class="font-size-h6 font-w700 mb-0">
                        <a href="${ laroute.route('product.detail', {produk : item.slug }) }" target="_blank">${ item.nama }</a>
                    </div>
                    <span class="text-primary">${ item.kategori.nama }</span>
                </div>
            </div>
        </td>
        <td>
            ${ item.harga }
        </td>
        <td>
            ${ item.stok }
        </td>
        <td class="text-center">
            <div class="btn-group" role="group">
                <button type="button" class="btn btn-secondary btn-sm dropdown-toggle" id="actionButton-${ item.id }" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                    OPSI
                </button>
                <div class="dropdown-menu" aria-labelledby="actionButton-${ item.id }">
                    <a class="dropdown-item" href="${ laroute.route('admin.produk.edit', { id : item.id }) }">
                        <i class="si si-note mr-5"></i>Ubah
                    </a>
                    <a class="dropdown-item btn-hapus" href="javascript:void(0)" onClick="hapus('`+ item.id +`');">
                        <i class="si si-trash mr-5"></i>Hapus
                    </a>
                </div>
            </div>
        </td>
    </tr>`;

    return row;
}