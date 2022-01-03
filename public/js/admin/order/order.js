$(function () {
    var table = $('#list-order').DataTable({
        processing: true,
        serverSide: true,
        ajax: {
          url: laroute.route('admin.order'),
          data: function (d) {
                d.status = $('#status').val(),
                d.search = $('input[type="search"]').val()
            }
        },
        columns: [
            {
                data: 'tgl_transaksi',
                name: 'tgl_transaksi'
            },
            {
                data: 'invoice_no',
                name: 'invoice_no'
            },
            {
                data: 'status',
                name: 'status'
            },
            {
                data: 'produk',
                name: 'produk'
            },
            {
                data: 'total_frm',
                name: 'total_frm'
            },
            {
                data: 'customer',
                name: 'customer'
            },
            // {
            //     data: 'aksi',
            //     name: 'aksi',
            //     orderable: false,
            //     searchable: false
            // },
        ]
    })
    $(".pencarian").keyup(function(){
        table.draw()
    })
})

// $(document).on('click', '.menu_tabs', function () {
//     status = $(this).attr("data-status");
//     $tabs = $(this);
//     // $tabs.addClass("active");
//     $('.menu_tabs').each(function(index, currentElement) {
//         if($(this).attr("data-status") !== status)
//         {
//             $(this).removeClass('active');
//         }else
//         {
//             $tabs.addClass('active');
//         }
//     });
//     $.ajax({
//         type: "GET",
//         url: laroute.route('mitra.order'),
//         data: {
//             status: status,
//             tgl_mulai : '17-18-2020',
//             tgl_akhir : '17-18-2020',
//         },
//         success: function (response) {
//             $('#tbl_variasi tbody').html(response.html);
//         }
//     });
//     window.history.replaceState(null, null, "?status="+status);
// });