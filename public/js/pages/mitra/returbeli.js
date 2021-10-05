$(document).ready(function () {
    moment.locale('id');
    var start =  moment().startOf('week');
    var end = moment();

    $('#tgl_range').daterangepicker({
        startDate: start,
        endDate: end,
        ranges: {
            'Hari Ini': [moment(), moment()],
            'Kemarin': [moment().subtract(1, 'days'), moment().subtract(1, 'days')],
            '7 Hari Yang Lalu': [moment().subtract(6, 'days'), moment()],
            '30 Hari Yang Lalu': [moment().subtract(29, 'days'), moment()],
            'Bulan Sekarang': [moment().startOf('month'), moment().endOf('month')],
            'Bulan Lalu': [moment().subtract(1, 'month').startOf('month'), moment().subtract(1, 'month').endOf('month')]
        },
        locale : {
            "customRangeLabel": "Custom Tanggal",
        }
    }, laksanakan);

    laksanakan(start, end);

    $(document).on('change', '#tgl_range', function () {
        var tgl_mulai = $('#tgl_mulai').val();
        var tgl_akhir = $('#tgl_akhir').val();
        var page = 1;
        $('#hidden_page').val(1)
        fetch_data(page, tgl_mulai, tgl_akhir);
    });

    $(document).on('click', '.sorting', function () {
        var column_name = $(this).data('column_name');
        var order_type = $(this).data('sorting_type');
        var reverse_order = '';
        if (order_type == 'asc') {
            $(this).data('sorting_type', 'desc');
            reverse_order = 'desc';
            clear_icon();
            $('#' + column_name + '_icon').html('<span class="glyphicon glyphicon-triangle-bottom"></span>');
        }
        if (order_type == 'desc') {
            $(this).data('sorting_type', 'asc');
            reverse_order = 'asc';
            clear_icon
            $('#' + column_name + '_icon').html('<span class="glyphicon glyphicon-triangle-top"></span>');
        }
        $('#hidden_column_name').val(column_name);
        $('#hidden_sort_type').val(reverse_order);
        var page = $('#hidden_page').val();
        var query = $('#serach').val();
        fetch_data(page, reverse_order, column_name, query);
    });

    $(document).on('click', '.pagination a', function (event) {
        event.preventDefault();
        var page = $(this).attr('href').split('page=')[1];
        $('#hidden_page').val(page);
        var tgl_mulai = $('#tgl_mulai').val();
        var tgl_akhir = $('#tgl_akhir').val();

        $('li').removeClass('active');
        $(this).parent().addClass('active');
        fetch_data(page, tgl_mulai, tgl_akhir);
    });

    $(document).on('click', 'table#list-riwayat tbody.data_transaksi tr', function () {
        var transaksi_id = $('table#list-riwayat tbody tr').data('transaksi_id');
        document.location = laroute.route('mitra.returbeli.detail', { transaksi_id : transaksi_id });
    });
});


function laksanakan(start, end) {
    if(start.format('D MMMM, YYYY') == end.format('D MMMM, YYYY'))
    {
        tampil = start.format('D MMMM, YYYY');
    }else{
        tampil = start.format('D MMMM, YYYY') + ' - ' + end.format('D MMMM, YYYY');
    }
    $('input#tgl_mulai').val(start.format('YYYY-MM-DD'));
    $('input#tgl_akhir').val(end.format('YYYY-MM-DD'));
    $('#tgl_range').val(tampil);
}

function clear_icon() {
    $('#id_icon').html('');
    $('#post_title_icon').html('');
}

function fetch_data(page, tgl_mulai, tgl_akhir) {
    $.ajax({
        url: laroute.route('mitra.returbeli.index'),
        data: {
            page: page,
            tgl_mulai : tgl_mulai,
            tgl_akhir : tgl_akhir,
        },
        success: function (response) {
            $('.tot_transaksi').html(response.total_transaksi);
            $('.jml_pembelian').html(response.penjualan_kotor);
            $('table#list-riwayat tbody').remove();
            $('table#list-riwayat tfoot').remove();
            $('table#list-riwayat thead').after(response.html);
        }
    })
}