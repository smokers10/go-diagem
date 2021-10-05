$(document).ready(function() {
    var voucher_type = '';
    $("#field-type").on("change", function(event) {
        voucher_type = $(this).val();
        $.ajax({
            type:"POST",
            url: laroute.route('admin.voucher.form_view'),
            data: {
                voucher_tipe : voucher_type
            },
            success: function(data){
                $('#voucher').html(data);

                $('#field-start_date').datetimepicker({
                    format: 'DD-MM-YYYY HH:mm',
                    useCurrent : true,
                    locale : moment.locale('id'),
                    minDate: moment(),
                    icons: {
                        time: 'fa fa-clock',
                        date: 'fa fa-calendar',
                        up: 'fa fa-arrow-up',
                        down: 'fa fa-arrow-down',
                        previous: 'fa fa-chevron-left',
                        next: 'fa fa-chevron-right',
                        today: 'fa fa-calendar-check',
                        clear: 'fa fa-trash-alt',
                        close: 'fa fa-times-circle'
                    }
                });
                $('#field-end_date').datetimepicker({
                    format: 'DD-MM-YYYY HH:mm',
                    useCurrent : true,
                    locale : moment.locale('id'),
                    minDate: moment(),
                    icons: {
                        time: 'fa fa-clock',
                        date: 'fa fa-calendar',
                        up: 'fa fa-arrow-up',
                        down: 'fa fa-arrow-down',
                        previous: 'fa fa-chevron-left',
                        next: 'fa fa-chevron-right',
                        today: 'fa fa-calendar-check',
                        clear: 'fa fa-trash-alt',
                        close: 'fa fa-times-circle'
                    }
                });
                $("#field-start_date").on("dp.change", function (e) {
                    $('#field-end_date').data("DateTimePicker").minDate(e.date);
                });
            }
        });
    });
    
    $('#add_produk').click(function () {
        $('#list-produk').DataTable({
            processing: true,
            serverSide: true,
            ajax: laroute.route('admin.produk'),
            ordering: false,
            columns: [
                {
                    data: 'nama',
                    name: 'nama'
                },
                {
                    data: 'produk',
                    name: 'produk'
                },
                {
                    data: 'aksi',
                    name: 'aksi',
                    orderable: false,
                    searchable: false
                },
            ]
        });
    });

    $("#form-voucher").validate({
        onfocusout: function(element) {
            $(element).valid()
            if ($(element).valid()) {
                $('#form-voucher').find('button:submit').prop('disabled', false);  
            } else {
                $('#form-voucher').find('button:submit').prop('disabled', 'disabled');
            }
        },    
        errorClass: "invalid-feedback font-size-sm animated fadeInDown",
        errorElement: "div",
        errorPlacement: function (e, n) {
            jQuery(n).parents(".form-group").find('div.invalid-feedback').html(e);
        },
        highlight: function (e) {
            jQuery(e).closest(".form-group").removeClass("is-invalid").addClass("is-invalid");
        },
        success: function (e) {
            jQuery(e).closest(".form-group").removeClass("is-invalid").addClass("is-valid");
            jQuery(e).closest(".form-group").removeClass("is-invalid"), jQuery(e).remove();
        },
        rules: {
            coupon_code: {
                required: true,
            },
            discount: {
                required: true,
            },
            min_buy: {
                required: true,
            },
            kuota: {
                required: true,
            },
        },
        messages: {
            coupon_code: {
                required: "Kode Voucher Wajib Diisi!",
            },
            discount: {
                required: 'Nilai Diskon Wajib Diisi!',
            },
            min_buy: {
                required: "Minimal Belanja Wajib Diisi!",
            },
            kuota: {
                required: 'Kuota Voucher Wajib Diisi!',
            },
        },
        submitHandler: function (form) {
            var fomr = $('form#form-voucher')[0];
            var formData = new FormData(fomr);
            $.ajax({
                type: 'POST',
                url: laroute.route('admin.voucher.simpan'),
                data: formData,
                cache: false,
                contentType: false,
                processData: false,
                beforeSend: function(){
                    Swal.fire({
                        title: 'Tunggu Sebentar...',
                        text: ' ',
                        imageUrl: laroute.url('assets/img/loading.gif', ['']),
                        showConfirmButton: false,
                        allowOutsideClick: false,
                    });
                },
                success: function (response) {
                    if (response.fail == false) {
                        Swal.fire({
                            title: "Berhasil",
                            text: "Anda Akan Segera Dialihkan!",
                            timer: 3000,
                            showConfirmButton: false,
                            icon: 'success'
                        });
                        window.setTimeout(function () {
                            location.reload();
                        }, 1500);
                    } else {
                        Swal.fire({
                            title: "Gagal",
                            text: "Periksa Form Input!",
                            timer: 3000,
                            showConfirmButton: false,
                            icon: 'error'
                        });
                        for (control in response.errors) {
                            $('#login-' + control).addClass('is-invalid');
                            $('#error-' + control).html(response.errors[control]);
                        }
                    }
                }
            });
            return false;
        }
    });
});